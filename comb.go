package main

import (
	"fmt"
)

var n, k, ig byte
var c []byte

func next_comb() {
	var i, j byte

	if ig == 0 {
		for i = 0; i < k; i++ {
			c[i] = i + 1
		}
		ig = 1
	} else {
		for i = k; i > 0; i-- {
			if c[i-1] < n-k+i {
				c[i-1]++
				for j = i; j < k; j++ {
					c[j] = c[j-1] + 1
				}
				return
			}
		}
		ig = 0
	}
}

func main() {
	ig = 0
	fmt.Print("n=")
	fmt.Scan(&n)
	fmt.Print("k=")
	fmt.Scan(&k)

	c = make([]byte, k)
	next_comb()
	var count = 0;
	for ig == 1 {
		count++;
		fmt.Printf("%3d. ",count)
		var i byte;
		for i = 0; i < k; i++ {
			fmt.Printf("%3d",c[i])
		}
		fmt.Printf("\n")
		next_comb()
	}
}
