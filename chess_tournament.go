/*
 * Copyright (C) 2014 Anca Emanuel
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 2 as
 * published by the Free Software Foundation.
 */

package main

import (
	"fmt"
)

var n byte
var d [250][250]byte

func generate_tournament() {
	var i, j, k byte

	for i = 2; i < n; i++ {
		for j = 1; j < i; j++ {
			k = i + j - 2
			if k >= n {
				k -= n - 1
			}
			d[j][i] = k
		}
	}
	d[1][n] = n - 1
	k = n / 2
	for i = 2; i <= k; i++ {
		d[i][n] = 2 * (i - 1)
		d[i+k-1][n] = 2*(i-1) - 1
	}
}

func main() {
	fmt.Println("Chess players problem. Each player plays with all others one game per day.")
	fmt.Print("n=")
	fmt.Scan(&n)

	if n > 200 {
		return
	}

	if n%2 == 1 {
		n++
		fmt.Printf("%d represents pause\n", n)
	}


	generate_tournament()

	var z byte
	var i, j byte
	for z = 1; z < n; z++ {
		fmt.Printf("Day%2d:", z)
		for i = 2; i <= n; i++ {
			for j = 1; j < i; j++ {
				if d[j][i] == z {
					fmt.Printf("%2d %2d, ", j, i)
				}
			}
		}
		fmt.Printf("\n")
	}
}
