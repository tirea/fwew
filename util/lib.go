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

// Package util handles general program stuff. lib.go handles common functions.
package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"unicode"
)

// Contains returns true if anything in q is also in s
func Contains(s []string, q []string) bool {
	if len(q) == 0 || len(s) == 0 {
		return false
	}
	for _, x := range q {
		for _, y := range s {
			if y == x {
				return true
			}
		}
	}
	return false
}

// ContainsStr returns true if q is in s
func ContainsStr(s []string, q string) bool {
	if len(q) == 0 || len(s) == 0 {
		return false
	}
	for _, x := range s {
		if q == x {
			return true
		}
	}
	return false
}

// DeleteEmpty "deletes" all empty string entries in s
// actually returns a new string slice containing all non-empty strings in s
func DeleteEmpty(s []string) []string {
	if len(s) == 0 {
		return s
	}
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// Index return the index of q in s
func Index(s []string, q string) int {
	for i, v := range s {
		if v == q {
			return i
		}
	}
	return -1
}

// IndexStr return the index of q in s
func IndexStr(s string, q rune) int {
	for i, v := range s {
		if v == q {
			return i
		}
	}
	return -1
}

// IsLetter returns true if s is an alphabet character or apostrophe
func IsLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || r == '\'' || r == '‘' {
			return true
		}
	}
	return false
}

// Reverse returns the reversed version of s
func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

// StripChars strips all the characters in chr out of str
func StripChars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

// Valid validates range of integers for input
func Valid(input int64, reverse bool) bool {
	const (
		maxIntDec int64 = 32767
		maxIntOct int64 = 77777
	)
	if reverse {
		if 0 <= input && input <= maxIntDec {
			return true
		}
		return false
	}
	if 0 <= input && input <= maxIntOct {
		return true
	}
	return false
}

// DownloadDict downloads the latest released version of the dictionary file
func DownloadDict() error {
	var (
		filepath = Text("dictionary")
		url      = Text("dictURL")
	)
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(Text("dlSuccess"))
	return nil
}
