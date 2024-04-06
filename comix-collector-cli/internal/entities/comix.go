package entities

import "fmt"

type ComixData struct {
	Url      string   `json:"url"`
	Keywords []string `json:"keywords"`
}

type Comix struct {
	Id         int    `json:"num"`
	ImgUrl     string `json:"img"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
}

type ComixEntry struct {
	Id int
	Cd ComixData
}

func (ce ComixEntry) String() string {
	return fmt.Sprintf("%d: {\n%s\n}", ce.Id, ce.Cd)
}
