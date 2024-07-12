package main

import (
	"log"
)

func binsertionsort2b(lns []string) []string {
	n := len(lns)
	if n == 1 {
		return lns
	}
	for i := 0; i < n; i++ {
		for j := i; j > 0 && lns[j-1] > lns[j]; j-- {
			lns[j], lns[j-1] = lns[j-1], lns[j]
		}
	}
	return lns
}

// bostic
func rsort2b(lns []string, recix int) []string {
	var piles = make([][]string, 256)
	var nc int
	nl := len(lns)

	if nl == 0 {
		log.Fatal("rsort2b: 0 len []string: ", recix)
	}
	if nl < THRESHOLD {
		return binsertionsort2b(lns)
	}

	// deal []string into piles
	for i, _ := range lns {
		var c int

		if len(lns[i]) == 0 {
			log.Fatal("rsort2b 0 length string")
		}
		if recix >= len(lns[i]) {
			c = 0
		} else {
			c = int(lns[i][recix])
		}
		piles[c] = append(piles[c], string(lns[i]))
		if len(piles[c]) == 1 {
			nc++ // number of piles so far
		}
	}

	// sort the piles
	if nc == 1 {
		return binsertionsort2b(lns)
	}
	for i, _ := range piles {
		if len(piles[i]) == 0 {
			continue
		}

		// sort pile
		if len(piles[i]) < THRESHOLD {
			piles[i] = binsertionsort2b(piles[i])
		} else {
			piles[i] = rsort2b(piles[i], recix+1)
		}
		nc--
		if nc == 0 {
			break
		}
	}

	// combine the sorted piles
	var slns []string
	for i, _ := range piles {
		for j, _ := range piles[i] {
			slns = append(slns, piles[i][j])
		}
	}
	return slns
}
