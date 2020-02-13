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
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/c-bata/go-prompt"
)

// Global
const (
	idField  int    = 0  // dictionary.txt line Field 0 is Database ID
	lcField  int    = 1  // dictionary.txt line field 1 is Language Code
	navField int    = 2  // dictionary.txt line field 2 is Na'vi word
	ipaField int    = 3  // dictionary.txt line field 3 is IPA data
	infField int    = 4  // dictionary.txt line field 4 is Infix location data
	posField int    = 5  // dictionary.txt line field 5 is Part of Speech data
	defField int    = 6  // dictionary.txt line field 6 is Local definition
	srcField int    = 7  // dictionary.txt line field 7 is Source data
	stsField int    = 8  // dictionary.txt line field 8 is Stressed syllable #
	sylField int    = 9  // dictionary.txt line field 9 is syllable breakdown
	ifdField int    = 10 // dictionary.txt line field 10 is dot-style infix data
	space    string = " "
)

// flags / options
var (
	configuration            Config
	configure, filename      *string
	language, posFilter      *string
	showInfixes, showIPA     *bool
	showInfDots, showDashed  *bool
	showVersion, showSource  *bool
	useAffixes, numConvert   *bool
	markdown, debug, reverse *bool
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
	if w0 == "nga" && w1 == "ngey" {
		return 1.0
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

	badChars := strings.Split("` ~ @ # $ % ^ & * ( ) [ ] { } < > _ / . , ; : ! ? | + \\", space)
	// remove all the sketchy chars from arguments
	for _, c := range badChars {
		word = strings.Replace(word, c, "", -1)
	}
	// normalize tìftang character
	word = strings.Replace(word, "’", "'", -1)
	word = strings.Replace(word, "‘", "'", -1)
	// No Results if empty string after removing sketch chars
	if len(word) == 0 {
		return results
	}

	// Prepare file for searching
	dictData, err := os.Open(Text("dictionary"))
	if err != nil {
		fmt.Println(Text("noDataError"))
		os.Exit(1)
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
					for _, w := range strings.Split(defString, space) {
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
					if *debug {
						fmt.Printf("Target: %s | Line: %s | [%f]\n", word, fields[navField], s)
					}
					if s < 0.50 && !strings.HasSuffix(strings.ToLower(word), "eyä") {
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
		fmt.Println(Text("dictCloseError"))
		os.Exit(1)
	}
	return results
}

func doUnderline(w Word) string {
	var (
		out              string
		mdUnderline      string
		shUnderlineA     string
		shUnderlineB     string
		dashed           string
		dSlice           []string
		stressedIndex    int
		stressedSyllable string
		err              error
	)

	if !strings.Contains(w.Syllables, "-") {
		out = w.Syllables
		return out
	}
	mdUnderline = "__"
	shUnderlineA = "\033[4m"
	shUnderlineB = "\033[0m"
	dashed = fmt.Sprintf("%s", w.Syllables)
	dSlice = strings.Split(dashed, "-")
	stressedIndex, err = strconv.Atoi(w.Stressed)
	if err != nil {
		fmt.Println(Text("invalidNumericError"))
		os.Exit(1)
	}
	stressedSyllable = dSlice[stressedIndex-1]

	if strings.Contains(stressedSyllable, " ") {
		tmp := strings.Split(stressedSyllable, " ")
		if *markdown {
			tmp[0] = mdUnderline + tmp[0] + mdUnderline
		} else {
			tmp[0] = shUnderlineA + tmp[0] + shUnderlineB
		}
		stressedSyllable = strings.Join(tmp, " ")
		dSlice[stressedIndex-1] = stressedSyllable
		out = strings.Join(dSlice, "-")
	} else {
		if *markdown {
			dSlice[stressedIndex-1] = mdUnderline + stressedSyllable + mdUnderline
		} else {
			dSlice[stressedIndex-1] = shUnderlineA + stressedSyllable + shUnderlineB
		}
		out = strings.Join(dSlice, "-")
	}

	return out
}

func printResults(results []Word) {
	if len(results) != 0 {
		var (
			out      string
			mdBold   = "**"
			mdItalic = "*"
			newline  = "\n"
			valNull  = "NULL"
		)

		for i, w := range results {
			num := fmt.Sprintf("[%d]", i+1)
			nav := fmt.Sprintf("%s", w.Navi)
			ipa := fmt.Sprintf("[%s]", w.IPA)
			pos := fmt.Sprintf("%s", w.PartOfSpeech)
			inf := fmt.Sprintf("%s", w.InfixLocations)
			def := fmt.Sprintf("%s", w.Definition)
			src := fmt.Sprintf("    %s: %s\n", Text("src"), w.Source)
			syl := doUnderline(w)
			ifd := fmt.Sprintf("%s", w.InfixDots)

			if *markdown {
				nav = mdBold + nav + mdBold
				pos = mdItalic + pos + mdItalic
			}

			out += num + space + nav + space

			if *showIPA {
				out += ipa + space
			}

			if *showInfixes && w.InfixLocations != valNull {
				out += inf + space
			}

			if *showDashed {
				out += "(" + syl
				if *showInfDots && w.InfixDots != valNull {
					out += "," + space
				} else {
					out += ")" + space
				}
			}

			if *showInfDots && w.InfixDots != valNull {
				if !*showDashed {
					out += "("
				}
				out += ifd + ")" + space

			}

			out += pos + space + def + newline

			if *useAffixes && len(w.Affixes) > 0 {
				for key, value := range w.Affixes {
					out += fmt.Sprintf("    %s: %s\n", key, value)
				}
			}

			if *showSource && w.Source != "" {
				out += src
			}
		}

		out += newline

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
		flagList = strings.Split(arg, space)
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
		case f == "id":
			*showInfDots = !*showInfDots
		case f == "s":
			*showDashed = !*showDashed
		case f == "src":
			*showSource = !*showSource
		case f == "a":
			*useAffixes = !*useAffixes
		case f == "n":
			*numConvert = !*numConvert
		case f == "m":
			*markdown = !*markdown
		case f == "d":
			*debug = !*debug
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
			err = fmt.Errorf("%s (%s)", Text("noOptionError"), f)
			fmt.Println(err)
			fmt.Println()
		}
	}
	if err == nil {
		var out string
		out += Text("set") + space
		out += "[ "
		if *reverse {
			out += "r "
		}
		if *showInfDots {
			out += "id "
		}
		if *showDashed {
			out += "s "
		}
		if *showInfixes {
			out += "i "
		}
		if *showIPA {
			out += "ipa "
		}
		if *showSource {
			out += "src "
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
		if *debug {
			out += "d "
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
		fmt.Printf("%s [%s] [%s]\n", Text("bin"), Text("options"), Text("w_words"))
		fmt.Printf("%s:\n", Text("options"))
		flag.PrintDefaults()
	}
	flag.Usage()
}

func syllableCount(w Word) int {
	var numSyllables int
	var vowels = []string{"a", "ä", "e", "i", "ì", "o", "u", "ll", "rr"}
	for _, p := range vowels {
		numSyllables += strings.Count(strings.ToLower(w.Navi), p)
	}
	return numSyllables
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
		case Text("w_words"):
			s, err := strconv.Atoi(spec)
			if err != nil {
				fmt.Printf("%s (%s)", Text("invalidNumericError"), spec)
			}
			switch cond {
			case Text("c_first"):
				if len(subset) >= s {
					return subset[0 : s-1]
				}
				return subset
			case Text("c_last"):
				if len(subset) >= s {
					return subset[len(subset)-s:]
				}
				return subset
			}
		case Text("w_pos"):
			switch cond {
			case Text("c_starts"):
				if strings.HasPrefix(w.PartOfSpeech, spec) {
					results = append(results, w)
				}
			case Text("c_ends"):
				if strings.HasSuffix(w.PartOfSpeech, spec) {
					results = append(results, w)
				}
			case Text("c_is"):
				if w.PartOfSpeech == spec {
					results = append(results, w)
				}
			case Text("c_has"):
				if strings.Contains(w.PartOfSpeech, spec) {
					results = append(results, w)
				}
			case Text("c_like"):
				if Glob(spec, w.PartOfSpeech) {
					results = append(results, w)
				}
			case Text("c_not-starts"):
				if !strings.HasPrefix(w.PartOfSpeech, spec) {
					results = append(results, w)
				}
			case Text("c_not-ends"):
				if !strings.HasSuffix(w.PartOfSpeech, spec) {
					results = append(results, w)
				}
			case Text("c_not-is"):
				if w.PartOfSpeech != spec {
					results = append(results, w)
				}
			case Text("c_not-has"):
				if !strings.Contains(w.PartOfSpeech, spec) {
					results = append(results, w)
				}
			case Text("c_not-like"):
				if !Glob(spec, w.PartOfSpeech) {
					results = append(results, w)
				}
			}
		case Text("w_word"):
			switch cond {
			case Text("c_starts"):
				if strings.HasPrefix(w.Navi, spec) {
					results = append(results, w)
				}
			case Text("c_ends"):
				if strings.HasSuffix(w.Navi, spec) {
					results = append(results, w)
				}
			case Text("c_has"):
				if strings.Contains(w.Navi, spec) {
					results = append(results, w)
				}
			case Text("c_like"):
				if Glob(spec, w.Navi) {
					results = append(results, w)
				}
			case Text("c_not-starts"):
				if !strings.HasPrefix(w.Navi, spec) {
					results = append(results, w)
				}
			case Text("c_not-ends"):
				if !strings.HasSuffix(w.Navi, spec) {
					results = append(results, w)
				}
			case Text("c_not-has"):
				if !strings.Contains(w.Navi, spec) {
					results = append(results, w)
				}
			case Text("c_not-like"):
				if !Glob(spec, w.Navi) {
					results = append(results, w)
				}
			}
		case Text("w_syllables"):
			ispec, err := strconv.Atoi(spec)
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
			case "!=":
				if syllableCount(w) != ispec {
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
		fmt.Println(Text("noDataError"))
		os.Exit(1)
	}
	count = 0
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
		fmt.Println(Text("dictCloseError"))
		os.Exit(1)
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
	// /list pos starts v
	// /list pos ends m.
	// /list pos has svin.
	// /list pos is v.
	// /list pos like *
	// /list pos not-starts v
	// /list pos not-ends m.
	// /list pos not-has svin.
	// /list pos not-is v.
	// /list pos not-like *
	// /list word starts ft
	// /list word ends ang
	// /list word has ts
	// /list word like *
	// /list word not-starts ft
	// /list word not-ends ang
	// /list word not-has ts
	// /list word not-like *
	// /list words first 20
	// /list words last 30
	// /list syllables > 1
	// /list syllables = 2
	// /list syllables <= 3

	dictData, err0 := os.Open(Text("dictionary"))
	if err0 != nil {
		fmt.Println(Text("noDataError"))
		os.Exit(1)
	}
	count = 1
	numLines = countLines()
	scanner := bufio.NewScanner(dictData)
	for scanner.Scan() {
		line := scanner.Text()
		fields = strings.Split(line, "\t")
		if fields[lcField] == *language {
			switch what {
			case Text("w_pos"):
				spec = strings.ToLower(spec)
				pos := strings.ToLower(fields[posField])
				switch cond {
				case Text("c_starts"):
					if strings.HasPrefix(pos, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_ends"):
					if strings.HasSuffix(pos, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_is"):
					if pos == spec {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_has"):
					if strings.Contains(pos, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_like"):
					if Glob(spec, pos) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_not-starts"):
					if !strings.HasPrefix(pos, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_not-ends"):
					if !strings.HasSuffix(pos, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_not-is"):
					if pos != spec {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_not-has"):
					if !strings.Contains(pos, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_not-like"):
					if !Glob(spec, pos) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				}
			case Text("w_word"):
				spec = strings.ToLower(spec)
				word := strings.ToLower(fields[navField])
				switch cond {
				case Text("c_starts"):
					if strings.HasPrefix(word, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_ends"):
					if strings.HasSuffix(word, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_has"):
					if strings.Contains(word, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_like"):
					if Glob(spec, word) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_not-starts"):
					if !strings.HasPrefix(word, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_not-ends"):
					if !strings.HasSuffix(word, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_not-has"):
					if !strings.Contains(word, spec) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_not-like"):
					if !Glob(spec, word) {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				}
			case Text("w_words"):
				s, err1 := strconv.Atoi(spec)
				if err1 != nil {
					fmt.Printf("%s (%s)\n", Text("invalidNumericError"), spec)
					os.Exit(1)
				}
				switch cond {
				case Text("c_first"):
					if count <= s {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				case Text("c_last"):
					if count >= numLines-s && count <= numLines {
						result = InitWordStruct(result, fields)
						results = append(results, result)
					}
				}
				count++
			case Text("w_syllables"):
				result = InitWordStruct(result, fields)
				ispec, err2 := strconv.Atoi(spec)
				if err2 != nil {
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
				case "!=":
					if syllableCount(result) != ispec {
						results = append(results, result)
					}
				}
			}
		}
	}
	err3 := dictData.Close()
	if err3 != nil {
		fmt.Println(Text("dictCloseError"))
		os.Exit(1)
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
				if j < k && results != nil {
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
		fmt.Println(Text("noDataError"))
		os.Exit(1)
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
				if j < k && results != nil {
					result = InitWordStruct(result, fields)
					results[j] = result
				}
			}
			i++
		}
	}
	err = dictData.Close()
	if err != nil {
		fmt.Println(Text("dictCloseError"))
		os.Exit(1)
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
		setArg  string
		confArg string
		k       int
		err     error
	)
	sc = strings.Split(s, space)
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
	case "/set", "/unset":
		setArg = strings.Join(args, space)
		setFlags(setArg, argsMode)
	// aliases for /set
	case "/a", "/id", "/s", "/i", "/ipa", "/l", "/n", "/p", "/r", "/src":
		for _, c := range command {
			if c != '/' {
				setArg += string(c)
			}
		}
		if nargs > 0 {
			setArg += space
		}
		setArg += strings.Join(args, space)
		setFlags(setArg, argsMode)
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
			if exprs != nil {
				subset := listWords(exprs[1])
				if len(exprs) > 2 {
					for _, expr := range exprs[2:] {
						subset = listWordsSubset(expr, subset)
					}
				}
				subset = listWordsSubset(exprs[0], subset)
				printResults(subset)
			}
		} else {
			fmt.Println()
		}
	case "/random":
		// k
		if args != nil && nargs == 1 && args[0] != "" {
			if args[0] == Text("n_random") {
				printResults(random(-1337))
			} else {
				k, err = strconv.Atoi(args[0])
				if err != nil {
					fmt.Printf("%s (%s)\n", Text("invalidNumericError"), args[0])
					os.Exit(1)
				}
				printResults(random(k))
			}
			// k where what cond spec [and what cond spec...]
		} else if args != nil && nargs >= 5 && args[1] == "where" {
			fargs := args[2:]
			nFargs := len(fargs)
			if nFargs == 3 {
				if args[0] == Text("n_random") {
					printResults(randomSubset(-1337, listWords(fargs)))
				} else {
					k, err = strconv.Atoi(args[0])
					if err != nil {
						fmt.Printf("%s (%s)\n", Text("invalidNumericError"), args[0])
						os.Exit(1)
					}
					printResults(randomSubset(k, listWords(fargs)))
				}
			} else if nFargs > 3 {
				//validate length of fargs, needs to be 4n-1
				valid := false
				// 10 nested query triplets is insane overkill
				for n := 1; n <= 10; n++ {
					if 4*n-1 == nFargs {
						valid = true
						break
					}
				}
				if !valid {
					fmt.Println()
					return
				}
				for i := 0; i < nFargs; i += 4 {
					exprs = append(exprs, fargs[i:i+3])
				}
				if exprs != nil {
					subset := listWords(exprs[0])
					for _, expr := range exprs[1:] {
						subset = listWordsSubset(expr, subset)
					}
					if args[0] == Text("n_random") {
						k = -1337
					} else {
						k, err = strconv.Atoi(args[0])
						if err != nil {
							fmt.Printf("%s (%s)\n", Text("invalidNumericError"), args[0])
							os.Exit(1)
						}
					}
					printResults(randomSubset(k, subset))
				}
			}
		} else {
			fmt.Println()
		}
	case "/lenition", "/len":
		fmt.Println(Text("lenTable"))
	case "/update":
		err := DownloadDict()
		if err != nil {
			fmt.Println(Text("downloadError"))
			os.Exit(1)
		}
		Version.DictBuild = SHA1Hash(Text("dictionary"))
	case "/version":
		fmt.Println(Version)
	case "/config":
		confArg = strings.Join(args, space)
		configuration = WriteConfig(confArg)
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
	// Infix locations in dot notation
	showInfDots = flag.Bool("id", false, Text("usageID"))
	// IPA flag, opt to show IPA data
	showIPA = flag.Bool("ipa", false, Text("usageIPA"))
	// Show syllable breakdown / stress
	showDashed = flag.Bool("s", false, Text("usageS"))
	// Source flag, opt to show source data
	showSource = flag.Bool("src", false, Text("usageSrc"))
	// Filter part of speech flag, opt to filter by part of speech
	posFilter = flag.String("p", configuration.PosFilter, Text("usageP"))
	// Attempt to find all matches using affixes
	useAffixes = flag.Bool("a", configuration.UseAffixes, Text("usageA"))
	// Convert numbers
	numConvert = flag.Bool("n", false, Text("usageN"))
	// Markdown formatting
	markdown = flag.Bool("m", false, Text("usageM"))
	// Input file / Fwewscript
	filename = flag.String("f", "", Text("usageF"))
	// Configuration editing
	configure = flag.String("c", "", Text("usageC"))
	// Debug mode
	debug = flag.Bool("d", configuration.DebugMode, Text("usageD"))
	flag.Parse()
	argsMode = flag.NArg() > 0
	fileMode = len(*filename) > 0

	if *showVersion {
		fmt.Println(Version)
		if flag.NArg() == 0 {
			os.Exit(0)
		}
	}

	if *configure != "" {
		configuration = WriteConfig(*configure)
		os.Exit(0)
	}

	if fileMode { // FILE MODE
		if *markdown {
			// restrict Discord users to cwd
			*filename = "./" + *filename
		}
		inFile, err := os.Open(*filename)
		if err != nil {
			fmt.Printf("%s (%s)\n", Text("noFileError"), *filename)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(inFile)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.HasPrefix(line, "#") && line != "" {
				if !*markdown {
					fmt.Printf("cmd %s\n", line)
				}
				executor(line)
			}
		}
		err = inFile.Close()
		if err != nil {
			fmt.Println(Text("fileCloseError"))
			os.Exit(1)
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
