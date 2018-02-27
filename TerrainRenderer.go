package main

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type TerrainRenderer struct {
	shader TerrainShader
}

func newTerrainRenderer(shader TerrainShader, projectionMatrix mgl32.Mat4) TerrainRenderer {
	shader.start()
	shader.loadProjectionMatrix(projectionMatrix)
	shader.stop()
	return TerrainRenderer{shader}
}

func (t *TerrainRenderer) render(terrains []Terrain) {
	for _, terrain := range terrains {
		t.prepareTerrain(terrain)
		t.loadModelMatrix(terrain)
		gl.DrawElements(gl.TRIANGLES, terrain.model.vertexCount, gl.UNSIGNED_INT, gl.PtrOffset(0))
		t.unbindTexturedModel()
	}
}

func (t *TerrainRenderer) prepareTerrain(terrain Terrain) {
	rawModel := terrain.model
	gl.BindVertexArray(rawModel.vaoID)
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.EnableVertexAttribArray(2)
	texture := terrain.texture
	t.shader.loadShineVariables(texture.shineDamper, texture.reflectivity)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture.textureID)
}

func (t *TerrainRenderer) unbindTexturedModel() {
	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
	gl.DisableVertexAttribArray(2)
	gl.BindVertexArray(0)
}

func (t *TerrainRenderer) loadModelMatrix(terrain Terrain) {
	transformationMatrix := createTransformationMatrix(mgl32.Vec3{terrain.x, 0, terrain.z}, 0, 0, 0, 1)
	t.shader.loadTransformationMatrix(transformationMatrix)
}