package lib

import mgl "github.com/go-gl/mathgl/mgl32"

type Cube struct {
	Vertices Vertices
	Indices  Indices
}

func NewCube(baseColor mgl.Vec4) *Cube {
	vertices := Vertices{
		Vertex{
			Position:          mgl.Vec3{1, -1, 1},
			Color:             baseColor,
			Normal:            mgl.Vec3{1, -1, 1},
			TextureCoordinate: mgl.Vec2{1, 1},
		},
		Vertex{
			Position:          mgl.Vec3{1, 1, 1},
			Color:             baseColor,
			Normal:            mgl.Vec3{1, 1, 1},
			TextureCoordinate: mgl.Vec2{1, 0},
		},
		Vertex{
			Position:          mgl.Vec3{-1, 1, 1},
			Color:             baseColor,
			Normal:            mgl.Vec3{-1, 1, 1},
			TextureCoordinate: mgl.Vec2{0, 0},
		},
		Vertex{
			Position:          mgl.Vec3{-1, -1, 1},
			Color:             baseColor,
			Normal:            mgl.Vec3{-1, -1, 1},
			TextureCoordinate: mgl.Vec2{0, 1},
		},
		Vertex{
			Position:          mgl.Vec3{1, -1, -1},
			Color:             baseColor,
			Normal:            mgl.Vec3{1, -1, -1},
			TextureCoordinate: mgl.Vec2{0, 1},
		},
		Vertex{
			Position:          mgl.Vec3{1, 1, -1},
			Color:             baseColor,
			Normal:            mgl.Vec3{1, 1, -1},
			TextureCoordinate: mgl.Vec2{0, 0},
		},
		Vertex{
			Position:          mgl.Vec3{-1, 1, -1},
			Color:             baseColor,
			Normal:            mgl.Vec3{-1, 1, -1},
			TextureCoordinate: mgl.Vec2{1, 0},
		},
		Vertex{
			Position:          mgl.Vec3{-1, -1, -1},
			Color:             baseColor,
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

	return &Cube{vertices, indices}
}
