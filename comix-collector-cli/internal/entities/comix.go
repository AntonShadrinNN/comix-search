package entities

import "fmt"

// A ComixData represents comic data with a URL and keywords.
type ComixData struct {
	Url      string   `json:"url"`      // represents link to a comic
	Keywords []string `json:"keywords"` // represents keywords from this comic
}

// Representation of a comix from xkcd
type Comix struct {
	Id         int    `json:"num"`
	ImgUrl     string `json:"img"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
}

// The same as ComixData but with id
type ComixEntry struct {
	Id int
	Cd ComixData
}

// Pretty-print for comix-entry
func (ce ComixEntry) String() string {
	return fmt.Sprintf("%d: {\n%s\n}", ce.Id, ce.Cd)
}
