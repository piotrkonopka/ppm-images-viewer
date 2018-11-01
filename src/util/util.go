package util

import (
    "os"
    "io/ioutil"
    "strings"
    "image"
    "image/draw"
    "errors"
    "github.com/lmittmann/ppm"
    "github.com/go-gl/gl/v4.1-core/gl"
)

const (
    galleryPath = "gallery/"
    fileExtension = ".ppm"
)

type Texture struct {
	handle uint32
	target uint32
	texUnit uint32
}

var errUnsupportedStride = errors.New("unsupported stride, only 32-bit colors supported")
var errTextureNotBound = errors.New("texture not bound")

func GetImageFromFile (filePath string) (image.Image, error) {
    file, err := os.Open(galleryPath + filePath)
    checkError(err)
    defer file.Close()

    return ppm.Decode(file)
}

func GetFilesList() []string {    
    files, err := ioutil.ReadDir(galleryPath)
    checkError(err)

    var filesList []string
    for _, file := range files {
        if name := file.Name(); strings.HasSuffix(name, fileExtension) {
            filesList = append(filesList, name)
        }
    }

    return filesList
}

func GetTextureFromImage(img image.Image) (*Texture, error) {
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, errUnsupportedStride
	}

	var handle uint32
	gl.GenTextures(1, &handle)

	target      := uint32(gl.TEXTURE_2D)
	internalFmt := int32(gl.SRGB_ALPHA)
	format      := uint32(gl.RGBA)
	width       := int32(rgba.Rect.Size().X)
	height      := int32(rgba.Rect.Size().Y)
	pixType     := uint32(gl.UNSIGNED_BYTE)
	dataPtr     := gl.Ptr(rgba.Pix)

	texture := Texture{
		handle:handle,
		target:target,
	}

	texture.Bind(gl.TEXTURE0)
	defer texture.UnBind()

	gl.TexParameteri(texture.target, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(texture.target, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(texture.target, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(texture.target, gl.TEXTURE_MAG_FILTER, gl.LINEAR) 

	gl.TexImage2D(target, 0, internalFmt, width, height, 0, format, pixType, dataPtr)
	gl.GenerateMipmap(texture.handle)

	return &texture, nil
}

func (tex *Texture) Bind(texUnit uint32) {
	gl.ActiveTexture(texUnit)
	gl.BindTexture(tex.target, tex.handle)
	tex.texUnit = texUnit
}

func (tex *Texture) UnBind() {
	tex.texUnit = 0
	gl.BindTexture(tex.target, 0)
}

func (tex *Texture) SetUniform(uniformLoc int32) error {
	if tex.texUnit == 0 {
		return errTextureNotBound
	}
	gl.Uniform1i(uniformLoc, int32(tex.texUnit - gl.TEXTURE0))
	return nil
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
