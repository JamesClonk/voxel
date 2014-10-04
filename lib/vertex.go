package lib

import (
	"unsafe"

	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	mgl "github.com/go-gl/mathgl/mgl32"
)

type Vertex struct {
	Position          mgl.Vec3
	Color             mgl.Vec4
	Normal            mgl.Vec3
	TextureCoordinate mgl.Vec2
}

type Vertices []Vertex

const (
	VERTEX_SIZE                      = int(unsafe.Sizeof(Vertex{}))
	VERTEX_OFFSET_POSITION           = 0
	VERTEX_OFFSET_COLOR              = VERTEX_OFFSET_POSITION + int(unsafe.Sizeof(mgl.Vec3{}))
	VERTEX_OFFSET_NORMAL             = VERTEX_OFFSET_COLOR + int(unsafe.Sizeof(mgl.Vec4{}))
	VERTEX_OFFSET_TEXTURE_COORDINATE = VERTEX_OFFSET_NORMAL + int(unsafe.Sizeof(mgl.Vec3{}))
)

func init() {
	if int(glh.Sizeof(gl.FLOAT))*12 != VERTEX_SIZE {
		panic("wrong vertex size!")
	}
}

func (v Vertex) Size() int {
	return VERTEX_SIZE
}

func (v Vertex) OffsetPosition() int {
	return VERTEX_OFFSET_POSITION
}

func (v Vertex) OffsetColor() int {
	return VERTEX_OFFSET_COLOR
}

func (v Vertex) OffsetNormal() int {
	return VERTEX_OFFSET_NORMAL
}

func (v Vertex) OffsetTextureCoordinate() int {
	return VERTEX_OFFSET_TEXTURE_COORDINATE
}
