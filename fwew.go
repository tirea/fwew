//	Fwew is free software: you can redistribute it and/or modify
// 	it under the terms of the GNU General Public License as published by
// 	the Free Software Foundation, either version 3 of the License, or
// 	(at your option) any later version.
//
//	Fwew is distributed in the hope that it will be useful,
//	but WITHOUT ANY WARRANTY; without even implied warranty of
//	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//	GNU General Public License for more details.
//
//	You should have received a copy of the GNU General Public License
//	along with Fwew.  If not, see http://gnu.org/licenses/

//	Package main obviously contains all the stuff for the main program
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"fwew/util"
	"io/ioutil"
	"os"
	"strings"
)

// Global
var debug *bool

type Config struct {
	Language  string `json:"language"`
	PosFilter string `json:"posFilter"`
}

func stripChars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func fwew(word string, lc string, posFilter string, reverse bool) [][]string {
	const (
		lcField  int = 1 // dictionary.tsv line field 1 is Language Code
		posField int = 5 // dictionary.tsv line field 5 is Part of Speech data
		defField int = 6 // dictionary.tsv line field 6 is Local definition
	)
	var results [][]string
	var fields []string
	var defString string

	// Searching for Local word, just search for it...
	word = strings.ToLower(word)

	// Prepare file for searching
	dictData, err := os.Open(util.Text("dictionary"))
	defer dictData.Close()
	if err != nil {
		fmt.Println(errors.New(util.Text("noDataError")))
		os.Exit(1)
	}
	scanner := bufio.NewScanner(dictData)

	// Go through each line and see what we can find
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		// Store the fields of the line into fields array in lowercase
		fields = strings.Split(line, "\t")

		if reverse {
			if posFilter == util.Text("defaultFilter") {
				if strings.Contains(fields[lcField], lc) {
					defString = stripChars(fields[defField], ",;")
					for _, w := range strings.Split(defString, " ") {
						if w == word {
							results = append(results, fields)
						}
					}
				}
			} else {
				if strings.Contains(fields[lcField], lc) && strings.Contains(fields[posField], posFilter) {
					defString = stripChars(fields[defField], ",;")
					for _, w := range strings.Split(defString, " ") {
						if w == word {
							results = append(results, fields)
						}
					}
				}
			}
		} else {
			if strings.Contains(fields[lcField], lc) && strings.Contains(line, "\t"+word+"\t") {
				results = append(results, fields)
				break
			}
		}
	}

	return results
}

func printResults(results [][]string, reverse bool, showInfixes bool, showIPA bool) {
	const (
		navField int = 2 // dictionary.tsv line field 2 is Na'vi word
		ipaField int = 3 // dictionary.tsv line field 3 is IPA data
		infField int = 4 // dictionary.tsv line field 4 is Infix location data
		posField int = 5 // dictionary.tsv line field 5 is Part of Speech data
		defField int = 6 // dictionary.tsv line field 6 is Local definition
	)
	var nav, ipa, inf, pos, def string

	if len(results) != 0 {

		for i, r := range results {
			nav = r[navField]
			ipa = "[" + r[ipaField] + "]"
			inf = r[infField]
			pos = r[posField]
			def = r[defField]

			fmt.Print("[", i+1, "] ")

			fmt.Print(pos + " ")
			if reverse {
				fmt.Print(nav + " ")
			} else {
				fmt.Print(def + " ")
			}
			if showIPA {
				fmt.Print(ipa + " ")
			}
			if showInfixes {
				fmt.Print(inf + " ")
			}
			if reverse {
				fmt.Println("(" + def + ")\n")
			} else {
				fmt.Println("(" + nav + ")\n")
			}
		}

	} else {
		fmt.Println(util.Text("none"))
	}
}

func setFlags(arg string, debug, r, i, ipa *bool, l, p *string) {
	const start int = 4
	flagList := strings.Split(arg[start:len(arg)-1], ",")
	for _, f := range flagList {
		switch {
		case f == "debug":
			*debug = true
		case f == "r":
			*r = true
		case f == "i":
			*i = true
		case f == "ipa":
			*ipa = true
		case strings.HasPrefix(f, "l="):
			*l = f[2:]
		case strings.HasPrefix(f, "p="):
			*p = f[2:]
		}
	}
	fmt.Println("<!", flagList, "set >\n")
}

func unsetFlags(arg string, debug, r, i, ipa *bool) {
	const start int = 6
	flagList := strings.Split(arg[6:len(arg)-1], ",")
	for _, f := range flagList {
		switch f {
		case "debug":
			*debug = false
		case "r":
			*r = false
		case "i":
			*i = false
		case "ipa":
			*ipa = false
		}
	}
	fmt.Println("<!", flagList, "unset >\n")
}

func LoadConfig() {
	confile, e := ioutil.ReadFile(util.Text("config"))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var config Config
	json.Unmarshal(confile, &config)
	util.SetText("language", config.Language)
	util.SetText("defaultFilter", config.PosFilter)
}

func main() {
	var results [][]string
	var language, posFilter *string
	var showVersion, showInfixes, showIPA, reverse *bool

	LoadConfig()

	// Debug flag, for verbose probing output
	debug = flag.Bool("debug", false, util.Text("usageDebug"))
	// Version flag, for displaying version data
	showVersion = flag.Bool("v", false, util.Text("usageV"))
	// Reverse direction flag, for local_lang -> Na'vi lookups
	reverse = flag.Bool("r", false, util.Text("usageR"))
	// Language specifier flag
	language = flag.String("l", util.Text("language"), util.Text("usageL"))
	// Infixes flag, opt to show infix location data
	showInfixes = flag.Bool("i", false, util.Text("usageI"))
	// IPA flag, opt to show IPA data
	showIPA = flag.Bool("ipa", false, util.Text("usageIPA"))
	// Show part of speech flag
	posFilter = flag.String("p", util.Text("defaultFilter"), util.Text("usageP"))
	flag.Parse()

	if *showVersion {
		fmt.Println(util.Text("name") + " " + util.Text("version") + "\n" + util.Text("dictVersion") + "\n")
		if flag.NArg() == 0 {
			os.Exit(0)
		}
	}

	// ARGS MODE
	if flag.NArg() > 0 {
		for _, arg := range flag.Args() {
			if strings.HasPrefix(arg, "set[") {
				setFlags(arg, debug, reverse, showInfixes, showIPA, language, posFilter)
			} else if strings.HasPrefix(arg, "unset[") {
				unsetFlags(arg, debug, reverse, showInfixes, showIPA)
			} else {
				results = fwew(arg, *language, *posFilter, *reverse)
				printResults(results, *reverse, *showInfixes, *showIPA)
			}
		}

		// INTERACTIVE MODE
	} else {
		fmt.Println(util.Text("header"))

		for {
			fmt.Print("Fwew:> ")

			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			input = strings.Trim(input, "\n")

			// catch EOF error
			if err != nil {
				fmt.Println()
				os.Exit(0)
			}

			if input != "" {
				if strings.HasPrefix(input, "set[") {
					setFlags(input, debug, reverse, showInfixes, showIPA, language, posFilter)
				} else if strings.HasPrefix(input, "unset[") {
					unsetFlags(input, debug, reverse, showInfixes, showIPA)
				} else {
					results = fwew(input, *language, *posFilter, *reverse)
					printResults(results, *reverse, *showInfixes, *showIPA)
				}
			} else {
				fmt.Println()
			}
		}
	}
}
