package main

import "testing"

func TestPrefix(t *testing.T) {
	w0 := Word{
		Target:       "tsatute",
		Attempt:      "tute",
		PartOfSpeech: "n.",
		Affixes:      map[string][]string{},
	}
	w0 = prefix(w0)
	if w0.Attempt != w0.Target {
		t.Errorf("Got %s, Want %s", w0.Attempt, w0.Target)
	}
}

func TestSuffix(t *testing.T) {
	w0 := Word{
		Target:       "eltuti",
		Attempt:      "eltu",
		PartOfSpeech: "n.",
		Affixes:      map[string][]string{},
	}
	w0 = suffix(w0)
	if w0.Attempt != w0.Target {
		t.Errorf("Got %s, Want %s", w0.Attempt, w0.Target)
	}
}

func TestInfix(t *testing.T) {
	w0 := Word{
		Target:         "täpeykiyevarängon",
		Attempt:        "taron",
		InfixLocations: "t<0><1>ar<2>on",
		PartOfSpeech:   "vtr.",
		Affixes:        map[string][]string{},
	}
	w0 = infix(w0)
	if w0.Attempt != w0.Target {
		t.Errorf("Got %s, Want %s", w0.Attempt, w0.Target)
	}
}

func TestReconstruct(t *testing.T) {
	w0 := Word{
		Target:       "tolaron",
		Attempt:      "taron",
		PartOfSpeech: "vtr.",
		Affixes:      map[string][]string{},
	}
	w0 = Reconstruct(w0)
	if w0.Attempt != w0.Target {
		t.Errorf("Got %s, Want %s", w0.Attempt, w0.Target)
	}
}
