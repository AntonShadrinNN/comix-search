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
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := jsondb.New(cfg.DbFile)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	cr := repo.New(db)

	stemmer := words.NewSnowballStemmer(words.English, getStopWords)

	client := xkcd.NewClient(ctx, cfg.Url)
	app := app.NewApp(cr, stemmer, client)

	err = cli.Run(ctx, app)
	if err != nil {
		panic(err)
	}
}
