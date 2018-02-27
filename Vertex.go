package main

import "github.com/go-gl/mathgl/mgl32"

const (
	NO_INDEX = -1
)

type Vertex struct {
	position mgl32.Vec3
	textureIndex int
	normalIndex int
	//duplicateVertex Vertex NOT ALLOWED TO HAVE RECURSION
	index int
	length float32
}

func newVertex(index int, position mgl32.Vec3) Vertex {
	return Vertex{position, NO_INDEX, NO_INDEX, /*nil,*/ index, position.Len()}
}

func (v *Vertex) isSet() bool {
	return v.textureIndex!=NO_INDEX && v.normalIndex!=NO_INDEX
}

func (v *Vertex) hasSameTextureAndNormal(textureIndexOther, normalIndexOther int) bool {
	return textureIndexOther==v.textureIndex && normalIndexOther==v.normalIndex
}

func (v *Vertex) setTextureIndex(textureIndex int) {
	v.textureIndex = textureIndex
}

func (v *Vertex) setNormalIndex(normalIndex int) {
	v.normalIndex = normalIndex
}

func (v *Vertex) setDuplicateVertex(duplicateVertex Vertex) {
	//v.duplicateVertex = duplicateVertex
}