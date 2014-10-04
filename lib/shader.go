package lib

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

type Shader struct {
	Program    gl.Program
	Ortho      gl.UniformLocation
	Model      gl.UniformLocation
	View       gl.UniformLocation
	Projection gl.UniformLocation
	Normal     gl.UniformLocation
}

func NewShader(vertexShaderSource, fragmentShaderSource string) *Shader {
	vertexShader := glh.Shader{gl.VERTEX_SHADER, vertexShaderSource}
	fragmentShader := glh.Shader{gl.FRAGMENT_SHADER, fragmentShaderSource}
	program := glh.NewProgram(vertexShader, fragmentShader)
	program.Use()
	glh.OpenGLSentinel()

	shader := Shader{Program: program}
	shader.SetUniformLocations()
	program.Unuse()
	glh.OpenGLSentinel()

	return &shader
}

func (s *Shader) Bind() {
	s.Program.Use()
}

func (s *Shader) Unbind() {
	s.Program.Unuse()
}

func (s *Shader) SetUniformLocations() {
	s.Ortho = s.Program.GetUniformLocation("ortho")
	s.Model = s.Program.GetUniformLocation("model")
	s.View = s.Program.GetUniformLocation("view")
	s.Projection = s.Program.GetUniformLocation("projection")
	s.Normal = s.Program.GetUniformLocation("normalMatrix")
}
