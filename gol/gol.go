package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	//Make a new world as a list of lists, copying previous world data
	newworld := make([][]byte, p.imageHeight)
	for i := 0; i < p.imageHeight; i++ {
		newworld[i] = make([]byte, p.imageHeight)
		copy(newworld[i], world[i])
	}
	//Counting the values of neighboring cells
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			count := int(world[(y+p.imageWidth-1)%p.imageWidth][(x+p.imageHeight-1)%p.imageHeight])
			count += int(world[(y+p.imageWidth-1)%p.imageWidth][(x+p.imageHeight)%p.imageHeight])
			count += int(world[(y+p.imageWidth-1)%p.imageWidth][(x+p.imageHeight+1)%p.imageHeight])
			count += int(world[(y+p.imageWidth)%p.imageWidth][(x+p.imageHeight-1)%p.imageHeight])
			count += int(world[(y+p.imageWidth)%p.imageWidth][(x+p.imageHeight+1)%p.imageHeight])
			count += int(world[(y+p.imageWidth+1)%p.imageWidth][(x+p.imageHeight-1)%p.imageHeight])
			count += int(world[(y+p.imageWidth+1)%p.imageWidth][(x+p.imageHeight)%p.imageHeight])
			count += int(world[(y+p.imageWidth+1)%p.imageWidth][(x+p.imageHeight+1)%p.imageHeight])
			count = count / 255
			//Death / reanimation logic for the current cell
			if count == 3 {
				newworld[y][x] = 255
			}
			if world[y][x] != 0 {
				if count < 2 {
					newworld[y][x] = 0
				} else if count == 2 || count == 3 {
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
	//Given the current image count all alive cells and return
	cells := []cell{}
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] != 0 {
				liveCell := cell{x, y}
				cells = append(cells, liveCell)
			}
		}
	}
	return cells
}
