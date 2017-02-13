# fwew

The Best Na'vi Dictionary on the Command Line

Fwew is the improved, faster, cleaner, more modern successor to vrrtepcli (github.com/tirea/vrrtepcli). Fwew was written using the Go Programming Language and is a cross-platform binary text-based program for dictionary lookups. Primarily for the Na'vi language (learnnavi.org)

## Install

### Compile and Install from Source

This option is mostly for Contributors and Developers.

You will need the [GO Programming Language](https://golang.org/) installed. If you don't have it or don't want to download it, see the next section.

```
cd $GOPATH                   # GOPATH should be set in ~/.bashrc as something like export GOPATH=$HOME/go
go get github.com/tirea/fwew # Pretty much same as git clone but puts stuff where it needs to be
make                         # to just compile
make install                 # to compile and install
```

### Install directly from .zip

If you don't have Go installed, you don't need it. You can just download the program here from GitHub in a .zip file then install it, without compiling.

- Download the [master.zip](https://github.com/tirea/fwew/archive/master.zip) file
- Extract the files
- Copy the bin/fwew binary to somewhere in your $PATH (the Makefile uses /usr/local/bin/)
- Copy the .fwew/ folder into your $HOME folder

## Un-install

### Using Makefile

```
make uninstall
```

### Otherwise

- Remove/delete the fwew binary from /usr/local/bin/ or wherever you put it
- Remove/delete the .fwew/ folder from your $HOME folder

## Command Line Arguments & Flags

### Search Na'vi Word(s) Using CLI Args

Run fwew with a word or list of words to look up:

```
fwew tirea
fwew oe tirea lu
```

Don't forget to escape ' somehow:

```
fwew \'a\'aw
fwew "'a'aw"
```

How to search a "__ si" verb:

```
echo "eltur tìtxen si" | fwew
echo "tìkangkem si" | fwew
```

### Search an English Word

Run fwew with the -r flag to reverse the lookup direction:

```
fwew -r test
fwew -r=true test
```

### Use a language other than English

Run fwew with the -l flag to specify the language:

```
fwew -l de "lì'fya"
fwew -l=sv lì\'fya
```

### Displaying IPA and Infix Location Data

Use flags -ipa and -i respectively:

```
fwew -ipa tireapängkxo
fwew -i taron
fwew -ipa -i plltxe
fwew -i -ipa käteng
```

### Filter Words by Part of Speech

Use -p flag followed by the part of speech abbreviation as found in any Na'vi dictionary. Most useful in -r=true (reverse lookup) mode to narrow down results when many are returned.

```
fwew -r in -p adp
fwew -r test -p=vtr
```

### Display Dictionary Version

```
fwew -v
fwew -v -r word
```

### Debug Mode

This is helpful if you encounter an error. This option allows the user to send the developer (me) detailed information about what the program is doing to make it easier to fix.

```
fwew -debug -r test
fwew -debug -r test > debugfile.txt
fwew -debug > debugfile.txt
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
$ fwew -i -ipa
Fwew 1.3-BETA (01 FEB 2017) by Tirea Aean
Crossplatform Na'vi Dictionary Search
fwew -h for usage, see README

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

Fwew:> ^C
```

## Input & Output Files

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

Settings for Fwew are stored in a plain-text JSON file in the .fwew/ directory.

config.json:

```
{
    "language": "eng",
    "posFilter": "all"
}
```

The default language to use when looking up words is "eng" and can be changed here. This is useful for people who don't want to hake to keep typing all the time this, for example:

```
fwew -l de
fwew -l=de
```

The default part of speech filter is "all" and can be changed here. This is useful for people who want to repeatedly run fwew searching for words of all the same part of speech. It avoids repeatedly typing, for example:

```
fwew -p n
fwew -p vtr
```

If you edit the config file and set your own defaults, you can override the config file settings using the set[] command keyword as shown above.
