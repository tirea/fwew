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
	"util"
	"os"
	"strings"
)

// global
var PROG_DEBUG bool = false

// some minimal exception handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// get the Database ID of a Na'vi root word
func getNavID(w string) string {
	word := "\t" + strings.ToLower(w) + "\t"
	navID := "-1"
	line := ""
	metaWordsData, err := os.Open(txt.Text("METAWORDS"))
	check(err)
	scanner := bufio.NewScanner(metaWordsData)
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Contains(line, word) {
			navID = line[0:strings.Index(line, "\t")]
		}
	}
	return navID
}

// get the Database ID of a Local word by Language
// typically returns many matches
func getLocID(w string, l string) []string {
	word := strings.ToLower(w)
	locIDs := []string{}
	locID := "-1"
	line := ""

	localizedData, err := os.Open(txt.Text("LOCALIZED"))
	check(err)

	scanner := bufio.NewScanner(localizedData)
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.ToLower(line)
		fields := strings.Split(line, "\t")
		if len(fields) ==4 {
			field_def := fields[2]
			field_arr := strings.Split(field_def, " ")
			field_lng := fields[1]
			if field_lng == l {
				if PROG_DEBUG {
					fmt.Println("<DEBUG:getLocID() word>" + word + "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() l>" + l + "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() line>" + line + "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() fields>", fields, "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() field_def>"+ field_def+ "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() field_arr>", field_arr, "</DEBUG>")
					fmt.Println("<DEBUG:getLocID() fild_lng>"+ field_lng+ "</DEBUG>")
				}
				if len(field_arr) == 1 && field_def == word {
					locID = line[0:strings.Index(line, "\t")]
					locIDs = append(locIDs, locID)
					if PROG_DEBUG {
						fmt.Println("<DEBUG:getLocID() >!MATCH!</DEBUG>")
						fmt.Println("<DEBUG:getLocID() locID>" + locID + "</DEBUG>")
						fmt.Println("<DEBUG:getLocID() locIDs>", locIDs, "</DEBUG>")
					}
					//break
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
	if PROG_DEBUG { fmt.Println("<DEBUG:getLocID() RETURNING locIDs>", locIDs, "</DEBUG>") }
	return locIDs
}

// get POS, Na'vi Word, IPA, Infixes, for given ID
func getDataByID(id string) (string, string, string, string) {
	metaData, err := os.Open(txt.Text("METAWORDS"))
	check(err)
	word, ipa, inf, pos := "", "", "", ""
	if PROG_DEBUG {
		fmt.Println("<DEBUG:getDataByID() id>" + id + "</DEBUG>")
	}
	scanner := bufio.NewScanner(metaData)
	for scanner.Scan() {
		line := scanner.Text()
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
	return pos, word, ipa, inf
}

// get Local word for given ID
func getLocalWordByID(id string, l string) string {
	localData, err := os.Open(txt.Text("LOCALIZED"))
	check(err)
	localWord := ""
	scanner := bufio.NewScanner(localData)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")
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
				return localWord
			}
		}
	}
	return ""
}

func main() {
	// cli flags
	DEBUG := flag.Bool("DEBUG", false, "allow fmt probing")
	flag_l := flag.String("l", txt.Text("DEFAULT_LANGUAGE"), "use specified language. \n"+
		"\tValid values: "+txt.Text("LANGUAGES"))
	flag_i := flag.Bool("i", false, "Display infix location data")
	flag_ipa := flag.Bool("ipa", false, "Display IPA data")
	flag_r := flag.Bool("r", false, "Reverse the lookup direction from Na'vi->Local to Local->Na'vi")

	flag.Parse()

	PROG_DEBUG = *DEBUG

	// vars
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

	// Na'vi -> Local lookup or Local -> Na'vi lookup
	if !*flag_r && flag.NArg() > 0 {
		if *DEBUG {
			fmt.Println("<DEBUG !*flag_r flag.NArg()!=0>Normal lookup direction | Args</DEBUG>")
		}

		for i := 0; i < flag.NArg(); i++ {

			// set vars
			wpos, word, wipa, infx = getDataByID(getNavID(flag.Args()[i]))
			dbid = getNavID(word)
			lwrd = getLocalWordByID(dbid, lang)

			// print out the results
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
	} else if *flag_r && flag.NArg() > 0 {
		if *DEBUG {
			fmt.Println("<DEBUG *flag_r flag.NArg()>0>Reverse lookup direction | Args</DEBUG>")
		}

		for i := 0; i < flag.NArg(); i++ {
			// set vars
			lwrd = flag.Args()[i]
			dbls = getLocID(lwrd, lang)
			if *DEBUG {
				fmt.Println("<DEBUG:main() dbls>",dbls,"</DEBUG>")
			}
			for i := 0; i < len(dbls); i++ {
				dbid = dbls[i]
				wpos, word, wipa, infx = getDataByID(dbid)
				lwds = append(lwds, lwrd)

				//print results
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
			fmt.Println("")
		}
		if len(dbls) == 0 {
			fmt.Println("")
		}
	}

	// INTERACTIVE MODE

	// Na'vi -> local lookup or Local -> Na'vi lookup
	if !*flag_r && flag.NArg() == 0 {

		// print the program Header text
		fmt.Println(txt.Text("HEADTEXT"))
		if *DEBUG {
			fmt.Println("<DEBUG=true></DEBUG>")
			fmt.Println("<DEBUG !*flag_r flag.NArg()==0>Normal lookup direction | Interactive</DEBUG>")
		}

		fmt.Println("")

		// set vars
		// read word from cli
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Fwew:> ")
		word, _ := reader.ReadString('\n')
		word = strings.Trim(word, "\n")
		if word != "" {
			wpos, _, wipa, infx = getDataByID(getNavID(word))
			dbid = getNavID(word)
			lwrd = getLocalWordByID(dbid, lang)
		} else {
			wpos = ""
			wipa = "[]"
			infx = "\\N"
			dbid = "-1"
			lwrd = ""
		}
		// print out the results
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

	} else if *flag_r && flag.NArg() == 0 {

		// print the program Header text
		fmt.Println(txt.Text("HEADTEXT"))
		if *DEBUG {
			fmt.Println("<DEBUG=true></DEBUG>")
			fmt.Println("<DEBUG *flag_r flag.NArg()==0>Reverse lookup direction | Interactive</DEBUG>")
		}
		fmt.Println("")

		//read word from cli
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Fwew:> ")
		word, _ := reader.ReadString('\n')
		word = strings.Trim(word, "\n")

		// set vars
		lwrd = word
		dbls = getLocID(lwrd, lang)
		for i := 0; i < len(dbls); i++ {
			dbid = dbls[i]
			wpos, word, wipa, infx = getDataByID(dbid)
			lwds = append(lwds, lwrd)

			//print results
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
		fmt.Println("")
	}

}
