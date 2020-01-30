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
		} else if len(csvCmds) == 1 {
			fmt.Println("printing newline")
		}
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	if d.GetWordBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	s := []prompt.Suggest{
		{Text: "/set", Description: Text("/setDesc")},
		{Text: "/unset", Description: Text("/unsetDesc")},
		{Text: "/list", Description: Text("/listDesc")},
		{Text: "/random", Description: Text("/randomDesc")},
		{Text: "/update", Description: Text("/updateDesc")},
		{Text: "/commands", Description: Text("/commandsDesc")},
		{Text: "/config", Description: Text("/configDesc")},
		{Text: "/version", Description: Text("/versionDesc")},
		{Text: "/help", Description: Text("/helpDesc")},
		{Text: "/exit", Description: Text("/exitDesc")},
		{Text: "/quit", Description: Text("/exitDesc")},
		{Text: "/q", Description: Text("/exitDesc")},
		{Text: "/r", Description: Text("usageR")},
		{Text: "/i", Description: Text("usageI")},
		{Text: "/ipa", Description: Text("usageIPA")},
		{Text: "/n", Description: Text("usageN")},
		{Text: "/a", Description: Text("usageA")},
		{Text: "/s", Description: Text("usageS")},
		{Text: "r", Description: Text("usageR")},
		{Text: "i", Description: Text("usageI")},
		{Text: "ipa", Description: Text("usageIPA")},
		{Text: "n", Description: Text("usageN")},
		{Text: "and", Description: Text("andDesc")},
		{Text: "a", Description: Text("usageA")},
		{Text: "m", Description: Text("usageM")},
		{Text: "s", Description: Text("usageS")},
		{Text: "c", Description: Text("/configDesc")},
		{Text: "l=de", Description: Text("l=deDesc")},
		{Text: "l=eng", Description: Text("l=engDesc")},
		{Text: "l=est", Description: Text("l=estDesc")},
		{Text: "l=fr", Description: Text("l=frDesc")},
		{Text: "l=hu", Description: Text("l=huDesc")},
		{Text: "l=nl", Description: Text("l=nlDesc")},
		{Text: "l=pl", Description: Text("l=plDesc")},
		{Text: "l=ru", Description: Text("l=ruDesc")},
		{Text: "l=sv", Description: Text("l=svDesc")},
		{Text: "pos", Description: Text("posDesc")},
		{Text: "word", Description: Text("wordDesc")},
		{Text: "words", Description: Text("wordsDesc")},
		{Text: "syllables", Description: Text("syllablesDesc")},
		{Text: "random", Description: Text("randomDesc")},
		{Text: "where", Description: Text("whereDesc")},
		{Text: "starts", Description: Text("startsDesc")},
		{Text: "ends", Description: Text("endsDesc")},
		{Text: "like", Description: Text("likeDesc")},
		{Text: "first", Description: Text("firstDesc")},
		{Text: "last", Description: Text("lastDesc")},
		{Text: "has", Description: Text("hasDesc")},
		{Text: "is", Description: Text("isDesc")},
		{Text: ">=", Description: Text(">=Desc")},
		{Text: ">", Description: Text(">Desc")},
		{Text: "<=", Description: Text("<=Desc")},
		{Text: "<", Description: Text("<Desc")},
		{Text: "=", Description: Text("=Desc")},
		{Text: "language", Description: Text("languageDesc")},
		{Text: "posFilter", Description: Text("posFilterDesc")},
		{Text: "useAffixes", Description: Text("useAffixesDesc")},
		{Text: "debugMode", Description: Text("debugModeDesc")},
		{Text: "all", Description: Text("allDesc")},
		{Text: "true", Description: Text("trueDesc")},
		{Text: "false", Description: Text("falseDesc")},
		{Text: "not-starts", Description: Text("not-startsDesc")},
		{Text: "not-ends", Description: Text("not-endsDesc")},
		{Text: "not-like", Description: Text("not-likeDesc")},
		{Text: "not-has", Description: Text("not-hasDesc")},
		{Text: "not-is", Description: Text("not-isDesc")},
		{Text: "!=", Description: Text("!=Desc")},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
