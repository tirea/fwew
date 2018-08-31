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

// Package main obviously contains all the stuff for the main program
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tirea/fwew/affixes"
	"github.com/tirea/fwew/config"
	"github.com/tirea/fwew/numbers"
	"github.com/tirea/fwew/util"
)

// Global
const (
	// idField  int = 0 // dictionary.txt line Field 0 is Database ID
	lcField  int = 1 // dictionary.txt line field 1 is Language Code
	navField int = 2 // dictionary.txt line field 2 is Na'vi word
	//ipaField int = 3 // dictionary.txt line field 3 is IPA data
	//infField int = 4 // dictionary.txt line field 4 is Infix location data
	posField int = 5 // dictionary.txt line field 5 is Part of Speech data
	defField int = 6 // dictionary.txt line field 6 is Local definition
)

// flags / options
var (
	configuration            config.Config
	language, posFilter      *string
	showVersion, showInfixes *bool
	showIPA, reverse         *bool
	useAffixes, numConvert   *bool
	markdown                 *bool
)

func fwew(word string) []affixes.Word {
	var (
		result    affixes.Word
		results   []affixes.Word
		fields    []string
		defString string
		added     bool
	)

	badChars := strings.Split("` ~ @ # $ % ^ & * ( ) [ ] { } < > _ / . , ; : ! ? | + \\", " ")
	// remove all the sketchy chars from arguments
	for _, c := range badChars {
		word = strings.Replace(word, c, "", -1)
	}
	// No Results if empty string after removing sketch chars
	if len(word) == 0 {
		return results
	}
	// Prepare file for searching
	dictData, err := os.Open(util.Text("dictionary"))
	defer dictData.Close()
	if err != nil {
		fmt.Println(errors.New(util.Text("noDataError")))
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(dictData)

	// Go through each line and see what we can find
	for scanner.Scan() {
		line := scanner.Text()
		// Store the fields of the line into fields array
		fields = strings.Split(line, "\t")
		// Put the stuff from fields into the Word struct
		result = affixes.InitWordStruct(result, fields)

		// Looking for Local word in Definition field
		if *reverse {
			// whole-word matching
			defString = util.StripChars(fields[defField], ",;")
			if fields[lcField] == *language {
				if *posFilter == "all" || fields[posField] == *posFilter {
					for _, w := range strings.Split(defString, " ") {
						if strings.ToLower(w) == strings.ToLower(word) && !added {
							results = append(results, result)
							added = true
						}
					}
				}
			}
			added = false

			// Looking for Na'vi word in Na'vi field
		} else {
			if fields[lcField] == *language {
				if strings.ToLower(fields[navField]) == strings.ToLower(word) {
					results = append(results, result)
					if !*useAffixes {
						break
					}
				} else if *useAffixes {
					result.Target = word
					result = affixes.Reconstruct(result)
					if result.ID != "-1" {
						results = append(results, result)
					}
					// reset these to not catch the next word
					result.Target = ""
					result.Attempt = ""
				}
			}
		}
	}

	return results
}

func printResults(results []affixes.Word) {
	if len(results) != 0 {
		var out string

		for i, w := range results {
			num := fmt.Sprintf("[%d] ", i+1)
			nav := fmt.Sprintf("%s", w.Navi)
			ipa := fmt.Sprintf("[%s] ", w.IPA)
			pos := fmt.Sprintf("%s", w.PartOfSpeech)
			inf := fmt.Sprintf("%s ", w.InfixLocations)
			def := fmt.Sprintf("%s\n", w.Definition)

			if *markdown {
				nav = "**" + nav + "** "
				pos = "*" + pos + "* "
			} else {
				nav += " "
				pos += " "
			}

			out += num
			out += nav
			if *showIPA {
				out += ipa
			}
			if *showInfixes && w.InfixLocations != "\\N" {
				out += inf
			}
			out += pos
			out += def
			if *useAffixes && len(w.Affixes) > 0 {
				for key, value := range w.Affixes {
					out += fmt.Sprintf("    %s: %s\n", key, value)
				}
			}
		}
		out += fmt.Sprintf("\n")

		fmt.Printf(out)

	} else {
		fmt.Println(util.Text("none"))
	}
}

func setFlags(arg string, argsMode bool) {
	var (
		flagList []string
		err      error
		langs    = strings.Split(util.Text("languages"), ", ")
	)
	if argsMode {
		start := util.IndexStr(arg, '[') + 1
		flagList = strings.Split(arg[start:len(arg)-1], ",")
	} else {
		flagList = strings.Split(arg, " ")
	}
	for _, f := range flagList {
		switch {
		case f == "":
		case f == "r":
			*reverse = !*reverse
		case f == "i":
			*showInfixes = !*showInfixes
		case f == "ipa":
			*showIPA = !*showIPA
		case f == "a":
			*useAffixes = !*useAffixes
		case f == "n":
			*numConvert = !*numConvert
		case f == "m":
			*markdown = !*markdown
		case strings.HasPrefix(f, "l="):
			if util.ContainsStr(langs, f[2:]) {
				*language = f[2:]
			} else {
				err = fmt.Errorf("%s: %s (%s: %s)", util.Text("invalidLanguageError"), f[2:], util.Text("options"), util.Text("languages"))
				fmt.Println(err)
				fmt.Println()
			}
		case strings.HasPrefix(f, "p="):
			*posFilter = f[2:]
		default:
			err = fmt.Errorf("%s: %s", util.Text("noOptionError"), f)
			fmt.Println(err)
			fmt.Println()
		}
	}
	if err == nil {
		fmt.Printf("%s r=%t i=%t ipa=%t a=%t n=%t m=%t l=%s p=%s\n\n", util.Text("set"), *reverse, *showInfixes, *showIPA, *useAffixes, *numConvert, *markdown, *language, *posFilter)
	}
}

func printHelp() {
	flag.Usage = func() {
		fmt.Printf("%s: ", util.Text("usage"))
		fmt.Printf("%s [%s] [%s]\n", util.Text("bin"), util.Text("options"), util.Text("words"))
		fmt.Printf("%s:\n", util.Text("options"))
		flag.PrintDefaults()
	}
	flag.Usage()
}

func syllableCount(w affixes.Word) int64 {
	// syllable dot counter
	var sdc int64 = 0
	for _, char := range w.IPA {
		if char == '.' {
			sdc += 1
		}
	}
	return sdc + 1
}

func listWords(args []string) {
	var (
		result  affixes.Word
		results []affixes.Word
		fields  []string
		what    = args[0]
		cond    = args[1]
		spec    = args[2]
	)
	// /list what cond spec
	// /list pos has svin.
	// /list pos is v.
	// /list word starts ft
	// /list word ends ang
	// /list word has ts
	// /list syllables > 1
	// /list syllables = 2
	// /list syllables <= 3

	// result = affixes.InitWordStruct(result, fields)
	// results = append(results, result)

	dictData, err := os.Open(util.Text("dictionary"))
	defer dictData.Close()
	if err != nil {
		fmt.Println(errors.New(util.Text("noDataError")))
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(dictData)
	for scanner.Scan() {
		line := scanner.Text()
		fields = strings.Split(line, "\t")
		if fields[lcField] == *language {
			switch what {
			case "pos":
				if cond == "is" {
					if fields[posField] == spec {
						result = affixes.InitWordStruct(result, fields)
						results = append(results, result)
					}
				} else if cond == "has" {
					if strings.Contains(fields[posField], spec) {
						result = affixes.InitWordStruct(result, fields)
						results = append(results, result)
					}
				}
			case "word":
				if cond == "starts" {
					if strings.HasPrefix(fields[navField], spec) {
						result = affixes.InitWordStruct(result, fields)
						results = append(results, result)
					}
				} else if cond == "ends" {
					if strings.HasSuffix(fields[navField], spec) {
						result = affixes.InitWordStruct(result, fields)
						results = append(results, result)
					}
				} else if cond == "has" {
					if strings.Contains(fields[navField], spec) {
						result = affixes.InitWordStruct(result, fields)
						results = append(results, result)
					}
				}
			case "syllables":
				result = affixes.InitWordStruct(result, fields)
				ispec, err := strconv.ParseInt(spec, 10, 64)
				if err != nil {
					fmt.Println(util.Text("invalidDecimalError"))
					return
				}
				switch cond {
				case "<":
					if syllableCount(result) < ispec {
						results = append(results, result)
					}
				case "<=":
					if syllableCount(result) <= ispec {
						results = append(results, result)
					}
				case "=":
					if syllableCount(result) == ispec {
						results = append(results, result)
					}
				case ">=":
					if syllableCount(result) >= ispec {
						results = append(results, result)
					}
				case ">":
					if syllableCount(result) > ispec {
						results = append(results, result)
					}
				}
			}
		}
	}
	printResults(results)
}

func slashCommand(s string, argsMode bool) {
	var (
		sc      []string
		command string
		args    []string
	)
	sc = strings.Split(s, " ")
	command = sc[0]
	if len(sc) > 1 {
		args = sc[1:]
	}
	switch command {
	case "/help":
		printHelp()
	case "/commands":
		fmt.Println(util.Text("slashCommandHelp"))
	case "/set":
		setFlags(strings.Join(args, " "), argsMode)
	case "/unset":
		setFlags(strings.Join(args, " "), argsMode)
	case "/list":
		if len(args) == 3 {
			listWords(args)
		} else {
			fmt.Println()
		}
	case "/update":
		util.DownloadDict()
	case "/quit", "/exit", "/q", "/wc":
		os.Exit(0)
	default:
		fmt.Println()
	}
}

func main() {
	var (
		results  []affixes.Word
		argsMode bool
	)
	configuration = config.ReadConfig()
	// Version flag, for displaying version data
	showVersion = flag.Bool("v", false, util.Text("usageV"))
	// Reverse direction flag, for local_lang -> Na'vi lookups
	reverse = flag.Bool("r", false, util.Text("usageR"))
	// Language specifier flag
	language = flag.String("l", configuration.Language, util.Text("usageL"))
	// Infixes flag, opt to show infix location data
	showInfixes = flag.Bool("i", false, util.Text("usageI"))
	// IPA flag, opt to show IPA data
	showIPA = flag.Bool("ipa", false, util.Text("usageIPA"))
	// Show part of speech flag
	posFilter = flag.String("p", configuration.PosFilter, util.Text("usageP"))
	// Attempt to find all matches using affixes
	useAffixes = flag.Bool("a", configuration.UseAffixes, util.Text("usageA"))
	// Convert numbers
	numConvert = flag.Bool("n", false, util.Text("usageN"))
	// Markdown formatting
	markdown = flag.Bool("m", false, util.Text("usageM"))
	flag.Parse()
	argsMode = flag.NArg() > 0

	if *showVersion {
		fmt.Println(util.Version)
		if flag.NArg() == 0 {
			os.Exit(0)
		}
	}

	// ARGS MODE
	if argsMode {
		for _, arg := range flag.Args() {
			arg = strings.Replace(arg, "’", "'", -1)
			if strings.HasPrefix(arg, "set[") && strings.HasSuffix(arg, "]") {
				setFlags(arg, argsMode)
			} else if strings.HasPrefix(arg, "unset[") && strings.HasSuffix(arg, "]") {
				setFlags(arg, argsMode)
			} else {
				if *numConvert {
					fmt.Println(numbers.Convert(arg, *reverse))
				} else {
					results = fwew(arg)
					printResults(results)
				}
			}
		}

		// INTERACTIVE MODE
	} else {
		fmt.Println(util.Text("header"))

		for {
			fmt.Print(util.Text("prompt"))

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()
			input = strings.Replace(input, "’", "'", -1)

			// catch EOF error
			if err := scanner.Err(); err != nil {
				fmt.Println()
				os.Exit(0)
			}

			if input != "" {
				if strings.HasPrefix(input, "/") {
					slashCommand(input, argsMode)
				} else {
					if *numConvert {
						fmt.Println(numbers.Convert(input, *reverse))
					} else {
						results = fwew(input)
						printResults(results)
					}
				}
			} else {
				fmt.Println()
			}
		}
	}
}
