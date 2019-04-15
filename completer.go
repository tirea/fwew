//	This file is part of Fwew.
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

// Package main contains all the things. completer.go contains the ingredients for prompt completion.
package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"strings"
)

func executor(cmds string) {
	csvCmds := strings.Split(cmds, ",")
	for _, cmd := range csvCmds {
		cmd = strings.Trim(cmd, " ")
		if cmd != "" {
			if strings.HasPrefix(cmd, "/") {
				slashCommand(cmd, false)
			} else {
				if *numConvert {
					fmt.Println(Convert(cmd, *reverse))
				} else {
					printResults(fwew(cmd))
				}
			}
		} else {
			fmt.Println()
		}
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	if d.GetWordBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	s := []prompt.Suggest{
		{Text: "/set", Description: "set option(s)"},
		{Text: "/unset", Description: "unset option(s)"},
		{Text: "/list", Description: "list entries satisfying given condition(s)"},
		{Text: "/random", Description: "list random entries"},
		{Text: "/update", Description: "update the dictionary data file"},
		{Text: "/commands", Description: "show commands help"},
		{Text: "/help", Description: "show usage help"},
		{Text: "/exit", Description: "end program"},
		{Text: "/quit", Description: "end program"},
		{Text: "/q", Description: "end program"},
		{Text: "r", Description: Text("usageR")},
		{Text: "i", Description: Text("usageI")},
		{Text: "ipa", Description: Text("usageIPA")},
		{Text: "n", Description: Text("usageN")},
		{Text: "a", Description: Text("usageA")},
		{Text: "m", Description: Text("usageM")},
		{Text: "s", Description: Text("usageS")},
		{Text: "l=de", Description: "Deutsch"},
		{Text: "l=eng", Description: "English"},
		{Text: "l=est", Description: "Eesti"},
		{Text: "l=hu", Description: "Magyar"},
		{Text: "l=nl", Description: "Nederlands"},
		{Text: "l=pl", Description: "Polski"},
		{Text: "l=ru", Description: "Русский"},
		{Text: "l=sv", Description: "Svenska"},
		{Text: "pos", Description: "part of speech"},
		{Text: "word", Description: "word"},
		{Text: "words", Description: "words"},
		{Text: "syllables", Description: "syllables"},
		{Text: "random", Description: "random number"},
		{Text: "where", Description: "add condition to random"},
		{Text: "starts", Description: "field starts with"},
		{Text: "ends", Description: "field ends with"},
		{Text: "like", Description: "field matches wildcard expression"},
		{Text: "first", Description: "list oldest words"},
		{Text: "last", Description: "list newest words"},
		{Text: "has", Description: "all matches of condition"},
		{Text: "is", Description: "exact matches of condition"},
		{Text: ">=", Description: "syllable count greater or equal"},
		{Text: ">", Description: "syllable count greater"},
		{Text: "<=", Description: "syllable count less or equal"},
		{Text: "<", Description: "syllable count less"},
		{Text: "=", Description: "syllable count equal"},
		{Text: "and", Description: "add condition to narrow search"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
