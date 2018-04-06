package asn

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	dbMaxRetries    = 5
	dbRetryInterval = 500 * time.Millisecond
	dbSoftError     = "40P01"
)

// internal interface to wrap SQL Rows for testing
type asn2NameRowScanner interface {
	Next() bool
	Scan(dest ...interface{}) error
}

// Logger is used for log messages
type Logger interface {
	Infof(prefix, format string, v ...interface{})
	Errorf(prefix, format string, v ...interface{})
}

// Hold error from db
// From https://github.com/lib/pq/blob/master/error.go
type pgError struct {
	Severity         string
	Code             ErrorCode
	Message          string
	Detail           string
	Hint             string
	Position         string
	InternalPosition string
	InternalQuery    string
	Where            string
	Schema           string
	Table            string
	Column           string
	DataTypeName     string
	Constraint       string
	File             string
	Line             string
	Routine          string
}

func (err pgError) Error() string {
	return "pq: " + err.Message
}

type ErrorCode string

// Store manages fetching, storing, and updating ASN data from Postgres.
type Store struct {
	log                Logger
	logPrefix          string
	lock               sync.RWMutex // used when swapping
	pgDB               *sql.DB
	asn2Name           map[uint32]string   // ASN -> name mapping; protected by lock
	name2ASNs          map[string][]uint32 // ASN name -> ids; protected by lock
	lcName2ASNs        map[string][]uint32 // lower-cased ASN name -> ids; protected by lock
	currentDataVersion string              // version of the currently-loaded data; protected by lock
}

// NewStore creates a new ASN.
func NewStore(log Logger, logPrefix string, pgDB *sql.DB) (*Store, error) {
	return &Store{
		log:         log,
		logPrefix:   logPrefix,
		lock:        sync.RWMutex{},
		pgDB:        pgDB,
		name2ASNs:   make(map[string][]uint32),
		lcName2ASNs: make(map[string][]uint32),
		asn2Name:    make(map[uint32]string),
	}, nil
}

func (s *Store) fetchASNData() (*sql.Rows, error) {
	// fetch the updated records outside transaction
	s.log.Infof(s.logPrefix, "Checked database for updated ASN data. New data is available - fetching it.")
	rows, err := s.pgDB.Query("SELECT id, description FROM mn_asn")
	return rows, err
}

func (s *Store) GetASNData() (*sql.Rows, error) {
	for retries := 0; ; retries++ {
		rows, err := s.fetchASNData()
		if err != nil {
			if pgerr, ok := err.(*pgError); ok {
				if pgerr.Code != dbSoftError {
					return nil, fmt.Errorf("Error fetching ASN information from database: %s", err)
				}
			} else {
				return nil, fmt.Errorf("Error fetching ASN information from database, not not get type: %s", err)
			}
		} else {
			return rows, err
		}
		if retries >= dbMaxRetries {
			return nil, fmt.Errorf("Error fetching ASN information from database (max retries exceeded): %s", err)
		}
		time.Sleep(dbRetryInterval * time.Duration(retries+1))
	}
}

// UpdateFromDB updates the internal store from the database.
// This is threadsafe, and will block until done. The underlying
func (s *Store) UpdateFromDB() error {
	// fetch the version in a read-lock
	s.lock.RLock()
	currentDataVersion := s.currentDataVersion
	s.lock.RUnlock()

	start := time.Now()

	// check the database to get the latest import version
	newestVersion, err := s.fetchLatestVersion()
	if err != nil {
		return fmt.Errorf("Error fetching latest ASN dataset version: %s", err)
	}
	if currentDataVersion == newestVersion {
		s.log.Infof(s.logPrefix, "Checked database for updated ASN data. Nothing to do: local cache is already current.")
		return nil
	}

	rows, err := s.GetASNData()
	if err != nil {
		return fmt.Errorf("Error fetching ASN information from database: %s", err)
	} else if rows != nil {
		defer func() {
			if err := rows.Close(); err != nil {
				s.log.Errorf(s.logPrefix, "Error closing rows after querying mn_asn.id|description")
			}
		}()
	}

	recordCount, err := s.consumeASNRows(newestVersion, rows)
	if err != nil {
		return err
	}

	s.log.Infof(s.logPrefix, "Imported/Updated %d ASN records from database in %s", recordCount, time.Now().Sub(start))
	return nil
}

