package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

var (
	// ago = flag.Duration("ago", 0, "time ago")
	ago = flag.String("ago", "0", "time ago")
)

func main() {
	flag.Parse()
	err := run(flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	for _, arg := range args {
		f, err := os.OpenFile(arg, os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return errors.Wrapf(err, "cannot open %s", arg)
		}
		defer f.Close()
		d, _ := time.ParseDuration("-" + *ago)
		t := time.Now().Add(d)
		os.Chtimes(f.Name(), t, t)
	}
	return nil
}
