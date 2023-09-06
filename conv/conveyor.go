package conv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Conveyor() error {
	//defer utils.TimeTrack(time.Now(), "simple mode")

	wordsCh := make(chan string)
	counterCh := make(chan int)

	go counter(counterCh)
	go wordCheck(counterCh, wordsCh)
	do(wordsCh)

	return nil
}

func do(words chan<- string) {
	FILES := []string{
		"./data/first.txt",
		"./data/second.txt",
		"./data/third.txt",
		"./data/fifth.txt",
		"./data/fourth.txt",
	}
	//var waitGroup sync.WaitGroup
	//fmt.Printf("%#v\n", waitGroup)

	for _, path := range FILES {
		//waitGroup.Add(1)

		go func() {
			//defer waitGroup.Done()
			f, err := os.Open(path)

			if err != nil {
				fmt.Println(err)
			}

			defer f.Close()

			scanner := bufio.NewScanner(f)
			scanner.Split(bufio.ScanWords)

			for scanner.Scan() {
				words <- scanner.Text()
			}

			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}

		}()

	}
	//waitGroup.Wait()

}

func wordCheck(nums chan<- int, words <-chan string) {
	for w := range words {
		if strings.ToLower(w) == "the" {
			nums <- 1
		}
	}

}

func counter(nums <-chan int) {
	var COUNTER int
	for num := range nums {
		COUNTER = COUNTER + num
	}

	fmt.Println("Total", COUNTER)
}
