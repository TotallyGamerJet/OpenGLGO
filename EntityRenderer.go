package main

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type EntityRenderer struct {
	shader StaticShader
}

func newEntityRenderer(shader StaticShader, projectionMatrix mgl32.Mat4) EntityRenderer {
	shader.start()
	shader.loadProjectionMatrix(projectionMatrix)
	shader.stop()
	return EntityRenderer{shader}
}

func (r *EntityRenderer) render(entities map[TexturedModel][]Entity) {
	for model, batch := range entities {
		r.prepareTexturedModel(model)
		for _, entity := range batch {
			r.prepareInstance(entity)
			gl.DrawElements(gl.TRIANGLES, model.rawModel.vertexCount, gl.UNSIGNED_INT, gl.PtrOffset(0))
		}
		r.unbindTexturedModel()
	}
}

func (r *EntityRenderer) prepareTexturedModel(model TexturedModel) {
	rawModel := model.rawModel
	gl.BindVertexArray(rawModel.vaoID)
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.EnableVertexAttribArray(2)
	texture := model.texture
	if texture.hasTransparency {
		disableCulling()
	}
	r.shader.loadFakeLighting(texture.useFakeLighting)
	r.shader.loadShineVariables(texture.shineDamper, texture.reflectivity)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture.textureID)
}

func (r *EntityRenderer) unbindTexturedModel() {
	enableCulling()
	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
	gl.DisableVertexAttribArray(2)
	gl.BindVertexArray(0)
}

func (r *EntityRenderer) prepareInstance(entity Entity) {
	transformationMatrix := createTransformationMatrix(entity.position, entity.rotX, entity.rotY, entity.rotZ, entity.scale)
	r.shader.loadTransformationMatrix(transformationMatrix)
}

/*func (r *Renderer) render(entity Entity, shader StaticShader) {
	model := entity.model
	rawModel := model.rawModel
	gl.BindVertexArray(rawModel.vaoID)
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.EnableVertexAttribArray(2)

	transformationMatrix := createTransformationMatrix(entity.position, entity.rotX, entity.rotY, entity.rotZ, entity.scale)
	shader.loadTransformationMatrix(transformationMatrix)
	texture := model.texture
	shader.loadShineVariables(texture.shineDamper, texture.reflectivity)
	//gl.ActiveTexture(gl.TEXTURE0)
	//gl.BindTexture(gl.TEXTURE_2D, texturerdModel.texture.textureID)
	gl.DrawElements(gl.TRIANGLES, rawModel.vertexCount, gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
	gl.DisableVertexAttribArray(2)

	gl.BindVertexArray(0)
}*/