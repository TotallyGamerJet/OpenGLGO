package main

import (
	"os"
	"github.com/go-gl/mathgl/mgl32"
	"bufio"
	"strings"
	"strconv"
	"io"
)

const (
	res_loc = "res/"
)

func readLine(r *bufio.Reader) string {
	var line string
	s, _, err := r.ReadLine()
	check(err)
	line = string(s[:])
	return line
}

func loadOBJ(objFileName string) ModelData {
	f, err := os.Open(res_loc + objFileName + ".obj")
	check(err)
	defer f.Close()

	r := bufio.NewReader(f)


	var line string
	var s []byte
	var vertices []Vertex
	var textures []mgl32.Vec2
	var normals []mgl32.Vec3
	var indices []int

	for {
		s, _, err = r.ReadLine()
		check(err)
		line = string(s[:])
		if strings.HasPrefix(line, "v ") {
			currentLine := strings.Split(line, " ")
			x, err := strconv.ParseFloat(currentLine[1], 64)
			check(err)
			y, err := strconv.ParseFloat(currentLine[2], 64)
			check(err)
			z, err := strconv.ParseFloat(currentLine[3], 64)
			check(err)
			vertex := mgl32.Vec3{float32(x), float32(y), float32(z)}
			newVertex := newVertex(len(vertices), vertex)
			vertices = append(vertices, newVertex)
		} else if strings.HasPrefix(line, "vt ") {
			currentLine := strings.Split(line, " ")
			x, err := strconv.ParseFloat(currentLine[1], 64)
			check(err)
			y, err := strconv.ParseFloat(currentLine[2], 64)
			check(err)
			texture := mgl32.Vec2{float32(x), float32(y)}
			textures = append(textures, texture)
		} else if strings.HasPrefix(line, "vn ") {
			currentLine := strings.Split(line, " ")
			x, err := strconv.ParseFloat(currentLine[1], 64)
			check(err)
			y, err := strconv.ParseFloat(currentLine[2], 64)
			check(err)
			z, err := strconv.ParseFloat(currentLine[3], 64)
			check(err)
			normal := mgl32.Vec3{float32(x), float32(y), float32(z)}
			normals = append(normals, normal)
		} else if strings.HasPrefix(line, "f ") {
			break
		}
	}
	for err != io.EOF && strings.HasPrefix(line, "f ") {
		currentLine := strings.Split(line, " ")
		vertex1 := strings.Split(currentLine[1], "/")
		vertex2 := strings.Split(currentLine[2], "/")
		vertex3 := strings.Split(currentLine[3], "/")
		vertices, indices = processVertex(vertex1, vertices, indices)
		vertices, indices = processVertex(vertex2, vertices, indices)
		vertices, indices = processVertex(vertex3, vertices, indices)
		s, _, err = r.ReadLine()
		check(err)
		line = string(s[:])
	}


	return ModelData{}
}

func processVertex(vertex []string, vertices []Vertex, indices []int) ([]Vertex, []int){
	index, err := strconv.ParseInt(vertex[0], 10, 64)
	check(err)
	index--
	currentVertex := vertices[index]
	textureIndex, err := strconv.ParseInt(vertex[1], 10, 64)
	check(err)
	normalIndex, err := strconv.ParseInt(vertex[2], 10, 64)
	check(err)
	if !currentVertex.isSet() {
		currentVertex.setTextureIndex(int(textureIndex))
		currentVertex.setNormalIndex(int(normalIndex))
		indices = append(indices, int(index))
	} else {
		vertices, indices = dealWithAlreadyProcessedVertex(currentVertex, int(textureIndex), int(normalIndex), vertices, indices)
	}
	return vertices, indices
}

func dealWithAlreadyProcessedVertex(previousVertex Vertex, newTextureIndex, newNormalIndex int, vertices []Vertex, indices []int) ([]Vertex, []int) {
	if previousVertex.hasSameTextureAndNormal(newTextureIndex, newNormalIndex) {
		indices = append(indices, previousVertex.index)
	} else {
		//anotherVertex := previousVertex
	}
	return vertices, indices
}