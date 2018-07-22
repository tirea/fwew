# fwew

The Best Na'vi Dictionary on the Command Line

Fwew is the improved, faster, cleaner, more modern successor to [vrrtepcli](https://github.com/tirea/vrrtepcli). Fwew was written using the Go Programming Language and is a cross-platform binary text-based program for dictionary lookups. Primarily for the [Na'vi language](https://learnnavi.org). See the [LearnNavi Forum thread](https://forum.learnnavi.org/projects/fwew-a-better-crossplatform-navi-dictionary-terminal-app/).

## Install

### Compile and install from source code

This option is mostly for Contributors and Developers.

You will need the [GO Programming Language](https://golang.org/) and [Git](https://git-scm.com/) installed. If you don't have these and don't want to download/install them, see the next section, `Install program from downloaded .zip`.

Run the following commands from inside a Terminal (Linux, MacOS ONLY)

```
cd $HOME                        # Start at home folder
mkdir -p go                     # Make a folder for all Go source code
export GOPATH=$HOME/go          # Set GOPATH variable to newly created go folder
go get github.com/tirea/fwew    # Pretty much same as git clone but puts stuff where it needs to be
cd go/src/github.com/tirea/fwew # Go to where the code is before trying to build it
make                            # to just compile
make install                    # to compile and install
```

For Windows users, instead run the following from inside a Powershell:

```
cd $HOME                             # Start at home folder
go get github.com/tirea/fwew         # Pretty much same as git clone but puts stuff where it needs to be
cd .\go\src\github.com\tirea\fwew    # Go to where the code is before trying to build it
go build -o .\bin\fwew.exe .\fwew.go # compile
cp -Recurse .\.fwew $HOME\           # copy data file folder to your user's home folder
```

On Windows, you will want to add the path to `fwew.exe` to your `Path` Environment Variable:

- On the taskbar where it says "Type here to search", type `Path` and press Enter
- In the Environment Variables window that opens, highlight the row that says `Path` in the top part
- Click `Edit...` button
- Click `New` button
- Type or paste in the text field, `%GOPATH%\src\github.com\tirea\fwew\bin`
- Click `OK` button, `OK` button again

### Install program from downloaded .zip

If you don't have Go or Git installed, you don't need to. You can just download the pre-built program here from GitHub in a .zip file then install it, without compiling it yourself.

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

Users familiar with the language can disable this feature and make fwew runtime faster in two ways (Note that this means that only root words can be searched.):

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

Use `-p` flag followed by the part of speech abbreviation as found in any Na'vi dictionary. Most useful in `-r=true` (reverse lookup) mode to narrow down results when many are returned.

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

You can search even quicker without re-running the program to update what information you want to see. Use the set[] and unset[] keywords to update the search options. Even on the command line! To set or unset multiple options at once, separate them with a comma. Language and Part of Speech Filter cannot be unset, just set to another value. The default values are l=eng and p=all

```
fwew -r -ipa test unset[r,ipa] wou set[l=de,i,ipa] taron
fwew fmetok set[i] omum unset[i] set[r,l=sv] hej
```

## REPL

There is also an interactive mode, activated when no words are present in the command line arguments: All flags are set to default values: -v=false -r=false -l=eng -i=false -ipa=false -p="all" The Read-Evaluate-Print-Loop will forever ask you for input until you either type Ctrl+C or Ctrl+D.

```
fwew
fwew -i -ipa
```

set[] and unset[] commands also work in the REPL. One command per line, and only the command on the line.

```
fwew -i -ipa
Fwew - Na'vi Dictionary Search - by Tirea Aean
`fwew -h` for usage, `fwew -v` for version. See README

Fwew:> eltu
[1] n. brain [ˈɛl.tu] (eltu)

Fwew:> unset[ipa]
<! [ipa] unset >

Fwew:> set[l=de,r]
<! [l=de r] set >

Fwew:> wald
[1] n. na'rìng (wald)

Fwew:> set[i,ipa,p=vtr]
<! [i ipa p=vtr] set >

Fwew:> essen
[1] vtr. yom [j·om] y<1><2><3>om (essen, speisen, fressen)

Fwew:>
```

Use set[] with empty list of flags to show all current set flag values.

```
fwew
Fwew - Na'vi Dictionary Search - by Tirea Aean
`fwew -h` for usage, `fwew -v` for version. See README

Fwew:> set[]
<! Currently set: debug=false, r=false, i=false, ipa=false, l=eng, p=all >

Fwew:> set[debug,i,ipa]
<! [debug i ipa] set >

Fwew:> set[]
<! Currently set: debug=true, r=false, i=true, ipa=true, l=eng, p=all >

Fwew:>
```

## Input & Output Files (Linux / MacOS)

You can make a text file containing all the words you want to search and all the flag settings.

input.txt:

```
eltu
set[r,p=adp]
on
unset[r]
set[p=all]
prrkxentrrkrr
set[l=sv]
tìfmetok nitram
set[i,ipa]
taron omum inan
unset[i,ipa]
```

pass this file to fwew:

```
fwew `cat input.txt`
```

Fwew output:

```
[1] n. brain (eltu)

<! [r p=adp] set >

[1] adp. mì (in, on)

[2] adp. sìn (on, onto)

<! [r] unset >

<! [p=all] set >

[1] n. day time smoking pleasure, vibrating tongue (na'vi idiom) (prrkxentrrkrr)

<! [l=sv] set >

[1] n. test (tìfmetok)

[1] adj. lycklig, glad (om folk) (nitram)

<! [i ipa] set >

[1] vtr. jaga [ˈt·a.ɾ·on] t<1><2>ar<3>on (taron)

[1] vtr. veta, känna till [·o.ˈm·um] <1><2>om<3>um (omum)

[1] vtr. läsa (tex. skogen), få kunskap ifrån sinnesintryck [·i.ˈn·an] <1><2>in<3>an (inan)

<! [i ipa] unset >
```

You can also direct the output of Fwew into a new text file.

```
fwew `cat input.txt` > output.txt
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

The default language to use when looking up words is `"eng"` and can be changed here. This is useful for people who don't want to hake to keep typing all the time this, for example:

```
fwew -l de
fwew -l=de
```

The default part of speech filter is `"all"` and can be changed here. This is useful for people who want to repeatedly run fwew searching for words of all the same part of speech. It avoids repeatedly typing, for example:

```
fwew -p n.
fwew -p vtr.
```

If you're familiar with the language and only ever need to search root words, you can set `"useAffixes"` to `false`, speeding up the program runtime by not trying to break down words to find results. This avoids repeatedly typing, for example:

```
fwew -a=false taron
fwew -a=false
```

The default value of DebugMode is `false` and can be changed here. DebugMode being set to `true` will cause a monstrous mountain of text to flood your Terminal or Powershell on every `fwew` run. The point of it all is to see where something went wrong in the logic. This option is mostly only useful to Contributors, Developers, and Users who want to report a bug. The `-debug` command line flag was removed in favor of having this option in the config file.

If you edit the config file to set your own defaults, you can override the config file settings using the set[] command keyword as shown above.
