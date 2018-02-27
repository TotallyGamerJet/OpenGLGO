package main

type ModelData struct {
	vertices, textureCoords, normals []float32
	indices []int
	furthestPoint float32
}

func newModelData(vertices, textureCoords, normals []float32, indices []int, furthestPoint float32) ModelData {
	return ModelData{vertices, textureCoords, normals, indices,furthestPoint}
}