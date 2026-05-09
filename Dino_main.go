package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(1500, 500, "Chrome Dino(12.03.2026-)")

	rl.SetTargetFPS(60)

	texture := rl.LoadTexture("assets/sprite.png")
	defer rl.UnloadTexture(texture)
	BULUTLAR := load_bulutlar(texture)
	//QUSHLAR := load_qushlar(texture)
	KAKTUS_QUSH := load_kaktus_qush(texture)
	DINOLAR := load_dinolar(texture)
	//	RESET := load_reset(texture)
	YER := load_yer(texture)

	//
	var fremeCounterQush int
	var fremeCounterDinoP int
	var fremeCounterDinoT int
	//var TimeInterval float32 = float32(rl.GetRandomValue(10, 15)) / 10
	var SpawnDistance float32 = 750
	SpawnDistance -= 5
	for !rl.WindowShouldClose() {
		fremeCounterQush++
		fremeCounterDinoT++
		fremeCounterDinoP++

		if fremeCounterQush%20 < 10 {
			KAKTUS_QUSH[4].source.X = 260
		} else {
			KAKTUS_QUSH[4].source.X = 350
		}
		//
		if rl.IsKeyDown(rl.KeyDown) {
			if DINOLAR[3].dest.Y >= DINOLAR[0].dest.Y-1 { ///  ishlamadi nimaga????
				if fremeCounterDinoP%20 < 10 {
					DINOLAR[8].source.X = 1860
				} else {
					DINOLAR[8].source.X = 1985
				}
			}
		}
		if fremeCounterDinoT%20 < 10 {
			DINOLAR[3].source.X = 1514
		} else {
			DINOLAR[3].source.X = 1602
		}

		if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyUp) {
			if DINOLAR[3].dest.Y+DINOLAR[3].dest.Height <= YER[0].dest.Y {
				DINOLAR[3].speed.Y = -10
			}
		}
		if DINOLAR[3].dest.Y+DINOLAR[3].dest.Height >= YER[0].dest.Y {
			DINOLAR[3].speed.Y = 0
			DINOLAR[3].dest.Y = YER[0].dest.Y - DINOLAR[3].dest.Height - 1
		}
		DINOLAR[3].speed.Y += 0.4
		DINOLAR[3].dest.Y += DINOLAR[3].speed.Y

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for i := range BULUTLAR {
			BULUTLAR[i].MoveX()
			BULUTLAR[i].LoopX(1500)
			BULUTLAR[i].Draw()
		}

		if SpawnDistance <= 0 {
			// random object tanlash
			index := rl.GetRandomValue(0, int32(len(KAKTUS_QUSH)-2))

			KAKTUS_QUSH[index].dest.X = 1500 + float32(rl.GetRandomValue(0, 600))

			// yangi masofa beramiz
			SpawnDistance = float32(rl.GetRandomValue(600, 1000))
		}
		for i := range KAKTUS_QUSH {
			KAKTUS_QUSH[i].MoveX()
			KAKTUS_QUSH[i].GroundLoop(1500)
			if i != 5 {
				KAKTUS_QUSH[i].Draw()
			}
		}
		for i := range DINOLAR {
			//if i == 3 || i == 8 {
			//DINOLAR[i].Draw()
			//}
			if rl.IsKeyDown(rl.KeyDown) {
				i = 8
			} else if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyUp) {
				i = 3
			} else {
				i = 3
			}
			DINOLAR[i].Draw()
		}
		// for i := range RESET {
		// 	// Dino_rasmlar_reset[i].MoveX()
		// 	RESET[i].Draw()
		// }
		for i := range YER {
			YER[i].MoveX()
			YER[i].GroundLoop(1500)
			YER[i].Draw()
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
