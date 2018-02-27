package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v4.5-core/gl"
	"log"
	"fmt"
)

const (
	WIDTH = 1280
	HEIGHT = 720
	TITLE = "OpenGL in Golang"
)

type Display struct {
	window *glfw.Window
}

var (
	cameraXTemp, cameraYTemp, cameraZTemp float32
)

func newDisplay() *glfw.Window {

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 5)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(WIDTH, HEIGHT, TITLE, nil, nil)
	check(err)
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	return window
}

func (d *Display) updateDisplay() {
	glfw.SwapInterval(1)
	glfw.PollEvents()
	d.window.SwapBuffers()
}

func (d *Display) closeDisplay() {
	glfw.Terminate()
}

func processInput(window *glfw.Window) {

	if window.GetKey(glfw.KeyW) == glfw.Press {
		cameraZTemp -= 0.2
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		cameraZTemp += 0.2
	}
	if window.GetKey(glfw.KeyA) == glfw.Press {
		cameraXTemp -= 0.2
	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		cameraXTemp += 0.2
	}
	if window.GetKey(glfw.KeySpace) == glfw.Press {
		cameraYTemp += 0.2
	}
	if window.GetKey(glfw.KeyLeftShift) == glfw.Press {
		cameraYTemp -= 0.2
	}

}
