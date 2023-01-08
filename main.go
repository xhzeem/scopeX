package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

	// read URLs on stdin, and print if not excluded
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {

		switch filterMode {
			case 1: 	// Subdmoains Mode
				iSub := sc.Text()
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
				fmt.Println("IP Mode")
			
			case 3: 	// CNAME Mode
				fmt.Println("CNAME Mode")

			case 0:
				fallthrough
			default:
				fmt.Fprintln(os.Stderr, "--mode 1: filter by subdmoains, input [sub]")
				fmt.Fprintln(os.Stderr, "--mode 2: filter by IP address, input [sub, ip]")
				fmt.Fprintln(os.Stderr, "--mode 3: filter by cname, input: [sub, cname]")
				os.Exit(1)
		}

	}
}
