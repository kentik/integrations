package testing

import (
	"fmt"
	"testing"
	"github.com/kentik/eggs/pkg/logger"
)


// Testing implementations:

func NewTestContextL(lc logger.Context, t *testing.T) logger.ContextL {
	return &logger.ContextLImpl{
		Context: lc,
		L:       &logger.LoggerImpl{UL: &Test{T: t}},
	}
}

// Implements logger.Underlying
type Test struct {
	T *testing.T
}

func (l *Test) Debugf(lp string, f string, params ...interface{}) {
	l.T.Logf("%s DEBUG %s", lp, fmt.Sprintf(f, params...))
}
func (l *Test) Infof(lp string, f string, params ...interface{}) {
	l.T.Logf("%s INFO %s", lp, fmt.Sprintf(f, params...))
}
func (l *Test) Warnf(lp string, f string, params ...interface{}) {
	l.T.Logf("%s WARN %s", lp, fmt.Sprintf(f, params...))
}
func (l *Test) Errorf(lp string, f string, params ...interface{}) {
	l.T.Logf("%s ERROR %s", lp, fmt.Sprintf(f, params...))
}