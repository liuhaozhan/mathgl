// Copyright 2014 The go-gl/mathgl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl64

import (
	"math"
)

type Vec2 [2]float64
type Vec3 [3]float64
type Vec4 [4]float64

// Add performs element-wise addition between two vectors. It is equivalent to iterating
// over every element of v1 and adding the corresponding element of v2 to it.
func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v1[0] + v2[0], v1[1] + v2[1]}
}

// Add performs element-wise addition between two vectors. It is equivalent to iterating
// over every element of v1 and adding the corresponding element of v2 to it.
func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

// Add performs element-wise addition between two vectors. It is equivalent to iterating
// over every element of v1 and adding the corresponding element of v2 to it.
func (v1 Vec4) Add(v2 Vec4) Vec4 {
	return Vec4{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

// Sub performs element-wise subtraction between two vectors. It is equivalent to iterating
// over every element of v1 and subtracting the corresponding element of v2 from it.
func (v1 Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v1[0] - v2[0], v1[1] - v2[1]}
}

// Sub performs element-wise subtraction between two vectors. It is equivalent to iterating
// over every element of v1 and subtracting the corresponding element of v2 from it.
func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

// Sub performs element-wise subtraction between two vectors. It is equivalent to iterating
// over every element of v1 and subtracting the corresponding element of v2 from it.
func (v1 Vec4) Sub(v2 Vec4) Vec4 {
	return Vec4{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

// Mul performs a scalar multiplication between the vector and some constant value
// c. This is equivalent to iterating over every vector element and multiplying by c.
func (v1 Vec2) Mul(c float64) Vec2 {
	return Vec2{v1[0] * c, v1[1] * c}
}

// Mul performs a scalar multiplication between the vector and some constant value
// c. This is equivalent to iterating over every vector element and multiplying by c.
func (v1 Vec3) Mul(c float64) Vec3 {
	return Vec3{v1[0] * c, v1[1] * c, v1[2] * c}
}

// Mul performs a scalar multiplication between the vector and some constant value
// c. This is equivalent to iterating over every vector element and multiplying by c.
func (v1 Vec4) Mul(c float64) Vec4 {
	return Vec4{v1[0] * c, v1[1] * c, v1[2] * c, v1[3] * c}
}

// Dot returns the dot product of this vector with another. There are multiple ways
// to describe this value. One is the multiplication of their lengths and cos(theta) where
// theta is the angle between the vectors: v1.v2 = |v1||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise multiplication of all
// elements. So for instance, two Vec3s would yield v1.x * v2.x + v1.y * v2.y + v1.z * v2.z.
//
// This means that the dot product of a vector and itself is the square of its Len (within
// the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to pointing in the same
// direction. If both vectors are normalized, the value will be -1 for opposite pointing,
// one for same pointing, and 0 for perpendicular vectors.
func (v1 Vec2) Dot(v2 Vec2) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

// Dot returns the dot product of this vector with another. There are multiple ways
// to describe this value. One is the multiplication of their lengths and cos(theta) where
// theta is the angle between the vectors: v1.v2 = |v1||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise multiplication of all
// elements. So for instance, two Vec3s would yield v1.x * v2.x + v1.y * v2.y + v1.z * v2.z.
//
// This means that the dot product of a vector and itself is the square of its Len (within
// the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to pointing in the same
// direction. If both vectors are normalized, the value will be -1 for opposite pointing,
// one for same pointing, and 0 for perpendicular vectors.
func (v1 Vec3) Dot(v2 Vec3) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

// Dot returns the dot product of this vector with another. There are multiple ways
// to describe this value. One is the multiplication of their lengths and cos(theta) where
// theta is the angle between the vectors: v1.v2 = |v1||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise multiplication of all
// elements. So for instance, two Vec3s would yield v1.x * v2.x + v1.y * v2.y + v1.z * v2.z.
//
// This means that the dot product of a vector and itself is the square of its Len (within
// the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to pointing in the same
// direction. If both vectors are normalized, the value will be -1 for opposite pointing,
// one for same pointing, and 0 for perpendicular vectors.
func (v1 Vec4) Dot(v2 Vec4) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

// Len returns the vector's length. Note that this is NOT the dimension of
// the vector (len(v)), but the mathematical length. This is equivalent to the square
// root of the sum of the squares of all elements. E.G. for a Vec2 it's
// math.Hypot(v[0], v[1]).
func (v1 Vec2) Len() float64 {
	return float64(math.Hypot(float64(v1[0]), float64(v1[1])))
}

// Len returns the vector's length. Note that this is NOT the dimension of
// the vector (len(v)), but the mathematical length. This is equivalent to the square
// root of the sum of the squares of all elements. E.G. for a Vec2 it's
// math.Hypot(v[0], v[1]).
func (v1 Vec3) Len() float64 {
	return float64(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2])))
}

// Len returns the vector's length. Note that this is NOT the dimension of
// the vector (len(v)), but the mathematical length. This is equivalent to the square
// root of the sum of the squares of all elements. E.G. for a Vec2 it's
// math.Hypot(v[0], v[1]).
func (v1 Vec4) Len() float64 {
	return float64(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2] + v1[3]*v1[3])))
}

