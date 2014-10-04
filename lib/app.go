package lib

import (
	"log"
	"runtime"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glh"
)

type App struct {
	Window       *glfw.Window
	Width        int
	Height       int
	Ratio        float32
	Title        string
	ViewportFunc func(*App)
	DrawFunc     func(*App)
	KeyFunc      func(*glfw.Window, glfw.Key, int, glfw.Action, glfw.ModifierKey)
	MouseFunc    func(*glfw.Window, glfw.MouseButton, glfw.Action, glfw.ModifierKey)
	CursorFunc   func(*glfw.Window, float64, float64)
	ErrorFunc    func(glfw.ErrorCode, string)
}

func NewSimpleApp(width, height int, title string, drawFunc func(*App)) *App {
	return NewApp(width, height, title, UpdateViewport, drawFunc, OnKeyDown, OnMouseDown, OnMouseMove, OnError)
}

func NewApp(width, height int, title string, viewportFunc func(*App), drawFunc func(*App), keyFunc func(*glfw.Window, glfw.Key, int, glfw.Action, glfw.ModifierKey), mouseFunc func(*glfw.Window, glfw.MouseButton, glfw.Action, glfw.ModifierKey), cursorFunc func(*glfw.Window, float64, float64), errorFunc func(glfw.ErrorCode, string)) *App {
	runtime.LockOSThread()

	if !glfw.Init() {
		panic("can't init glfw!")
	}
	glfw.SetErrorCallback(errorFunc)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	window.SetKeyCallback(keyFunc)
	window.SetMouseButtonCallback(mouseFunc)
	window.SetCursorPositionCallback(cursorFunc)

	if gl.Init() != 0 {
		panic("can't init glew!")
	}
	gl.GetError()

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.LineWidth(3)

	return &App{
		Window:       window,
		Width:        width,
		Height:       height,
		Ratio:        float32(width) / float32(height),
		Title:        title,
		ViewportFunc: viewportFunc,
		DrawFunc:     drawFunc,
		KeyFunc:      keyFunc,
		MouseFunc:    mouseFunc,
		CursorFunc:   cursorFunc,
	}
}

func (a *App) Start() {
	for !a.Window.ShouldClose() {
		a.ViewportFunc(a)

		gl.ClearColor(0.1, 0.1, 0.1, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		a.DrawFunc(a)
		glh.OpenGLSentinel()

		a.Window.SwapBuffers()
		glfw.PollEvents()
	}
}

func (a *App) Close() {
	a.Window.SetShouldClose(true)
}

func (a *App) Destroy() {
	glh.OpenGLSentinel()
	a.Window.Destroy()
	glfw.Terminate()
}

func UpdateViewport(a *App) {
	w, h := a.Window.GetFramebufferSize()

	gl.Viewport(0, 0, w, h)

	a.Width = w
	a.Height = h
	a.Ratio = float32(w) / float32(h)
}

func OnKeyDown(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mod glfw.ModifierKey) {
	log.Printf("Key [%v], Scancode [%v], Action [%v], Modifier [%v]\n", key, scancode, action, mod)

	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}

func OnMouseDown(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	log.Printf("Mouse Button [%v], Action [%v], Modifier [%v]\n", button, action, mod)
}

func OnMouseMove(window *glfw.Window, x, y float64) {
	log.Printf("Mouse Position [%.0f, %.0f]\n", x, y)
}

func OnError(err glfw.ErrorCode, description string) {
	log.Printf("%v: %v\n", err, description)
}
