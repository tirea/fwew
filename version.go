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

// Package main contains all the things. version.go handles program version.
package main

import (
	"fmt"
	"strconv"
)

type version struct {
	Major, Minor, Patch int
	Label               string
	Name                string
	DictVersion         float64
	DictBuild           string
}

// Version is a printable version struct containing program version information
var Version = version{
	4, 2, 0,
	"dev",
	"Fkewa Fkio",
	14.4,
	"",
}

func (v version) String() string {
	if v.Label != "" {
		return fmt.Sprintf("%s %d.%d.%d-%s \"%s\"\ndictionary %s (EE %s)",
			Text("name"), v.Major, v.Minor, v.Patch, v.Label, v.Name, v.DictBuild, strconv.FormatFloat(v.DictVersion, 'f', -1, 64))
	}

	return fmt.Sprintf("%s %d.%d.%d \"%s\"\ndictionary %s (EE %s)",
		Text("name"), v.Major, v.Minor, v.Patch, v.Name, v.DictBuild, strconv.FormatFloat(v.DictVersion, 'f', -1, 64))
}
