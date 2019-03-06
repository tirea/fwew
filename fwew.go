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

// Package main contains all the things
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/c-bata/go-prompt"
)

// Global
const (
	idField  int = 0 // dictionary.txt line Field 0 is Database ID
	lcField  int = 1 // dictionary.txt line field 1 is Language Code
	navField int = 2 // dictionary.txt line field 2 is Na'vi word
	ipaField int = 3 // dictionary.txt line field 3 is IPA data
	infField int = 4 // dictionary.txt line field 4 is Infix location data
	posField int = 5 // dictionary.txt line field 5 is Part of Speech data
	defField int = 6 // dictionary.txt line field 6 is Local definition
	srcField int = 7 // dictionary.txt line field 7 is Source data
)

// flags / options
var (
	configuration            Config
	language, posFilter      *string
	showVersion, showInfixes *bool
	showIPA, reverse         *bool
	showSource               *bool
	useAffixes, numConvert   *bool
	markdown                 *bool
	filename                 *string
)

func intersection(a, b string) (c string) {
	m := make(map[rune]bool)
	for _, r := range a {
		m[r] = true
	}
	for _, r := range b {
		if _, ok := m[r]; ok {
			c += string(r)
		}
	}
	return
}

func similarity(w0, w1 string) float64 {
	if w0 == w1 {
		return 1.0
	}
	if len(w0) > len(w1)+1 {
		return 0.0
	}
	vowels := "aäeiìoulr"
	w0v := intersection(w0, vowels)
	w1v := intersection(w1, vowels)
	wis := intersection(w0, w1)
	wav := intersection(w0v, w1)
	if len(w0v) > len(w1v) {
		return 0.0
	}
	if len(wav) == 0 {
		return 0.0
	}
	scc := len(wis)
	iratio := float64(scc) / float64(len(w0))
	lratio := float64(len(w0)) / float64(len(w1))
	return (iratio + lratio) / 2
}

