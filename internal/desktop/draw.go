package desktop

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/paw1a/digit-recognition/pkg/algebra"
	"image/color"
	"math"
	"strconv"
)

type Canvas struct {
	app *Application

	texture rl.Texture2D
	pixels  []color.RGBA

	currentDrawDigit   string
	currentModelOutput []float64
}

func (c *Canvas) DrawPointWithMouse() {
	mousePrevX := float32(rl.GetMouseX()) - rl.GetMouseDelta().X
	mousePrevY := float32(rl.GetMouseY()) - rl.GetMouseDelta().Y

	for k := 0; k < 50; k++ {
		x := int(mousePrevX + float32(k)*rl.GetMouseDelta().X/50 - 50)
		y := int(mousePrevY + float32(k)*rl.GetMouseDelta().Y/50 - 50)

		if x > 0 && x < RectSize && y > 0 && y < RectSize {
			for i := algebra.Max(y-CursorRadius, 0); i < algebra.Min(y+CursorRadius, RectSize-1); i++ {
				for j := algebra.Max(x-CursorRadius, 0); j < algebra.Min(x+CursorRadius, RectSize-1); j++ {
					if (x-j)*(x-j)+(y-i)*(y-i) < CursorRadius*CursorRadius {
						c.pixels[i*RectSize+j] = WhiteColor
					}
				}
			}
		}
	}
}

func (c *Canvas) DrawPaintArea() {
	rl.UpdateTexture(c.texture, c.pixels)
	rl.DrawTexture(c.texture, 50, 50, WhiteColor)

	rl.DrawRectangleLinesEx(rl.Rectangle{
		X:      50,
		Y:      50,
		Width:  RectSize,
		Height: RectSize,
	}, 7, GrayColor)
}

func (c *Canvas) DrawRecognitionArea() {
	rl.DrawRectangleLinesEx(rl.Rectangle{
		X:      50,
		Y:      70 + RectSize,
		Width:  RectSize / 2,
		Height: RectSize / 2,
	}, 7, GrayColor)

	digit, _ := strconv.Atoi(c.currentDrawDigit)

	rl.DrawText("OUTPUT:", 70+RectSize/2, 70+RectSize, 40, GrayColor)
	rl.DrawText(fmt.Sprintf("%.2f%%", c.currentModelOutput[digit]*100),
		70+RectSize/2, 110+RectSize, 40, GrayColor)

	rl.DrawText(c.currentDrawDigit, 100, 80+RectSize, RectSize/2, WhiteColor)
}

func (c *Canvas) DrawProbabilitiesArea() {
	rl.DrawText("DRAW DIGIT", 50, 10, 40, GrayColor)
	rl.DrawText("PROBABILITIES", RectSize+400, 10, 40, GrayColor)

	for i := 0; i < 10; i++ {
		rl.DrawText(strconv.Itoa(i), RectSize+400, int32(55+i*40), 45, GrayColor)
		rl.DrawRectangle(RectSize+450, int32(55+i*40),
			int32(math.Max(c.currentModelOutput[i]*300, 5)), 40, GrayColor)
	}
}

func (c *Canvas) DrawTrainingArea() {
	rl.DrawText("TRAINING", RectSize+400, 70+RectSize, 40, GrayColor)
	if c.app.mod.Trained {
		rl.DrawText("DONE", RectSize+400, 110+RectSize, 40, GrayColor)
		rl.DrawRectangle(RectSize+400, RectSize+150, 400, 40, GrayColor)
	} else {
		rl.DrawText(fmt.Sprintf("EPOCH: %d", c.app.mod.TrainingState.CurrentEpoch+1),
			RectSize+400, 110+RectSize, 40, GrayColor)
		progressBarSize := int32(400 * (float32(c.app.mod.TrainingState.CurrentIteration) /
			float32(c.app.mod.TrainingState.DatasetSize)))

		rl.DrawRectangle(RectSize+400, RectSize+150, progressBarSize, 40, GrayColor)
	}
}

func NewCanvas(app *Application) *Canvas {
	canvas := Canvas{
		app:                app,
		pixels:             make([]color.RGBA, RectSize*RectSize),
		currentDrawDigit:   "",
		currentModelOutput: make([]float64, 10),
	}

	image := rl.NewImage(make([]byte, RectSize*RectSize*4),
		RectSize, RectSize, 1, rl.UncompressedR8g8b8a8)
	canvas.texture = rl.LoadTextureFromImage(image)

	rl.UpdateTexture(canvas.texture, canvas.pixels)

	return &canvas
}
