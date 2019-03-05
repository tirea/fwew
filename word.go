package main

import "fmt"

// Word is a struct that contains all the data about a given word
type Word struct {
	ID             string
	LangCode       string
	Navi           string
	Target         string
	Attempt        string
	IPA            string
	InfixLocations string
	PartOfSpeech   string
	Definition     string
	Source         string
	Affixes        map[string][]string
}

func (w Word) String() string {
	// this string only doesn't get translated or called from util.Text() because they're var names
	return fmt.Sprintf("Id: %s\nLangCode: %s\nNavi: %s\nTarget: %s\nAttempt: %s\nIPA: %s\nInfixLocations: %s\nPartOfSpeech: %s\nDefinition: %s\nAffixes: %v\n",
		w.ID, w.LangCode, w.Navi, w.Target, w.Attempt, w.IPA, w.InfixLocations, w.PartOfSpeech, w.Definition, w.Affixes)
}

// InitWordStruct is basically a constructer for Word struct
func InitWordStruct(w Word, dataFields []string) Word {
	//const (
	//	idField  int = 0 // dictionary.txt line Field 0 is Database ID
	//	lcField  int = 1 // dictionary.txt line field 1 is Language Code
	//	navField int = 2 // dictionary.txt line field 2 is Na'vi word
	//	ipaField int = 3 // dictionary.txt line field 3 is IPA data
	//	infField int = 4 // dictionary.txt line field 4 is Infix location data
	//	posField int = 5 // dictionary.txt line field 5 is Part of Speech data
	//	defField int = 6 // dictionary.txt line field 6 is Local definition
	//	srcField int = 7 // dictionary.txt line field 7 is Source data
	//)
	w.ID = dataFields[idField]
	w.LangCode = dataFields[lcField]
	w.Navi = dataFields[navField]
	w.IPA = dataFields[ipaField]
	w.InfixLocations = dataFields[infField]
	w.PartOfSpeech = dataFields[posField]
	w.Definition = dataFields[defField]
	w.Source = dataFields[srcField]
	w.Affixes = map[string][]string{}

	return w
}

// CloneWordStruct is basically a copy constructor for Word struct
func CloneWordStruct(w Word) Word {
	var nw Word
	nw.ID = w.ID
	nw.LangCode = w.LangCode
	nw.Navi = w.Navi
	nw.Target = w.Target
	nw.Attempt = w.Attempt
	nw.IPA = w.IPA
	nw.InfixLocations = w.InfixLocations
	nw.PartOfSpeech = w.PartOfSpeech
	nw.Definition = w.Definition
	nw.Source = w.Source
	nw.Affixes = make(map[string][]string)
	for k := range w.Affixes {
		nw.Affixes[k] = make([]string, len(w.Affixes[k]))
		copy(nw.Affixes[k], w.Affixes[k])
	}
	return nw
}
