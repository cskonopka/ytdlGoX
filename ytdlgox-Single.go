package main

import (
	"github.com/codeskyblue/go-sh"
	"sync"
)

var wg sync.WaitGroup 

func ripYT (vidLink string) {
	sh.Command("youtube-dl", "-f", "mp4", vidLink).Run()
	wg.Done()
}

func main() {
	wg.Add(1)
	go ripYT("https://www.youtube.com/watch?v=ygS_Cwng7IU")
	wg.Wait()
}