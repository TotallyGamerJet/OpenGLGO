package main

type TexturedModel struct {
	rawModel RawModel
	texture ModelTexture
}

func newTexturedModel(model RawModel, texture ModelTexture) TexturedModel {
	return TexturedModel{model, texture}
}
