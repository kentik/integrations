package main

import (
	"encoding/base64"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	flags "github.com/jessevdk/go-flags"
	"github.com/kentik/killswitch"
)

const (
	TIME_VERIFY_LOOP_DEBUG = 1 * time.Second
	TIME_VERIFY_LOOP       = 1 * time.Hour
	DEFAULT_TICKS          = 8760
)

type Args struct {
	Ticks   uint32 `short:"t" description:"number of ticks (be default 1 hour each)" required:"true"`
	LogFile string `short:"l" description:"log file to write ticks to" required:"true"`
	Create  bool   `short:"c" description:"create tick file" optional:"true"`
	Show    bool   `short:"s" description:"show crypted pk" optional:"true"`
	Debug   bool   `short:"d" description:"run in debug mode" optional:"true"`
	Verify  bool   `short:"v" description:"verify license" optional:"true"`
}

func main() {
	args := &Args{}

	parser := flags.NewParser(args, flags.PassDoubleDash|flags.HelpFlag)
	if _, err := parser.Parse(); err != nil {
		switch err.(*flags.Error).Type {
		case flags.ErrHelp:
			parser.WriteHelp(os.Stderr)
			os.Exit(1)
		default:
			log.Fatal(err)
		}
	}

	if args.Show {
		pk, err := showPrivateKey()
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%s\n", base64.StdEncoding.EncodeToString(pk))
		}
		os.Exit(0)
	}

	var kill *killswitch.Killer
	var err error

	// Only in the case where we are creating do this.
	if args.Ticks == 0 && args.Create {
		args.Ticks = DEFAULT_TICKS
	}

	if args.Verify {
		kill, err = killswitch.NewDefaultKiller(args.LogFile)
	} else {
		signer, err := buildPrivateKeySigner()
		if err != nil {
			log.Fatalf("Error building private key signer: %s", err)
		}
		if args.Debug {
			kill, err = killswitch.NewKiller(args.Ticks, args.LogFile, TIME_VERIFY_LOOP_DEBUG, signer)
		} else {
			kill, err = killswitch.NewKiller(args.Ticks, args.LogFile, TIME_VERIFY_LOOP, signer)
		}
	}

	if err != nil {
		log.Fatalf("Could not init killer: %v", err)
	}

	if args.Create {
		err := kill.Create()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s, Kentik License Expires on %v", args.LogFile, kill.ExpireTime())
		os.Exit(0)
	}

	if args.Verify {
		log.Printf("%s, Kentik License Expires on %v", args.LogFile, kill.ExpireTime())
		os.Exit(0)
	}

	go kill.Kill()

	exit := make(chan bool)
	s := make(chan os.Signal, 2)
	signal.Notify(s, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-exit:
			// Anything to do here?
			os.Exit(0)
		case sig := <-s:
			switch sig {
			case syscall.SIGQUIT:
				go func() { exit <- true }()
			case syscall.SIGINT:
				go func() { exit <- true }()
			case syscall.SIGTERM:
				go func() { exit <- true }()
			}
		}
	}
}
