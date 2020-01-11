# TODO

### BUGS

- fix /random n to first do all the where filtering before truncating 
  to n number of words;
  allows /random n where a b c and d e f [and g h i...]
  to actually yield results
- fix /list to give actually specified number of words even after filtering

### Testing

- more code coverage
- more cases

### FEATURES

- Add proper nouns to dict
- add all affixes and their definitions to the dictionary file
- implement /list tag [is|has] tag
  modal, ofp, nfp, uncountable, irregular infix, body part, family,
  waterfalls, senses, si-verbs
- implement /list prefixes [pos] [pro|unpro|all]
- implement /list infixes [pos] [pro|unpro|all]
- implement /list suffixes [pos] [pro|unpro|all]
- implement /lenition ( /len )
- merge fr, hu, est into dict

### Refactoring & Optimization

- how can affixes.go be more efficient?

**Always remember to update util/version.go if major/minor changes are made to what it produces or how it works**
