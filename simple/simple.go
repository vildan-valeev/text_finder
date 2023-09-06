package simple

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text_finder/utils"
	"time"
)

func Simple() error {
	defer utils.TimeTrack(time.Now(), "simple mode")

	var COUNTER int

	FILES := []string{
		"./data/first.txt",
		"./data/second.txt",
		"./data/third.txt",
		"./data/fifth.txt",
		"./data/fourth.txt",
	}

	for _, path := range FILES {
		f, err := os.Open(path)

		if err != nil {
			fmt.Println(err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		var fileCounter int
		for scanner.Scan() {
			if strings.ToLower(scanner.Text()) == "the" {
				COUNTER++
				fileCounter++
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("FileCounter", fileCounter)
	}

	fmt.Println("Total", COUNTER)

	return nil
}

/*
Run mode is simple.
Total 666876
2023/09/07 01:05:01 simple mode took 917.042371ms
*/
