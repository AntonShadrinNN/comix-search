package app

import "github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/entities"

// Storager represents any storage. K is key in storage, V is value, E is full entry
type Storager[K comparable, V any, E any] interface {
	Create(key K, value V) error   // Create new entry
	Read(key K) (V, error)         // Get value by key
	GetLastWrittenId() (K, error)  // Get last written id
	GetWrittenCount() (int, error) // Get number of entries in databaes
	ReadN(n int) ([]E, error)      // Get N entries from database
	ReadAll() ([]E, error)         // Get all entries from database
}

// A Stemmer reduces inflected words to their word stem
type Stemmer interface {
	Stem(words []string) ([]string, error)
}

// ComixClient represents any client to fetch comixes with
type ComixClient interface {
	GetComixesCount() (int, error)                 // General number of comixes on a resourse
	FetchComixById(id int) (entities.Comix, error) // Fetch concrete comix
}

// ComixDataRepo implements repository pattern for ComixData usage
type ComixDataRepo interface {
	Create(id int, cd entities.ComixData) error // Write to storage
	Read(id int) (entities.ComixData, error)    // Read from storage
	ReadN(n int) ([]entities.ComixEntry, error) // Read n entries from storage
	ReadAll() ([]entities.ComixEntry, error)    // Read all entries from storage
	GetLastWrittenId() (int, error)             // Get last written id
	GetWrittenCount() (int, error)              // Get a number of written entries
}

// AppRepo implements repository pattern for App usage
type AppRepo interface {
	Create(id int, cd entities.ComixData) error
	Read(id int) (entities.ComixData, error)
	ReadN(n int) ([]entities.ComixEntry, error)
	ReadAll() ([]entities.ComixEntry, error)
	Stem(sentence string) ([]string, error)
	GetComixesCount() (int, error)
	GetLastWrittenId() (int, error)
	FetchComixById(id int) (entities.Comix, error)
	GetWrittenCount() (int, error)
}
