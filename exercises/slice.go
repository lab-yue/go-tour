package main

func pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for y := range img {
		img[y] = make([]uint8, dx)
		for x := range img[y] {
			img[y][x] = uint8(x + y)
		}
	}
	return img
}

func main() {
	pic.Show(pic)
}
