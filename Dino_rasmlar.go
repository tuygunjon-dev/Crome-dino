package main

import rl "github.com/gen2brain/raylib-go/raylib"

func load_reset(texture rl.Texture2D) []Object {
	reset := Object{
		texture: texture,
		source:  rl.Rectangle{X: 0, Y: 0, Width: 73, Height: 94},
		dest:    rl.Rectangle{X: 700, Y: 50, Width: 88, Height: 94},
	}
	return []Object{reset}
}

func load_yer(texture rl.Texture2D) []Object {
	yer2 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 500, Y: 370, Width: 2000, Height: 20},
		dest:    rl.Rectangle{X: 1500, Y: 350, Width: 1500, Height: 20},
		speed:   rl.Vector2{X: 4, Y: 0},
	}
	yer1 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 500, Y: 370, Width: 2000, Height: 20},
		dest:    rl.Rectangle{X: 0, Y: 350, Width: 1500, Height: 20},
		speed:   rl.Vector2{X: 4, Y: 0},
	}
	return []Object{yer1, yer2}
}

func load_dinolar(texture rl.Texture2D) []Object {
	dino0 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 75, Y: 5, Width: 88, Height: 94},
		dest:    rl.Rectangle{X: 515, Y: 256, Width: 88, Height: 94},
		speed:   rl.Vector2{X: 0, Y: 0},
	}
	dino1 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 1338, Y: 0, Width: 88, Height: 94},
		dest:    rl.Rectangle{X: 300, Y: 256, Width: 88, Height: 94},
		speed:   rl.Vector2{X: 0, Y: 0},
	}
	dino2 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 1426, Y: 0, Width: 88, Height: 94},
		dest:    rl.Rectangle{X: 410, Y: 256, Width: 88, Height: 94},
		speed:   rl.Vector2{X: 0, Y: 0},
	}
	dino3 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 1514, Y: 0, Width: 88, Height: 94},
		dest:    rl.Rectangle{X: 515, Y: 256, Width: 88, Height: 94},
		speed:   rl.Vector2{X: 0, Y: 0},
	}

	dino4 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 1602, Y: 0, Width: 88, Height: 94},
		dest:    rl.Rectangle{X: 625, Y: 256, Width: 88, Height: 94},
		speed:   rl.Vector2{X: 0, Y: 0},
	}
	dino5 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 1690, Y: 0, Width: 88, Height: 94},
		dest:    rl.Rectangle{X: 725, Y: 256, Width: 88, Height: 94},
		speed:   rl.Vector2{X: 0, Y: 0},
	}
	dino6 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 1778, Y: 0, Width: 88, Height: 94},
		dest:    rl.Rectangle{X: 825, Y: 256, Width: 88, Height: 94},
		speed:   rl.Vector2{X: 0, Y: 0},
	}
	dino7p := Object{
		texture: texture,
		source:  rl.Rectangle{X: 1860, Y: 20, Width: 123, Height: 73},
		dest:    rl.Rectangle{X: dino3.dest.X, Y: 280, Width: 123, Height: 73},
		speed:   rl.Vector2{X: 0, Y: 0},
	}
	dino8p := Object{
		texture: texture,
		source:  rl.Rectangle{X: 1985, Y: 20, Width: 123, Height: 73},
		dest:    rl.Rectangle{X: dino3.dest.X, Y: 280, Width: 123, Height: 73},
		speed:   rl.Vector2{X: 0, Y: 0},
	}

	return []Object{dino0, dino1, dino2, dino3, dino4, dino5, dino6, dino7p, dino8p}
}

func load_kaktus_qush(texture rl.Texture2D) []Object {

	kaktuscha1 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 515, Y: 0, Width: 67, Height: 80},
		dest:    rl.Rectangle{X: 200, Y: 280, Width: 65, Height: 70},
		speed:   rl.Vector2{X: 4, Y: 0},
	}
	kaktuscha2 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 515, Y: 0, Width: 33, Height: 80},
		dest:    rl.Rectangle{X: 600, Y: 280, Width: 33, Height: 70},
		speed:   rl.Vector2{X: 4, Y: 0},
	}
	kaktus2 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 850, Y: 0, Width: 103, Height: 100},
		dest:    rl.Rectangle{X: 1000, Y: 256, Width: 88, Height: 94},
		speed:   rl.Vector2{X: 4, Y: 0},
	}
	kaktus1 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 803, Y: 0, Width: 47, Height: 100},
		dest:    rl.Rectangle{X: 1300, Y: 256, Width: 44, Height: 94},
		speed:   rl.Vector2{X: 4, Y: 0},
	}
	qush1 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 260, Y: 5, Width: 88, Height: 94},
		dest:    rl.Rectangle{X: 0, Y: 200, Width: 88, Height: 94},
		speed:   rl.Vector2{X: 4, Y: 0},
	}
	qush2 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 350, Y: 5, Width: 88, Height: 94},
		dest:    rl.Rectangle{X: 100, Y: 200, Width: 88, Height: 94},
		speed:   rl.Vector2{X: 4, Y: 0},
	}
	return []Object{kaktus1, kaktus2, kaktuscha1, kaktuscha2, qush1, qush2}
}

// func load_qushlar(texture rl.Texture2D) []Object {

// 	qush1 := Object{
// 		texture: texture,
// 		source:  rl.Rectangle{X: 260, Y: 5, Width: 88, Height: 94},
// 		dest:    rl.Rectangle{X: 0, Y: 200, Width: 88, Height: 94},
// 		speed:   rl.Vector2{X: 3, Y: 0},
// 	}
// 	qush2 := Object{
// 		texture: texture,
// 		source:  rl.Rectangle{X: 350, Y: 5, Width: 88, Height: 94},
// 		dest:    rl.Rectangle{X: 100, Y: 200, Width: 88, Height: 94},
// 		speed:   rl.Vector2{X: 3, Y: 0},
// 	}
// 	return []Object{qush1, qush2}
// }

func load_bulutlar(texture rl.Texture2D) []Object {

	bulut1 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 170, Y: 0, Width: 88, Height: 70},
		dest:    rl.Rectangle{X: 50, Y: 50, Width: 88, Height: 40},
		speed:   rl.Vector2{X: 1, Y: 0},
	}
	bulut2 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 170, Y: 0, Width: 88, Height: 70},
		dest:    rl.Rectangle{X: 500, Y: 50, Width: 110, Height: 70},
		speed:   rl.Vector2{X: 1, Y: 0},
	}
	bulut3 := Object{
		texture: texture,
		source:  rl.Rectangle{X: 170, Y: 0, Width: 88, Height: 70},
		dest:    rl.Rectangle{X: 1300, Y: 50, Width: 140, Height: 90},
		speed:   rl.Vector2{X: 1, Y: 0},
	}

	return []Object{bulut1, bulut2, bulut3}
}
