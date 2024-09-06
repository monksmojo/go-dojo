### Practical 5

#### practical Five description

#### Practical Five:

---

Write a Go function `printNameCounts` that takes a slice of strings `names` and returns a map of maps. The outer map will be keyed by the first letter of each name in the input slice. The inner map will be keyed by the name itself and will contain the number of times that name appears in the input slice.

For example, if the input slice is `["Alice", "Archie", "Bob", "Bob", "Boyne", "Carol", "Charlie", "David", "Eve"]`, the function should return the following map:

Output:

```go
map[rune]map[string]int{
  "A": {"Alice": 1,"Archie": 1},
  "B": {"Bob": 2,"Boyne": 1},
  "C": {"Carol": 1,"Charlie": 1},
  "D": {"David": 1},
  "E": {"Eve": 1},
}
```

---

The function `PrintNameCounts` takes a list of names, counts the occurrences of each name, and
prints the result.

```go
func PrintNameCounts() {
	names := []string{"Alice", "Archie", "Bob", "Bob", "Boyne", "Carol", "Charlie", "David", "Eve"}
	namesDictionary := getNameCounts(names)
	fmt.Println(namesDictionary)
}
```
---

The function `getNameCounts` takes in a slice of strings and returns a map that counts the
occurrences of each name based on the first letter of the name.
```go
func getNameCounts(names []string) map[byte]map[string]int {
	namesDictionary := make(map[byte]map[string]int)
	for _, name := range names {
		_, ok := namesDictionary[name[0]]
		if !ok {
			namesDictionary[name[0]] = make(map[string]int)
		}
		namesDictionary[name[0]][name]++
	}
	return namesDictionary
}
```
---
