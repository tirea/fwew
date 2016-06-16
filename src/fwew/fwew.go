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
	"os"
	"strings"
	"regexp"
	"util"
)

// global
var PROG_DEBUG bool = false
var WORD_HAS []string

// some minimal exception handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// get the Database ID of a Na'vi root word
// only seemsm to support verb infix stripping at the moment
func getNavID(w string) string {

	// initialize stuffs
	w = strings.ToLower(w)
	word := "\t" + w + "\t"
	navID := "-1"
	line := ""
	fields := []string{}
	inf := "-1"
	pos := "-1"
	metaWordsData, err := os.Open(txt.Text("METAWORDS"))
	check(err)
	scanner := bufio.NewScanner(metaWordsData)

	// look for the word
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.ToLower(line)
		fields = strings.Split(line, "\t")
		inf = fields[3]
		pos = fields[4]

		// if it's a verb, prepare the infix regex
		if strings.HasPrefix(pos, "v") {
			inf = strings.Replace(inf,"<1>",txt.Text("INFIX_0"),1)
			inf = strings.Replace(inf,"<2>",txt.Text("INFIX_1"),1)
			inf = strings.Replace(inf,"<3>",txt.Text("INFIX_2"),1)
		}
		re, err := regexp.Compile(inf)
		check(err)

		// pull out all infixes used and stash them in the result array
		result := re.FindAllStringSubmatch(w, -1)

		// if user searched a root word and it's found, then just pull the ID
		if strings.Contains(line, word) {
			navID = line[0:strings.Index(line, "\t")]
			break
		// if infixes were found and it's actually a verb...
		} else if len(result) != 0 && strings.HasPrefix(pos, "v") {
			// ...and if the found infixed word ends with same letter as input
			if strings.HasSuffix(result[0][0], w[len(w)-1:]) {
				// ... and if found infixed word starts with same letter
				// ... then print out what was found and grab the ID
				if strings.HasPrefix(result[0][0], w[0:1]){
					navID = line[0:strings.Index(line, "\t")]
					fmt.Println(result)
				}
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
	locIDs := []string{}
	locID := "-1"
	line := ""
	fields := []string{}
	field_def := ""
	field_arr := []string{}
	field_lng := ""
	localizedData, err := os.Open(txt.Text("LOCALIZED"))
	check(err)
	scanner := bufio.NewScanner(localizedData)

	// look for matching words
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.ToLower(line)
		fields = strings.Split(line, "\t")

		// there should be 4 fields..
		if len(fields) == 4 {
			field_def = fields[2]
			field_arr = strings.Split(field_def, " ")
			field_lng = fields[1]

			// only try to grab the id from line using requested language
			if field_lng == l {
				if PROG_DEBUG {
					fmt.Println("<DEBUG:getLocID() word>" + word + "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() l>" + l + "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() line>" + line + "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() fields>", fields, "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() field_def>" + field_def + "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() field_arr>", field_arr, "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() fild_lng>" + field_lng + "</DEBUG>")
				}

				// single-word definition and happens to be what user searched
				if len(field_arr) == 1 && field_def == word {
					locID = line[0:strings.Index(line, "\t")]
					locIDs = append(locIDs, locID)
					if PROG_DEBUG {
						fmt.Println("<DEBUG:getLocID() >!MATCH!</DEBUG>")
						fmt.Println("<DEBUG:getLocID() locID>" + locID + "</DEBUG>")
						fmt.Println("<DEBUG:getLocID() locIDs>", locIDs, "</DEBUG>")
					}

				// multiple words in the local definition, search through each word
				} else if len(field_arr) > 1 {
					for i := 0; i < len(field_arr); i++ {
						if PROG_DEBUG {
							fmt.Println("<DEBUG:getLocID() field_arr[i]>" + field_arr[i] + "</DEBUG>")
						}
						if field_arr[i] == word || field_arr[i] == word+"," {
							if PROG_DEBUG {
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
	if PROG_DEBUG {
		fmt.Println("<DEBUG:getLocID() RETURNING locIDs>", locIDs, "</DEBUG>")
	}
	return locIDs
}

// get POS, Na'vi Word, IPA, Infixes, for given ID
func getDataByID(id string) (string, string, string, string) {

	// set up filestuffs
	metaData, err := os.Open(txt.Text("METAWORDS"))
	check(err)
	word, ipa, inf, pos := "", "", "", ""
	if PROG_DEBUG {
		fmt.Println("<DEBUG:getDataByID() id>" + id + "</DEBUG>")
	}
	scanner := bufio.NewScanner(metaData)
	
	// break up each line by field and capture all the things...
	for scanner.Scan() {
		line := scanner.Text()
		// ... but only if the line matches the requested ID
		if strings.HasPrefix(line, id) {
			fields := strings.Split(line, "\t")
			word = fields[1]
			ipa = "[" + fields[2] + "]"
			inf = fields[3]
			pos = fields[4]
			if PROG_DEBUG {
				fmt.Println("<DEBUG:getDataByID() line>" + line + "</DEBUG>")
				fmt.Println("<DEBUG:getDataByID() fields>", fields, "</DEBUG>")
				fmt.Println("<DEBUG:getDataByID() word>" + word + "</DEBUG>")
			}
			break
		}
	}
	if PROG_DEBUG {
		fmt.Println("<DEBUG:getDataByID() word>" + word + "</DEBUG>")
	}
	return pos, word, ipa, inf
}

// get Local word for given ID
func getLocalWordByID(id string, l string) string {
	
	//filestuffs
	localData, err := os.Open(txt.Text("LOCALIZED"))
	check(err)
	localWord := ""
	scanner := bufio.NewScanner(localData)

	// search through each line to match the requested ID and language
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")
		if len(fields) == 4 {
			field_wid := fields[0]
			field_lng := fields[1]
			field_def := fields[2]
			localWord = field_def
			if field_lng == l {
				if field_wid == id {
					if PROG_DEBUG {
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
	return txt.Text("NONE")
}

func main() {
	// cli flags
	DEBUG := flag.Bool("DEBUG", false, txt.Text("USAGEDEBUG"))
	flag_l := flag.String("l", txt.Text("DEFAULT_LANGUAGE"), txt.Text("USAGEFLAG_L"))
	flag_i := flag.Bool("i", false, txt.Text("USAGEFLAG_I"))
	flag_ipa := flag.Bool("ipa", false, txt.Text("USAGEFLAG_IPA"))
	flag_r := flag.Bool("r", false, txt.Text("USAGEFLAG_R"))

	flag.Parse()

	PROG_DEBUG = *DEBUG

	// initialize vars
	lang := *flag_l
	word := ""
	lwrd := ""
	lwds := []string{}
	dbid := ""
	dbls := []string{}
	wipa := ""
	infx := ""
	wpos := ""

	// ARGS MODE

	// Na'vi -> Local lookup
	if !*flag_r && flag.NArg() > 0 {
		if PROG_DEBUG {
			fmt.Println("<DEBUG !*flag_r flag.NArg()!=0>Normal lookup direction | Args</DEBUG>")
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
		if PROG_DEBUG {
			fmt.Println("<DEBUG *flag_r flag.NArg()>0>Reverse lookup direction | Args</DEBUG>")
		}

		for i := 0; i < flag.NArg(); i++ {
			lwrd = flag.Args()[i]
			dbls = getLocID(lwrd, lang)
			if PROG_DEBUG {
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
			fmt.Println(txt.Text("NONE"))
		}
		fmt.Println("")
	}

	// INTERACTIVE MODE

	// Na'vi -> local lookup
	if !*flag_r && flag.NArg() == 0 {

		// print the program Header text
		fmt.Println(txt.Text("HEADTEXT"))
		if PROG_DEBUG {
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

		// print the program Header text
		fmt.Println(txt.Text("HEADTEXT"))
		if PROG_DEBUG {
			fmt.Println("<DEBUG=true></DEBUG>")
			fmt.Println("<DEBUG *flag_r flag.NArg()==0>Reverse lookup direction | Interactive</DEBUG>")
		}
		fmt.Println("")

		//read word from cli
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Fwew:> ")
		word, _ := reader.ReadString('\n')
		word = strings.Trim(word, "\n")

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
			fmt.Println(txt.Text("NONE"))
		}
		fmt.Println("")
	}

}
