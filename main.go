package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/OlegZapara/asd-2/alg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const squareSize = 250

var (
	game = &Game{
		matrix:   [alg.N][alg.N]int{},
		finished: false,
	}
	blackColor  = color.RGBA{R: 119, G: 149, B: 86, A: 255}
	whiteColor  = color.RGBA{R: 235, G: 236, B: 208, A: 255}
	finishColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	queen       = Queen{}
	start       = make(chan bool)
)

type Game struct {
	matrix   [alg.N][alg.N]int
	mode     string
	finished bool
	delay    int
}

type Queen struct {
	image *ebiten.Image
	scale float64
}

func (q *Queen) Draw(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(q.scale, q.scale)
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(q.image, op)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		select {
		case start <- true:
		default:
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y, row := range g.matrix {
		for x, cell := range row {
			c := whiteColor
			if (x+y)%2 == 1 {
				c = blackColor
			}
			if g.finished && g.matrix[y][x] == 1 {
				c = finishColor
			}
			vector.DrawFilledRect(screen, float32(x*squareSize), float32(y*squareSize), float32(squareSize), float32(squareSize), c, false)

			if cell == 1 {
				queen.Draw(screen, x*squareSize, y*squareSize)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return len(g.matrix[0]) * squareSize, len(g.matrix) * squareSize
}

func (g *Game) solve() {
	<-start
	pos := [2]int{0, 0}
	delay := time.Duration(g.delay) * time.Millisecond
	switch g.mode {
	case "dfs":
		fmt.Println("Starting LDFS")
		alg.LDFS(&g.matrix, 0, pos, delay)
	case "bfs":
		fmt.Println("Starting RBFS")
		alg.RBFS(&g.matrix, 0, 0, pos, delay)
	default:
		log.Fatal("Invalid mode")
	}
	fmt.Printf("Result: F1 = %d - number of queens attacking each other\n", alg.F1(&g.matrix))
	g.finished = true
}

func main() {
	mode := flag.String("mode", "dfs", "mode to solve the problem (dfs or bfs)")
	delay := flag.Int("delay", 5, "delay between steps in milliseconds")
	flag.Parse()
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <mode>")
	}
	game.mode = *mode
	game.delay = *delay
	queen.image, _, _ = ebitenutil.NewImageFromFile("queen.png")
	queen.scale = float64(squareSize) / float64(queen.image.Bounds().Dx())

	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("8 Queens Problem")

	go game.solve()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
