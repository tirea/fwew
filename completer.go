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
		{Text: "/set", Description: Text("/set-desc")},
		{Text: "/unset", Description: Text("/unset-desc")},
		{Text: "/list", Description: Text("/list-desc")},
		{Text: "/random", Description: Text("/random-desc")},
		{Text: "/update", Description: Text("/update-desc")},
		{Text: "/commands", Description: Text("/commands-desc")},
		{Text: "/config", Description: Text("/config-desc")},
		{Text: "/version", Description: Text("/version-desc")},
		{Text: "/help", Description: Text("/help-desc")},
		{Text: "/exit", Description: Text("/exit-desc")},
		{Text: "/quit", Description: Text("/exit-desc")},
		{Text: "/q", Description: Text("/exit-desc")},
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
		{Text: "and", Description: Text("and-desc")},
		{Text: "a", Description: Text("usageA")},
		{Text: "m", Description: Text("usageM")},
		{Text: "s", Description: Text("usageS")},
		{Text: "c", Description: Text("/config-desc")},
		{Text: "l=de", Description: Text("l=de-desc")},
		{Text: "l=eng", Description: Text("l=eng-desc")},
		{Text: "l=est", Description: Text("l=est-desc")},
		{Text: "l=hu", Description: Text("l=hu-desc")},
		{Text: "l=nl", Description: Text("l=nl-desc")},
		{Text: "l=pl", Description: Text("l=pl-desc")},
		{Text: "l=ru", Description: Text("l=ru-desc")},
		{Text: "l=sv", Description: Text("l=sv-desc")},
		{Text: "pos", Description: Text("pos-desc")},
		{Text: "word", Description: Text("word-desc")},
		{Text: "words", Description: Text("words-desc")},
		{Text: "syllables", Description: Text("syllables-desc")},
		{Text: "random", Description: Text("random-desc")},
		{Text: "where", Description: Text("where-desc")},
		{Text: "starts", Description: Text("starts-desc")},
		{Text: "ends", Description: Text("ends-desc")},
		{Text: "like", Description: Text("like-desc")},
		{Text: "first", Description: Text("first-desc")},
		{Text: "last", Description: Text("last-desc")},
		{Text: "has", Description: Text("has-desc")},
		{Text: "is", Description: Text("is-desc")},
		{Text: ">=", Description: Text(">=-desc")},
		{Text: ">", Description: Text(">-desc")},
		{Text: "<=", Description: Text("<=-desc")},
		{Text: "<", Description: Text("<-desc")},
		{Text: "=", Description: Text("=-desc")},
		{Text: "not-starts", Description: Text("not-starts-desc")},
		{Text: "not-ends", Description: Text("not-ends-desc")},
		{Text: "not-like", Description: Text("not-like-desc")},
		{Text: "not-has", Description: Text("not-has-desc")},
		{Text: "not-is", Description: Text("not-is-desc")},
		{Text: "!=", Description: Text("!=-desc")},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
