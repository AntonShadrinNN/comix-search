package app

import (
	"strings"

	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/entities"
)

type App struct {
	comixRepo ComixDataRepo
	stemmer   Stemmer
	client    ComixClient
}

func NewApp(cr ComixDataRepo, s Stemmer, c ComixClient) App {
	return App{
		comixRepo: cr,
		stemmer:   s,
		client:    c,
	}
}

func (a App) Create(id int, cd entities.ComixData) error {
	return a.comixRepo.Create(id, cd)
}

func (a App) Read(id int) (entities.ComixData, error) {
	return a.comixRepo.Read(id)
}

func (a App) ReadN(n int) ([]entities.ComixEntry, error) {
	return a.comixRepo.ReadN(n)
}

func (a App) ReadAll() ([]entities.ComixEntry, error) {
	return a.comixRepo.ReadAll()
}

func (a App) Stem(sentence string) ([]string, error) {
	words := strings.Split(sentence, " ")
	return a.stemmer.Stem(words)
}

func (a App) GetComixesCount() (int, error) {
	return a.client.GetComixesCount()
}

func (a App) GetLastWrittenId() (int, error) {
	return a.comixRepo.GetLastWrittenId()
}

func (a App) FetchComixById(id int) (entities.Comix, error) {
	return a.client.FetchComixById(id)
}

func (a App) GetWrittenCount() (int, error) {
	return a.comixRepo.GetWrittenCount()
}