// fetch the most recent version
func (s *Store) fetchLatestVersion() (string, error) {
	// check the database to get the latest version
	rows, err := s.pgDB.Query("SELECT MAX(edate) FROM mn_dataset_version WHERE dataset_name='asn'")
	if err != nil {
		return "", fmt.Errorf("Error querying for most recent ASN import: %s", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			s.log.Errorf(s.logPrefix, "Error closing rows after querying mn_dataset_version.edate")
		}
	}()

	mostRecentVersion := ""
	if rows.Next() {
		err = rows.Scan(&mostRecentVersion)
		if err != nil {
			return "", fmt.Errorf("Error scanning edate from mn_dataset_version: %s", err)
		}
	}
	return mostRecentVersion, nil
}

// consume the results of the (asnID, description) SQL query, updating the Store's state
// when done. It is this method's responsibility to handle locking the internal state.
func (s *Store) consumeASNRows(version string, rows asn2NameRowScanner) (int, error) {
	newASN2Name := make(map[uint32]string)
	newName2ASNs := make(map[string][]uint32)
	newLCName2ASNs := make(map[string][]uint32)
	asnID := uint32(0)
	asnName := ""
	recordCount := 0
	for rows.Next() {
		if err := rows.Scan(&asnID, &asnName); err != nil {
			return 0, fmt.Errorf("Error scanning a row of ASN data: %s", err)
		}

		// add ID -> name lookup
		newASN2Name[asnID] = asnName

		// name -> []ID lookup
		idsForName, ok := newName2ASNs[asnName]
		if !ok {
			idsForName = make([]uint32, 0)
		}
		idsForName = append(idsForName, asnID)
		newName2ASNs[asnName] = idsForName

		// lower-case name -> []ID lookup
		lcName := strings.ToLower(asnName)
		idsForLCName, ok := newLCName2ASNs[strings.ToLower(lcName)]
		if !ok {
			idsForLCName = make([]uint32, 0)
		}
		idsForLCName = append(idsForLCName, asnID)
		newLCName2ASNs[lcName] = idsForLCName

		recordCount++
	}

	// Note: since we didn't use a transaction, it's possible that we've just loaded new data with
	// the old version. This isn't a big deal - data will just be stale till next poll.

	// swap in the new data in a write-lock
	s.lock.Lock()
	s.currentDataVersion = version
	s.asn2Name = newASN2Name
	s.name2ASNs = newName2ASNs
	s.lcName2ASNs = newLCName2ASNs
	s.lock.Unlock()

	return recordCount, nil
}

// AllData returns all ASN->Name data. The returned map is never modified.
func (s *Store) AllData() map[uint32]string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.asn2Name
}

// ASNsByName returns a list of ASN IDs for the input name/description. If 'compare' is true,
// full name must match. Case sensitivity determined by matchCase.
// -1 is used for ASNs that couldn't be found
func (s *Store) ASNsByName(name string, compare bool, matchCase bool) []uint32 {
	// get working instances of the maps so we can let go of the lock ASAP
	s.lock.RLock()
	name2ASNs := s.name2ASNs
	lcName2ASNs := s.lcName2ASNs
	s.lock.RUnlock()

	var asns []uint32
	if compare {
		// match full string
		if matchCase {
			// exact case
			if matches, ok := name2ASNs[name]; ok {
				asns = append(asns, matches...)
			}
		} else {
			// case-insensitive
			lcName := strings.ToLower(name)
			if matches, ok := lcName2ASNs[lcName]; ok {
				asns = append(asns, matches...)
			}
		}
	} else {
		// partial match - need to loop
		if matchCase {
			// exact case
			for k := range name2ASNs {
				if strings.Contains(k, name) {
					asns = append(asns, name2ASNs[k]...)
				}
			}
		} else {
			// case insensitive
			lcName := strings.ToLower(name)
			for k := range lcName2ASNs {
				if strings.Contains(k, lcName) {
					asns = append(asns, lcName2ASNs[k]...)
				}
			}
		}
	}

	return asns
}

// NamesByASNs looks up ASN names for an array of IDs. Return value
// must line up with the request, with empty values for those that could
// not be found.
func (s *Store) NamesByASNs(asns []uint32) []string {
	// get working instances of the map so we can let go of the lock ASAP
	s.lock.RLock()
	asn2Name := s.asn2Name
	s.lock.RUnlock()

	ret := make([]string, len(asns))
	for i := 0; i < len(asns); i++ {
		ret[i] = asn2Name[asns[i]]
	}
	return ret
}
