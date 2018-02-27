package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"strconv"
	"os"
	"bufio"
	"strings"
	"io"
)

func loadOBJModel(filename string, loader Loader) RawModel {
	f, err := os.Open(filename + ".obj")
	check(err)
	defer f.Close()

	r := bufio.NewReader(f)

	var line string
	var currentLine []string
	var s []byte

	var vertices []mgl32.Vec3
	var textures []mgl32.Vec2
	var normals []mgl32.Vec3
	var indices []int
	var verticesArray []float32
	var textureArray []float32
	var normalsArray []float32

	var x, y, z float64
	var vertex1, vertex2, vertex3 []string

	for {
		s, _, err = r.ReadLine()
		check(err)
		line = string(s[:])
		currentLine = strings.Split(line, " ")

		if strings.HasPrefix(line, "v ") {
			x, err = strconv.ParseFloat(currentLine[1], 64)
			check(err)
			y, err = strconv.ParseFloat(currentLine[2], 64)
			check(err)
			z, err = strconv.ParseFloat(currentLine[3], 64)
			check(err)

			vertices = append(vertices, mgl32.Vec3{float32(x), float32(y), float32(z)})
		} else if strings.HasPrefix(line, "vt ") {
			x, err = strconv.ParseFloat(currentLine[1], 64)
			check(err)
			y, err = strconv.ParseFloat(currentLine[2], 64)
			check(err)

			textures = append(textures, mgl32.Vec2{float32(x), float32(y)})
		} else if strings.HasPrefix(line, "vn ") {
			x, err = strconv.ParseFloat(currentLine[1], 64)
			check(err)
			y, err = strconv.ParseFloat(currentLine[2], 64)
			check(err)
			z, err = strconv.ParseFloat(currentLine[3], 64)
			check(err)

			normals = append(normals, mgl32.Vec3{float32(x), float32(y), float32(z)})
		} else if strings.HasPrefix(line, "f ") {
			textureArray = make([]float32, len(vertices) * 2)
			normalsArray = make([]float32, len(vertices) * 3)
			break
		}
	}

	for err != io.EOF {
		if !strings.HasPrefix(line, "f ") {
			s, _, err = r.ReadLine()
			check(err)
			line = string(s[:])
			continue
		}
		currentLine = strings.Split(line, " ")

		vertex1 = strings.Split(currentLine[1], "/")
		vertex2 = strings.Split(currentLine[2], "/")
		vertex3 = strings.Split(currentLine[3], "/")

		indices, textureArray, normalsArray = processVertex2(vertex1, indices, textures, normals, textureArray, normalsArray)
		indices, textureArray, normalsArray = processVertex2(vertex2, indices, textures, normals, textureArray, normalsArray)
		indices, textureArray, normalsArray = processVertex2(vertex3, indices, textures, normals, textureArray, normalsArray)
		s, _, err = r.ReadLine()
		line = string(s[:])
	}

	for _, vertex := range vertices {
		verticesArray = append(verticesArray, vertex.X(), vertex.Y(), vertex.Z())
	}

	return loader.loadToVAO(verticesArray, textureArray, normalsArray, indices)
}

func processVertex2(vertexData []string , indices []int, textures []mgl32.Vec2,
	normals[]mgl32.Vec3,  textureArray,  normalsArray []float32) ([]int, []float32, []float32) {
	currentVertexPointer, err := strconv.ParseInt(vertexData[0], 10, 64)
	check(err)
	currentVertexPointer--
	indices = append(indices, int(currentVertexPointer))
	texCoord, err := strconv.ParseInt(vertexData[1], 10, 64)
	check(err)
	currentTex := textures[texCoord - 1]
	textureArray[currentVertexPointer * 2] = currentTex.X()
	textureArray[currentVertexPointer * 2 + 1] = 1 - currentTex.Y()
	normCoord, err := strconv.ParseInt(vertexData[2], 10,64)
	check(err)
	currentNorm := normals[normCoord - 1]
	normalsArray[currentVertexPointer * 3] = currentNorm.X()
	normalsArray[currentVertexPointer * 3 + 1] = currentNorm.Y()
	normalsArray[currentVertexPointer * 3 + 2] = currentNorm.Z()
	return indices, textureArray, normalsArray
}
