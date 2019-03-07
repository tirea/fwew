# fwew
[![Build Status](https://travis-ci.com/tirea/fwew.svg?branch=master)](https://travis-ci.com/tirea/fwew) 
[![License: GPL v2](https://img.shields.io/badge/License-GPL%20v2-blue.svg)](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html)

The Best Na'vi Dictionary on the Command Line

Fwew is written using the Go Programming Language and is a cross-platform binary text-based program for dictionary lookups. 
Primarily for the [Na'vi language](https://learnnavi.org). 
See the [LearnNavi Forum thread](https://forum.learnnavi.org/projects/fwew-a-better-crossplatform-navi-dictionary-terminal-app/).

## Install

### Compile and install from source code

This option is mostly for Contributors and Developers. Or people who like to compile stuff themselves.

You will need the [GO Programming Language](https://golang.org/) and [Git](https://git-scm.com/) installed. 
If you don't have these and don't want to download/install them, skip to one of the following sections.

#### Linux or macOS

Run the following commands from inside a Terminal:

```
cd $HOME                                 # Start at home folder
mkdir -p go/src                          # Make a folder for all Go source code
export GOPATH=$HOME/go                   # Set GOPATH variable to newly created go folder
```
You may want to add that last line in your shell config file (~/.bashrc, ~/.zshrc, or similar)
```
cd go/src/                               # This is where we will download the fwew source code
go get github.com/c-bata/go-prompt       # Install the go-prompt library (fwew depends on this library)
git clone https://github.com/tirea/fwew  # Download the code
cd fwew                                  # Go to where the code is before trying to build it
make                                     # to just compile
make install                             # to compile and install
```

#### Windows

Run the following from inside a Powershell:

```
cd $HOME                                # Start at home folder
mkdir go\src                            # Make a folder for all Go source code
setx GOPATH %USERPROFILE%\go            # Set GOPATH variable to newly created go folder
```
Open a new Powershell so that last line takes effect, then continue:
```
cd go\src\                              # This is where we will download the fwew source code
go get github.com/c-bata/go-prompt      # Install the go-prompt library (fwew depends on this library)
git clone https://github.com/tirea/fwew # Download the code
cd fwew                                 # Go to where the code is before trying to build it
go build -o $HOME\fwew.exe .\fwew.go    # compile
cp -Recurse .\.fwew $HOME\              # copy data file folder to your user's home folder
```

### Docker

If you have [Docker](https://www.docker.com) installed, you can either bulid an image using the Dockerfile in this
project, or you can pull the latest pre-built one from dockerhub.

#### install via building docker container from Dockerfile

It's as easy as 1, 2, 3 (if you have docker installed and running):

1. `git clone https://github.com/tirea/fwew`
2. `cd fwew`
3. `make docker`

#### Install via docker container from dockerhub

Even easier: a minimized container of fwew is available on dockerhub:

Pull the container:
```
docker pull tirea/fwew:latest
```

Run the containier:
```
docker run -it --rm tirea/fwew
```

Type `/q` to quit running the program. See the **REPL** and **/commands** sections below for more information.

Command line arguments work as expected; just append your arguments to the end of the `docker run` statement above. 
See the **Command Line Arguments & Flags** section below for more about fwew options and argumets.

### Install program from downloaded .zip

If you don't have Go or Git or Docker installed, you don't need to. You can just download the pre-built program here 
from GitHub in a .zip file then install and/or run it, without compiling it yourself.

**Note**: these binaries are not typically kept up to date with every release.

Windows/MacOS/Linux:

- Download the [master.zip](https://github.com/tirea/fwew/archive/master.zip) file
- Extract the files
- Copy the `.fwew` folder into your user's home folder

Linux/MacOS ONLY:

- Depending on your OS, copy the `bin/linux/fwew` or `bin/mac/fwew` file to your user's home folder
- Add this text your shell config file
(`~/.bashrc` or `~/.profile` or `~/.zshrc` or whatever):
`export PATH=$PATH:$HOME`

Windows ONLY:

- Copy the `bin\windows\fwew.exe` file to your user's home folder

## Uninstall

### Using Makefile

If you're on Linux/MacOS and did `Compile and install from source code` and want to now uninstall Fwew:

From Terminal where `tirea/fwew/Makefile` is, run:
```
make uninstall
```

### Otherwise

- Remove/delete the `fwew` or `fwew.exe` binary from wherever you put it or installed it to
- Remove/delete the `.fwew/` folder from your home folder

## Command Line Arguments & Flags

### Search Na'vi Word(s) Using CLI Args

Run fwew with a word or list of words to look up:

```
fwew tirea
fwew oe tirea lu
```

Don't forget to escape apostrophe `'` by either using `\` before each `'`, or surrounding whole word with quotes:

```
fwew \'a\'aw
fwew "'a'aw"
```

Search a `"__ si"` verb by enclosing all parts of the word within quotes:

```
fwew "eltur tìtxen si"
fwew "tìkangkem si"
```

### Affix parsing

Fwew parses and displays affixes used to construct the input word by default.

Users familiar with the language can disable this feature and make fwew runtime faster in two ways 
(Note that this means that only root words can be searched.):

Use the `-a=false` flag

```
fwew -a=false taron
fwew -a=false
```

Or set `useAffixes` to false in the config file. (See Configuration file section at the end of this README)

### Search an English Word

Run fwew with the `-r` flag to reverse the lookup direction:

```
fwew -r test
fwew -r=true test
```

### Use a language other than English

Run fwew with the `-l` flag to specify the language:

```
fwew -l de "lì'fya"
fwew -l=sv lì\'fya
```

### Displaying IPA and Infix Location Data

Use flags `-ipa` and `-i` respectively:

```
fwew -ipa tireapängkxo
fwew -i taron
fwew -ipa -i plltxe
fwew -i -ipa käteng
```

### Filter Words by Part of Speech

Use `-p` flag followed by the part of speech abbreviation as found in any Na'vi dictionary. 
Most useful in `-r=true` (reverse lookup) mode to narrow down results when many are returned.

```
fwew -r -p adp. in
fwew -r -p=vtr. test
```

### Display Dictionary Version

```
fwew -v
fwew -v -r word
```

### Set and Unset Flags

You can search even quicker without re-running the program to update what information you want to see. 
Use the set[] and unset[] keywords to update the search options. Even on the command line! To set or unset 
multiple options at once, separate them with a comma. Language and Part of Speech Filter cannot be unset, just 
set to another value. The default values are l=eng and p=all

```
fwew -r -ipa test unset[r,ipa] wou set[l=de,i,ipa] taron
fwew fmetok set[i] omum unset[i] set[r,l=sv] hej
```

## REPL

There is also an interactive mode, activated when no words are present in the command line arguments: 
All flags are set to default values: -v=false -r=false -l=eng -i=false -ipa=false -p="all", unless overridden on 
the command line. Fwew will continuously prompt you for input.

```
fwew
fwew -i -ipa
```

setting options also works in the REPL/Interactive mode. Here however, it's a slash-command, `/set`. 
flags to be set are separated by space. `/unset` is actually just an alias to `/set`, since what happens here is 
options are toggled on/off. One command per line, and only the command on the line.

Sample Output of `fwew -i -ipa`:
```
~~> eltu
[1] eltu [ˈɛl.tu] n. brain

~~> /unset ipa
set [ i a l=eng p=all ]

~~> /set l=de r
set [ r i a l=de p=all ]

~~> wald
[1] na'rìng n. Wald

~~> /set i ipa p=vtr.
set [ r ipa a l=de p=vtr. ]

~~> essen
[1] yom [j·om] vtr. essen, speisen, fressen

```

Use `/set` with empty list of flags to show all current set flag values.

```
~~> /set
set [ a l=eng p=all ]

~~> /set i ipa
set [ i ipa a l=eng p=all ]

~~> /set
set [ i ipa a l=eng p=all ]
```

### /commands 

While in interactive mode, the following commands are available and can be seen by running the `/commands` command:
```
/set       show currently set options, or set given options (separated by space)
/unset     unset given options (separated by space)
/list      list all words that meet given criteria
/random    display given number of random entries
/update    download and update the dictionary file
/commands  show this commands help text
/help      show main help text
/exit      exit/quit the program (aliases /quit /q /wc)
```

Note that as of 3.0.0-dev, Tab-completion is possible by typing a command partially then pressing tab to complete it, 
and additionally pressing tab as nessary to complete the command you are looking for.
Note also that as of 3.0.0-dev, history is recorded each time you enter a command. Use the up and down arrows to cycle 
through your history and reuse a previous command or modify a previous command before running again.

### /set and /unset

`/set` and `/unset` allow options to be toggled or set while within `fwew`. Here are the available options:

```
a       use affix-recognition
i       display infix location data
ipa     display IPA data
l=de    use German language
l=eng   use English language
l=est   use Estonian language
l=hu    use Hungarian language
l=nl    use Dutch language
l=pl    use Polish language
l=ru    use Russian language
l=sv    use Swedish language
m       format output in markdown for bold and italic (mostly only useful for fwew-discord bot)
n       convert numbers octal<->decimal
p=[pos] search for word(s) with specified part of speech abbreviation
r       reverse the lookup direction from Na'vi->local to local->Na'vi
s       display source data

```

use any of these you wish to set, and separate them with spaces

```
/set ipa i l=eng p=all r
/unset i r
```

### /list

`/list` is a powerful search feature of `fwew` that allows you to list all of the words that satisfy a set of given conditions.

The syntax is as follows (cond is short for condition, spec is short for specification):

```
/list what cond spec
/lits what cond spec and what cond spec [and what cond spec...]
```

`what` can be any one of the following:

```
pos          part of speech of na'vi word
word         na'vi word
syllables    number of syllables in the na'vi word
words        selection of na'vi words
```

`cond` depends on the `what`. Here are the conditions that apply to each `what`:

pos:
```
has    part of speech has the following character sequence anywhere
is     part of speech is exactly the following character sequence
```
word:
```
starts    word starts with the following character sequence
ends      word ends with the following character sequence
has       word has the following character sequence anywhere
```
syllables:
```
<     less than the following number
<=    less than or equal to the following number
=     exactly equal to the following number
>=    greater than or equal to the following number
>     greater than the following number
```
words:
```
first    the first consecutive words in the datafile (chronologically oldest words) 
last     the last consecutive words in the datafile (chronologically newest words)
```

`spec` depends on the `cond`. Here are the specifications that apply to each `cond`:

`has`, `is`, `starts`, and `ends` all expect a character sequence to come next.

`<`, `<=`, `=`, `>=`, `>`, `first`, and `last` all expect a number to come next.

#### Examples of /list

List all modal verbs:
```
/list pos has v and pos has m.
```
List all stative verbs:
```
/list pos has svin.
```
List all nouns that start with tì:
```
/list word starts tì and pos is n.
```
List all 3 syllable transitive verbs:
```
/list syllables = 3 and pos has vtr.
```
List the newest 25 words in the language:
```
/list words last 25
```

### /random

`/random` is a random entry generator that generates the a given number (or random number!) of random entries. 
It also features a `where` clause in which the `what cond spec` syntax from `/list` is supported to narrow down what kinds of random entries you get.

#### Examples of /random

List 10 random entires
```
/random 10
```
List 5 random transitive verbs
```
/random 5 where pos has vtr
```
List a random number of random words
```
/random random
```
List a random number of nouns
```
/random random where pos is n.
```

### /update

Running `/update` will download and update the fwew dictionary file.

### /commands and /help

`/commands` shows the list of commands and examples
`/help` shows the program options / command line flags

### /exit /quit /q

Any of these will quit the program from within.

## Input & Output Files

You can make a text file containing all the words you want to search and all the flag settings. Each thing on its own line.

input.txt:

```
eltu
/set r p=adp.
on
/unset r
/set p=all
prrkxentrrkrr
/set l=sv
tìfmetok
nitram
/set i ipa
taron
omum
inan
/unset i ipa
```

pass this file to fwew:

```
fwew -f input.txt
```

Fwew output:

```
cmd eltu
[1] eltu n. brain

cmd /set r p=adp.
cmd on
[1] mì+ adp. in, on
[2] sìn adp. on, onto

cmd /unset r
cmd /set p=all
cmd prrkxentrrkrr
[1] prrkxentrrkrr n. day time smoking pleasure, vibrating tongue (Na'vi idiom)

cmd /set l=sv
cmd tìfmetok
[1] tìfmetok n. test

cmd nitram
[1] nitram adj. lycklig, glad (om folk)

cmd /set i ipa
cmd taron
[1] taron [ˈt·a.ɾ·on] t<0><1>ar<2>on vtr. jaga

cmd omum
[1] omum [·o.ˈm·um] <0><1>om<2>um vtr. veta, känna till

cmd inan
[1] inan [·i.ˈn·an] <0><1>in<2>an vtr. läsa (tex. skogen), få kunskap ifrån sinnesintryck

cmd /unset i ipa
```

You can also direct the output of Fwew into a new text file.

```
fwew -f input.txt > output.txt
```

## Configuration file

Settings for Fwew are stored in a plain-text JSON file in the `.fwew/` directory.

`config.json`:

```
{
    "language": "eng",
    "posFilter": "all",
    "useAffixes": true,
    "DebugMode": false
}
```

The default language to use when looking up words is `"eng"` and can be changed here. This is useful for people who 
don't want to hake to keep typing all the time this, for example:

```
fwew -l de
fwew -l=de
```

The default part of speech filter is `"all"` and can be changed here. This is useful for people who want to repeatedly 
run fwew searching for words of all the same part of speech. It avoids repeatedly typing, for example:

```
fwew -p n.
fwew -p vtr.
```

If you're familiar with the language and only ever need to search root words, you can set `"useAffixes"` to `false`, 
speeding up the program runtime by not trying to break down words to find results. This avoids repeatedly typing, for example:

```
fwew -a=false taron
fwew -a=false
```

The default value of DebugMode is `false` and can be changed here. DebugMode being set to `true` will cause a monstrous 
mountain of text to flood your Terminal or Powershell on every `fwew` run. The point of it all is to see where something 
went wrong in the logic. This option is mostly only useful to Contributors, Developers, and Users who want to report a bug. 
The `-debug` command line flag was removed in favor of having this option in the config file.

If you edit the config file to set your own defaults, you can override the config file settings using the set[] command keyword as shown above.
