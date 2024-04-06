package repo

import (
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/app"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/entities"
)

type ComixDataRepo struct {
	db app.Storager[int, entities.ComixData, entities.ComixEntry]
}

func New(db app.Storager[int, entities.ComixData, entities.ComixEntry]) ComixDataRepo {
	return ComixDataRepo{
		db: db,
	}
}

func (cdr ComixDataRepo) Create(id int, cd entities.ComixData) error {
	err := cdr.db.Create(id, cd)
	if err != nil {
		return err
	}
	return nil
}

func (cdr ComixDataRepo) Read(id int) (entities.ComixData, error) {
	comix, err := cdr.db.Read(id)
	if err != nil {
		return entities.ComixData{}, err
	}

	return comix, nil
}

func (cdr ComixDataRepo) ReadN(n int) ([]entities.ComixEntry, error) {
	comixes, err := cdr.db.ReadN(n)
	if err != nil {
		return nil, err
	}
	return comixes, nil
}

func (cdr ComixDataRepo) ReadAll() ([]entities.ComixEntry, error) {
	return cdr.db.ReadAll()
}

func (cdr ComixDataRepo) GetLastWrittenId() (int, error) {
	return cdr.db.GetLastWrittenId()
}

func (cdr ComixDataRepo) GetWrittenCount() (int, error) {
	return cdr.db.GetWrittenCount()
}
