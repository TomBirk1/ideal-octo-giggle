package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if theydie(x, y, world, p) && world[x][y] == 255 {
				world[y][x] = 0
			}
			if theyreanimate(x, y, world, p) && world[x][y] == 0 {
				world[y][x] = 255
			}
		}
	}
	p.turns += 1
	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	cells := []cell{}
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] == 255 {
				liveCell := cell{y, x}
				cells = append(cells, liveCell)
			}
		}
	}
	return cells
}

func checkNumNeighbors(x, y int, world [][]byte, p golParams) int {
	count := 0
	count = int(world[(y+p.imageWidth-1)%p.imageWidth][(x+p.imageHeight-1)%p.imageHeight] + world[(y+p.imageWidth-1)%p.imageWidth][(x+p.imageHeight)%p.imageHeight] + world[(y+p.imageWidth-1)%p.imageWidth][(x+p.imageHeight+1)%p.imageHeight] + world[(y+p.imageWidth)%p.imageWidth][(x+p.imageHeight-1)%p.imageHeight] + world[(y+p.imageWidth)%p.imageWidth][(x+p.imageHeight+1)%p.imageHeight] + world[(y+p.imageWidth+1)%p.imageWidth][(x+p.imageHeight-1)%p.imageHeight] + world[(y+p.imageWidth+1)%p.imageWidth][(x+p.imageHeight)%p.imageHeight] + world[(y+p.imageWidth+1)%p.imageWidth][(x+p.imageHeight+1)%p.imageHeight])
	return count
}
func theydie(x, y int, world [][]byte, p golParams) bool {
	count := checkNumNeighbors(x, y, world, p)
	if count < 2 {
		return true
	} else if count == 2 || count == 3 {
		return false
	} else {
		return false
	}
}
func theyreanimate(x, y int, world [][]byte, p golParams) bool {
	if checkNumNeighbors(x, y, world, p) == 3 {
		return true
	} else {
		return false
	}
}
