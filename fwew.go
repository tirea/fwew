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
	"github.com/tirea/fwew/util"
	"os"
	"strings"
)

// global
var DEBUG bool
var WORD_HAS []string
var MW_FIELD_ID int = 0
var MW_FIELD_NAV int = 1
var MW_FIELD_IPA int = 2
var MW_FIELD_INF int = 3
var MW_FIELD_POS int = 4
var LW_FIELD_ID int = 0
var LW_FIELD_LC int = 1
var LW_FIELD_DEF int = 2
var LWFIELD_POS int = 3
var LW_NUM_FIELDS int = 4
var MW_NUM_FIELDS int = 5

// some minimal exception handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// simple containment check function
func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

// get the Database ID of a Na'vi root word
// only seemsm to support verb infix stripping at the moment
func getNavID(w string) string {

	// declare / initialize stuffs
	w = strings.ToLower(w)
	word := "\t" + w + "\t"
	var navID string
	var line string
	var fields []string
	//var nav string
	var pre string
	var inf string
	//var suf string
	var pos string
	var result [][]string

	metaWordsData, err := os.Open(util.Text("METAWORDS"))
	check(err)
	scanner := bufio.NewScanner(metaWordsData)

	// look for the word
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.ToLower(line)
		fields = strings.Split(line, "\t")
		//nav = fields[MW_FIELD_NAV]
		pre = fields[MW_FIELD_NAV]
		inf = fields[MW_FIELD_INF]
		//suf = fields[MW_FIELD_NAV]
		pos = fields[MW_FIELD_POS]

		// if it's a verb, prepare the infix regex
		if strings.HasPrefix(pos, "v") {
			result = util.Prefix(w, pre, pos)
			result = util.Infix(w, inf)
			//result = util.Suffix(w, suf, pos)
		} else {
			result = util.Prefix(w, pre, pos)
			//result = util.Suffix(w, suf, pos)
		}

		if DEBUG {
			fmt.Println("<DEBUG:getNavID() result>",result,"</DEBUG>")
		}

		// if user searched a root word and it's found, then just pull the ID
		if strings.Contains(line, word) {
			navID = line[0:strings.Index(line, "\t")]
			break
		// if affixes were found...
		} else if len(result) != 0 {
			// ...and if the found infixed VERB ends+starts with same letter as input (THIS LOGIC ONLY WORKS FOR INFIXES!)
			if strings.HasPrefix(pos, "v") && strings.HasSuffix(result[0][0], w[len(w)-1:]) && strings.HasPrefix(result[0][0], w[0:1]) {
				// ... then print out what was found and grab the ID
				navID = line[0:strings.Index(line, "\t")]
				fmt.Println(result)
				break
			} else if len(result[0]) > 1 && !stringInSlice("",result[0]) {
				navID = line[0:strings.Index(line, "\t")]
				fmt.Println(result)
			}
		}
	}
	return navID
}

// get the Database ID of a Local word by Language
// typically returns many matches
func getLocID(w string, l string) []string {

	// initialize some stuffs
	word := strings.ToLower(w)
	var locIDs []string
	var locID string
	var line string
	var fields []string
	var field_def string
	var field_arr []string
	var field_lng string

	localizedData, err := os.Open(util.Text("LOCALIZED"))
	check(err)
	scanner := bufio.NewScanner(localizedData)

	// look for matching words
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.ToLower(line)
		fields = strings.Split(line, "\t")

		// there should be 4 fields..
		if len(fields) == LW_NUM_FIELDS {
			field_def = fields[LW_FIELD_DEF]
			field_arr = strings.Split(field_def, " ")
			field_lng = fields[LW_FIELD_LC]

			// only try to grab the id from line using requested language
			if field_lng == l {
				if DEBUG {
					fmt.Println("<DEBUG:getLocID() word>" + word + "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() l>" + l + "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() line>" + line + "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() fields>", fields, "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() field_def>" + field_def + "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() field_arr>", field_arr, "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() field_lng>" + field_lng + "</DEBUG>")
				}

				// single-word definition and happens to be what user searched
				if len(field_arr) == 1 && field_def == word {
					locID = line[0:strings.Index(line, "\t")]
					locIDs = append(locIDs, locID)
					if DEBUG {
						fmt.Println("<DEBUG:getLocID() >!MATCH!</DEBUG>")
						fmt.Println("<DEBUG:getLocID() locID>" + locID + "</DEBUG>")
						fmt.Println("<DEBUG:getLocID() locIDs>", locIDs, "</DEBUG>")
					}

					// multiple words in the local definition, search through each word
				} else if len(field_arr) > 1 {
					for i := 0; i < len(field_arr); i++ {
						if DEBUG {
							fmt.Println("<DEBUG:getLocID() field_arr[i]>" + field_arr[i] + "</DEBUG>")
						}
						if field_arr[i] == word || field_arr[i] == word+"," {
							if DEBUG {
								fmt.Println("<DEBUG:getLocID() contains>l and *word*</DEBUG>")
							}
							locID = line[0:strings.Index(line, "\t")]
							locIDs = append(locIDs, locID)
						}
					}
				}
			}
		}
	}
	if DEBUG {
		fmt.Println("<DEBUG:getLocID() RETURNING locIDs>", locIDs, "</DEBUG>")
	}
	return locIDs
}

