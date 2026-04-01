// 9c12f0fb-fe8b-4ad1-bdef-36b19307216e
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

// ANSI 256-color bright palette (Elmer-style vivid patchwork)
var colors = []string{
	"\033[38;5;196m", // red
	"\033[38;5;46m",  // green
	"\033[38;5;21m",  // blue
	"\033[38;5;226m", // yellow
	"\033[38;5;201m", // magenta/pink
	"\033[38;5;51m",  // cyan
	"\033[38;5;208m", // orange
	"\033[38;5;15m",  // white
	"\033[38;5;93m",  // purple
	"\033[38;5;163m", // hot pink
	"\033[38;5;82m",  // lime
	"\033[38;5;33m",  // dodger blue
}

const reset = "\033[0m"
const bold = "\033[1m"
const black = "\033[38;5;0m"
const gray = "\033[38;5;245m"
const darkGray = "\033[38;5;238m"

// The elephant shape as a grid. Each cell is 2 chars wide (██).
// 1 = body, 2 = eye, 0 = empty
// Approx 20 wide x 16 tall grid

var elephantShape = [][]int{
	//  0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19
	{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0}, // 0  top of back
	{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}, // 1
	{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}, // 2
	{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}, // 3  head starts
	{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, // 4
	{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 0, 0}, // 5  eye row
	{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, // 6
	{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}, // 7
	{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}, // 8  body
	{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}, // 9
	{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}, // 10
	{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0}, // 11
	{0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0}, // 12 legs
	{0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0}, // 13
	{0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0}, // 14
}

// Trunk shapes (appended to the right side of the elephant)
// Each is a list of (row, col_offset_from_17) with value
type trunkCell struct {
	row, col int
}

var trunkDown = []trunkCell{
	{5, 18}, {5, 19},
	{6, 18}, {6, 19},
	{7, 17}, {7, 18},
	{8, 17},
	{9, 16}, {9, 17},
	{10, 16},
	{11, 15}, {11, 16},
}

var trunkUp = []trunkCell{
	{3, 17}, {3, 18},
	{2, 18}, {2, 19},
	{1, 19},
	{0, 19},
}

var trunkMid = []trunkCell{
	{5, 18}, {5, 19},
	{6, 18}, {6, 19},
	{7, 18}, {7, 19},
	{8, 19},
}

// Tail cells
var tailCells = []trunkCell{
	{7, 1},
	{8, 0},
	{9, 0},
}

// Generate a patchwork color map for the elephant
func generateColorMap(seed int) [15][20]int {
	r := rand.New(rand.NewSource(int64(seed)))
	var cm [15][20]int
	for row := 0; row < 15; row++ {
		for col := 0; col < 20; col++ {
			cm[row][col] = r.Intn(len(colors))
		}
	}
	return cm
}

func buildFrame(trunkType int, legShift int, colorMap [15][20]int) []string {
	// Create a canvas 17 rows x 22 cols
	rows := 17
	cols := 22
	type cell struct {
		ch    string
		color string
	}
	canvas := make([][]cell, rows)
	for i := range canvas {
		canvas[i] = make([]cell, cols)
		for j := range canvas[i] {
			canvas[i][j] = cell{"  ", ""}
		}
	}

	// Draw main body (offset by 1 row for head room on trunk-up)
	offsetR := 1
	for r := 0; r < 15; r++ {
		for c := 0; c < 20; c++ {
			if elephantShape[r][c] > 0 {
				cr := r + offsetR
				if cr < rows && c < cols {
					if elephantShape[r][c] == 2 {
						canvas[cr][c] = cell{"⬤ ", "\033[38;5;232m"} // dark eye
					} else {
						canvas[cr][c] = cell{"██", colors[colorMap[r][c]]}
					}
				}
			}
		}
	}

	// Apply leg shift for dancing (-1, 0, +1)
	if legShift != 0 {
		for r := 13; r <= 15; r++ {
			if r >= rows {
				continue
			}
			if legShift > 0 {
				for c := cols - 1; c > 0; c-- {
					canvas[r][c] = canvas[r][c-1]
				}
				canvas[r][0] = cell{"  ", ""}
			} else {
				for c := 0; c < cols-1; c++ {
					canvas[r][c] = canvas[r][c+1]
				}
				canvas[r][cols-1] = cell{"  ", ""}
			}
		}
	}

	// Draw trunk
	var trunk []trunkCell
	switch trunkType {
	case 0:
		trunk = trunkDown
	case 1:
		trunk = trunkUp
	case 2:
		trunk = trunkMid
	}
	for _, tc := range trunk {
		cr := tc.row + offsetR
		if cr >= 0 && cr < rows && tc.col < cols {
			ci := colorMap[tc.row%15][tc.col%20]
			canvas[cr][tc.col] = cell{"██", colors[ci]}
		}
	}

	// Draw tail
	for _, tc := range tailCells {
		cr := tc.row + offsetR
		if cr >= 0 && cr < rows && tc.col < cols {
			ci := colorMap[tc.row%15][tc.col%20]
			canvas[cr][tc.col] = cell{"██", colors[ci]}
		}
	}

	// Render to strings
	lines := make([]string, rows)
	for r := 0; r < rows; r++ {
		line := ""
		for c := 0; c < cols; c++ {
			line += canvas[r][c].color + canvas[r][c].ch
		}
		lines[r] = line + reset
	}
	return lines
}

var music = [8]string{
	"  🎵 ✨ 🎶        ✨ 🎵   ",
	"     🎶   ✨  🎵 ✨       ",
	"  ✨   🎵     🎶    ✨    ",
	"    🎶 ✨  🎵   ✨  🎶    ",
	"  🎵    🎶   ✨   🎵      ",
	"     ✨   🎵 🎶 ✨         ",
	"  🎶 ✨      🎵    🎶     ",
	"    🎵   ✨ 🎶   ✨  🎵   ",
}

var ground = [2]string{
	` ░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░`,
	` ▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒░▒`,
}

type danceStep struct {
	trunk    int // 0=down, 1=up, 2=mid
	legShift int // -1, 0, +1
	seed     int
}

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	fmt.Print("\033[2J\033[?25l")

	defer func() {
		fmt.Print("\033[?25h")
		fmt.Printf("\n  %s🐘  Elmer takes a bow! Goodbye! 🐘%s\n\n", bold, reset)
	}()

	// Choreography
	steps := []danceStep{
		{0, 0, 42},   // stand
		{2, 1, 42},   // trunk mid, lean right
		{0, 0, 42},   // stand
		{2, -1, 42},  // trunk mid, lean left
		{0, 0, 42},   // stand
		{1, 0, 42},   // trunk UP
		{1, 1, 42},   // trunk UP, lean right
		{1, -1, 42},  // trunk UP, lean left
		{0, 0, 100},  // stand (new colors!)
		{2, -1, 100}, // trunk mid, lean left
		{0, 0, 100},  // stand
		{2, 1, 100},  // trunk mid, lean right
		{0, 0, 100},  // stand
		{1, 0, 100},  // trunk UP
		{1, -1, 100}, // trunk UP lean left
		{1, 1, 100},  // trunk UP lean right
	}

	tick := 0
	for {
		select {
		case <-sig:
			return
		default:
			step := steps[tick%len(steps)]
			cm := generateColorMap(step.seed)
			frame := buildFrame(step.trunk, step.legShift, cm)

			fmt.Print("\033[H")
			trumpet := ""
			if step.trunk == 1 {
				trumpet = "  📢 TOOT!"
			}
			fmt.Printf("\n  %s\033[38;5;213m 🐘  P A T C H W O R K   E L E P H A N T  🐘 %s%s\n",
				bold, trumpet, reset)
			fmt.Printf("  %s%s%s\n\n", gray, music[tick%8], reset)

			for _, line := range frame {
				fmt.Printf("    %s\n", line)
			}

			fmt.Printf("    %s%s%s\n", darkGray, ground[tick%2], reset)
			fmt.Printf("\n  %s♫ Ctrl+C to stop the music ♫%s\n", gray, reset)

			tick++
			time.Sleep(400 * time.Millisecond)
		}
	}
}
