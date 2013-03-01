package main

import (
	"fmt"
	"github.com/Jragonmiris/mathgl"
	"github.com/Jragonmiris/mathgl/examples/opengl-tutorial/helper"
	"github.com/Jragonmiris/mathgl/examples/opengl-tutorial/indexer"
	"github.com/Jragonmiris/mathgl/examples/opengl-tutorial/input"
	"github.com/Jragonmiris/mathgl/examples/opengl-tutorial/objloader"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"os"
)

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.FsaaSamples, 4)
	glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
	glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 3)
	glfw.OpenWindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	if err := glfw.OpenWindow(1024, 768, 0, 0, 0, 0, 32, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	glfw.SetSwapInterval(0)

	//gl.GlewExperimental(true)
	gl.Init()     // Can't find gl.GLEW_OK or any variation, not sure how to check if this worked
	gl.GetError() // ignore error, since we're telling it to use CoreProfile above, we get "invalid enumerant" (GLError 1280) which freaks the OpenGLSentinel out

	glfw.SetWindowTitle("Tutorial 09")

	glfw.Enable(glfw.StickyKeys)
	glfw.Disable(glfw.MouseCursor) // Not in the original tutorial, but IMO it SHOULD be there
	glfw.SetMousePos(1024.0/2.0, 768.0/2.0)

	gl.ClearColor(0., 0., 0.4, 0.)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.Enable(gl.CULL_FACE)

	camera := input.NewCamera()

	vertexArray := gl.GenVertexArray()
	defer vertexArray.Delete()
	vertexArray.Bind()

	prog := helper.MakeProgram("StandardShading.vertexshader", "StandardShading.fragmentshader")
	defer prog.Delete()

	matrixID := prog.GetUniformLocation("MVP")
	viewMatrixID := prog.GetUniformLocation("V")
	modelMatrixID := prog.GetUniformLocation("M")

	texture := helper.MakeTextureFromTGA("uvmap.tga") // Had to convert to tga, go-gl is missing the texture method for DDS right now
	defer texture.Delete()
	texSampler := prog.GetUniformLocation("myTextureSampler")

	meshObj := objloader.LoadObject("suzanne.obj")

	indices, indexedVertices, indexedUVs, indexedNormals := indexer.IndexVBO(meshObj.Vertices, meshObj.UVs, meshObj.Normals)

	vertexBuffer := gl.GenBuffer()
	defer vertexBuffer.Delete()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(indexedVertices)*3*4, indexedVertices, gl.STATIC_DRAW)

	uvBuffer := gl.GenBuffer()
	defer uvBuffer.Delete()
	uvBuffer.Bind(gl.ARRAY_BUFFER)
	// And yet, the weird length stuff doesn't seem to matter for UV or normal
	gl.BufferData(gl.ARRAY_BUFFER, len(indexedUVs)*2*4, indexedUVs, gl.STATIC_DRAW)

	normBuffer := gl.GenBuffer()
	defer normBuffer.Delete()
	normBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(indexedNormals)*3*4, indexedNormals, gl.STATIC_DRAW)

	elementBuffer := gl.GenBuffer()
	defer elementBuffer.Delete()
	elementBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*2, indices, gl.STATIC_DRAW) // NOTE: a GL_UNSIGNED_SHORT is 16-bits

	lightID := prog.GetUniformLocation("LightPosition_worldspace")
	lastTime := glfw.Time()
	nbFrames := 0
	// Equivalent to a do... while
	for ok := true; ok; ok = (glfw.Key(glfw.KeyEsc) != glfw.KeyPress && glfw.WindowParam(glfw.Opened) == gl.TRUE && glfw.Key('Q') != glfw.KeyPress) {

		currTime := glfw.Time()
		nbFrames++
		if currTime-lastTime >= 1.0 {
			fmt.Printf("%f ms/frame\n", 1000.0/float64(nbFrames))
			nbFrames = 0
			lastTime += 1.0
		}

		func() {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			prog.Use()
			defer gl.ProgramUnuse()

			view, proj := camera.ComputeViewPerspective()
			model := mathgl.Ident4f()

			MVP := proj.Mul4(view).Mul4(model)
			//mvpArray := mvp.AsCMOArray(mathgl.FLOAT32).([16]float32)
			//vArray := view.AsCMOArray(mathgl.FLOAT32).([16]float32)
			//mArray := model.AsCMOArray(mathgl.FLOAT32).([16]float32)

			matrixID.UniformMatrix4fv(false, MVP)
			viewMatrixID.UniformMatrix4fv(false, view)
			modelMatrixID.UniformMatrix4fv(false, model)

			lightID.Uniform3f(4., 4., 4.)

			gl.ActiveTexture(gl.TEXTURE0)
			texture.Bind(gl.TEXTURE_2D)
			defer texture.Unbind(gl.TEXTURE_2D)
			texSampler.Uniform1i(0)

			vertexAttrib := gl.AttribLocation(0)
			vertexAttrib.EnableArray()
			defer vertexAttrib.DisableArray()
			vertexBuffer.Bind(gl.ARRAY_BUFFER)
			defer vertexBuffer.Unbind(gl.ARRAY_BUFFER)
			vertexAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)

			uvAttrib := gl.AttribLocation(1)
			uvAttrib.EnableArray()
			defer uvAttrib.DisableArray()
			uvBuffer.Bind(gl.ARRAY_BUFFER)
			defer uvBuffer.Unbind(gl.ARRAY_BUFFER)
			uvAttrib.AttribPointer(2, gl.FLOAT, false, 0, nil)

			normAttrib := gl.AttribLocation(2)
			normAttrib.EnableArray()
			defer normAttrib.DisableArray()
			normBuffer.Bind(gl.ARRAY_BUFFER)
			defer normBuffer.Unbind(gl.ARRAY_BUFFER)
			normAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)

			elementBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
			defer elementBuffer.Unbind(gl.ELEMENT_ARRAY_BUFFER)

			gl.DrawElements(gl.TRIANGLES, len(indices), gl.UNSIGNED_SHORT, nil)

			glfw.SwapBuffers()
		}() // Defers unbinds and disables to here, end of the loop
	}

}

