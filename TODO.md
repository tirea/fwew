# TODO

## Road Map
- Fix bugs in affixes.go
- Experiment with sqlite3 database (challenge: affix recognition without additional data)
- Data file Updates (auto?)

## Affixes Algorithm

### Current Reconstruct() Process
- For each line in dictionary.txt do if nav field not input:
- infix(w) if verb
- prefix(w)
- lenite(w) if w.Attempt and w.Target start with different char
- suffix(w)
- lenite(w)
- set w.ID = -1 and give up

### Current General Affix Function Process
- Compile a Regexp based on w.Attempt and the affixes matching for
- Match the Regexp against w.Target
- Based on above, determine which affixes appear to be used in w.Target
- Add the above affixes to w.Attempt
- Add the array of above to w.Affixes
- return w

### New Reconstruct() Process
- TODO

### New General Affix Function Process
- TODO

## Always remember to update util/version.go if major/minor changes are made to what it produces or how it works
