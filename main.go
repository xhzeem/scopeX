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
	Ip  []string `json:"ip"`
	cnm []string `json:"cnm"`
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
			
			case 2: 	// IP Mode
				iLine := strings.Split(line.Text(), " ")
				iSub := iLine[0]
				iIPList := strings.Split(iLine[1], ",")

				matched := false

				for _, iIP := range iIPList {
					matched = false
					for _, fIP := range inputJSON.Ip {
						// Check if the IP is excluded
						if iIP == fIP {
							matched = true
							break
						}
					}
				}

				if !matched {
					fmt.Println(iSub)
				}
			
			case 3: 	// CNAME Mode
				fmt.Println("CNAME Mode")

			case 0:
				fallthrough
			default:
				fmt.Fprintln(os.Stderr, "Modes Available:")
				fmt.Fprintln(os.Stderr, "Mode 1: filter by subdmoains, input [sub]")
				fmt.Fprintln(os.Stderr, "Mode 2: filter by IP address, input [sub ip,ip]")
				fmt.Fprintln(os.Stderr, "Mode 3: filter by cname, input: [sub cname,cname]")
				os.Exit(1)
		}

	}
}
