package json

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"

	"golang.org/x/exp/maps"
)

type JsonDb[K cmp.Ordered, E any, V any] struct {
	filePath string
}

func NewDb[K cmp.Ordered, E any, V any](filePath string) (JsonDb[K, V, E], error) {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(filePath)
		if err != nil {
			return JsonDb[K, V, E]{}, err
		}
		f.Close()
	}

	return JsonDb[K, V, E]{
		filePath: filePath,
	}, nil
}

func (db JsonDb[K, E, V]) Create(key K, value V) error {
	data, err := os.ReadFile(db.filePath)
	if err != nil {
		return err
	}

	allRecords := make(map[K]V)
	if len(data) != 0 {
		err = json.Unmarshal(data, &allRecords)
		if err != nil && err != io.EOF {
			return err
		}
	}
	allRecords[key] = value
	jsonData, err := json.MarshalIndent(allRecords, "", "\t")
	if err != nil {
		return err
	}
	os.WriteFile(db.filePath, jsonData, 0600)

	return nil
}

func (db JsonDb[K, E, V]) GetLastWrittenId() (K, error) {
	m, err := db.ReadAll()
	if err == io.EOF {
		return *new(K), nil
	}
	if err != nil {
		return *new(K), err
	}
	keys := maps.Keys(m)
	return slices.Max(keys), nil
}

func (db JsonDb[K, E, V]) ReadAll() (map[K]V, error) {
	file, err := os.OpenFile(db.filePath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dec := json.NewDecoder(file)

	var m map[K]V
	err = dec.Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
func (db JsonDb[K, E, V]) Read(key K) (V, error) {
	m, err := db.ReadAll()
	if err != nil {
		return *new(V), err
	}
	if val, ok := m[key]; ok {
		return val, nil
	}

	return *new(V), fmt.Errorf("Not found")
}

func (db JsonDb[K, E, V]) GetWrittenCount() (int, error) {
	m, err := db.ReadAll()
	if err != nil {
		return -1, err
	}
	return len(maps.Keys(m)), nil
}