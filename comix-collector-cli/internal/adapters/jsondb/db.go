package jsondb

import (
	"sort"

	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/entities"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/pkg/database/json"
)

// JsonDb is a wrapper over json.JsonDb to interract with domain entities
type JsonDb struct {
	jsondb json.JsonDb[int, entities.ComixEntry, entities.ComixData]
}

// New creates new JsonDb
func New(filepath string) (JsonDb, error) {
	db, err := json.NewDb[int, entities.ComixData, entities.ComixEntry](filepath)
	if err != nil {
		return JsonDb{}, err
	}
	return JsonDb{
		jsondb: db,
	}, nil
}

// Create creates new entry in a database
func (db JsonDb) Create(id int, cd entities.ComixData) error {
	return db.jsondb.Create(id, cd)
}

// Read reads entry from database
func (db JsonDb) Read(id int) (entities.ComixData, error) {
	return db.jsondb.Read(id)
}

// ReadN returns first n entries from database.
// If n is greater than number of entries, then n entries will be returned
func (db JsonDb) ReadN(n int) ([]entities.ComixEntry, error) {
	res, err := db.ReadAll()
	if err != nil {
		return nil, err
	}
	return res[:min(n, len(res))], nil
}

// ReadAll reads all entries from database and sorts it by id
func (db JsonDb) ReadAll() ([]entities.ComixEntry, error) {
	m, err := db.jsondb.ReadAll()
	if err != nil {
		return nil, err
	}
	var res []entities.ComixEntry
	for k, v := range m {
		ce := entities.ComixEntry{
			Id: k,
			Cd: v,
		}
		res = append(res, ce)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Id < res[j].Id
	})
	return res, nil
}

// GetLastWrittenId returns last id written in database
func (db JsonDb) GetLastWrittenId() (int, error) {
	return db.jsondb.GetLastWrittenId()
}

// GetWrittenCount return number of comixes written to database
func (db JsonDb) GetWrittenCount() (int, error) {
	return db.jsondb.GetLastWrittenId()
}
