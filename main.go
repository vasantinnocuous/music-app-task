package main

import (
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

type cmdFlags struct {
	filePath, releasedBefore *string
	minTrackCount            *int
}

func main() {
	flags := initCmdFlags()
	resp := parseFile(flags.filePath)
	filters := getFilters(flags)
	filteredResponse := applyFilters(resp.Records, filters)
	printReleases(filteredResponse)
}

// Parsing command line arguments
func initCmdFlags() cmdFlags {
	var flags cmdFlags
	flags.filePath = flag.String("filepath", "/tmp", "Filepath to read data from about records")
	flags.minTrackCount = flag.Int("minTrackCount", 0, "Minimum number of tracks")
	flags.releasedBefore = flag.String("releasedBefore", "2099.12.31", "Date to before the records are asked for")
	flag.Parse()
	for !flag.Parsed() {
		// Wait till flags being parsed
	}
	return flags
}

// Parsing xml file content
func parseFile(filePath *string) *Records {
	fileContents, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("Couldn't read file contents due to %v", err)
	}
	resp := new(Records)
	err = resp.Unmarshal(fileContents)
	if err != nil {
		log.Fatalf("Couldn't unmarshal xml due to %v", err)
	}
	return resp
}

// Encodes data back to xml, and outputs on the standard console
func printReleases(releases []Release) {
	err := xml.NewEncoder(os.Stdout).Encode(&MatchingReleases{Releases: releases})
	if err != nil {
		log.Fatalf("Couldn't encode releases due to %v", err)
	}

}
