package main

import(
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const(
	FOV = 70
	NEAR_PLANE = 0.1
	FAR_PLANE = 1000
	RED = 0.5
	GREEN = 0.8
	BLUE = 0.9
)

type MasterRenderer struct {
	shader StaticShader
	renderer EntityRenderer

	terrainRenderer TerrainRenderer
	terrainShader TerrainShader

	entities map[TexturedModel][]Entity
	terrains []Terrain
}

func newMasterRenderer() MasterRenderer {
	gl.Enable(gl.CULL_FACE) //enableCulling()
	gl.CullFace(gl.BACK)

	shader := newStaticShader()
	terrainShader := newTerrainShader()

	projectionMatrix := createProjectionMatrix()

	renderer := newEntityRenderer(shader, projectionMatrix)
	terrainRenderer := newTerrainRenderer(terrainShader, projectionMatrix)

	return MasterRenderer{shader, renderer, terrainRenderer, terrainShader, make(map[TexturedModel][]Entity), make([]Terrain, 100)}
}

func enableCulling() {
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
}

func disableCulling() {
	gl.Disable(gl.CULL_FACE)
}

func (m *MasterRenderer) render(sun Light, camera Camera) {
	m.prepare()

	m.shader.start()
	m.shader.loadSkyColour(RED, GREEN, BLUE)
	m.shader.loadLight(sun)
	m.shader.loadViewMatrix(camera)
	m.renderer.render(m.entities)
	m.shader.stop()

	m.terrainShader.start()
	m.shader.loadSkyColour(RED, GREEN, BLUE)
	m.terrainShader.loadLight(sun)
	m.terrainShader.loadViewMatrix(camera)
	m.terrainRenderer.render(m.terrains)
	m.terrainShader.stop()

	m.terrains = nil
	m.entities = make(map[TexturedModel][]Entity)
}

func (m *MasterRenderer) processTerrain(terrain Terrain) {
	m.terrains = append(m.terrains, terrain)
}

func (m *MasterRenderer) processEntity(entity Entity) {
	entityModel := entity.model
	m.entities[entityModel] = append(m.entities[entityModel], entity)
}

func (m *MasterRenderer) cleanUp() {
	m.terrainShader.cleanUp()
	m.shader.cleanUp()
}

func (m *MasterRenderer) prepare() {
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(RED, GREEN, BLUE, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}


func createProjectionMatrix() mgl32.Mat4 {
	var projectionMatrix mgl32.Mat4
	projectionMatrix = mgl32.Ident4()
	aspectRatio := float32(WIDTH)/HEIGHT
	projectionMatrix = mgl32.Perspective(mgl32.DegToRad(FOV), aspectRatio, NEAR_PLANE, FAR_PLANE)
	return projectionMatrix
}