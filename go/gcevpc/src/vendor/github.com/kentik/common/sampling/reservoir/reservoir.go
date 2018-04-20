package reservoir

import (
	"math/rand"

	"github.com/joeshaw/gengen/generic"
)

type reservoirSample struct {
	weight uint64
	item   generic.T
}

// A Reservoir is a single instance of a weighted reservoir sampling.  The creator specifies
// a target weight, then Add()s an arbitrary number of weighted elements.  On Close(), the
// Reservoir returns a randomly-selected slice of elements whose weight is as close as possible
// to the target weight.  (If the caller calls only Add(), never AddMust(), and the weight argument
// to Add() is always one, this reduces to Knuth's "Algorithm R" -- the weights and AddMust are
// present to allow us to handle unparsed flow-messages and ipfix/v9 templates.)
//
// A caller *could* create only a single reservoir, and generate a single sample over the caller's
// lifetime, but our services are more likely to generate a series of these at a timed interval,
// creating one of these at the start of the interval, adding to it during the interval, then
// closing it at the end of the interval and creating a new one for the next interval.
type Reservoir struct {
	targetWeight  uint64
	currentWeight uint64
	musts         []reservoirSample
	winners       []reservoirSample
	discardWeight uint64
	discardItems  uint64
	release       func(generic.T)
}

// NewReservoir creates a reservoir, specfying a target weight to sample, an optional previously-used
// Reservoir for sizing hints, and an optional function to be called on elements that are discarded
// during the course of sampling.
//
// If the caller passes in the previously-used reservoir, we'll pre-allocate the reservoir slices to
// be the same size as the previously-used one.  Assuming that the targetWeights of the two reservoirs
// are the same, and that the typical weights added to the two reservoirs are similar, this should
// allow us to do minimal reallocation of the slices after the first reservoir's lifetime.
func NewReservoir(targetWeight uint64, previous *Reservoir, release func(generic.T)) *Reservoir {

	mustCap := 0
	winnersCap := 0
	if previous != nil {
		mustCap = cap(previous.musts)
		winnersCap = cap(previous.winners)
	}

	return &Reservoir{
		targetWeight: targetWeight,
		musts:        make([]reservoirSample, 0, mustCap),
		winners:      make([]reservoirSample, 0, winnersCap),
		release:      release,
	}
}

// Add a single item to the reservoir, specifying its weight.
func (r *Reservoir) Add(item generic.T, weight uint64) {
	new := reservoirSample{weight: weight, item: item}

	if r.currentWeight+weight <= r.targetWeight {
		// there's room in the reservoir to add this message without an eviction
		r.winners = append(r.winners, new)
		r.currentWeight += weight
	} else {
		// evict as many times as we need to make room for the new element
		// (except that the new element may be the evicted element each time)
		for len(r.winners) > 0 {
			// evict a random message
			m := rand.Intn(len(r.winners) + int(r.discardItems) + 1)
			if m >= len(r.winners) {
				// The "evicted" message is the new one
				r.discardWeight += weight
				r.discardItems++
				if r.release != nil {
					r.release(new.item)
				}
				break
			} else {
				// evict the selected message, but replace it with the new one
				// only if that eviction frees up enough room
				if r.currentWeight+weight-r.winners[m].weight <= r.targetWeight {
					r.currentWeight += weight - r.winners[m].weight
					r.discardWeight += r.winners[m].weight
					r.discardItems++
					if r.release != nil {
						r.release(r.winners[m].item)
					}
					r.winners[m] = new
					break
				} else {
					// unfortunately, we've evicted a message that's smaller than the one
					// we want to add.  We've got to go through the loop again, but we
					// have to remove the hole in the list first.  Fortunately, we don't
					// care about order here, so just plug the hole with the last element.
					r.currentWeight -= r.winners[m].weight
					r.discardWeight += r.winners[m].weight
					r.discardItems++
					if r.release != nil {
						r.release(r.winners[m].item)
					}
					r.winners[m] = r.winners[len(r.winners)-1]
					r.winners = r.winners[:len(r.winners)-1]
				}
			}
		}
	}
}

// AddMust adds an element to the reservoir that can never be discarded; any element
// passed into AddMust *will* be an element of the slice returned from Close.  The
// element does have a weight, so it displaces elements added with Add.  If the sum
// of the weights passed to AddMust is greater than the original targetWeight, then
// Close may return an actualWeight greater than the original targetWeight (and only
// elements from the AddMust calls -- all Add'ed elements will have been discarded.)
//
// (This is a fairly clunky way to handle ipfix/v9 template, which are required
// for the correct interpretation of flow data, so we must always pass them through
// to the client.  Unfortunately, a single message can hold both template records and
// flow records, so when we AddMust the templates, we're also AddMust'ing the associated
// flow records, and they don't get properly sampled.  This is unfortunate, but
// many senders never put both templates and flow-records in the same packet, and
// others do it only for options-data, which we're going to throw away at the client
// anyway.  I don't expect it to be a significant problem; if it is, we'll deal with
// it down the line.)
func (r *Reservoir) AddMust(item generic.T, weight uint64) {
	new := reservoirSample{weight: weight, item: item}

	// This message is definitely making it into the must list, but we may need
	// to evict one or more regular messages to make room
	r.currentWeight += weight
	r.musts = append(r.musts, new)
	for len(r.winners) > 0 && r.currentWeight > r.targetWeight {
		// evict a random existing message, plugging the hole with the last element
		m := rand.Intn(len(r.winners))
		r.currentWeight -= r.winners[m].weight
		r.discardWeight += r.winners[m].weight
		r.discardItems++
		if r.release != nil {
			r.release(r.winners[m].item)
		}
		r.winners[m] = r.winners[len(r.winners)-1]
		r.winners = r.winners[:len(r.winners)-1]
	}
}

// Close a Reservoir, returning the list of elements that have made it through the sampling process,
// their cumulative weight, the number of elements that were discarded during sampling, and *their*
// cumulative weight.  A caller should never Add or AddMust to a Reservoir after calling Close; the
// only thing you can do with the Reservoir at that point is pass it as a sizing-hint to the next
// NewReservoir call.
func (r *Reservoir) Close() (winners []generic.T, actualWeight uint64, discardItems uint64, discardWeight uint64) {
	nWinners := len(r.musts) + len(r.winners)
	winners = make([]generic.T, 0, nWinners)
	for _, must := range r.musts {
		winners = append(winners, must.item)
	}
	for _, winner := range r.winners {
		winners = append(winners, winner.item)
	}

	return winners, r.currentWeight, r.discardItems, r.discardWeight

}
