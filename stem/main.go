package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

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
	flags := ParseFlags()

	err := flags.validateFlags()
	if err != nil {
		panic(err)
	}

	stemmer := SnowballStemmer{
		language:      English,
		stopWordsFunc: getStopWords,
	}
	stemmed, err := stemmer.Stem(strings.Split(flags.InitialString, " "))
	if err != nil {
		panic(err)
	}
	log.Printf(stemmed)
}
