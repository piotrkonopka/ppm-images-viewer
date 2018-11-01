package renderer

import (
    "unsafe"
    "fmt"
    "github.com/go-gl/gl/v4.1-core/gl"
    "github.com/go-gl/glfw/v3.2/glfw"
    "../util"
)

var (
	vertices = []float32{
		// top left
		-1.0, 1.0, 0.0,   // position
		1.0, 1.0, 1.0,    // Color
		1.0, 0.0,         // texture coordinates

		// top right
		1.0, 1.0, 0.0,
		1.0, 1.0, 1.0,
		0.0, 0.0,

		// bottom right
		1.0, -1.0, 0.0,
		1.0, 1.0, 1.0,
		0.0, 1.0,

		// bottom left
		-1.0, -1.0, 0.0,
		1.0, 1.0, 1.0,
		1.0, 1.0,
	}

	indices = []uint32{
		// rectangle
		0, 1, 2,  // top triangle
		0, 2, 3,  // bottom triangle
    }
)

func getTextureFromFile(window *glfw.Window, fileList []string, currentFile, numOfFiles int) *util.Texture {
    image, err := util.GetImageFromFile(fileList[currentFile])
    if err != nil {
        panic(err.Error())
    }

    texture, err := util.GetTextureFromImage(image)
    if err != nil {
        panic(err.Error())
    }

    newTitle := fmt.Sprintf("PPM Images Viewer [%d of %d]", currentFile+1, numOfFiles)
    window.SetTitle(newTitle)

    return texture
}

func Render(window *glfw.Window, program uint32) {
    fileList    := util.GetFilesList()
    numOfFiles := len(fileList)
    currentFile := 0
    texture := getTextureFromFile(window, fileList, currentFile, numOfFiles)

    gl.ClearColor(0, 0, 0, 0)
	vertexArrayObject := makeVertexArrayObject()

	for !window.ShouldClose() {
        if window.GetKey(glfw.KeyEscape) == glfw.Press {
              break
        }
        if window.GetKey(glfw.KeyRight) == glfw.Press {
            if window.GetKey(glfw.KeyRight) == glfw.Release {
                if (currentFile < numOfFiles-1) {
                    currentFile += 1
                } else {
                    currentFile = 0
                }
                texture = getTextureFromFile(window, fileList, currentFile, numOfFiles)
            }
        }
        if window.GetKey(glfw.KeyLeft) == glfw.Press {
            if window.GetKey(glfw.KeyLeft) == glfw.Release {
                if (currentFile > 0) {
                    currentFile -= 1
                } else {
                    currentFile = numOfFiles-1
                }
                texture = getTextureFromFile(window, fileList, currentFile, numOfFiles)
            }
        }
        draw(vertexArrayObject, window, program, texture)
	}
}

func makeVertexArrayObject() uint32 {
	var vertexArrayObject uint32
	gl.GenVertexArrays(1, &vertexArrayObject)
	gl.BindVertexArray(vertexArrayObject)

	var vertexBufferObject uint32
	gl.GenBuffers(1, &vertexBufferObject)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObject)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	var elementBufferObject uint32;
	gl.GenBuffers(1, &elementBufferObject)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementBufferObject)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	var stride int32 = 3*4 + 3*4 + 2*4
	var offset int = 0

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(0)
	offset += 3*4

	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(1)
	offset += 3*4

	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(2)

	gl.BindVertexArray(0)

    return vertexArrayObject
}

func draw(vertexArrayObject uint32, window *glfw.Window, program uint32, texture *util.Texture) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

    texture.Bind(gl.TEXTURE0)
    texture.SetUniform(gl.GetUniformLocation(program, gl.Str("texture0" + "\x00")))

	gl.BindVertexArray(vertexArrayObject)
    gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))
    gl.BindVertexArray(0)

    texture.UnBind()

	glfw.PollEvents()
	window.SwapBuffers()
}
