# TODO

### BUGS

- fix /list to give actually specified number of words even after filtering  
Repro:  
`/list words last 10 and pos not-is meme.`

### FEATURES

- add all affixes and their definitions to the dictionary file
- fix tsa-{adp.}, tsa-{case}
- implement `/list tag [is|has] tag`  
  Tag ideas:  
  modal, ofp, nfp, uncountable, ii, body, family,  
  waterfall, sense, si-verbs, F-word, Tsa-word,  
  emotion, lenition, color, flora, fauna, linguistics,  
  loan, Disney, time, weather, <eyk>, music...
- implement `/list prefixes {pos} {pro|unpro|all}`
- implement `/prefixes`
- implement `/list infixes {pos} {pro|unpro|all}`
- implement `/infixes`
- implement `/list suffixes {pos} {pro|unpro|all}`
- implement `/suffixes`
- implement `/lenition` (with alias: `/len` )
- merge fr, hu, est into dict

### IDEAS

- `/help <command>`
- `/list <what> <cond> <spec> limit <n>`
- Stress underlining rather than IPA default
- `/examples <word> [limit <n>]`
- `/define <jargony linguistics term>`
- SQLite DB?
- `/audio <Na'vi word(s)>`

### Testing

- more code coverage
- more test cases

### Refactoring & Optimization

- how can affixes.go be more efficient?

**Always remember to update util/version.go if major/minor changes are made 
to what it produces or how it works**
