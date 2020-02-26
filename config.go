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

// Package main contains all the things. config.go handles... the configuration file stuff. Probably.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Config is a struct designed to hold the values of the configuration file when loaded
type Config struct {
	Language    string `json:"language"`
	PosFilter   string `json:"posFilter"`
	UseAffixes  bool   `json:"useAffixes"`
	ShowInfixes bool   `json:"showInfixes"`
	ShowIPA     bool   `json:"showIPA"`
	ShowInfDots bool   `json:"showInfDots"`
	ShowDashed  bool   `json:"showDashed"`
	ShowSource  bool   `json:"showSource"`
	NumConvert  bool   `json:"numConvert"`
	Markdown    bool   `json:"markdown"`
	Reverse     bool   `json:"reverse"`
	DebugMode   bool   `json:"DebugMode"`
}

// ReadConfig reads a configuration file and puts the data into Config struct
func ReadConfig() Config {
	configFile, e := ioutil.ReadFile(Text("config"))
	if e != nil {
		fmt.Println(Text("fileError"))
		log.Fatal(e)
	}

	var config Config
	err := json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

// WriteConfig saves specified options to the config file
func WriteConfig(entry string) Config {
	var s []string
	const strTrue = "true"
	const strFalse = "false"
	config := ReadConfig()
	if strings.Contains(entry, " ") {
		s = strings.Split(entry, " ")
	} else if strings.Contains(entry, "=") {
		s = strings.Split(entry, "=")
	}

	// parse key-value pairs from user input and store as data in Config struct
	if len(s) == 2 {
		key := s[0]
		value := s[1]
		switch strings.ToLower(key) {
		case "language":
			if strings.Contains(Text("languages"), value) {
				config.Language = value
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "posfilter":
			if strings.Contains(Text("POSFilters"), value) {
				config.PosFilter = value
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "useaffixes":
			if value == strTrue {
				config.UseAffixes = true
			} else if value == strFalse {
				config.UseAffixes = false
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "showinfixes":
			if value == strTrue {
				config.ShowInfixes = true
			} else if value == strFalse {
				config.ShowInfixes = false
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "showipa":
			if value == strTrue {
				config.ShowIPA = true
			} else if value == strFalse {
				config.ShowIPA = false
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "showinfdots":
			if value == strTrue {
				config.ShowInfDots = true
			} else if value == strFalse {
				config.ShowInfDots = false
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "showdashed":
			if value == strTrue {
				config.ShowDashed = true
			} else if value == strFalse {
				config.ShowDashed = false
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "showsource":
			if value == strTrue {
				config.ShowSource = true
			} else if value == strFalse {
				config.ShowSource = false
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "numconvert":
			if value == strTrue {
				config.NumConvert = true
			} else if value == strFalse {
				config.NumConvert = false
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "markdown":
			if value == strTrue {
				config.Markdown = true
			} else if value == strFalse {
				config.Markdown = false
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "reverse":
			if value == strTrue {
				config.Reverse = true
			} else if value == strFalse {
				config.Reverse = false
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "debugmode":
			if value == strTrue {
				config.DebugMode = true
			} else if value == strFalse {
				config.DebugMode = false
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		default:
			fmt.Printf("%s: %s\n\n", Text("configOptionError"), key)
			return config
		}

		// convert Config struct to JSON
		jconf, err := json.Marshal(config)
		if err != nil {
			log.Fatal(err)
		}

		e := ioutil.WriteFile(Text("config"), jconf, 0644)
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println(Text("configSaved"))

	} else if len(s) == 0 {
		fmt.Println(config)
	} else {
		fmt.Println(Text("configSyntaxError"))
	}
	return config
}

func (c Config) String() string {
	var str string
	str += fmt.Sprintf("Language: %s\n", c.Language)
	str += fmt.Sprintf("PosFilter: %s\n", c.PosFilter)
	str += fmt.Sprintf("UseAffixes: %t\n", c.UseAffixes)
	str += fmt.Sprintf("ShowInfixes: %t\n", c.ShowInfixes)
	str += fmt.Sprintf("ShowIPA: %t\n", c.ShowIPA)
	str += fmt.Sprintf("ShowInfDots: %t\n", c.ShowInfDots)
	str += fmt.Sprintf("ShowDashed: %t\n", c.ShowDashed)
	str += fmt.Sprintf("ShowSource: %t\n", c.ShowSource)
	str += fmt.Sprintf("NumConvert: %t\n", c.NumConvert)
	str += fmt.Sprintf("Markdown: %t\n", c.Markdown)
	str += fmt.Sprintf("Reverse: %t\n", c.Reverse)
	str += fmt.Sprintf("DebugMode: %t\n", c.DebugMode)
	// this string only doesn't get translated or called from Text() because they're var names
	return str
}
