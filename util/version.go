package util

import "fmt"

type version struct {
	Major, Minor, Patch int
	Label               string
	Name                string
	Dict                string
}

var Version = version{1, 4, 0, "dev", "Eana Yayo", "Na'vi Dictionary 13.31 (07 JAN 2017)"}

var Build string

func (v version) String() string {
	if v.Label != "" {
		return fmt.Sprintf("Fwew version %d.%d.%d-%s \"%s\"\n%s\nGit commit hash: %s", v.Major, v.Minor, v.Patch, v.Label, v.Name, v.Dict, Build)
	} else {
		return fmt.Sprintf("Fwew version %d.%d.%d \"%s\"\n%s\nGit commit hash: %s", v.Major, v.Minor, v.Patch, v.Name, v.Dict, Build)
	}
}
