package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type ExObject struct {
	Sub []string `json:"sub"`
	Ip  []string `json:"ip"`
	cnm []string `json:"cnm"`
}

func main() {

	var inputJSONString string
	flag.StringVar(&inputJSONString, "f", "", "input JSON file for execlude object")

	var filterMode int
	flag.IntVar(&filterMode, "mode", 0, "filter mode from 1-4")

	flag.Parse()

	// Unmarshal the JSON data into a JSONObject
	var inputJSON ExObject
	if err := json.Unmarshal([]byte(inputJSONString), &inputJSON); err != nil {
		fmt.Println("Error parsing the JSON object:", err)
		os.Exit(1)
	}

	// Read the list of URLs from stdin
	inputSubs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading subdomains from stdin:", err)
		os.Exit(1)
	}

	// Split the list of URLs by newline
	subList := strings.Split(string(inputSubs), "\n")

	if filterMode == 0 {
		fmt.Fprintln(os.Stderr, "--mode 1: filter by subdmoains, input [sub]")
		fmt.Fprintln(os.Stderr, "--mode 2: filter by IP address, input [sub, ip]")
		fmt.Fprintln(os.Stderr, "--mode 3: filter by cname, input: [sub, cname]")
	} else if filterMode == 1 {

		// Iterate over the list of URLs
		for _, iSub := range subList {
			// Set matched to false initially
			matched := false

			// Iterate over the list of regexes in inputJSON.Sub
			for _, fSub := range inputJSON.Sub {
				fSubRegex, err := regexp.Compile(fSub+"$")
				if err != nil {
					fmt.Println("Error compiling regex:", err)
					os.Exit(1)
				}

				// Check if the iSub matches the fSub regex
				if fSubRegex.MatchString(iSub) {
					// If it does, set matched to true and break out of the loop
					matched = true
					break
				}
			}

			if !matched {
				fmt.Println(iSub)
			}
		}
	}
}
