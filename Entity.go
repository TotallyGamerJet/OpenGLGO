package main

import "github.com/go-gl/mathgl/mgl32"

type Entity struct {
	model TexturedModel
	position mgl32.Vec3
	rotX, rotY, rotZ, scale float32
}

func newEntity(model TexturedModel, position mgl32.Vec3, rotX, rotY, rotZ, scale float32) Entity {
	return Entity{model, position, rotX, rotY, rotZ, scale}
}

func (e *Entity) increasePosition(dx, dy, dz float32) {
	e.position = e.position.Add(mgl32.Vec3{dx, dy, dz})
}

func (e *Entity) increaseRotation(dx, dy, dz float32) {
	e.rotX += dx
	e.rotY += dy
	e.rotZ += dz
}