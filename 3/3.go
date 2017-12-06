package main

import "fmt"
import "math"
import "time"
import "os"
import "strconv"

// go run 3.go <input>

/*

"Debugging is twice as hard as writing the code in the first place.
 Therefore, if you write the code as cleverly as possible, you are,
 by definition, not smart enough to debug it." -- Brian W. Kernighan

I probably should have just brute forced this, it took me ages to track
down the off-by-one error I had for even squares *some* of the time...

*/

func main() {

	start := time.Now()

	target, _ := strconv.Atoi(os.Args[1])
	squareSize := int(math.Ceil(math.Sqrt(float64(target))))
	startVal := (squareSize-1)*(squareSize-1) + 1

	dirX := 0
	dirY := 0
	min := 0
	max := 0
	posX := 0
	posY := 0

	if (squareSize % 2) == 1 {
		min = -1 * (squareSize - 1) / 2
		max = (squareSize - 1) / 2
		dirX = 0
		dirY = -1
		posX = min
		posY = max

		fmt.Println("odd", squareSize)
		fmt.Println("start:", startVal, "min:", min, "max:", max)

	} else {
		min = (-1 * (squareSize) / 2) + 1
		max = (squareSize) / 2
		dirX = 0
		dirY = 1
		posX = max
		posY = min

		fmt.Println("even", squareSize)
		fmt.Println("start:", startVal, "min:", min, "max:", max)

	}

	for i := startVal; i < target; i++ {
		fmt.Println(i, posX, posY)

		if dirY == -1 && posY == min {
			dirY = 0
			dirX = 1
		}
		if dirY == 1 && posY == max {
			dirY = 0
			dirX = -1
		}
		if dirX == -1 && posX == min {
			dirY = -1
			dirX = 0
		}
		if dirX == 1 && posX == max {
			dirY = 1
			dirX = 0
		}

		posX += dirX
		posY += dirY
	}

	fmt.Println("result:", posX, posY, (math.Abs(float64(posX)) + math.Abs(float64(posY))))

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed, "ms")
}
