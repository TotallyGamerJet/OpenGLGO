package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

func toRadian(degree float32) float32 {
	return degree * math.Pi / 180
}

func createTransformationMatrix(translation mgl32.Vec3, rx, ry, rz, scale float32) mgl32.Mat4 {
	matrix := mgl32.Ident4()
	matrix = matrix.Mul4(mgl32.Translate3D(translation.X(), translation.Y(), translation.Z()))
	matrix = matrix.Mul4(mgl32.HomogRotate3DX(toRadian(rx)))
	matrix = matrix.Mul4(mgl32.HomogRotate3DY(toRadian(ry)))
	matrix = matrix.Mul4(mgl32.HomogRotate3DZ(toRadian(rz)))
	matrix = matrix.Mul4(mgl32.Scale3D(scale, scale, scale))
	return matrix
}

func createViewMatrix(c Camera) mgl32.Mat4 {
	matrix := mgl32.Ident4()
	matrix = matrix.Mul4(mgl32.HomogRotate3DX(toRadian(c.pitch)))
	matrix = matrix.Mul4(mgl32.HomogRotate3DY(toRadian(c.yaw)))
	matrix = matrix.Mul4(mgl32.Translate3D(-c.position.X(), -c.position.Y(), -c.position.Z()))
	return matrix
}