// http://tour.golang.org/#36
package main

/*
Implement Pic. It should return a slice of length dy, each element of which is a
slice of dx 8-bit unsigned integers. When you run the program, it will display
your picture, interpreting the integers as grayscale (well, bluescale) values.
*/
import "code.google.com/p/go-tour/pic"

func pixelValue(i, j int) uint8 {
	// A pleasant gradient.
	return uint8((i + j) / 2)
}

func Pic(dx, dy int) [][]uint8 {
	// Make a slice of len dy made up of slices of uint8's ([]uint8)
	result := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		// Make a slice of len dx
		temp := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			temp[j] = pixelValue(i, j)
		}
		result[i] = temp
	}

	return result
}

func main() {
	pic.Show(Pic)
}
