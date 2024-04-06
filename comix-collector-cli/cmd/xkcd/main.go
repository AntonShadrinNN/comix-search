/*
xkcdparse parses xkcd.com and saves data to storage with stemmed output.
It is also possible to display fetched data on a screen.

Usage:

	xkcdparse [flags]

The flags are:

	-o
	    Print fetched data on a screen.
	-n
		Limit to print, e.g. -n 10 means to print only 10 first comixes.
		Setting this parameter to -1 is equal to not use this flag.
		Must be combined with -o, otherwise has no effect
	-t
		Number of threads to parse with. Default to 1.
*/
package main

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/adapters/jsondb"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/adapters/repo"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/app"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/config"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/ports/cli"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/pkg/words"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/pkg/xkcd"
)

// getStopWords returns most frequent english words
// Currently data is taken from open bases
func getStopWords() ([]string, error) {
	resp, err := http.Get("https://gist.githubusercontent.com/rg089/35e00abf8941d72d419224cfd5b5925d/raw/12d899b70156fd0041fa9778d657330b024b959c/stopwords.txt")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func main() {
	// Parse yaml config or set default values
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// Initialize database connection. If already exists, it wouldn't be recreated
	db, err := jsondb.New(cfg.DbFile)
	if err != nil {
		panic(err)
	}

	// Initialize ComixDataRepository
	cr := repo.New(db)

	// Initialize stemmer
	stemmer := words.NewSnowballStemmer(words.English, getStopWords)

	ctx := context.Background()
	// Initialize client to parse xkcd
	client := xkcd.NewClient(ctx, cfg.Url)

	// Build app object
	app := app.NewApp(cr, stemmer, client)

	// Run cli
	err = cli.Run(ctx, app)
	if err != nil {
		panic(err)
	}
}
