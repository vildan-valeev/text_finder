package conv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"text_finder/utils"
	"time"
)

var COUNTER uint32

var signal chan struct{}

func Conveyor() error {
	defer utils.TimeTrack(time.Now(), "conv mode")

	wordsCh := make(chan string)
	//counterCh := make(chan int)

	signal = make(chan struct{})

	//go counter(counterCh)
	//go wordCheck(counterCh, wordsCh)
	do(wordsCh)
	fmt.Println(COUNTER)
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
	var waitGroup sync.WaitGroup
	//fmt.Printf("%#v\n", waitGroup)

	for _, path := range FILES {
		waitGroup.Add(1)
		go func(p string) {
			defer waitGroup.Done()

			f, err := os.Open(p)

			if err != nil {
				fmt.Println(err)
			}

			defer f.Close()

			scanner := bufio.NewScanner(f)
			scanner.Split(bufio.ScanWords)

			for scanner.Scan() {
				//words <- scanner.Text()
				if strings.ToLower(scanner.Text()) == "the" {
					atomic.AddUint32(&COUNTER, 1)
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
			//waitGroup.Done()
		}(path)

	}
	waitGroup.Wait()

}

//func wordCheck(nums chan<- int, words <-chan string) {
//	for w := range words {
//		//TODO: ",the", "the," обработка со знаками
//		if strings.ToLower(w) == "the" {
//			nums <- 1
//		}
//	}
//	close(nums)
//
//}
//
//func counter(nums <-chan int) {
//	for _ = range nums {
//
//		atomic.AddUint32(&COUNTER, 1)
//	}
//}

/*
/snap/go/10319/bin/go build -race -o /home/boom/GolandProjects/text_finder/start /home/boom/GolandProjects/text_finder/main.go #gosetup
/home/boom/GolandProjects/text_finder/start -s conv
Run mode is conv.
666876
2023/09/18 13:13:10 simple mode took 3.246455345s

Process finished with the exit code 0
*/
