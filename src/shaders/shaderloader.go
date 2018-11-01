package shaderloader

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
    shadersPath = "src/shaders/"
	vertexFilePath   = "vertexshader.glsl"
	fragmentFilePath = "fragmentshader.glsl"
)

func LoadShaders() (uint32, uint32) {
	vertexShader, err := compileShaderFromFile(vertexFilePath, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShaderFromFile(fragmentFilePath, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	return vertexShader, fragmentShader
}

func readShaderFile(filePath string) string {
	data, err := ioutil.ReadFile(shadersPath+filePath)
	if err != nil {
		log.Fatal(err)
	}

	return string(data) + "\x00"
}

func compileShaderFromFile(filePath string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	source := readShaderFile(filePath)
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	defer free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		return handleShaderCompileError(shader, filePath)
	}

	return shader, nil
}

func handleShaderCompileError(shader uint32, filePath string) (uint32, error) {
	var logLength int32
	gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

	log := strings.Repeat("\x00", int(logLength+1))
	gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

	return 0, fmt.Errorf("failed to compile %v: %v", filePath, log)
}
