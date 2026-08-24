package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: qr-code-generator <text> [out.svg]")
		os.Exit(1)
	}
	grid := buildGrid(os.Args[1])
	svg := renderSVG(grid)
	w := os.Stdout
	if len(os.Args) > 2 {
		f, err := os.Create(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
		w = f
	}
	w.WriteString(svg)
	fmt.Fprintf(os.Stderr, "%dx%d grid\n", len(grid), len(grid))
}

func buildGrid(text string) [][]bool {
	g := make([][]bool, 21)
	for i := range g {
		g[i] = make([]bool, 21)
	}
	drawFinder(g, 0, 0)
	drawFinder(g, 14, 0)
	drawFinder(g, 0, 14)
	pos := 0
	for _, ch := range []byte(text) {
		for bit := 7; bit >= 0; bit-- {
			set := (ch>>uint(bit))&1 == 1
			r := pos / 19
			c := pos % 19
			if r < 21 && c+1 < 21 {
				g[r][c+1] = set
			}
			pos++
		}
	}
	return g
}

func drawFinder(g [][]bool, r, c int) {
	for dr := 0; dr < 7; dr++ {
		for dc := 0; dc < 7; dc++ {
			if r+dr < 21 && c+dc < 21 {
				border := dr == 0 || dr == 6 || dc == 0 || dc == 6
				inner := dr >= 2 && dr <= 4 && dc >= 2 && dc <= 4
				g[r+dr][c+dc] = border || inner
			}
		}
	}
}

func renderSVG(g [][]bool) string {
	var b strings.Builder
	n := len(g)
	sz := 8
	fmt.Fprintf(&b, "<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"%d\" height=\"%d\">\n", n*sz, n*sz)
	b.WriteString("<rect width=\"100%\" height=\"100%\" fill=\"white\"/>\n")
	for r, row := range g {
		for c, v := range row {
			if v {
				fmt.Fprintf(&b, "<rect x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\" fill=\"black\"/>\n", c*sz, r*sz, sz, sz)
			}
		}
	}
	b.WriteString("</svg>\n")
	return b.String()
}
