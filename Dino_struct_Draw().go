package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Object struct {
	texture rl.Texture2D
	source  rl.Rectangle
	dest    rl.Rectangle
	origin  rl.Vector2
	speed   rl.Vector2
}

func (o *Object) Draw() {
	rl.DrawTexturePro(o.texture, o.source, o.dest, o.origin, 0, rl.White)
}
