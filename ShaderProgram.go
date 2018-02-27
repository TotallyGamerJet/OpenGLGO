package main

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"strings"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
)

type ShaderProgram struct {
	programID, vertexID, fragmentID uint32
	uniformLocations map[string]int32
}

func newShaderProgram(vertexFile, fragmentFile string, bindAttributes, getAllUniformLocations func(ShaderProgram)) ShaderProgram {
	vertexShaderID, err := loadShader(vertexFile, gl.VERTEX_SHADER)
	check(err)
	fragmentShaderID, err := loadShader(fragmentFile, gl.FRAGMENT_SHADER)
	check(err)
	programID := gl.CreateProgram()
	gl.AttachShader(programID, vertexShaderID)
	gl.AttachShader(programID, fragmentShaderID)

	program := ShaderProgram{programID, vertexShaderID, fragmentShaderID, make(map[string]int32)}

	bindAttributes(program)

	gl.LinkProgram(programID)
	gl.ValidateProgram(programID)

	getAllUniformLocations(program)
	return program
}

func (p *ShaderProgram) start() {
	gl.UseProgram(p.programID)
}

func (p *ShaderProgram) stop() {
	gl.UseProgram(0)
}

func (p *ShaderProgram) getUniformLocation(uniformName string) int32 {
	return gl.GetUniformLocation(p.programID, gl.Str(uniformName + "\x00"))
}

func (p *ShaderProgram) bindAttribute(attribute uint32, variableName string) {
	gl.BindAttribLocation(p.programID, attribute, gl.Str(variableName + "\x00"))
}

func (p *ShaderProgram) loadFloat(loc int32, value float32) {
	gl.Uniform1f(loc, value)
}

func (p *ShaderProgram) loadVector(loc int32, vector mgl32.Vec3) {
	gl.Uniform3f(loc, vector.X(),vector.Y(),vector.Z())
}

func (p *ShaderProgram) loadBoolean(loc int32, value bool) {
	var toLoad float32 = 0
	if value {
		toLoad = 1
	}
	gl.Uniform1f(loc, toLoad)
}

func (p *ShaderProgram) loadMatrix(loc int32, matrix mgl32.Mat4) {
	gl.UniformMatrix4fv(loc, 1, false, &matrix[0])
}

func (p *ShaderProgram) cleanUp() {
	p.stop()
	gl.DetachShader(p.programID, p.vertexID)
	gl.DetachShader(p.programID, p.fragmentID)
	gl.DeleteShader(p.vertexID)
	gl.DeleteShader(p.fragmentID)
	gl.DeleteProgram(p.programID)
}

func loadShader(file string, shaderType uint32) (uint32, error) {
	//file, err := ioutil.ReadFile(filename)
	//check(err)
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(string(file))
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0,fmt.Errorf("failed to compile %v: %v", file, log)
	}

	return shader, nil
}
