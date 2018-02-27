package main

const (
	SIZE = 800
	VERTEX_COUNT = 128
)

type Terrain struct {
	x, z float32
	model RawModel
	texture ModelTexture
}

func newTerrain(gridX, gridZ int, loader Loader, texture ModelTexture) Terrain {
	x := float32(gridX * SIZE)
	z := float32(gridZ * SIZE)
	model := generateTerrain(loader)
	return Terrain{x, z, model, texture}
}

func generateTerrain(loader Loader) RawModel {
	count := VERTEX_COUNT * VERTEX_COUNT
	vertices := make([]float32, count * 3)
	normals := make([]float32, count * 3)
	textureCoords := make([]float32, count * 2)
	indices := make([]int, 6*(VERTEX_COUNT-1)*(VERTEX_COUNT-1))
	vertexPointer := 0
	for i := 0; i < VERTEX_COUNT; i++ {
		for j := 0; j < VERTEX_COUNT; j++ {
			vertices[vertexPointer*3] = float32(j)/float32(VERTEX_COUNT-1) * SIZE
			vertices[vertexPointer*3+1] = 0
			vertices[vertexPointer*3+2] = float32(i)/float32(VERTEX_COUNT-1) * SIZE
			normals[vertexPointer*3] = 0
			normals[vertexPointer*3+1] = 1
			normals[vertexPointer*3+2] = 0
			textureCoords[vertexPointer*2] = float32(j)/float32(VERTEX_COUNT-1)
			textureCoords[vertexPointer*2+1] = float32(i)/float32(VERTEX_COUNT-1)
			vertexPointer++
		}
	}
	pointer := 0
	for gz := 0; gz < VERTEX_COUNT - 1; gz++ {
		for gx := 0; gx <VERTEX_COUNT - 1; gx++ {
			topLeft := (gz*VERTEX_COUNT)+gx
			topRight := topLeft + 1
			bottomLeft := ((gz+1)*VERTEX_COUNT)+gx
			bottomRight := bottomLeft + 1
			indices[pointer] = topLeft
			pointer++
			indices[pointer] = bottomLeft
			pointer++
			indices[pointer] = topRight
			pointer++
			indices[pointer] = topRight
			pointer++
			indices[pointer] = bottomLeft
			pointer++
			indices[pointer] = bottomRight
			pointer++
		}
	}
	return loader.loadToVAO(vertices, textureCoords, normals, indices)
}
