package client

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	"strings"
	"sync"
	"time"
)

type ResolverCacheItem struct {
	ip  net.IP
	set time.Time
}

type Resolver struct {
	c         *dns.Client
	r         string
	suffix    string
	suffixLen int
	cache     map[string]ResolverCacheItem
	cacheTTL  time.Duration
	cacheLock sync.RWMutex
}

const (
	DialTimeout  = 500 * time.Millisecond
	ReadTimeout  = 1000 * time.Millisecond
	WriteTimeout = 1000 * time.Millisecond
	DefaultTTL   = 3600 * time.Second
	Retries      = 3
)

func NewResolver(r string, suffix string) (*Resolver, error) {

	if !strings.HasPrefix(suffix, ".") {
		suffix = "." + suffix
	}

	return &Resolver{
		c: &dns.Client{
			Net:          "udp",
			DialTimeout:  DialTimeout,
			ReadTimeout:  ReadTimeout,
			WriteTimeout: WriteTimeout,
		},
		cacheTTL:  DefaultTTL,
		r:         r,
		suffix:    dns.Fqdn(suffix),
		suffixLen: len(dns.Fqdn(suffix)),
		cache:     make(map[string]ResolverCacheItem),
	}, nil
}

func (r *Resolver) DeleteFromCache(host string) {
	r.cacheLock.Lock()
	defer r.cacheLock.Unlock()
	if _, ok := r.cache[host]; ok {
		delete(r.cache, host)
	}
}

func (r *Resolver) GetFromCache(host string) (net.IP, bool) {
	r.cacheLock.RLock()
	defer r.cacheLock.RUnlock()

	if ans, ok := r.cache[host]; ok {
		return ans.ip, false
	}

	return nil, false
}

func (r *Resolver) SetToCache(host string, ans net.IP) {
	r.cacheLock.Lock()
	defer r.cacheLock.Unlock()
	r.cache[host] = ResolverCacheItem{
		ip:  ans,
		set: time.Now(),
	}
}

func (r *Resolver) lookupLocal(host string) (net.IP, error) {

	hostFull := host
	if r.suffixLen > 0 {
		if !strings.HasSuffix(host, r.suffix[0:r.suffixLen-1]) {
			hostFull = hostFull + r.suffix
		}
	}

	if ips, err := net.LookupHost(hostFull); err == nil {
		ipFinal := net.ParseIP(ips[0])
		r.SetToCache(host, ipFinal)
		return ipFinal, nil
	} else {
		return nil, err
	}
}

func (r *Resolver) Lookup(host string) (net.IP, error) {

	if r == nil {
		return nil, fmt.Errorf("No valid connection")
	}

	// If ip, return IP.
	if ipn := net.ParseIP(host); ipn != nil {
		return ipn, nil
	}

	// Cache check for first
	if ans, exp := r.GetFromCache(host); ans != nil && exp == false {
		return ans, nil
	} else if exp == false {
		r.DeleteFromCache(host)
	}

	// Get from local resolver instead, if possible
	if ip, err := r.lookupLocal(host); err == nil {
		return ip, err
	}

	if r.r == "" {
		return nil, fmt.Errorf("No valid answers")
	} // If there is an external resolver specified, ask it now.

	m := new(dns.Msg)
	if r.suffixLen > 0 {
		if strings.HasSuffix(host, r.suffix[0:r.suffixLen-1]) {
			if strings.HasSuffix(host, ".") {
				m.SetQuestion(host, dns.TypeA)
			} else {
				m.SetQuestion(host+".", dns.TypeA)
			}
		} else {
			m.SetQuestion(host+r.suffix, dns.TypeA)
		}
	} else {
		m.SetQuestion(host, dns.TypeA)
	}

	for i := 0; i < Retries; i++ {
		if in, _, err := r.c.Exchange(m, r.r); err != nil {
			// Nothing, just let retry
		} else {
			if in != nil && in.Rcode != dns.RcodeSuccess {
				return nil, fmt.Errorf("Could not get a valid response for %s", host)
			}

			if len(in.Answer) == 0 {
				return nil, fmt.Errorf("No valid answers")
			} else {
				if t, ok := in.Answer[0].(*dns.A); ok {
					r.SetToCache(host, t.A)
					return t.A, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("Could not connect to dns server: %s", r.r)
}
