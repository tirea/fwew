# fwew
A better Crossplatform Na'vi Dictionary Terminal App

Fwew is the improved, clean, more modern successor to vrrtepcli (github.com/tirea/vrrtepcli). Fwew was written using the Go Programming Language and is a cross-platform binary text-based program for dictionary lookups. Primarily for the Na'vi language (learnnavi.org)

Installation
-----

### Linux Installation (Mac and Windows coming soon, switch to the windows or mac-osx branch for those OS)
Simply run the install script:

	./install.sh
	
Usage
-----

You'll need a new Terminal open first.
use cd command to change to the folder where the fwew program is downloaded

typical example:

	cd Downloads/fwew/bin

Linux users will need to type ./ on the beginning of all the following example commands in this README.
Unless they have already installed using the install.sh script.

### Search a Na'vi word
Simply run fwew with a word or list of words to look up:

	fwew tirea
	fwew oe tirea lu

Don't forget to escape ' somehow:

	fwew \'a\'aw
	fwew "'a'aw"

Ways to search a "____" si verb:

	s="eltur tìtxen si"; echo "$s" | fwew
	echo "eltur tìtxen si" | fwew

There is also an interactive mode, in case you forgot to put words as args:
	
	fwew

### Search an English word
Run fwew with the -r flag to reverse the lookup direction:

	fwew -r test
	fwew -r=true

### Use a language other than English
Run fwew with the -l flag to specify the language:

	fwew -l de
	fwew -l="pl"
	fwew -l=sv

### Displaying IPA and Infix location data
Use flags -ipa and -i respectively:

	fwew -ipa tireapängkxo
	fwew -i taron
	fwew -ipa -i plltxe
	fwew -i -ipa käteng

### Debug Mode 
This is helpful if you encounter an error. 
This option allows the user to send the developer (me) detailed information about what the program is doing to make it easier to fix.

	fwew -DEBUG -r test
	fwew -DEBUG -r test > debugfile.txt
	fwew -DEBUG > debugfile.txt

Examples
-----

## Typical usage
look up the Na'vi word for "brain" (from English)

	fwew -r brain

look up the English word for yayotsrul

	fwew yayotsrul

look up English translation, IPA, and infixes for yemstokx

	fwew -ipa -i yemstokx

### Use German, Show infixes and IPA, Reverse lookup, look up essen
	fwew -l de -i -ipa -r essen
    or:
	fwew -l=de -ipa -i -r essen

### Use Polish, show IPA for kemlì'u
	fwew -l=pl -ipa kemlì\'u
    or:
	fwew -ipa -l pl "kemlì'u"
