package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleOnce struct{}

var singleInstanceOnce *singleOnce

func getInstanceOnce() *singleOnce {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating singleOnce instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("SingleOnce instance already created.")
	}

	return singleInstanceOnce
}
