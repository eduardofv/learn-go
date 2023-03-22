//https://go.dev/tour/moretypes/18

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var ret [dx][dy]uint8
	
/*	for i := 0; i < dy; i++ {
	
	}
*/
	return ret
}

func main() {
	pic.Show(Pic)
}

