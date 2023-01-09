package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

//A type shld have a primary responsiblity and a reason to change

// Journal primary responsbility store entries
type Journal struct {
	entries []string
}

var entryCount = 0

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// Specific to journa;
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

//Seperation of concerns,
//Persistance, seperation of concerns (can be influenced other type of structs are saved), have common settings basically resuable component sort of
//Take out the concerns and save it somewhere so we can control, this place in a single setting

// Persist different line seperator for Journal
var LineSeperator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeperator)), 0644)
}

func (j *Journal) LoadFromWeb(url *url.URL) {}

func main() {
	J1 := Journal{}

	J1.AddEntry("I am sleepy")
	J1.AddEntry("I am tired")

	fmt.Println(J1.String())
	//Persistance
	SaveToFile(&J1, "./Solid-Design-Principles/journal.txt")

}
