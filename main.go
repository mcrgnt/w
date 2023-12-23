package w

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

func CreateWindow(w, h int, title string) (window *glfw.Window, err error) {
	if err = glfw.Init(); err != nil {
		return
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err = glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		return
	}

	window.MakeContextCurrent()

	return
}

func Draw(_ *glfw.Window) {

}
