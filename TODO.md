# TODO

### BUGS

- fix /random n to first do all the where filtering before truncating  
Repro:  
`/random 5 where pos is n. and word starts l`
- fix /list to give actually specified number of words even after filtering  
Repro:  
`/list words last 10 and pos not-is meme.`

### FEATURES

- add all forms of pesu from Horen 3.4.1.1 to the dictionary file
- add all affixes and their definitions to the dictionary file
- implement `/list tag [is|has] tag`  
  Tag ideas:  
  modal, ofp, nfp, uncountable, ii, body, family,  
  waterfall, sense, si-verbs, F-word, Tsa-word,  
  emotion, lenition, color, flora, fauna, linguistics,  
  loan, Disney, time, weather, music...
- implement `/list prefixes {pos} {pro|unpro|all}`
- implement `/prefixes`
- implement `/list infixes {pos} {pro|unpro|all}`
- implement `/infixes`
- implement `/list suffixes {pos} {pro|unpro|all}`
- implement `/suffixes`
- implement `/lenition` (with alias: `/len` )
- merge fr, hu, est into dict

### Testing

- more code coverage
- more test cases

### Refactoring & Optimization

- how can affixes.go be more efficient?

**Always remember to update util/version.go if major/minor changes are made 
to what it produces or how it works**