func fwew(word string) []Word {
	var (
		result    Word
		results   []Word
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
	// hardcoded hack for tseyä
	if word == "tseyä" {
		result = InitWordStruct(result, []string{
			"5268", "eng", "tsaw", "t͡saw", "NULL", "dem., pn.", "that, it (as intransitive subject)",
			"http://forum.learnnavi.org/language-updates/some-glossed-over-words/msg254625/#msg254625 (03 Jul 2010)",
		})
		result.Affixes[Text("suf")] = []string{"yä"}
		results = append(results, result)
		return results
	}
	// Prepare file for searching
	dictData, err := os.Open(Text("dictionary"))
	if err != nil {
		fmt.Println(errors.New(Text("noDataError")))
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(dictData)

	// Go through each line and see what we can find
	for scanner.Scan() {
		line := scanner.Text()
		// Store the fields of the line into fields array
		fields = strings.Split(line, "\t")
		// Put the stuff from fields into the Word struct
		result = InitWordStruct(result, fields)

		if fields[lcField] == *language {
			// Looking for Local word in Definition field
			if *reverse {
				// whole-word matching
				defString = StripChars(fields[defField], ",;")
				if *posFilter == "all" || fields[posField] == *posFilter {
					for _, w := range strings.Split(defString, " ") {
						if strings.ToLower(w) == strings.ToLower(word) && !added {
							results = append(results, result)
							added = true
						}
					}
				}
				added = false

				// Looking for Na'vi word in Na'vi field
			} else {
				// be able to search with or without --word+ marking
				fields[navField] = strings.Replace(fields[navField], "+", "", -1)
				fields[navField] = strings.Replace(fields[navField], "--", "", -1)
				word = strings.Replace(word, "--", "", -1)
				if strings.ToLower(fields[navField]) == strings.ToLower(word) {
					results = append(results, result)
					if !*useAffixes {
						break
					}
				} else if *useAffixes {
					// skip words that obviously won't work
					s := similarity(fields[navField], word)
					if s < 0.50 {
						//fmt.Printf("Target: %s | Line: %s | [%f]\n", word, fields[navField], s)
						continue
					}
					result.Target = word
					result = Reconstruct(result)
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
	err = dictData.Close()
	if err != nil {
		fmt.Println(errors.New(Text("dictCloseError")))
		log.Fatal(err)
	}
	return results
}

func printResults(results []Word) {
	if len(results) != 0 {
		var out string

		for i, w := range results {
			num := fmt.Sprintf("[%d] ", i+1)
			nav := fmt.Sprintf("%s", w.Navi)
			ipa := fmt.Sprintf("[%s] ", w.IPA)
			pos := fmt.Sprintf("%s", w.PartOfSpeech)
			inf := fmt.Sprintf("%s ", w.InfixLocations)
			def := fmt.Sprintf("%s\n", w.Definition)
			src := fmt.Sprintf("    %s: %s\n", Text("src"), w.Source)

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
			if *showInfixes && w.InfixLocations != "\\N" && w.InfixLocations != "NULL" {
				out += inf
			}
			out += pos
			out += def
			if *useAffixes && len(w.Affixes) > 0 {
				for key, value := range w.Affixes {
					out += fmt.Sprintf("    %s: %s\n", key, value)
				}
			}
			if *showSource && w.Source != "" {
				out += src
			}
		}
		out += fmt.Sprintf("\n")

		fmt.Print(out)

	} else {
		fmt.Println(Text("none"))
	}
}

func setFlags(arg string, argsMode bool) {
	var (
		flagList []string
		err      error
		langs    = strings.Split(Text("languages"), ", ")
	)
	if argsMode {
		start := IndexStr(arg, '[') + 1
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
		case f == "s":
			*showSource = !*showSource
		case f == "a":
			*useAffixes = !*useAffixes
		case f == "n":
			*numConvert = !*numConvert
		case f == "m":
			*markdown = !*markdown
		case strings.HasPrefix(f, "l="):
			if ContainsStr(langs, f[2:]) {
				*language = f[2:]
			} else {
				err = fmt.Errorf("%s: %s (%s: %s)", Text("invalidLanguageError"), f[2:], Text("options"), Text("languages"))
				fmt.Println(err)
				fmt.Println()
			}
		case strings.HasPrefix(f, "p="):
			*posFilter = f[2:]
		default:
			err = fmt.Errorf("%s: %s", Text("noOptionError"), f)
			fmt.Println(err)
			fmt.Println()
		}
	}
	if err == nil {
		var out string
		out += Text("set") + " "
		out += "[ "
		if *reverse {
			out += "r "
		}
		if *showInfixes {
			out += "i "
		}
		if *showIPA {
			out += "ipa "
		}
		if *showSource {
			out += "s "
		}
		if *useAffixes {
			out += "a "
		}
		if *numConvert {
			out += "n "
		}
		if *markdown {
			out += "m "
		}
		out += fmt.Sprintf("l=%s p=%s", *language, *posFilter)
		out += " ]\n"
		if len(*filename) == 0 {
			fmt.Println(out)
		}
	}
}

func printHelp() {
	flag.Usage = func() {
		fmt.Printf("%s: ", Text("usage"))
		fmt.Printf("%s [%s] [%s]\n", Text("bin"), Text("options"), Text("words"))
		fmt.Printf("%s:\n", Text("options"))
		flag.PrintDefaults()
	}
	flag.Usage()
}

func syllableCount(w Word) int64 {
	// syllable dot counter
	var sdc int64 = 0
	for _, char := range w.IPA {
		if char == '.' {
			sdc += 1
		}
	}
	return sdc + 1
}

func listWordsSubset(args []string, subset []Word) []Word {
	var (
		results []Word
		what    = args[0]
		cond    = args[1]
		spec    = args[2]
	)
	// /list word starts tì and pos is n.
	// /list syllables > 2 and pos is v.
	for _, w := range subset {
		switch what {
		case "pos":
			if cond == "is" {
				if w.PartOfSpeech == spec {
					results = append(results, w)
				}
			} else if cond == "has" {
				if strings.Contains(w.PartOfSpeech, spec) {
					results = append(results, w)
				}
			}
		case "word":
			if cond == "starts" {
				if strings.HasPrefix(w.Navi, spec) {
					results = append(results, w)
				}
			} else if cond == "ends" {
				if strings.HasSuffix(w.Navi, spec) {
					results = append(results, w)
				}
			} else if cond == "has" {
				if strings.Contains(w.Navi, spec) {
					results = append(results, w)
				}
			}
		case "syllables":
			ispec, err := strconv.ParseInt(spec, 10, 64)
			if err != nil {
				fmt.Println(Text("invalidDecimalError"))
				return nil
			}
			switch cond {
			case "<":
				if syllableCount(w) < ispec {
					results = append(results, w)
				}
			case "<=":
				if syllableCount(w) <= ispec {
					results = append(results, w)
				}
			case "=":
				if syllableCount(w) == ispec {
					results = append(results, w)
				}
			case ">=":
				if syllableCount(w) >= ispec {
					results = append(results, w)
				}
			case ">":
				if syllableCount(w) > ispec {
					results = append(results, w)
				}
			}
		}
	}
	return results
}

func countLines() int {
	var (
		count  int
		fields []string
	)
	dictData, err := os.Open(Text("dictionary"))
	if err != nil {
		fmt.Println(errors.New(Text("noDataError")))
		log.Fatal(err)
	}
	count = 1
	scanner := bufio.NewScanner(dictData)
	for scanner.Scan() {
		line := scanner.Text()
		fields = strings.Split(line, "\t")
		if fields[lcField] == *language {
			count++
		}
	}
	err = dictData.Close()
	if err != nil {
		fmt.Println(errors.New(Text("dictCloseError")))
		log.Fatal(err)
	}
	return count
}

func listWords(args []string) []Word {
	var (
		result   Word
		results  []Word
		fields   []string
		what     = args[0]
		cond     = args[1]
		spec     = args[2]
		count    int
		numLines int
	)
	// /list what cond spec
	// /list pos has svin.
	// /list pos is v.
	// /list word starts ft
	// /list word ends ang
	// /list word has ts
	// /list words first 20
	// /list words last 30
	// /list syllables > 1
	// /list syllables = 2
	// /list syllables <= 3

	dictData, err := os.Open(Text("dictionary"))
	if err != nil {
		fmt.Println(errors.New(Text("noDataError")))
		log.Fatal(err)
	}
	count = 1
	numLines = countLines()
	scanner := bufio.NewScanner(dictData)
	for scanner.Scan() {
		line := scanner.Text()
		fields = strings.Split(line, "\t")
		if fields[lcField] == *language {
			switch what {
			case "pos":
				spec = strings.ToLower(spec)
				if cond == "is" {
					if fields[posField] == spec {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				} else if cond == "has" {
					if strings.Contains(fields[posField], spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				}
			case "word":
				spec = strings.ToLower(spec)
				word := strings.ToLower(fields[navField])
				if cond == "starts" {
					if strings.HasPrefix(word, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				} else if cond == "ends" {
					if strings.HasSuffix(word, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				} else if cond == "has" {
					if strings.Contains(word, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				}
			case "words":
				s, err := strconv.Atoi(spec)
				if err != nil {
					log.Fatal(err)
				}
				if cond == "first" {
					if count <= s {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				} else if cond == "last" {
					if count >= numLines-s && count <= numLines {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				}
				count++
			case "syllables":
				result = InitWordStruct(result, fields)
				ispec, err := strconv.ParseInt(spec, 10, 64)
				if err != nil {
					fmt.Println(Text("invalidDecimalError"))
					return nil
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
	err = dictData.Close()
	if err != nil {
		fmt.Println(errors.New(Text("dictCloseError")))
		log.Fatal(err)
	}
	return results
}

func randomSubset(k int, subset []Word) []Word {
	var (
		results []Word
		i       int
		r       *rand.Rand
	)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if k == -1337 {
		k = r.Intn(r.Intn(countLines()))
	} else if k < 1 {
		return results
	}
	for _, w := range subset {
		if w.LangCode == *language {
			if i < k {
				results = append(results, w)
			} else {
				j := r.Intn(i)
				if j < k {
					results[j] = w
				}
			}
			i++
		}
	}
	return results
}

func random(k int) []Word {
	var (
		results []Word
		result  Word
		fields  []string
		i       int
		r       *rand.Rand
	)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if k == -1337 {
		k = r.Intn(r.Intn(countLines()))
	} else if k < 1 {
		return results
	}
	dictData, err := os.Open(Text("dictionary"))
	if err != nil {
		fmt.Println(errors.New(Text("noDataError")))
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(dictData)
	for scanner.Scan() {
		line := scanner.Text()
		fields = strings.Split(line, "\t")
		if fields[lcField] == *language {
			if i < k {
				result = InitWordStruct(result, fields)
				results = append(results, result)
			} else {
				j := r.Intn(i)
				if j < k {
					result = InitWordStruct(result, fields)
					results[j] = result
				}
			}
			i++
		}
	}
	err = dictData.Close()
	if err != nil {
		fmt.Println(errors.New(Text("dictCloseError")))
		log.Fatal(err)
	}
	return results
}

func slashCommand(s string, argsMode bool) {
	var (
		sc      []string
		command string
		args    []string
		exprs   [][]string
		nargs   int
	)
	sc = strings.Split(s, " ")
	sc = DeleteEmpty(sc)
	command = sc[0]
	if len(sc) > 1 {
		args = sc[1:]
		nargs = len(args)
	}
	switch command {
	case "/help":
		printHelp()
	case "/commands":
		fmt.Println(Text("slashCommandHelp"))
	case "/set":
		setFlags(strings.Join(args, " "), argsMode)
	case "/unset":
		setFlags(strings.Join(args, " "), argsMode)
	case "/list":
		// word starts tì
		if nargs == 3 {
			printResults(listWords(args))
			// word starts tì and pos is n. and syllables > 10
		} else if nargs > 3 {
			//validate length of args, needs to be 4n-1
			valid := false
			// 10 nested query triplets is insane overkill
			for n := 1; n <= 10; n++ {
				if 4*n-1 == nargs {
					valid = true
					break
				}
			}
			if !valid {
				fmt.Println()
				return
			}
			for i := 0; i < len(args); i += 4 {
				exprs = append(exprs, args[i:i+3])
			}
			subset := listWords(exprs[0])
			for _, expr := range exprs[1:] {
				subset = listWordsSubset(expr, subset)
			}
			printResults(subset)
		} else {
			fmt.Println()
		}
	case "/random":
		// k
		if nargs == 1 && args[0] != "" {
			if args[0] == "random" {
				printResults(random(-1337))
			} else {
				k, err := strconv.Atoi(args[0])
				if err != nil {
					log.Fatal(err)
				}
				printResults(random(k))
			}
			// k where what cond spec [and what cond spec...]
		} else if nargs >= 5 && args[1] == "where" {
			fargs := args[2:]
			nFargs := len(fargs)
			if nFargs == 3 {
				if args[0] == "random" {
					printResults(randomSubset(-1337, listWords(fargs)))
				} else {
					k, err := strconv.Atoi(args[0])
					if err != nil {
						log.Fatal(err)
					}
					printResults(randomSubset(k, listWords(fargs)))
				}
			}
		} else {
			fmt.Println()
		}
	case "/update":
		err := DownloadDict()
		if err != nil {
			log.Fatal(err)
		}
	case "/quit", "/exit", "/q", "/wc":
		os.Exit(0)
	default:
		fmt.Println()
	}
}

func main() {
	var (
		argsMode bool
		fileMode bool
	)
	configuration = ReadConfig()
	// Version flag, for displaying version data
	showVersion = flag.Bool("v", false, Text("usageV"))
	// Reverse direction flag, for local_lang -> Na'vi lookups
	reverse = flag.Bool("r", false, Text("usageR"))
	// Language specifier flag
	language = flag.String("l", configuration.Language, Text("usageL"))
	// Infixes flag, opt to show infix location data
	showInfixes = flag.Bool("i", false, Text("usageI"))
	// IPA flag, opt to show IPA data
	showIPA = flag.Bool("ipa", false, Text("usageIPA"))
	// Source flag, opt to show source data
	showSource = flag.Bool("s", false, Text("usageS"))
	// Filter part of speech flag, opt to filter by part of speech
	posFilter = flag.String("p", configuration.PosFilter, Text("usageP"))
	// Attempt to find all matches using affixes
	useAffixes = flag.Bool("a", configuration.UseAffixes, Text("usageA"))
	// Convert numbers
	numConvert = flag.Bool("n", false, Text("usageN"))
	// Markdown formatting
	markdown = flag.Bool("m", false, Text("usageM"))
	filename = flag.String("f", "", Text("usageF"))
	flag.Parse()
	argsMode = flag.NArg() > 0
	fileMode = len(*filename) > 0

	if *showVersion {
		fmt.Println(Version)
		if flag.NArg() == 0 {
			os.Exit(0)
		}
	}

	if fileMode { // FILE MODE
		inFile, err := os.Open(*filename)
		if err != nil {
			fmt.Println(errors.New(Text("noFileError")))
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(inFile)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.HasPrefix(line, "#") && line != "" {
				fmt.Printf("cmd %s\n", line)
				executor(line)
			}
		}
		err = inFile.Close()
		if err != nil {
			fmt.Println(errors.New(Text("fileCloseError")))
			log.Fatal(err)
		}
	} else if argsMode { // ARGS MODE
		for _, arg := range flag.Args() {
			arg = strings.Replace(arg, "’", "'", -1)
			if strings.HasPrefix(arg, "set[") && strings.HasSuffix(arg, "]") {
				setFlags(arg, argsMode)
			} else if strings.HasPrefix(arg, "unset[") && strings.HasSuffix(arg, "]") {
				setFlags(arg, argsMode)
			} else {
				executor(arg)
			}
		}
	} else { // INTERACTIVE MODE
		fmt.Println(Text("header"))

		p := prompt.New(executor, completer,
			prompt.OptionTitle(Text("name")),
			prompt.OptionPrefix(Text("prompt")),
			prompt.OptionSelectedDescriptionTextColor(prompt.DarkGray),
		)
		p.Run()
	}
}
