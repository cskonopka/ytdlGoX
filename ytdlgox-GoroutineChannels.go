package main
import (
       "fmt"
       "sync"
       "github.com/codeskyblue/go-sh"
)
var wg sync.WaitGroup

func toggle(ch <-chan string) {
       for k := range ch { 
              fmt.Println(k)
              sh.Command("youtube-dl", "-f", "mp4", k).Run()
       }
       wg.Done()
}

func main() {
    ytVideos := []string{
       "https://www.youtube.com/watch?v=ygS_Cwng7IU", 
       "https://www.youtube.com/watch?v=8i5ljNff2Tw", 
       "https://www.youtube.com/watch?v=uo14xGYwWd4",
    }
       ch1 := make(chan string)
       ch2 := make(chan string)
       ch3 := make(chan string)
       wg.Add(3)
       defer wg.Wait()
       go toggle(ch2)
       go toggle(ch3)
       go toggle(ch1)
       for i:=0;i<1;i++ {
              ch1<-ytVideos[i]
              ch2<-ytVideos[i+1]
              ch3<-ytVideos[i+2]
              fmt.Println(i)
       }
       close(ch1)
       close(ch2)
       close(ch3)
}