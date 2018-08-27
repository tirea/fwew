# TODO

## Road Map
- Fix (more) bugs in affixes.go
- Implement:
    - root adjustment on affixing:
        - zenke // e comes back from zene in cases <3> => ats, uy
    - merges:
        - pxen (pxe+ 'en, not as root word) // works as \*pxeen
        - pen (pe+ 'en, ay+pxen) // works as aypen, \*aypen
        - poltxe // works as \*pollltxe
        - molte // works as \*molllte
    - vowel changes:
        - ngeyä, peyä, tseyä
        - meylan, pxeylan, peylan, etc.

**Always remember to update util/version.go if major/minor changes are made to what it produces or how it works**
