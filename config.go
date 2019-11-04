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
	Language   string `json:"language"`
	PosFilter  string `json:"posFilter"`
	UseAffixes bool   `json:"useAffixes"`
	DebugMode  bool   `json:"DebugMode"`
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

func WriteConfig(entry string) Config {
	var s []string
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
			if value == "true" {
				config.UseAffixes = true
			} else if value == "false" {
				config.UseAffixes = false
			} else {
				fmt.Printf("%s %s: %s\n\n", Text("configValueError"), key, value)
				return config
			}
		case "debugmode":
			if value == "true" {
				config.DebugMode = true
			} else if value == "false" {
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
	// this string only doesn't get translated or called from Text() because they're var names
	return fmt.Sprintf("Language: %s\nPosFilter: %s\nUseAffixes: %t\nDebugMode: %t\n", c.Language, c.PosFilter, c.UseAffixes, c.DebugMode)
}
