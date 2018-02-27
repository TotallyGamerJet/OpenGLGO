package main

import (
	"runtime"
	"github.com/go-gl/mathgl/mgl32"
	"math/rand"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	win := newDisplay()
	display := Display{win}
	defer win.Destroy()

	loader := newLoader()
	defer loader.cleanUp()

	grassTextureModel := newTexturedModel(loadOBJModel("grassModel", loader), newModelTexture(loader.loadTexture("grass")))
	grassTextureModel.texture.setHasTransparency(true)
	grassTextureModel.texture.setuseFakeLighting(true)

	fernTextureModel := newTexturedModel(loadOBJModel("fern", loader), newModelTexture(loader.loadTexture("fern")))
	fernTextureModel.texture.setHasTransparency(true)

	var entityList []Entity
	for i:=0;i<100;i++ {
		entityList = append(entityList, newEntity(grassTextureModel, mgl32.Vec3{float32(rand.Intn(300)), 0, float32(-rand.Intn(300))}, 0, 0, 0, 1))
		entityList = append(entityList, newEntity(fernTextureModel, mgl32.Vec3{float32(rand.Intn(300)), 0, float32(-rand.Intn(300))}, 0, 0, 0, 1))
	}

	light := newLight(mgl32.Vec3{0,1000,-5}, mgl32.Vec3{1,1,1})

	terrain := newTerrain(0, 0, loader, newModelTexture(loader.loadTexture("grassy2")))
	terrain2 := newTerrain(-1, 0, loader, newModelTexture(loader.loadTexture("grassy2")))
	terrain3 := newTerrain(-1, -1, loader, newModelTexture(loader.loadTexture("grassy2")))
	terrain4 := newTerrain(0, -1, loader, newModelTexture(loader.loadTexture("grassy2")))

	camera := newCamera()

	renderer := newMasterRenderer()
	defer renderer.cleanUp()

	for !display.window.ShouldClose() {
		processInput(display.window)
		camera.move()

		for _, entity := range entityList {
			renderer.processEntity(entity)
		}
		renderer.processTerrain(terrain)
		renderer.processTerrain(terrain2)
		renderer.processTerrain(terrain3)
		renderer.processTerrain(terrain4)

		renderer.render(light, camera)
		display.updateDisplay()
	}

}

func check(err error) {
	if err !=  nil {
		panic(err)
	}
}