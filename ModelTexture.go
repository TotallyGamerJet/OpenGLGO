package main

type ModelTexture struct {
	textureID uint32
	shineDamper float32
	reflectivity float32
	hasTransparency bool
	useFakeLighting bool
}

func newModelTexture(id uint32) ModelTexture {
	return ModelTexture{id, 1, 0, false, false}
}

func (m *ModelTexture) setuseFakeLighting(useFakeLighting bool) {
	m.useFakeLighting = useFakeLighting
}


func (m *ModelTexture) setHasTransparency(hasTransparency bool) {
	m.hasTransparency = hasTransparency
}

func (m *ModelTexture) setShineDamper(damper float32) {
	m.shineDamper = damper
}

func (m *ModelTexture) setReflectivity(reflectivity float32) {
	m.reflectivity = reflectivity
}