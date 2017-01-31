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
	"errors"
	"flag"
	"fmt"
	"fwew/util"
	"os"
	"strings"
)

// Global
var debug *bool

func fwew(word string, lc string, reverse bool) [][]string {
	var lcField int = 1  // dictionary.tsv line field 1 is Language Code
	var defField int = 6 // dictionary.tsv line field 6 is Local definition
	var results [][]string
	var fields []string

	// Searching for Local word, just search for it...
	word = strings.ToLower(word)

	// Prepare file for searching
	dictData, err := os.Open(util.Text("DICTIONARY"))
	defer dictData.Close()
	if err != nil {
		fmt.Println(errors.New(util.Text("ERR_MISSING_DATAFILE")))
		os.Exit(1)
	}
	scanner := bufio.NewScanner(dictData)

	// Go through each line and see what we can find
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		// Store the fields of the line into fields array in lowercase
		fields = strings.Split(line, "\t")

		if reverse {
			if strings.Contains(fields[defField], word) && strings.Contains(fields[lcField], lc) {
				results = append(results, fields)
				if *debug { fmt.Println(results) }
			}
		} else {
			if strings.Contains(line, "\t"+word+"\t") && strings.Contains(fields[lcField], lc) {
				results = append(results, fields)
				break
			}
		}
	}

	return results
}

func printResults(results [][]string, reverse bool, showInfixes bool, showIPA bool) {
	var navField int = 2 // dictionary.tsv line field 2 is Na'vi word
	var ipaField int = 3 // dictionary.tsv line field 3 is IPA data
	var infField int = 4 // dictionary.tsv line field 4 is Infix location data
	var posField int = 5 // dictionary.tsv line field 5 is Part of Speech data
	var defField int = 6 // dictionary.tsv line field 6 is Local definition
	//TODO: infixes.tsv fields
	var nav, ipa, inf, pos, def string

	if len(results) != 0 {

		for _, r := range results {
			nav = r[navField]
			ipa = "[" + r[ipaField] + "]"
			inf = r[infField]
			pos = r[posField]
			def = r[defField]

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
		fmt.Println(util.Text("NORESULTS"))
	}
}

func main() {
	var language *string
	var showVersion, showInfixes, showIPA, reverse *bool
	// Debug flag, for verbose probing output
	debug = flag.Bool("DEBUG", false, util.Text("USAGEDEBUG"))
	// Version flag, for displaying version data
	showVersion = flag.Bool("v", false, util.Text("USAGEFLAG_V"))
	// Language specifier flag
	language = flag.String("l", util.Text("DEFAULT_LANGUAGE"), util.Text("USAGEFLAG_L"))
	// Infixes flag, opt to show infix location data
	showInfixes = flag.Bool("i", false, util.Text("USAGEFLAG_I"))
	// IPA flag, opt to show IPA data
	showIPA = flag.Bool("ipa", false, util.Text("USAGEFLAG_IPA"))
	// Show part of speech flag
	//flag_pos := flag.String("pos", "", util.Text("USAGEFLAG_POS")) //TODO
	// Reverse direction flag, for local_lang -> Na'vi lookups
	reverse = flag.Bool("r", false, util.Text("USAGEFLAG_R"))
	flag.Parse()

	var results [][]string
	var input string

	if *showVersion {
		fmt.Println(util.Text("NAME") + " " + util.Text("VERSION") + "\n" + util.Text("DICTVERSION") + "\n")
		if flag.NArg() == 0 {
			os.Exit(0)
		}
	}

	// ARGS MODE
	if flag.NArg() > 0 {
		for _, arg := range flag.Args() {
			results = fwew(arg, *language, *reverse)
			printResults(results, *reverse, *showInfixes, *showIPA)
		}

		// INTERACTIVE MODE
	} else {
		fmt.Println(util.Text("HEADTEXT"))
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Fwew:> ")
		input, _ = reader.ReadString('\n')
		input = strings.Trim(input, "\n")

		if input != "" {
			results = fwew(input, *language, *reverse)
			printResults(results, *reverse, *showInfixes, *showIPA)
		} else {
			fmt.Println()
		}
	}

}
