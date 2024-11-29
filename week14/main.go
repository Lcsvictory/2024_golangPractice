package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	var names []string
	var counts []int
	for _, line := range lines {
		flag := false
		for i, name := range names {
			if name == line {
				counts[i]++
				flag = true
			}
		}
		if !flag {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("Votes for %s: %d\n", name, counts[i])
	}

	// counts := make(map[string]int)
	// for _, line := range lines {
	// 	counts[line]++
	// }
	// for name, count := range counts {
	// 	fmt.Printf("Votes for %s: %d\n", name, count)
	// }
}
