package repo

import (
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/app"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/entities"
)

// ComixDataRepo is an implementation of repository pattern
type ComixDataRepo struct {
	db app.Storager[int, entities.ComixData, entities.ComixEntry] // any database
}

// New creates new repo instantiation
func New(db app.Storager[int, entities.ComixData, entities.ComixEntry]) ComixDataRepo {
	return ComixDataRepo{
		db: db,
	}
}

// Create creates new entry in underlying database
func (cdr ComixDataRepo) Create(id int, cd entities.ComixData) error {
	err := cdr.db.Create(id, cd)
	if err != nil {
		return err
	}
	return nil
}

// Read returns entry from database
func (cdr ComixDataRepo) Read(id int) (entities.ComixData, error) {
	comix, err := cdr.db.Read(id)
	if err != nil {
		return entities.ComixData{}, err
	}

	return comix, nil
}

// ReadN returns n entries from database
func (cdr ComixDataRepo) ReadN(n int) ([]entities.ComixEntry, error) {
	comixes, err := cdr.db.ReadN(n)
	if err != nil {
		return nil, err
	}
	return comixes, nil
}

// ReadAll returns all entries from
func (cdr ComixDataRepo) ReadAll() ([]entities.ComixEntry, error) {
	return cdr.db.ReadAll()
}

// GetLastWrittenId returns last id written to database
func (cdr ComixDataRepo) GetLastWrittenId() (int, error) {
	return cdr.db.GetLastWrittenId()
}

// GetWrittenCount returns number of comixes written to database
func (cdr ComixDataRepo) GetWrittenCount() (int, error) {
	return cdr.db.GetWrittenCount()
}
