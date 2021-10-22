package main

import "fmt"

func calculateNextState(p golParams, world [][]byte) [][]byte {
	sum := 0
	newworld := make([][]byte, p.imageHeight)
	for i := 0; i < p.imageWidth; i++ {
		newworld[i] = make([]byte, p.imageWidth)
		copy(newworld[i], world[i])
	}

	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			/*count := int(world[(y+p.imageWidth-1)%p.imageWidth][(x+p.imageHeight-1)%p.imageHeight])
			count += int(world[(y+p.imageWidth-1)%p.imageWidth][(x+p.imageHeight)%p.imageHeight])
			count += int(world[(y+p.imageWidth-1)%p.imageWidth][(x+p.imageHeight+1)%p.imageHeight])
			count += int(world[(y+p.imageWidth)%p.imageWidth][(x+p.imageHeight-1)%p.imageHeight])
			count += int(world[(y+p.imageWidth)%p.imageWidth][(x+p.imageHeight+1)%p.imageHeight])
			count += int(world[(y+p.imageWidth+1)%p.imageWidth][(x+p.imageHeight-1)%p.imageHeight])
			count += int(world[(y+p.imageWidth+1)%p.imageWidth][(x+p.imageHeight)%p.imageHeight])
			count += int(world[(y+p.imageWidth+1)%p.imageWidth][(x+p.imageHeight+1)%p.imageHeight])
			count = count / 255
			sum = count */

			sum = howmanyneighbors(y, x, world, p)
			if sum == 3 {
				newworld[y][x] = 255
			}
			if world[y][x] != 0 {
				if sum < 2 {
					newworld[y][x] = 0
				} else if sum == 2 || sum == 3 {
					newworld[y][x] = 255
				} else {
					newworld[y][x] = 0
				}
			}

		}
	}
	return newworld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	cells := []cell{}
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] != 0 {
				liveCell := cell{x, y}
				cells = append(cells, liveCell)
			}
		}
	}
	fmt.Println(cells)
	return cells
}
func howmanyneighbors(y, x int, world [][]byte, p golParams) int {
	alive := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if world[(y+i+p.imageHeight)%p.imageHeight][(x+j+p.imageWidth)%p.imageWidth] == 255 {
				alive += 1
			}
		}
	}
	if world[y][x] == 255 {
		alive -= 1
	}
	return alive
}

/*func checkNumNeighbors(y, x int, world [][]byte, p golParams) int {
	count := 0
	count += int(world[((x + p.imageWidth - 1) % p.imageWidth)][(y+p.imageHeight-1)%p.imageHeight])
	count += int(world[((x + p.imageWidth - 1) % p.imageWidth)][(y+p.imageHeight)%p.imageHeight])
	count += int(world[((x + p.imageWidth - 1) % p.imageWidth)][(y+p.imageHeight+1)%p.imageHeight])
	count += int(world[((x + p.imageWidth) % p.imageWidth)][(y+p.imageHeight-1)%p.imageHeight])
	count += int(world[((x + p.imageWidth) % p.imageWidth)][(y+p.imageHeight+1)%p.imageHeight])
	count += int(world[((x + p.imageWidth + 1) % p.imageWidth)][(y+p.imageHeight-1)%p.imageHeight])
	count += int(world[((x + p.imageWidth + 1) % p.imageWidth)][(y+p.imageHeight)%p.imageHeight])
	count += int(world[((x + p.imageWidth + 1) % p.imageWidth)][(y+p.imageHeight+1)%p.imageHeight])
	fmt.Print(count / 255)
	return count / 255
}

func theydie(y, x int, world [][]byte, p golParams) bool {
	count := checkNumNeighbors(y, x, world, p)
	if count < 2 {
		return true
	} else if count == 2 || count == 3 {
		return false
	} else {
		return false
	}
}
func theyreanimate(y, x int, world [][]byte, p golParams) bool {
	if checkNumNeighbors(y, x, world, p) == 3 {
		return true
	} else {
		return false
	}
}*/
