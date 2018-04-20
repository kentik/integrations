package log

import "log"

var verbose int

func Debugf(fmt string, v ...interface{}) {
	if verbose > 0 {
		log.Printf(fmt, v...)
	}
}

func SetVerbose(level int) {
	verbose = level
}
