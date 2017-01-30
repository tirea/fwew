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
	"flag"
	"fmt"
	"fwew/util"
	"os"
	"strings"
	"errors"
)

// Global vars
var DEBUG bool         // whether or not to produce debugging output
var FIELD_ID int = 0   // dictionary.tsv line field 0 is Database ID
var FIELD_LC int = 1   // dictionary.tsv line field 1 is Language Code
var FIELD_NAV int = 2  // dictionary.tsv line field 2 is Na'vi word
var FIELD_IPA int = 3  // dictionary.tsv line field 3 is IPA data
var FIELD_INF int = 4  // dictionary.tsv line field 4 is Infix location data
var FIELD_POS int = 5  // dictionary.tsv line field 5 is Part of Speech data
var FIELD_DEF int = 6  // dictionary.tsv line field 6 is Local definition
var NUM_FIELDS int = 7 // dictionary.tsv number of line fields is 7
//TODO: infixes.tsv fields?
var LANGUAGE string = util.Text("DEFAULT_LANGUAGE")

/* Fwew function
 * Params: word string, the user's word input
 *         lc string, the language code
 *         reverse bool, whether or not the lookup will be local word instead of Na'vi
 * Returns: []string, {ID, LC, NAV, IPA, INF, POS, DEF}
 * Search data file for the user's Na'vi word and return a slice with pertinent data
 */
func fwew(word string, lc string, reverse bool) [][]string {
	var results [][]string
	var fields []string

	// Searching for Local word, just search for it...
	word = strings.ToLower(word)

	// Prepare file for searching
	dictData, err := os.Open(util.Text("DICTIONARY"))
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
			if strings.Contains(fields[FIELD_DEF], word) && strings.Contains(fields[FIELD_LC], lc){
				results = append(results, fields)
				break
			}
		} else {
			if strings.Contains(line, "\t"+word+"\t") && strings.Contains(fields[FIELD_LC], lc) {
				results = append(results, fields)
			}
		}
	}

	return results
}

/* Main function
 * Params: none
 * Returns: void
 * Program flow: 
 *   Get user input (flags and word(s))
 *	 Search data file for word(s) to get data
 *   Format and print requested data
 *	 OR: Print version text or help text
 */
func main() {
	// CLI FLAGS
	// Debug flag, for verbose probing output
	PROG_DEBUG := flag.Bool("DEBUG", false, util.Text("USAGEDEBUG"))
	// Version flag, for displaying version data
	flag_v := flag.Bool("v", false, util.Text("USAGEFLAG_V"))
	// Language specifier flag
	flag_l := flag.String("l", util.Text("DEFAULT_LANGUAGE"), util.Text("USAGEFLAG_L"))
	// Infixes flag, opt to show infix location data
	flag_i := flag.Bool("i", false, util.Text("USAGEFLAG_I"))
	// IPA flag, opt to show IPA data
	flag_ipa := flag.Bool("ipa", false, util.Text("USAGEFLAG_IPA"))
	// Part of Show part of speech flag
//	flag_pos := flag.String("pos", "", util.Text("USAGEFLAG_POS")) //TODO
	// Reverse direction flag, for local_lang -> Na'vi lookups
	flag_r := flag.Bool("r", false, util.Text("USAGEFLAG_R"))
	flag.Parse()
	//set the global debugging bool to the cli flag value
	DEBUG = *PROG_DEBUG

	var results [][]string
	var input, nav, ipa, inf, pos, def string

	if *flag_v {
		fmt.Println(util.Text("NAME") + " " + util.Text("VERSION") + "\n" + util.Text("DICTVERSION") + "\n")
		
		if flag.NArg() == 0 {
			os.Exit(0)
		}
	}

	// ARGS MODE
	if flag.NArg() > 0 {

		for _, arg := range flag.Args() {
			results = fwew(arg, *flag_l, *flag_r)

			if len(results) != 0 {

				for _, r := range results {
					nav = r[FIELD_NAV]
					ipa = "[" + r[FIELD_IPA] + "]"
					inf = r[FIELD_INF]
					pos = r[FIELD_POS]
					def = r[FIELD_DEF]
				}

				fmt.Print(pos + " ")
				if *flag_r {
					fmt.Print(nav + " ")
				} else {
					fmt.Print(def + " ")
				}
				if *flag_ipa {
					fmt.Print(ipa + " ")
				}
				if *flag_i {
					fmt.Print(inf + " ")
				}
				if *flag_r {
					fmt.Println("(" + def + ")\n")
				} else {
					fmt.Println("(" + nav + ")\n")
				}

			} else {
				fmt.Println(util.Text("NORESULTS"))
			}
		}

	// INTERACTIVE MODE
	} else {
		fmt.Println(util.Text("HEADTEXT"))
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Fwew:> ")
		input, _ = reader.ReadString('\n')
		input = strings.Trim(input, "\n")

		results = fwew(input, *flag_l, *flag_r)

		if len(results) != 0 {

			for _, r := range results {
				nav = r[FIELD_NAV]
				ipa = "[" + r[FIELD_IPA] + "]"
				inf = r[FIELD_INF]
				pos = r[FIELD_POS]
				def = r[FIELD_DEF]
			}

			fmt.Print(pos + " ")
			if *flag_r {
				fmt.Print(nav + " ")
			} else {
				fmt.Print(def + " ")
			}
			if *flag_ipa {
				fmt.Print(ipa + " ")
			}
			if *flag_i {
				fmt.Print(inf + " ")
			}
			if *flag_r {
				fmt.Println("(" + def + ")\n")
			} else {
				fmt.Println("(" + nav + ")\n")
			}

		} else {
			fmt.Println(util.Text("NORESULTS"))
		}
	}

}