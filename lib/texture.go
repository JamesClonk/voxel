package lib

import (
	"image"
	"image/png"
	"os"
	"unsafe"

	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

type Texture struct {
	Texture       gl.Texture
	Width, Height int
	Image         *image.NRGBA
}

func init() {
	var u uint8 = 123
	if int(glh.Sizeof(gl.UNSIGNED_BYTE)) != int(unsafe.Sizeof(u)) {
		panic("wrong uint8 size!")
	}
}

func NewTexture(width, height int, filename string, mode int) *Texture {
	texture := gl.GenTexture()
	texture.Bind(gl.TEXTURE_2D)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, mode)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, mode)
	texture.Unbind(gl.TEXTURE_2D)
	glh.OpenGLSentinel()

	return &Texture{
		Texture: texture,
		Width:   width,
		Height:  height,
		Image:   loadTexture(filename),
	}
}

func (t *Texture) Bind() {
	t.Texture.Bind(gl.TEXTURE_2D)
}

func (t *Texture) Unbind() {
	t.Texture.Unbind(gl.TEXTURE_2D)
}

func (t *Texture) Update() {
	t.Bind()

	gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, t.Width, t.Height, gl.RGBA, gl.UNSIGNED_BYTE, t.Image.Pix)
	gl.GenerateMipmap(gl.TEXTURE_2D)

	t.Unbind()
}

func loadTexture(filename string) *image.NRGBA {
	texfile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer texfile.Close()

	img, err := png.Decode(texfile)
	if err != nil {
		panic(err)
	}

	return img.(*image.NRGBA)
}
