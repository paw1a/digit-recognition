package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/paw1a/digit-recognition/internal/dataset"
	"github.com/paw1a/digit-recognition/pkg/model"
	"image/color"
	"math"
	"os"
	"path"
	"strconv"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720

	RectSize     = 420
	CursorRadius = 15

	ModelDirectory = "model"
	ModelFilename  = "model.gob"
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

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}

func TrainModel(mod *model.Model) {
	if !dataset.DatasetExists() {
		err := dataset.DownloadDataset()
		if err != nil {
			fmt.Printf("failed to download dataset: %v\n", err)
			return
		}
	}

	err := os.MkdirAll(ModelDirectory, os.ModePerm)
	if err != nil {
		fmt.Printf("can't create model directory: %v", err)
		return
	}

	input, labels, err := dataset.LoadDataset(
		path.Join(dataset.Directory, dataset.TrainImages),
		path.Join(dataset.Directory, dataset.TrainLabels))

	if err != nil {
		fmt.Printf("failed to load dataset: %v\n", err)
		return
	}

	mod.Fit(input, labels)

	err = model.SerializeModel(path.Join(ModelDirectory, ModelFilename), mod)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	mod := new(model.Model)

	if _, err := os.Stat(path.Join(ModelDirectory, ModelFilename)); err != nil {
		mod = model.NewModel([]int{784, 250, 10}, 0.001, 5)
		go TrainModel(mod)
	} else {
		err := model.DeserializeModel(path.Join(ModelDirectory, ModelFilename), mod)
		if err != nil {
			fmt.Printf("failed to load model, new model was created: %v\n", err)
			mod = model.NewModel([]int{784, 250, 10}, 0.001, 5)
			go TrainModel(mod)
		}
	}

	digit := ""
	output := make([]float64, 10)

	rl.SetTraceLog(rl.LogNone)
	rl.InitWindow(ScreenWidth, ScreenHeight, "Digit Recognition")

	rl.SetTargetFPS(60)

	pixels := make([]color.RGBA, RectSize*RectSize)
	image := rl.NewImage(make([]byte, RectSize*RectSize*4),
		RectSize, RectSize, 1, rl.UncompressedR8g8b8a8)
	texture := rl.LoadTextureFromImage(image)

	rl.UpdateTexture(texture, pixels)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(DarkGrayColor)

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			mousePrevX := float32(rl.GetMouseX()) - rl.GetMouseDelta().X
			mousePrevY := float32(rl.GetMouseY()) - rl.GetMouseDelta().Y

			for k := 0; k < 50; k++ {
				x := int(mousePrevX + float32(k)*rl.GetMouseDelta().X/50 - 50)
				y := int(mousePrevY + float32(k)*rl.GetMouseDelta().Y/50 - 50)

				if x > 0 && x < RectSize && y > 0 && y < RectSize {
					for i := max(y-CursorRadius, 0); i < min(y+CursorRadius, RectSize-1); i++ {
						for j := max(x-CursorRadius, 0); j < min(x+CursorRadius, RectSize-1); j++ {
							if (x-j)*(x-j)+(y-i)*(y-i) < CursorRadius*CursorRadius {
								pixels[i*RectSize+j] = WhiteColor
							}
						}
					}
				}
			}
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			pixels = make([]color.RGBA, RectSize*RectSize)
		}

		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			sum := make([]float64, 28*28)
			for i := 0; i < RectSize*RectSize; i++ {
				y := i / (RectSize * 15)
				x := (i % RectSize) / 15

				if pixels[i].R > 0 {
					sum[y*28+x] += 1
				}
			}

			for i := 0; i < 28*28; i++ {
				sum[i] = float64(sum[i]) / (15 * 15)
			}

			resultDigit, resultOutput := mod.PredictDigit(sum)
			digit = strconv.Itoa(resultDigit)
			output = resultOutput
		}

		rl.UpdateTexture(texture, pixels)
		rl.DrawTexture(texture, 50, 50, WhiteColor)

		rl.DrawRectangleLinesEx(rl.Rectangle{
			X:      50,
			Y:      50,
			Width:  RectSize,
			Height: RectSize,
		}, 7, GrayColor)

		rl.DrawRectangleLinesEx(rl.Rectangle{
			X:      50,
			Y:      70 + RectSize,
			Width:  RectSize / 2,
			Height: RectSize / 2,
		}, 7, GrayColor)

		for i := 0; i < 10; i++ {
			rl.DrawText(strconv.Itoa(i), RectSize+500, int32(55+i*40), 45, GrayColor)
			rl.DrawRectangle(RectSize+550, int32(55+i*40), int32(math.Max(output[i]*300, 5)), 40, GrayColor)
		}

		rl.DrawText("OUTPUT:", 70+RectSize/2, 70+RectSize, 40, GrayColor)
		index, _ := strconv.Atoi(digit)
		rl.DrawText(fmt.Sprintf("%.2f%%", output[index]*100),
			70+RectSize/2, 110+RectSize, 40, GrayColor)

		rl.DrawText(digit, 100, 80+RectSize, RectSize/2, WhiteColor)

		rl.DrawText("DRAW DIGIT", 50, 10, 40, GrayColor)
		rl.DrawText("PROBABILITIES", RectSize+500, 10, 40, GrayColor)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
