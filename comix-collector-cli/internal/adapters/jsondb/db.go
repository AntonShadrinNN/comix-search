package jsondb

import (
	"sort"

	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/entities"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/pkg/database/json"
)

type Db struct {
	jsondb json.JsonDb[int, entities.ComixEntry, entities.ComixData]
}

func New(filepath string) (Db, error) {
	db, err := json.NewDb[int, entities.ComixData, entities.ComixEntry](filepath)
	if err != nil {
		return Db{}, err
	}
	return Db{
		jsondb: db,
	}, nil
}

func (db Db) Create(id int, cd entities.ComixData) error {
	return db.jsondb.Create(id, cd)
}

func (db Db) Read(id int) (entities.ComixData, error) {
	return db.jsondb.Read(id)
}

func (db Db) ReadN(n int) ([]entities.ComixEntry, error) {
	res, err := db.ReadAll()
	if err != nil {
		return nil, err
	}
	return res[:min(n, len(res))], nil
}

func (db Db) ReadAll() ([]entities.ComixEntry, error) {
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

func (db Db) GetLastWrittenId() (int, error) {
	res, err := db.jsondb.GetLastWrittenId()
	return res, err
}

func (db Db) GetWrittenCount() (int, error) {
	res, err := db.jsondb.GetLastWrittenId()
	return res, err
}
