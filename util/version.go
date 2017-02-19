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

// This util library handles all the version stuff
package util

import "fmt"

type version struct {
	Major, Minor, Patch int
	Label               string
	Name                string
	Dict                string
}

var Version = version{1, 5, 0, "dev", "Eana Yayo", "Na'vi Dictionary 13.31 (07 JAN 2017)"}

func (v version) String() string {
	if v.Label != "" {
		return fmt.Sprintf("Fwew version %d.%d.%d-%s \"%s\"\n%s", v.Major, v.Minor, v.Patch, v.Label, v.Name, v.Dict)
	} else {
		return fmt.Sprintf("Fwew version %d.%d.%d \"%s\"\n%s", v.Major, v.Minor, v.Patch, v.Name, v.Dict)
	}
}
