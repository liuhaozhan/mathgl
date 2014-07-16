package matstack

import (
	"errors"

	"github.com/go-gl/mathgl/mgl32"
)

// A MatStack is an OpenGL-style matrix stack,
// usually used for things like scenegraphs. This allows you
// to easily maintain matrix state per call level.
type MatStack []mgl32.Mat4

func NewMatStack() *MatStack {
	return &MatStack{mgl32.Ident4()}
}

// Copies the top element and pushes it on the stack.
func (ms *MatStack) Push() {
	(*ms) = append(*ms, (*ms)[len(*ms)-1])
}

// Removes the first element of the matrix from the stack, if there is only one element left
// there is an error.
func (ms *MatStack) Pop() error {
	if len(*ms) == 1 {
		errors.New("Cannot pop from mat stack, at minimum stack length of 1")
	}
	(*ms) = (*ms)[:len(*ms)-1]
}

// Right multiplies the current top of the matrix by the
// argument.
func (ms *MatStack) RightMul(m Mat4) {
	(*ms)[len(*ms)-1] = (*ms)[len(*ms)-1].Mul4(m)
}

// Left multiplies the current top of the matrix by the
// argument.
func (ms *MatStack) LeftMul(m Mat4) {
	(*ms)[len(*ms)-1] = (*ms)[len(*ms)-1].Mul4(m)
}

// Returns the top element.
func (ms *MatStack) Peek() Mat4 {
	return (*ms)[len(*ms)-1]
}

// Rewrites the top element of the stack with m
func (ms *MatStack) Load(m Mat4) {
	(*ms)[len(*ms)] = m
}

// A shortcut for Load(mgl.Ident4())
func (ms *MatStack) LoadIdent() {
	(*ms)[len(*ms)] = mgl32.Ident4()
}