// Normalize normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating point error),
// while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
func (v1 Vec2) Normalize() Vec2 {
	l := 1.0 / v1.Len()
	return Vec2{v1[0] * l, v1[1] * l}
}

// Normalize normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating point error),
// while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
func (v1 Vec3) Normalize() Vec3 {
	l := 1.0 / v1.Len()
	return Vec3{v1[0] * l, v1[1] * l, v1[2] * l}
}

// Normalize normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating point error),
// while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
func (v1 Vec4) Normalize() Vec4 {
	l := 1.0 / v1.Len()
	return Vec4{v1[0] * l, v1[1] * l, v1[2] * l, v1[3] * l}
}

// The vector cross product is an operation only defined on 3D vectors. It is equivalent to
// Vec3{v1[1]*v2[2]-v1[2]*v2[1], v1[2]*v2[0]-v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}.
// Another interpretation is that it's the vector whose magnitude is |v1||v2|sin(theta)
// where theta is the angle between v1 and v2.
//
// The cross product is most often used for finding surface normals. The cross product of vectors
// will generate a vector that is perpendicular to the plane they form.
//
// Technically, a generalized cross product exists as an "(N-1)ary" operation
// (that is, the 4D cross product requires 3 4D vectors). But the binary
// 3D (and 7D) cross product is the most important. It can be considered
// the area of a parallelogram with sides v1 and v2.
//
// Like the dot product, the cross product is roughly a measure of directionality.
// Two normalized perpendicular vectors will return a vector with a magnitude of
// 1.0 or -1.0 and two parallel vectors will return a vector with magnitude 0.0.
// The cross product is "anticommutative" meaning v1.Cross(v2) = -v2.Cross(v1),
// this property can be useful to know when finding normals,
// as taking the wrong cross product can lead to the opposite normal of the one you want.
func (v1 Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{v1[1]*v2[2] - v1[2]*v2[1], v1[2]*v2[0] - v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}
}

// ApproxEqual takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
func (v1 Vec2) ApproxEqual(v2 Vec2) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
func (v1 Vec3) ApproxEqual(v2 Vec3) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
func (v1 Vec4) ApproxEqual(v2 Vec4) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

