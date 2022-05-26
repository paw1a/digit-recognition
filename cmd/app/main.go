package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"math/rand"
	"strconv"
)

const (
	ScreenWidth  = 1080
	ScreenHeight = 720

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

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Digit Recognition")

	rl.SetTargetFPS(60)

	pixels := make([]color.RGBA, RectSize*RectSize)
	image := rl.NewImage(make([]byte, RectSize*RectSize),
		RectSize, RectSize, 1, rl.UncompressedR8g8b8a8)
	texture := rl.LoadTextureFromImage(image)

	rl.UpdateTexture(texture, pixels)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

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

		if rl.IsKeyPressed(rl.KeyEnter) {
			sum := make([]int, 28*28)
			for i := 0; i < RectSize*RectSize; i++ {
				y := i / (RectSize * 15)
				x := (i % RectSize) / 15

				if pixels[i].R > 0 {
					sum[y*28+x] += 1
				}
			}

			for i := 0; i < 28*28; i++ {
				sum[i] = int(255 * (float32(sum[i]) / (15 * 15)))
			}

			for i := 0; i < RectSize*RectSize; i++ {
				y := i / (RectSize * 15)
				x := (i % RectSize) / 15

				col := uint8(sum[y*28+x])
				pixels[i] = color.RGBA{
					R: col,
					G: col,
					B: col,
					A: 255,
				}
			}
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
			rl.DrawText(strconv.Itoa(i), RectSize+200, int32(50+i*50), 50, GrayColor)
			rl.DrawRectangle(RectSize+250, int32(50+i*50), int32(rand.Intn(300)), 42, GrayColor)
		}

		rl.DrawText("0", 100, 80+RectSize, RectSize/2, WhiteColor)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
