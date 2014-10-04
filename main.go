package main

import (
	"math"

	. "github.com/JamesClonk/voxel/lib"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	mgl "github.com/go-gl/mathgl/mgl32"
)

var shader *Shader
var mesh *Mesh
var texture *Texture
var time float64

const vertexShaderSource = `
	#version 130
		in vec3 position;
		in vec4 color;
		in vec3 normal;
		in vec2 textureCoordinate;

		varying vec2 texCoord;
		varying float diffuse;
		varying vec4 inColor;

		uniform mat4 model;
		uniform mat4 view;
		uniform mat4 projection;
		uniform mat3 normalMatrix;

		float doColor() {
			vec3 normalized  = normalize(normalMatrix * normalize(normal));
			vec3 light = normalize(vec3(1.0, 1.0, 1.0));
			return max(dot(normalized, light), 0.0);
		}

		void main()	{
			diffuse = doColor();
			inColor = color;
			texCoord = textureCoordinate;
			gl_Position = projection * view * model * vec4(position, 1);
		}
`

const fragmentShaderSource = `
	#version 130
		uniform sampler2D texture;

		varying vec2 texCoord;  
		varying float diffuse;
		varying vec4 inColor;

		void main() {
			gl_FragColor =  inColor * vec4(texture2D(texture, texCoord).rgb * diffuse, 1.0);
		}
`

func main() {
	app := NewSimpleApp(640, 480, "Voxel", draw)
	defer app.Destroy()

	shader = NewShader(vertexShaderSource, fragmentShaderSource)
	mesh = NewMesh(shader, NewCube(mgl.Vec4{1, 1, 1, 1}))
	texture = NewTexture("data/grass.png", gl.NEAREST)

	app.Start()
}

func draw(app *App) {
	time += 0.01

	shader.Bind()

	ortho := mgl.Ortho(-app.Ratio, app.Ratio, -1.0, 1.0, -1.0, 1.0)
	shader.Ortho.UniformMatrix4fv(false, ortho)

	// view and projection
	view := mgl.LookAtV(mgl.Vec3{0, 0, 5}, mgl.Vec3{0, 0, 0}, mgl.Vec3{0, 1, 0})
	projection := mgl.Perspective(math.Pi/3.0, app.Ratio, 0.1, -10.0)

	// send view and projection to shader
	shader.View.UniformMatrix4fv(false, view)
	shader.Projection.UniformMatrix4fv(false, projection)

	// transformation matrix for rotation
	//model := mgl.HomogRotate3D(float32(time), mgl.Vec3{0, 1, 0})
	//shader.Model.UniformMatrix4fv(false, model)

	// calculate normal matrix and send to shader
	//normal := view.Mul4(model).Mat3().Inv().Transpose()
	//shader.Normal.UniformMatrix3fv(false, normal)

	//mesh.SubBuffer()
	texture.Bind()
	mesh.Bind()

	length := float32(10)
	//total := length * length
	for i := float32(1); i <= length; i++ {
		for j := float32(1); j <= length; j++ {
			//index := i*length + j

			translate := mgl.Translate3D(-(length)+i*2, -2, -(length)+j*2-15)
			rotate := mgl.HomogRotate3D(0, mgl.Vec3{0, 0, 1})
			//scale := mgl.Scale3D(1-(index/total), 1-(index/total), 1-(index/total))

			//model := translate.Mul4(rotate).Mul4(scale)
			model := translate.Mul4(rotate)
			shader.Model.UniformMatrix4fv(false, model)

			//normal := view.Mul4(model).Mat3().Inv().Transpose()
			normal := view.Mat3()
			shader.Normal.UniformMatrix3fv(false, normal)

			mesh.DrawElements(gl.QUADS)
		}
	}

	mesh.Unbind()
	texture.Bind()

	shader.Unbind()

	glh.OpenGLSentinel()
}
