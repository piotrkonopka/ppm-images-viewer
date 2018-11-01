package main

import (
    "runtime"
    "github.com/go-gl/glfw/v3.2/glfw"
    "./src/program"
    "./src/renderer"
    "./src/window"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	runtime.LockOSThread()
}

func main() {
    wind := window.CreateWindow()
	defer glfw.Terminate()

    prog := program.CreateProgram()

	renderer.Render(wind, prog)
}