// get POS, Na'vi Word, IPA, Infixes, for given ID
func getDataByID(id string) (string, string, string, string) {

	if id == "" { return "", "", "", ""}

	// set up filestuffs
	metaData, err := os.Open(util.Text("METAWORDS"))
	check(err)
	scanner := bufio.NewScanner(metaData)

	var word string
	var ipa string
	var inf string
	var pos string

	if DEBUG {
		fmt.Println("<DEBUG:getDataByID() id>" + id + "</DEBUG>")
	}

	// break up each line by field and capture all the things...
	for scanner.Scan() {
		line := scanner.Text()
		// ... but only if the line matches the requested ID
		if strings.HasPrefix(line, id) {
			fields := strings.Split(line, "\t")
			word = fields[MW_FIELD_NAV]
			ipa = "[" + fields[MW_FIELD_IPA] + "]"
			inf = fields[MW_FIELD_INF]
			pos = fields[MW_FIELD_POS]
			if DEBUG {
				fmt.Println("<DEBUG:getDataByID() line>" + line + "</DEBUG>")
				fmt.Println("<DEBUG:getDataByID() fields>", fields, "</DEBUG>")
				fmt.Println("<DEBUG:getDataByID() word>" + word + "</DEBUG>")
			}
			break
		}
	}
	if DEBUG {
		fmt.Println("<DEBUG:getDataByID() word>" + word + "</DEBUG>")
	}
	return pos, word, ipa, inf
}

// get Local word for given ID
func getLocalWordByID(id string, l string) string {

	if id == "" || l == "" { return util.Text("NONE") }

	//filestuffs
	localData, err := os.Open(util.Text("LOCALIZED"))
	check(err)
	scanner := bufio.NewScanner(localData)

	var localWord string

	// search through each line to match the requested ID and language
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")
		if len(fields) == LW_NUM_FIELDS {
			field_wid := fields[LW_FIELD_ID]
			field_lng := fields[LW_FIELD_LC]
			field_def := fields[LW_FIELD_DEF]
			localWord = field_def
			if field_lng == l {
				if field_wid == id {
					if DEBUG {
						fmt.Println("<DEBUG:getLocalWordByID() line>" + line + "</DEBUG>")
						fmt.Println("<DEBUG:getLocalWordByID() fields>", fields, "</DEBUG>")
						fmt.Println("<DEBUG:getLocalWordByID() localWord>" + localWord + "</DEBUG>")
					}
					// match
					return localWord
				}
			}
		}
	}
	// no matches
	return util.Text("NONE")
}

