package main

import "github.com/go-gl/mathgl/mgl32"

type Camera struct {
	position mgl32.Vec3
	pitch, yaw float32
}

func newCamera() Camera {
	return Camera{mgl32.Vec3{0,3,0}, 0,0}
}

func (c *Camera) move() {
	//MOVE CAMERA & INPUT
	c.position = c.position.Add(mgl32.Vec3{cameraXTemp, cameraYTemp, cameraZTemp})
	cameraXTemp = 0
	cameraYTemp = 0
	cameraZTemp = 0
}