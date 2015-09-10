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

//	Package txt contains all the program strings used by Fwew

package txt

import (
	"os/user"
	"path/filepath"
)

var USER, err = user.Current()
var texts = map[string]string{}

func init() {
	texts["NAME"] = "Fwew"
	texts["VERSION"] = "1.00-BETA"
	texts["AUTHOR"] = "Tirea Aean"
	texts["USER"] = USER.Username
	texts["HOMEDIR"], _ = filepath.Abs(USER.HomeDir)
	texts["DATADIR"] = filepath.Join(texts["HOMEDIR"], ".fwew")
	texts["METAWORDS"] = filepath.Join(texts["DATADIR"], "metaWords.txt")
	texts["LOCALIZED"] = filepath.Join(texts["DATADIR"], "localizedWords.txt")
	texts["LANGUAGES"] = "de, eng, est, hu, nl, pl, ru, sv"
	texts["DEFAULT_LANGUAGE"] = "eng"
	texts["LANGUAGE"] = texts["DEFAULT_LANGUAGE"]
	texts["NONE"] = "- not found -"
	texts["HEADTEXT"] = texts["NAME"] + " " + texts["VERSION"] +
		" by " + texts["AUTHOR"] + "\n" +
		"Crossplatform dictionary search" + "\n" +
		"fwew -h for usage, see README"
}

func Text(s string) string {
	return texts[s]
}

func Set(s string, v string) {
	texts[s] = v
}
