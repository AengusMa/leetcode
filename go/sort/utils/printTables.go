package utils

import (
	"fmt"
	"unicode"
)

type bd struct {
	H  rune // BOX DRAWINGS HORIZONTAL
	V  rune // BOX DRAWINGS VERTICAL
	VH rune // BOX DRAWINGS VERTICAL AND HORIZONTAL
	HU rune // BOX DRAWINGS HORIZONTAL AND UP
	HD rune // BOX DRAWINGS HORIZONTAL AND DOWN
	VL rune // BOX DRAWINGS VERTICAL AND LEFT
	VR rune // BOX DRAWINGS VERTICAL AND RIGHT
	DL rune // BOX DRAWINGS DOWN AND LEFT
	DR rune // BOX DRAWINGS DOWN AND RIGHT
	UL rune // BOX DRAWINGS UP AND LEFT
	UR rune // BOX DRAWINGS UP AND RIGHT
}

var m = map[string]bd{
	"ascii":       bd{'-', '|', '+', '+', '+', '+', '+', '+', '+', '+', '+'},
	"box-drawing": bd{'─', '│', '┼', '┴', '┬', '┤', '├', '┐', '┌', '┘', '└'},
}

// Print 格式化输出表格
func Print(data [][]string) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			l := 0
			for _, r := range data[i][j] {
				if isHan(r) {
					l++
				}
			}
			fmt.Printf("|%-[1]*s", 15-l/2, data[i][j])
		}
		fmt.Printf("\n")
	}
}
func isHan(r rune) bool {
	return unicode.Is(unicode.Han, r)

}
