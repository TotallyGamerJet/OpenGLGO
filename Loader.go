package main

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"image"
	"fmt"

	"image/draw"
	_ "image/png"
	"github.com/gobuffalo/packr"
	"strings"
)

type Loader struct {
	vaos []*uint32
	vbos []*uint32
	textures []*uint32

}

func newLoader() Loader {
	return Loader{}
}

func (l *Loader) loadToVAO(positions, textureCoords, normals []float32, indices []int) RawModel {
	vaoID := l.createVAO()
	l.bindIndicesBuffer(indices)
	l.storeDataInAttributeList(0, 3, positions)
	l.storeDataInAttributeList(1, 2, textureCoords)
	l.storeDataInAttributeList(2, 3, normals)

	l.unbindVAO()
	return RawModel{vaoID, int32(len(indices))}
}

func (l *Loader) loadTexture(file string) uint32 {
	//imgFile, err := os.Open("res/" + file + ".png")
	box := packr.NewBox("./res")
	s, err := box.MustString(file + ".png")
	if err != nil {
		fmt.Errorf("texture %q not found on disk: %v", file, err)
		return 0
	}
	imgFile := strings.NewReader(s)
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		fmt.Errorf("unsupported stride")
		return 0
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	l.textures = append(l.textures, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT);
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture
}

func (l *Loader) createVAO() uint32{
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	l.vaos = append(l.vaos, &vao)
	gl.BindVertexArray(vao)
	return vao
}

func (l *Loader) storeDataInAttributeList(attribnum uint32, coordSize int32, data []float32) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)//bind
	l.vbos = append(l.vbos, &vbo)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	gl.VertexAttribPointer(attribnum, coordSize, gl.FLOAT, false, 0/**(5*4)*/, gl.PtrOffset(0))

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)//unbind
}

func (l *Loader) unbindVAO() {
	gl.BindVertexArray(0)
}

func (l *Loader) bindIndicesBuffer(indices []int) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)//bind
	l.vbos = append(l.vbos, &vbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
}

func (l *Loader) cleanUp() {
	for _, vao := range l.vaos {
		gl.DeleteVertexArrays(1, vao)
	}
	for _, vbo := range l.vbos {
		gl.DeleteBuffers(1, vbo)
	}
	for _, texture := range l.textures {
		gl.DeleteTextures(1, texture)
	}
}
