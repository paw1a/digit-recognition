package desktop

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/paw1a/digit-recognition/pkg/model"
	"image/color"
	"strconv"
)

type Application struct {
	mod    *model.Model
	canvas *Canvas
}

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
	FPS          = 60

	RectSize     = 420
	CursorRadius = 15
)

var (
	GrayColor = color.RGBA{
		R: 180,
		G: 180,
		B: 180,
		A: 255,
	}
	WhiteColor = color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	DarkGrayColor = color.RGBA{
		R: 50,
		G: 50,
		B: 50,
		A: 255,
	}
)

func (app *Application) Run() {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(DarkGrayColor)

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			app.canvas.DrawPointWithMouse()
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			app.canvas.pixels = make([]color.RGBA, RectSize*RectSize)
		}

		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			app.recognizeDigit()
		}

		app.canvas.DrawPaintArea()
		app.canvas.DrawRecognitionArea()
		app.canvas.DrawProbabilitiesArea()
		app.canvas.DrawTrainingArea()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func (app *Application) recognizeDigit() {
	sum := make([]float64, 28*28)
	for i := 0; i < RectSize*RectSize; i++ {
		y := i / (RectSize * 15)
		x := (i % RectSize) / 15

		if app.canvas.pixels[i].R > 0 {
			sum[y*28+x] += 1
		}
	}

	for i := 0; i < 28*28; i++ {
		sum[i] = float64(sum[i]) / (15 * 15)
	}

	resultDigit, resultOutput := app.mod.PredictDigit(sum)
	app.canvas.currentDrawDigit = strconv.Itoa(resultDigit)
	app.canvas.currentModelOutput = resultOutput
}

func NewApplication(mod *model.Model) *Application {
	rl.SetTraceLog(rl.LogNone)
	rl.InitWindow(ScreenWidth, ScreenHeight, "Digit Recognition")

	rl.SetTargetFPS(FPS)

	app := Application{
		mod: mod,
	}
	app.canvas = NewCanvas(&app)

	return &app
}
