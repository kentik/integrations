package killswitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	DEFAULT_TIME_CHECK_GRACE = 86400
	DEFAULT_TIME_LOOP        = 1 * time.Hour
)

// A Signer is can create signatures that verify against a public key.
type Signer interface {
	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	Sign(data []byte) ([]byte, error)
}

type Verifier interface {
	Verify(data []byte, sig []byte) error
}

type State struct {
	Signature []byte    `json:"s"`
	Issued    time.Time `json:"i"`
	End       time.Time `json:"e"`
}

func (s *State) BytesToSign() ([]byte, error) {
	return json.Marshal(&State{
		Issued: s.Issued,
		End:    s.End,
	})
}

type Killer struct {
	signer         Signer
	verifier       Verifier
	tickFile       string
	issueTime      time.Time
	endTime        time.Time
	timeVerifyLoop time.Duration
	graceSeconds   int
}

func NewDefaultKiller(file string) (*Killer, error) {
	kill, err := NewKiller(0, file, DEFAULT_TIME_LOOP, nil)
	if err != nil {
		return nil, err
	}

	// Verify right now in this case.
	if err := kill.Verify(); err != nil {
		return nil, err
	}

	go kill.Kill()
	return kill, nil
}

func NewKiller(ticks uint32, file string, timeVerifyLoop time.Duration, signer Signer) (*Killer, error) {
	k := Killer{
		tickFile:       file,
		issueTime:      time.Unix(0, 0),
		endTime:        time.Unix(0, 0),
		timeVerifyLoop: timeVerifyLoop,
		graceSeconds:   DEFAULT_TIME_CHECK_GRACE,
	}

	verifier, err := buildPublicKeyVerifier()
	if err != nil {
		return nil, err
	}

	k.signer = signer
	k.verifier = verifier

	if ticks > 0 {
		k.endTime = time.Now().UTC().Add(time.Duration(ticks) * timeVerifyLoop)
	}

	return &k, nil
}

// SetGraceSeconds sets the grace period in seconds - mostly for testing
func (k *Killer) SetGraceSeconds(graceSeconds int) {
	k.graceSeconds = graceSeconds
}

func (k *Killer) IssueTime() time.Time {
	return k.issueTime
}

func (k *Killer) ExpireTime() time.Time {
	return k.endTime
}

func (k *Killer) Create() error {
	if k.signer == nil {
		panic("Cannot create license without Signer")
	}
	os.Remove(k.tickFile)
	k.issueTime = time.Now().UTC()
	err := k.write()
	if err != nil {
		return err
	}

	return k.Verify()
}

func (k *Killer) Kill() {
	for {
		if err := k.Verify(); err != nil {
			log.Fatalf("Verification failed: %v", err)
		}
		time.Sleep(k.timeVerifyLoop)
	}
}

func (k *Killer) Verify() error {
	// Read in file
	data, err := ioutil.ReadFile(k.tickFile)
	if err != nil {
		return err
	}

	state := &State{}
	err = json.Unmarshal(data, state)
	if err != nil {
		return err
	}

	tosign, err := state.BytesToSign()
	if err != nil {
		return err
	}

	err = k.verifier.Verify(tosign, state.Signature)
	if err != nil {
		return err
	}

	// If the issue time is 0, load it
	if k.issueTime.IsZero() || k.issueTime.Unix() < 100 {
		k.issueTime = state.Issued
	}

	// If the end time is 0, load it
	if k.endTime.IsZero() || k.endTime.Unix() < 100 {
		k.endTime = state.End
	}

	now := time.Now().UTC()

	// Are we done?
	if now.Before(k.issueTime) {
		return fmt.Errorf("License not yet valid")
	}
	if now.After(k.endTime) {
		if now.After(k.endTime.Add(time.Second * time.Duration(k.graceSeconds))) {
			return fmt.Errorf("License Expired")
		}
		log.Printf("In Grace Period\n")
	}

	return nil
}

func (k *Killer) write() error {
	state := &State{
		Issued: k.issueTime,
		End:    k.endTime,
	}

	tosign, err := state.BytesToSign()
	if err != nil {
		return err
	}

	if k.signer == nil {
		return fmt.Errorf("Cannot write without Signer")
	}

	// Sign
	state.Signature, err = k.signer.Sign(tosign)
	if err != nil {
		return err
	}

	fullData, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(k.tickFile, fullData, os.FileMode(0600))
}