func main() {
	// cli flags
	PROG_DEBUG := flag.Bool("DEBUG", false, util.Text("USAGEDEBUG"))
	flag_v := flag.Bool("v", false, util.Text("USAGEFLAG_V"))
	flag_l := flag.String("l", util.Text("DEFAULT_LANGUAGE"), util.Text("USAGEFLAG_L"))
	flag_i := flag.Bool("i", false, util.Text("USAGEFLAG_I"))
	flag_ipa := flag.Bool("ipa", false, util.Text("USAGEFLAG_IPA"))
	flag_r := flag.Bool("r", false, util.Text("USAGEFLAG_R"))

	flag.Parse()

	DEBUG = *PROG_DEBUG

	// declare / initialize vars
	var word string
	var lwrd string
	var lwds []string
	var dbid string
	var dbls []string
	var wipa string
	var infx string
	var wpos string
	lang := *flag_l

	// ARGS MODE

	// Na'vi -> Local lookup
	if !*flag_r && flag.NArg() > 0 {
		if DEBUG {
			fmt.Println("<DEBUG !*flag_r flag.NArg()!=0>Normal lookup direction | Args</DEBUG>")
		}
		if *flag_v {
			fmt.Println(util.Text("NAME") + " " + util.Text("VERSION") + "\n" + util.Text("DICTVERSION") + "\n")
		}

		for i := 0; i < flag.NArg(); i++ {
			dbid = getNavID(flag.Args()[i])
			wpos, word, wipa, infx = getDataByID(dbid)
			lwrd = getLocalWordByID(dbid, lang)

			// print out the results, what of it was requested
			if !*flag_ipa && !*flag_i {
				fmt.Print(wpos, " ", lwrd, "\n")
			} else {
				fmt.Print(wpos, " ", lwrd, " ")
			}
			if *flag_ipa {
				if *flag_i {
					fmt.Print(wipa, " ")
				} else {
					fmt.Print(wipa, "\n")
				}
			}
			if *flag_i {
				fmt.Print(infx, "\n")
			}
			fmt.Println("")
		}

		// Local -> Na'vi lookup
	} else if *flag_r && flag.NArg() > 0 {
		if DEBUG {
			fmt.Println("<DEBUG *flag_r flag.NArg()>0>Reverse lookup direction | Args</DEBUG>")
		}
		if *flag_v {
			fmt.Println(util.Text("NAME") + " " + util.Text("VERSION") + "\n" + util.Text("DICTVERSION") + "\n")
		}

		for i := 0; i < flag.NArg(); i++ {
			lwrd = flag.Args()[i]
			dbls = getLocID(lwrd, lang)
			if DEBUG {
				fmt.Println("<DEBUG:main() dbls>", dbls, "</DEBUG>")
			}
			for i := 0; i < len(dbls); i++ {
				dbid = dbls[i]
				wpos, word, wipa, infx = getDataByID(dbid)
				lwds = append(lwds, lwrd)

				//print results, what of it was requested
				fmt.Print(wpos, " ", word, " ")
				if *flag_ipa {
					if *flag_i {
						fmt.Print(wipa, " ")
					} else {
						fmt.Print(wipa, "\n")
					}
				}
				if *flag_i {
					fmt.Print(infx, " ")
				}
				fmt.Println("(" + getLocalWordByID(dbid, lang) + ")")
			}
		}
		if len(dbls) == 0 {
			fmt.Println(util.Text("NONE"))
		}
		fmt.Println("")
	}

	// INTERACTIVE MODE

	// Na'vi -> local lookup
	if !*flag_r && flag.NArg() == 0 {
		if *flag_v {
			fmt.Println(util.Text("NAME") + " " + util.Text("VERSION") + "\n" + util.Text("DICTVERSION") + "\n")
			os.Exit(0)
		}

		// print the program Header text
		fmt.Println(util.Text("HEADTEXT"))

		if DEBUG {
			fmt.Println("<DEBUG=true></DEBUG>")
			fmt.Println("<DEBUG !*flag_r flag.NArg()==0>Normal lookup direction | Interactive</DEBUG>")
		}

		fmt.Println("")

		// read word from cli args
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Fwew:> ")
		word, _ := reader.ReadString('\n')
		word = strings.Trim(word, "\n")

		if word != "" {
			dbid = getNavID(word)
			wpos, _, wipa, infx = getDataByID(dbid)
			lwrd = getLocalWordByID(dbid, lang)
		} else {
			wpos = ""
			wipa = "[]"
			infx = "\\N"
			dbid = "-1"
			lwrd = ""
		}

		// print out the results, what of it was requested
		if !*flag_ipa && !*flag_i {
			fmt.Print(wpos, " ", lwrd, "\n")
		} else {
			fmt.Print(wpos, " ", lwrd, " ")
		}
		if *flag_ipa {
			if *flag_i {
				fmt.Print(wipa, " ")
			} else {
				fmt.Print(wipa, "\n")
			}
		}
		if *flag_i {
			fmt.Print(infx, "\n")
		}
		fmt.Println("")

		// Local -> Na'vi lookup
	} else if *flag_r && flag.NArg() == 0 {
		if *flag_v {
			fmt.Println(util.Text("NAME") + " " + util.Text("VERSION") + "\n" + util.Text("DICTVERSION") + "\n")
			os.Exit(0)
		}

		// print the program Header text
		fmt.Println(util.Text("HEADTEXT"))

		if DEBUG {
			fmt.Println("<DEBUG=true></DEBUG>")
			fmt.Println("<DEBUG *flag_r flag.NArg()==0>Reverse lookup direction | Interactive</DEBUG>")
		}

		fmt.Println("")

		//read word from cli
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Fwew:> ")
		word, _ := reader.ReadString('\n')
		word = strings.Trim(word, "\n")

		if word == "" {
			fmt.Println("\n")
			os.Exit(0)
		}

		lwrd = word
		dbls = getLocID(lwrd, lang)

		for i := 0; i < len(dbls); i++ {
			dbid = dbls[i]
			wpos, word, wipa, infx = getDataByID(dbid)
			lwds = append(lwds, lwrd)

			//print results, what of it was requested
			fmt.Print(wpos, " ", word, " ")
			if *flag_ipa {
				if *flag_i {
					fmt.Print(wipa, " ")
				} else {
					fmt.Print(wipa, "\n")
				}
			}
			if *flag_i {
				fmt.Print(infx, " ")
			}
			fmt.Println("(" + getLocalWordByID(dbid, lang) + ")")
		}
		if len(dbls) == 0 {
			fmt.Println(util.Text("NONE"))
		}
		fmt.Println("")
	}

}
