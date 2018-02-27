package main

type RawModel struct {
	vaoID uint32
	vertexCount int32
}

func newRawModel(vaoID uint32, vertexCount int32) RawModel {
	return RawModel{vaoID, vertexCount}
}