package renderer

import (
        "github.com/go-gl/gl/v4.1-core/gl"
        "github.com/go-gl/glfw/v3.2/glfw"
)

var (
	vertexBufferData = []float32{
		0, 0.5, 0,          // top
		-0.5, -0.5, 0,      // left
		0.5, -0.5, 0,       // right
	}
	colorBufferData = []float32{
                0.5,  0,  0,
                0,  0.5,  0,
                0,  0,  0.5,
	}
        sizeOfVertexBufferData = len(vertexBufferData)*4  // size on bytes (*4 = float32)
        sizeOfColorBufferData = len(colorBufferData)*4
        indicatesToRender = int32(len(vertexBufferData)/3)
)

func Render(window *glfw.Window, program uint32) {
        gl.ClearColor(0, 0, 0, 0)

	vertexArrayObject := makeVertexArrayObject()

	for !window.ShouldClose() {
                if window.GetKey(glfw.KeyEscape) == glfw.Press {
                      break
                }
		draw(vertexArrayObject, window, program)
	}
}

func makeVertexArrayObject() uint32 {
	var vertexBufferObject uint32
	gl.GenBuffers(1, &vertexBufferObject)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObject)
	gl.BufferData(gl.ARRAY_BUFFER, sizeOfVertexBufferData, gl.Ptr(vertexBufferData), gl.STATIC_DRAW)

	var colorBufferObject uint32
	gl.GenBuffers(1, &colorBufferObject)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorBufferObject)
	gl.BufferData(gl.ARRAY_BUFFER, sizeOfColorBufferData, gl.Ptr(colorBufferData), gl.STATIC_DRAW)

	var vertexArrayObject uint32
	gl.GenVertexArrays(1, &vertexArrayObject)
	gl.BindVertexArray(vertexArrayObject)

	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObject)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorBufferObject)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)

	return vertexArrayObject
}

func draw(vertexArrayObject uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.PointSize(5.0)
	gl.BindVertexArray(vertexArrayObject)
	gl.DrawArrays(gl.TRIANGLES, 0, indicatesToRender) // gl.POINTS

	glfw.PollEvents()
	window.SwapBuffers()
}
