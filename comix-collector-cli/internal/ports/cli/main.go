package cli

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/app"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/entities"
	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/pkg/xkcd"
)

func fetch(ctx context.Context, a app.App, limit int, threads int) error {
	var n int
	if limit != -1 {
		n = limit
	} else {
		var err error
		n, err = a.GetComixesCount()
		if err != nil {
			return err
		}
	}
	lastComix, err := a.GetLastWrittenId()
	if err != nil {
		return err
	}
	mux := &sync.Mutex{}
	wg := sync.WaitGroup{}
	comixChan := make(chan int)
	errChan := make(chan error)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case comixID, ok := <-comixChan:
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

					keywords, err := a.Stem(comix.Transcript + comix.Alt)
					if err != nil {
						errChan <- err
						continue
					}

					comixData := entities.ComixData{
						Url:      comix.ImgUrl,
						Keywords: keywords,
					}
					mux.Lock()
					err = a.Create(comixID, comixData)
					if err != nil {
						errChan <- err
					}
					mux.Unlock()
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

	for err := range errChan {
		log.Println(err)
	}

	return nil
}

func fetchWithOut(ctx context.Context, a app.App, limit int, threads int) error {
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

func printList(comixes []entities.ComixEntry) {
	for _, comix := range comixes {
		fmt.Println(comix)
	}
}

func Run(ctx context.Context, a app.App) error {
	f := ParseFlags()
	if f.Output {
		return fetchWithOut(ctx, a, f.Limit, f.Threads)
	} else {
		return fetch(ctx, a, -1, f.Threads)
	}
}
