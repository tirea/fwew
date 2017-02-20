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

// package config handles... the configuration file stuff. Probably.
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tirea/fwew/util"
)

type Config struct {
	Language   string `json:"language"`
	PosFilter  string `json:"posFilter"`
	UseAffixes bool   `json:"useAffixes"`
}

func ReadConfig() Config {
	configfile, e := ioutil.ReadFile(util.Text("config"))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
	}

	var config Config
	json.Unmarshal(configfile, &config)

	return config
}

func (c Config) String() string {
	return fmt.Sprintf("Language: %s\nPosFilter: %s\nUseAffixes: %t\n", c.Language, c.PosFilter, c.UseAffixes)
}
