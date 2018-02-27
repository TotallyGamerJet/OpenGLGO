package main

import "github.com/go-gl/mathgl/mgl32"

type Light struct {
	position mgl32.Vec3
	colour mgl32.Vec3
}

func newLight(position, colour mgl32.Vec3) Light {
	return Light{position, colour}
}