// ApproxThresholdEq takes in a threshold for comparing two floats, and uses it to do an
// element-wise comparison of the vector to another.
func (v1 Vec2) ApproxEqualThreshold(v2 Vec2, threshold float64) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxThresholdEq takes in a threshold for comparing two floats, and uses it to do an
// element-wise comparison of the vector to another.
func (v1 Vec3) ApproxEqualThreshold(v2 Vec3, threshold float64) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxThresholdEq takes in a threshold for comparing two floats, and uses it to do an
// element-wise comparison of the vector to another.
func (v1 Vec4) ApproxEqualThreshold(v2 Vec4, threshold float64) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxFuncEq takes in a func that compares two floats, and uses it to do an element-wise
// comparison of the vector to another. This is intended to be used with FloatEqualFunc
func (v1 Vec2) ApproxFuncEqual(v2 Vec2, eq func(float64, float64) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

// ApproxFuncEq takes in a func that compares two floats, and uses it to do an element-wise
// comparison of the vector to another. This is intended to be used with FloatEqualFunc
func (v1 Vec3) ApproxFuncEqual(v2 Vec3, eq func(float64, float64) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

// ApproxFuncEq takes in a func that compares two floats, and uses it to do an element-wise
// comparison of the vector to another. This is intended to be used with FloatEqualFunc
func (v1 Vec4) ApproxFuncEqual(v2 Vec4, eq func(float64, float64) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

// This is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v[0] or v.X() depending on personal preference.
func (v Vec2) X() float64 {
	return v[0]
}

// This is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v[0] or v.X() depending on personal preference.
func (v Vec2) Y() float64 {
	return v[1]
}

// This is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v[0] or v.X() depending on personal preference.
func (v Vec3) X() float64 {
	return v[0]
}

// This is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v[0] or v.X() depending on personal preference.
func (v Vec3) Y() float64 {
	return v[1]
}

// This is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v[0] or v.X() depending on personal preference.
func (v Vec3) Z() float64 {
	return v[2]
}

// This is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v[0] or v.X() depending on personal preference.
func (v Vec4) X() float64 {
	return v[0]
}

// This is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v[0] or v.X() depending on personal preference.
func (v Vec4) Y() float64 {
	return v[1]
}

// This is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v[0] or v.X() depending on personal preference.
func (v Vec4) Z() float64 {
	return v[2]
}

// This is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v[0] or v.X() depending on personal preference.
func (v Vec4) W() float64 {
	return v[3]
}

func (v Vec2) Vec3(z float64) Vec3 {
	return Vec3{v[0], v[1], z}
}

func (v Vec2) Vec4(z, w float64) Vec4 {
	return Vec4{v[0], v[1], z, w}
}

func (v Vec3) Vec2() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec3) Vec4(w float64) Vec4 {
	return Vec4{v[0], v[1], v[2], w}
}

func (v Vec4) Vec2() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec4) Vec3() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

// extracts the elements of the vector for direct value assignment
func (v Vec2) Elem() (x, y float64) {
	return v[0], v[1]
}

// extracts the elements of the vector for direct value assignment
func (v Vec3) Elem() (x, y, z float64) {
	return v[0], v[1], v[2]
}

// extracts the elements of the vector for direct value assignment
func (v Vec4) Elem() (x, y, z, w float64) {
	return v[0], v[1], v[2], v[3]
}

// Does the vector outer product
// of two vectors. The outer product produces an
// NxM matrix. E.G. a Vec2 * Vec3 = Mat2x3.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
func (v1 Vec2) OuterProd2(v2 Vec2) Mat2 {
	return Mat2{v1[0] * v2[0], v1[1] * v2[0], v1[0] * v2[1], v1[1] * v2[1]}
}

// Does the vector outer product
// of two vectors. The outer product produces an
// NxM matrix. E.G. a Vec2 * Vec3 = Mat2x3.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
func (v1 Vec2) OuterProd3(v2 Vec3) Mat2x3 {
	return Mat2x3{v1[0] * v2[0], v1[1] * v2[0], v1[0] * v2[1], v1[1] * v2[1], v1[0] * v2[2], v1[1] * v2[2]}
}

// Does the vector outer product
// of two vectors. The outer product produces an
// NxM matrix. E.G. a Vec2 * Vec3 = Mat2x3.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
func (v1 Vec2) OuterProd4(v2 Vec4) Mat2x4 {
	return Mat2x4{v1[0] * v2[0], v1[1] * v2[0], v1[0] * v2[1], v1[1] * v2[1], v1[0] * v2[2], v1[1] * v2[2], v1[0] * v2[3], v1[1] * v2[3]}
}

// Does the vector outer product
// of two vectors. The outer product produces an
// NxM matrix. E.G. a Vec2 * Vec3 = Mat2x3.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
func (v1 Vec3) OuterProd2(v2 Vec2) Mat3x2 {
	return Mat3x2{v1[0] * v2[0], v1[1] * v2[0], v1[2] * v2[0], v1[0] * v2[1], v1[1] * v2[1], v1[2] * v2[1]}
}

// Does the vector outer product
// of two vectors. The outer product produces an
// NxM matrix. E.G. a Vec2 * Vec3 = Mat2x3.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
func (v1 Vec3) OuterProd3(v2 Vec3) Mat3 {
	return Mat3{v1[0] * v2[0], v1[1] * v2[0], v1[2] * v2[0], v1[0] * v2[1], v1[1] * v2[1], v1[2] * v2[1], v1[0] * v2[2], v1[1] * v2[2], v1[2] * v2[2]}
}

// Does the vector outer product
// of two vectors. The outer product produces an
// NxM matrix. E.G. a Vec2 * Vec3 = Mat2x3.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
func (v1 Vec3) OuterProd4(v2 Vec4) Mat3x4 {
	return Mat3x4{v1[0] * v2[0], v1[1] * v2[0], v1[2] * v2[0], v1[0] * v2[1], v1[1] * v2[1], v1[2] * v2[1], v1[0] * v2[2], v1[1] * v2[2], v1[2] * v2[2], v1[0] * v2[3], v1[1] * v2[3], v1[2] * v2[3]}
}

// Does the vector outer product
// of two vectors. The outer product produces an
// NxM matrix. E.G. a Vec2 * Vec3 = Mat2x3.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
func (v1 Vec4) OuterProd2(v2 Vec2) Mat4x2 {
	return Mat4x2{v1[0] * v2[0], v1[1] * v2[0], v1[2] * v2[0], v1[3] * v2[0], v1[0] * v2[1], v1[1] * v2[1], v1[2] * v2[1], v1[3] * v2[1]}
}

// Does the vector outer product
// of two vectors. The outer product produces an
// NxM matrix. E.G. a Vec2 * Vec3 = Mat2x3.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
func (v1 Vec4) OuterProd3(v2 Vec3) Mat4x3 {
	return Mat4x3{v1[0] * v2[0], v1[1] * v2[0], v1[2] * v2[0], v1[3] * v2[0], v1[0] * v2[1], v1[1] * v2[1], v1[2] * v2[1], v1[3] * v2[1], v1[0] * v2[2], v1[1] * v2[2], v1[2] * v2[2], v1[3] * v2[2]}
}

// Does the vector outer product
// of two vectors. The outer product produces an
// NxM matrix. E.G. a Vec2 * Vec3 = Mat2x3.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
func (v1 Vec4) OuterProd4(v2 Vec4) Mat4 {
	return Mat4{v1[0] * v2[0], v1[1] * v2[0], v1[2] * v2[0], v1[3] * v2[0], v1[0] * v2[1], v1[1] * v2[1], v1[2] * v2[1], v1[3] * v2[1], v1[0] * v2[2], v1[1] * v2[2], v1[2] * v2[2], v1[3] * v2[2], v1[0] * v2[3], v1[1] * v2[3], v1[2] * v2[3], v1[3] * v2[3]}
}
