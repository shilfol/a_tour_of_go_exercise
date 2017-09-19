package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8,dy)
	
	for i:=0 ; i < dy ; i++ {
		tmp := make([]uint8 , dx)
		for j:=0 ; j < dx ; j++ {
			
			tmp[j] = uint8(i*j);
		}
		picture[i] = tmp
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
