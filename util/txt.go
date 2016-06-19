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

// This util library handles all program strings

package util

import (
	"os/user"
	"path/filepath"
)

var USER, err = user.Current()
var texts = map[string]string{}

func init() {
	texts["NAME"] = "Fwew"
	texts["VERSION"] = "1.03-BETA (19 JUN 2016)"
	texts["DICTVERSION"] = "Dictionary 13.284 (03 MAR 2016)"
	texts["AUTHOR"] = "Tirea Aean"
	texts["LANGUAGES"] = "de, eng, est, hu, nl, pl, ru, sv"
	texts["DEFAULT_LANGUAGE"] = "eng"
	texts["LANGUAGE"] = texts["DEFAULT_LANGUAGE"]
	texts["BASELANG"] = "Na'vi"
	texts["NONE"] = "- not found -"
	texts["USAGEDEBUG"] = "Show extremely verbose probing"
	texts["USAGEFLAG_V"] = "Show program & dictionary version numbers"
	texts["USAGEFLAG_L"] = "Use specified language \n\tValid values: "+ texts["LANGUAGES"]
	texts["USAGEFLAG_I"] = "Display infix location data"
	texts["USAGEFLAG_IPA"] = "Display IPA data"
	texts["USAGEFLAG_R"] = "Reverse the lookup direction from Na'vi->Local to Local->Na'vi"
	texts["USER"] = USER.Username
	texts["HOMEDIR"], _ = filepath.Abs(USER.HomeDir)
	texts["DATADIR"] = filepath.Join(texts["HOMEDIR"], ".fwew")
	texts["METAWORDS"] = filepath.Join(texts["DATADIR"], "metaWords.txt")
	texts["LOCALIZED"] = filepath.Join(texts["DATADIR"], "localizedWords.txt")
	texts["HEADTEXT"] = texts["NAME"]+" "+texts["VERSION"]+" by "+texts["AUTHOR"]+"\n"+
						"Crossplatform "+texts["BASELANG"]+" Dictionary Search"+"\n"+
						"fwew -h for usage, see README"
	texts["INFIX_0"] = "(äp|eyk|äpeyk)?"
	texts["INFIX_1"] = "(am|ìm|ìyev|ay|ìsy|asy|ol|er|iv|arm|ìrm|ìry|ary|alm|ìlm|ìly|aly|imv|iyev|ìy|irv|ilv|us|awn)?"
	texts["INFIX_2"] = "(ei|äng|eng|ats|uy)?"
}

func Text(s string) string {
	return texts[s]
}
