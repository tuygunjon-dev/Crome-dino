package main

import (
	"math/rand"
)

func (o *Object) MoveX() {
	o.dest.X -= o.speed.X
}

func (o *Object) LoopX(screenwidth float32) {
	if o.dest.X+o.dest.Width < 0 {
		o.dest.X = screenwidth + rand.Float32()*400
		// o.dest.Y = 40 + rand.Float32()*80
	}
}

//var TimeInterval float32 = float32(rl.GetRandomValue(10, 15)) / 10

func (o *Object) GroundLoop(screenwidth float32) {
	if o.dest.X+o.dest.Width < 0 {
		o.dest.X = screenwidth
	}
}

// /  GroundLoop2 kaktus va qushlar u.n
func (o *Object) GroundLoop2(screenwidth float32) {
	if o.dest.X+o.dest.Width < 0 {

		o.dest.X = screenwidth

		//o.dest.X = screenwidth + rand.Float32()*500
	}
}
