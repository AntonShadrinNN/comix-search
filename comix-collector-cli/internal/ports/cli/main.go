package cli

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/app"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/entities"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/pkg/xkcd"
	"github.com/schollz/progressbar/v3"
)

func setupProgressBar(size int) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(size,
		progressbar.OptionClearOnFinish(),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionFullWidth(),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stdout, "\n")
		}),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetDescription("[cyan] Updating database...[reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[white]❚[reset]",
			SaucerHead:    "[red]❚[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	return bar
}

// fetch fetches comixes limit number of comixes in some number of threads
func fetch(ctx context.Context, a app.AppRepo, limit int, threads int) error {
	var n int
	// Unlimited
	if limit != -1 {
		n = limit
	} else {
		var err error
		n, err = a.GetComixesCount()
		if err != nil {
			return err
		}
	}
	// Continue from last saved comic
	lastComix, err := a.GetLastWrittenId()
	if err != nil {
		return err
	}
	mux := &sync.Mutex{}
	wg := sync.WaitGroup{}
	comixChan := make(chan int)
	errChan := make(chan error)
	bar := setupProgressBar(n - lastComix - 1)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case comixID, ok := <-comixChan:
					// channel is closed
					if !ok {
						return
					}

					comix, err := a.FetchComixById(comixID)
					if err == xkcd.ErrNotFound {
						continue
					}
					if err != nil {
						errChan <- err
						continue
					}
					// Stem keywords
					keywords, err := a.Stem(comix.Transcript + comix.Alt)
					if err != nil {
						errChan <- err
						continue
					}

					comixData := entities.ComixData{
						Url:      comix.ImgUrl,
						Keywords: keywords,
					}
					// Save to database
					mux.Lock()
					err = a.Create(comixID, comixData)
					if err != nil {
						errChan <- err
						mux.Unlock()
						continue
					}
					mux.Unlock()
					err = bar.Add(1)
					if err != nil {
						errChan <- err
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	for i := lastComix; i <= n; i++ {
		comixChan <- i
	}
	close(comixChan)
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Wait for errors
	for err := range errChan {
		return err
	}

	return nil
}

// fetchWithOut is the same as fetch, but also prints output to console
func fetchWithOut(ctx context.Context, a app.AppRepo, limit int, threads int) error {
	err := fetch(ctx, a, limit, threads)
	if err != nil {
		return err
	}
	var comixes []entities.ComixEntry
	if limit == -1 {
		comixes, err = a.ReadAll()
	} else {
		comixes, err = a.ReadN(limit)
	}
	if err != nil {
		return err
	}
	printList(comixes)
	return nil
}

// Pretty-print list
func printList(comixes []entities.ComixEntry) {
	for _, comix := range comixes {
		fmt.Println(comix)
	}
}

// Run runs cli
func Run(ctx context.Context, a app.AppRepo) error {
	f, err := ParseFlags()
	if err != nil {
		return err
	}
	if f.Output {
		return fetchWithOut(ctx, a, f.Limit, f.Threads)
	} else {
		return fetch(ctx, a, -1, f.Threads)
	}
}
