//  This file is part of Fwew.
//  Fwew is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  Fwew is distributed in the hope that it will be usefil,
//  but WITHOUT ANY WARRANTY; without even implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with Fwew.  If not, see http://gnu.org/licenses/

// This util library handles all the debug probing output

package util

import (
  "fmt"
)

var head string = "<DEBUG:"
var mid string = ">"
var tail string = "</DEBUG>"

// output status of a bunch of string variables
func DebugSnap(progDebug bool, funcName string, varNames []string, varVals []string) {
  if progDebug {
    for i := 0; i < len(varNames); i++ {
      fmt.Println(head + funcName + varNames[i] + mid + varVals[i] + tail)
    }
  }
}

// output status of a single string variable
func DebugVar(progDebug bool, funcName string, varName string, varVal string) {
  if progDebug {
    fmt.Println(head + funcName + varName + mid + varVal + tail)
  }
}

// output status of a single array variable
func DebugArr(progDebug bool, funcName string, varName string, varVal []string) {
  if progDebug {
    fmt.Println(head + funcName + varName + mid, varVal, tail)
    //fmt.Println(varVal)
    //fmt.Println(tail)
  }
}

// output status of result [][]string 
func DebugResVar(progDebug bool, funcName string, varName string, varVal [][]string) {
  if progDebug {
    fmt.Println(head + funcName + varName + mid, varVal, tail) 
    //fmt.Println(varVal)
    //fmt.Println(tail)
  }
}