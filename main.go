package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/loosehose/ColumbusCLI/domain"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Initialize zerolog with console writer
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Create new parser object
	parser := argparse.NewParser("Columbus CLI", "Fetches API responses based on user input from columbus.elmasy.com")

	// Create string flag
	lookupArg := parser.String("l", "lookup", &argparse.Options{
		Required: false,
		Help:     "Lookup subdomains for domains",
	})
	startsArg := parser.String("s", "starts", &argparse.Options{
		Required: false,
		Help:     "Find domains that start with the given string",
	})
	tldArg := parser.String("t", "tld", &argparse.Options{
		Required: false,
		Help:     "Find TLDs for the given domain",
	})
	historyArg := parser.String("r", "history", &argparse.Options{
		Required: false,
		Help:     "DNS record history",
	})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// Call appropriate function based on the flag provided
	if *lookupArg != "" {
		domain.Lookup(*lookupArg)
	}
	if *startsArg != "" {
		domain.Starts(*startsArg)
	}
	if *tldArg != "" {
		domain.TLD(*tldArg)
	}
	if *historyArg != "" {
		domain.History(*historyArg)
	}
}

// Here, include the Lookup, Starts, TLD, and History functions
// that we defined previously.
