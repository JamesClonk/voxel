package lib

import (
	"unsafe"

	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	mgl "github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	Shader             *Shader
	Position           mgl.Vec3
	Rotation           mgl.Quat
	Scale              float32
	Vertices           Vertices
	Indices            Indices
	VertexArray        gl.VertexArray
	VertexArrayBuffer  gl.Buffer
	ElementArrayBuffer gl.Buffer
}

type Indices []int32

type Meshable interface {
	Vertices() Vertices
	Indices() Indices
}

func init() {
	var b byte = 255
	var i int32 = 1234567890
	if int(glh.Sizeof(gl.FLOAT))*4 != int(unsafe.Sizeof(mgl.Vec4{})) {
		panic("wrong float type!")
	} else if int(glh.Sizeof(gl.UNSIGNED_BYTE)) != int(unsafe.Sizeof(b)) {
		panic("wrong byte size!")
	} else if int(glh.Sizeof(gl.UNSIGNED_INT)) != int(unsafe.Sizeof(i)) {
		panic("wrong int size!")
	}
}

func NewMesh(shader *Shader, object Meshable) *Mesh {
	mesh := &Mesh{
		Shader:   shader,
		Position: mgl.Vec3{0, 0, 0},
		Rotation: mgl.Quat{1, mgl.Vec3{0, 0, 0}},
		Scale:    1,
		Vertices: object.Vertices(),
		Indices:  object.Indices(),
	}
	mesh.Buffer()

	return mesh
}

func (m *Mesh) Buffer() {
	m.setVertexArray()
	m.setVertexArrayBuffer(gl.DYNAMIC_DRAW)
	m.setElementArrayBuffer(gl.STATIC_DRAW)
	m.enableVertexAttributes()
	glh.OpenGLSentinel()

	m.VertexArray.Unbind()
	m.VertexArrayBuffer.Unbind(gl.ARRAY_BUFFER)
	m.ElementArrayBuffer.Unbind(gl.ELEMENT_ARRAY_BUFFER)
	glh.OpenGLSentinel()
}

func (m *Mesh) SubBuffer() {
	m.Bind()

	m.VertexArrayBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(m.Vertices)*VERTEX_SIZE, m.Vertices)
	m.VertexArrayBuffer.Unbind(gl.ARRAY_BUFFER)

	m.Unbind()
}

func (m *Mesh) Bind() {
	m.VertexArray.Bind()
}

func (m *Mesh) Unbind() {
	m.VertexArray.Unbind()
}

func (m *Mesh) Model() mgl.Mat4 {
	scale := mgl.Scale3D(m.Scale, m.Scale, m.Scale)
	rotate := m.Rotation.Mat4()
	translate := mgl.Translate3D(m.Position.X(), m.Position.Y(), m.Position.Z())

	return translate.Mul4(rotate).Mul4(scale)
}

func (m *Mesh) DrawArrays(mode gl.GLenum) {
	gl.DrawArrays(mode, 0, len(m.Vertices))
}

func (m *Mesh) DrawElements(mode gl.GLenum) {
	gl.DrawElements(mode, len(m.Indices), gl.UNSIGNED_INT, nil)
}

func (m *Mesh) setVertexArray() {
	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()
	m.VertexArray = vertexArray
}

func (m *Mesh) setVertexArrayBuffer(mode gl.GLenum) {
	vertexBuffer := gl.GenBuffer()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(m.Vertices)*VERTEX_SIZE, m.Vertices, mode)
	m.VertexArrayBuffer = vertexBuffer
}

func (m *Mesh) setElementArrayBuffer(mode gl.GLenum) {
	elementBuffer := gl.GenBuffer()
	elementBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(m.Indices)*int(glh.Sizeof(gl.UNSIGNED_INT)), m.Indices, mode)
	m.ElementArrayBuffer = elementBuffer
}

func (m *Mesh) enableVertexAttributes() {
	m.enableVertexAttribute("position", 3, VERTEX_SIZE, VERTEX_OFFSET_POSITION)
	m.enableVertexAttribute("color", 4, VERTEX_SIZE, VERTEX_OFFSET_COLOR)
	m.enableVertexAttribute("normal", 3, VERTEX_SIZE, VERTEX_OFFSET_NORMAL)
	m.enableVertexAttribute("textureCoordinate", 2, VERTEX_SIZE, VERTEX_OFFSET_TEXTURE_COORDINATE)
}

func (m *Mesh) enableVertexAttribute(name string, length uint, size int, offset int) {
	attrib := m.Shader.Program.GetAttribLocation(name)
	attrib.EnableArray()
	attrib.AttribPointer(length, gl.FLOAT, false, size, uintptr(offset))
}