/*func MakeTextureFromDDS(fname string) gl.Texture {
	var header [124]byte
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	var code [4]byte
	reader.Read(&code)
	if string(code) != "DDS " {
		panic("File code not correct")
	}

	reader.Read(&header)

	var height, width, linearSize, mipMapCount uint32
	var fourcc [4]byte

	buf := bytes.NewBuffer(header[8:12])
	binary.Read(buf, binary.BigEndian, &height)

	buf := bytes.NewBuffer(header[12:16])
	binary.Read(buf, binary.BigEndian, &width)

	buf := bytes.NewBuffer(header[16:20])
	binary.Read(buf, binary.BigEndian, &linearSize)

	buf := bytes.NewBuffer(header[24:28])
	binary.Read(buf, binary.BigEndian, &mipMapCount)

	buf := bytes.NewBuffer(header[80:84])
	binary.Read(buf, binary.BigEndian, &fourCC)

	var bufSize
	if mipMapCount > 1 {
		bufSize = linearSize * 2
	} else {
		bufSize = linearSize
	}

	buffer := make([]byte, bufSize)
	reader.Read(&buffer)

	var components uint32
	var format gl.GLEnum
	if string(fourcc) == "DXT1" {
		components = 3
	} else {
		components = 4
	}

	switch string(fourcc) {
	case "DXT1":
		format = gl.COMPRESSED_RGBA_S3TC_DXT1_EXT
	case "DXT3":
		format = gl.COMPRESSED_RGBA_S3TC_DXT3_EXT
	case "DXT5":
		format = gl.COMPRESSED_RGBA_S3TC_DXT5_EXT
	default:
		panic("No recognized DXT code")
	}

	tex := gl.GenTexture()

	tex.Bind(gl.TEXTURE_2D)
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)

	var blockSize uint32
	if format ==  gl.COMPRESSED_RGBA_S3TC_DXT1_EXT {
		blockSize = uint32(8)
	} else {
		blockSize = uint32(16)
	}

	offset := uint32(0)

	for level := uint32(0); level < mipMapCount && (width != 0 || height != 0); level++ {
		size := ((width+3)/4)*((height+3)/4)*blockSize
		gl.
	}
}*/
