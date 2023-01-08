package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"encoding/json"
)

type ExObject struct {
	Sub []string `json:"sub"`
	Dns  []string `json:"dns"`
}

func main() {

	var excludeObj string
	flag.StringVar(&excludeObj, "x", "", "Exclude JSON object")

	var filterMode int
	flag.IntVar(&filterMode, "m", 0, "Filter mode from 1-4")

	flag.Parse()

	// Unmarshal the JSON data into a JSONObject
	var inputJSON ExObject
	if err := json.Unmarshal([]byte(excludeObj), &inputJSON); err != nil {
		fmt.Println("Error parsing the JSON object:", err)
		os.Exit(1)
	}

	// read from stdin, and print if not excluded
	line := bufio.NewScanner(os.Stdin)
	
	for line.Scan() {
		switch filterMode {
			case 1: 	// Subdmoains Mode
				iSub := line.Text()
				matched := false
				// Iterate over the list of regexes in inputJSON.Sub
				for _, fSub := range inputJSON.Sub {
					fSubRegex, err := regexp.Compile("^" + fSub + "$")
					if err != nil {
						fmt.Println("Error compiling regex:", err)
						os.Exit(1)
					}

					// Check if the Sub matches the fSub regex
					if fSubRegex.MatchString(iSub) {
						matched = true
						break
					}
				}

				if !matched {
					fmt.Println(iSub)
				}
			
			case 2: 	// DNS Mode
				iLine := strings.Split(line.Text(), " ")
				iSub := iLine[0]
				iDNSList := strings.Split(iLine[1], ",")

				matched := false

				for _, iDNS := range iDNSList {
					matched = false
					for _, fDNS := range inputJSON.Dns {
						// Check if the DNS is excluded
						if iDNS == fDNS {
							matched = true
							break
						}
					}
				}

				if !matched {
					fmt.Println(iSub)
				}

			case 0:
				fallthrough
			default:
				fmt.Fprintln(os.Stderr, "Modes Available:")
				fmt.Fprintln(os.Stderr, "Mode 1: filter by subdmoains, input [sub]")
				fmt.Fprintln(os.Stderr, "Mode 2: filter by DNS record, input [sub ip,cname,txt]")
				os.Exit(1)
		}

	}
}
