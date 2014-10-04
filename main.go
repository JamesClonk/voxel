package main

import (
	"image"
	"image/png"
	"math"
	"os"

	. "github.com/JamesClonk/voxel/lib"
	"github.com/go-gl/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
)

var shader *Shader
var mesh *Mesh
var texture *Texture
var time float64

const vertexShaderSource = `
	#version 130
		in vec4 position;
		in vec4 color;
		in vec3 norm;
		in vec2 textureCoordinate;

		varying vec2 texCoord;
		varying float diffuse;
		varying vec4 inColor;

		uniform mat4 model;
		uniform mat4 view;
		uniform mat4 projection;
		uniform mat3 normal;

		float doColor() {
			vec3 normalized  = normalize(normal * normalize(norm));
			vec3 light = normalize(vec3(1.0, 1.0, 1.0));
			return max(dot(normalized, light), 0.0);
		}

		void main()	{
			diffuse = doColor();
			inColor = color;
			texCoord = textureCoordinate;
			gl_Position = projection * view * model * position;
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

	cube := Vertices{
		Vertex{
			Position:          mgl.Vec3{1, -1, 1},
			Color:             mgl.Vec4{1, 1, 0, 1},
			Normal:            mgl.Vec3{1, -1, 1},
			TextureCoordinate: mgl.Vec2{1, 1},
		},
		Vertex{
			Position:          mgl.Vec3{1, 1, 1},
			Color:             mgl.Vec4{0, 1, 0, 1},
			Normal:            mgl.Vec3{1, 1, 1},
			TextureCoordinate: mgl.Vec2{1, 0},
		},
		Vertex{
			Position:          mgl.Vec3{-1, 1, 1},
			Color:             mgl.Vec4{1, 1, 0, 1},
			Normal:            mgl.Vec3{-1, 1, 1},
			TextureCoordinate: mgl.Vec2{0, 0},
		},
		Vertex{
			Position:          mgl.Vec3{-1, -1, 1},
			Color:             mgl.Vec4{1, 0, 0, 1},
			Normal:            mgl.Vec3{-1, -1, 1},
			TextureCoordinate: mgl.Vec2{0, 1},
		},
		Vertex{
			Position:          mgl.Vec3{1, -1, -1},
			Color:             mgl.Vec4{0, 1, 0, 1},
			Normal:            mgl.Vec3{1, -1, -1},
			TextureCoordinate: mgl.Vec2{0, 1},
		},
		Vertex{
			Position:          mgl.Vec3{1, 1, -1},
			Color:             mgl.Vec4{0, 0, 1, 1},
			Normal:            mgl.Vec3{1, 1, -1},
			TextureCoordinate: mgl.Vec2{0, 0},
		},
		Vertex{
			Position:          mgl.Vec3{-1, 1, -1},
			Color:             mgl.Vec4{1, 0, 0, 1},
			Normal:            mgl.Vec3{-1, 1, -1},
			TextureCoordinate: mgl.Vec2{1, 0},
		},
		Vertex{
			Position:          mgl.Vec3{-1, -1, -1},
			Color:             mgl.Vec4{0, 0, 1, 1},
			Normal:            mgl.Vec3{-1, -1, -1},
			TextureCoordinate: mgl.Vec2{1, 1},
		},
	}
	/*
	       //6-------------/5
	     //  .           // |
	   //2--------------1   |
	   //    .          |   |
	   //    .          |   |
	   //    .          |   |
	   //    .          |   |
	   //    7.......   |   /4
	   //               | //
	   //3--------------/0
	*/

	indices := []int32{
		0, 1, 2, 3, // front
		7, 6, 5, 4, // back
		3, 2, 6, 7, // left
		4, 5, 1, 0, // right
		1, 5, 6, 2, // top
		4, 0, 3, 7, // bottom
	}

	shader = NewShader(vertexShaderSource, fragmentShaderSource)
	mesh = NewMesh(shader)
	mesh.Vertices = cube
	mesh.Indices = indices
	mesh.Buffer()

	textureData = loadTexture("texture.png")
	texture = NewTexture(24, 24, gl.NEAREST)

	app.Start()
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
	model := mgl.HomogRotate3D(float32(time), mgl.Vec3{0, 1, 0})
	shader.Model.UniformMatrix4fv(false, model)

	// calculate normal matrix and send to shader
	normal := view.Mul4(model).Mat3().Inv().Transpose()
	shader.Normal.UniformMatrix3fv(false, normal)

	gl.DrawElements(gl.QUADS, 24, gl.UNSIGNED_INT, nil)

	shader.Unbind()
}
