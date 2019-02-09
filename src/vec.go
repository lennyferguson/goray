package src

import "math"

const Delta = 0.0000001

type (

	//Vec3 is a Vec3
	Vec3 struct {
		arr [3]float32
	}

	//Vec4 is a Vec4
	Vec4 struct {
		arr [4]float32
	}

	// Vec is an interface
	Vec interface {
		ToSlice() []float32
		ToVec3() Vec3
		ToVec4() Vec4

		// Operators
		Add(other Vec) Vec
		Sub(other Vec) Vec
		Div(other Vec) Vec
		Mul(other Vec) Vec

		Dot(other Vec) float32
		Cross(other Vec) Vec

		// Are these needed?
		ScalarMul(c float32) Vec
		ScalarDiv(c float32) Vec
		ScalarAdd(c float32) Vec
		ScalarSub(c float32) Vec

		Equals(other Vec) bool
	}
)

func equalizeArrays(a []float32, b []float32) ([]float32, []float32) {
	if len(a) > len(b) {
		d := make([]float32, len(a))
		copy(d, b)
		return a, d
	} else if len(b) > len(a) {
		c := make([]float32, len(b))
		copy(c, a)
		return c, b
	} else {
		return a, b
	}
}

func toArray3(arr []float32) (ans [3]float32) {
	copy(ans[:], arr)
	return
}

func toArray4(arr []float32) (ans [4]float32) {
	copy(ans[:], arr)
	return
}

func map1(fn func(float32) float32, a []float32) []float32 {
	size := len(a)
	ans := make([]float32, size)
	for i, v := range a {
		ans[i] = fn(v)
	}
	return ans
}

func map2(fn func(float32, float32) float32, a, b []float32) []float32 {
	a, b = equalizeArrays(a, b)
	ans := make([]float32, len(a))
	for i := range a {
		ans[i] = fn(a[i], b[i])
	}
	return ans
}

func fold1(fn func(float32, float32) float32, start float32, a []float32) float32 {
	current := start
	for _, v := range a {
		current = fn(current, v)
	}
	return current
}

func fold2(fn func(float32, float32, float32) float32, start float32, a, b []float32) float32 {
	current := start
	a, b = equalizeArrays(a, b)
	for i := range a {
		current = fn(current, a[i], b[i])
	}
	return current
}

func equals(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(float64(a[i]-b[i])) > Delta {
			return false
		}
	}
	return true
}

func add(a, b []float32) []float32 {
	return map2(func(l float32, r float32) float32 {
		return l + r
	}, a, b)
}

func scalarAdd(a []float32, b float32) []float32 {
	return map1(func(v float32) float32 {
		return v + b
	}, a)
}

func sub(a, b []float32) []float32 {
	return map2(func(l float32, r float32) float32 {
		return l - r
	}, a, b)
}

func scalarSub(a []float32, b float32) []float32 {
	return map1(func(v float32) float32 {
		return v - b
	}, a)
}

func scalarMul(a []float32, b float32) []float32 {
	return map1(func(v float32) float32 {
		return v * b
	}, a)
}

func scalarDiv(a []float32, b float32) []float32 {
	return map1(func(v float32) float32 {
		return v / b
	}, a)
}

func mul(a, b []float32) []float32 {
	return map2(func(l float32, r float32) float32 {
		return l * r
	}, a, b)
}

func div(a, b []float32) []float32 {
	return map2(func(l float32, r float32) float32 {
		return l / r
	}, a, b)
}

func dot(a, b []float32) float32 {
	return fold2(func(c float32, u float32, v float32) float32 {
		return c + u*v
	}, 0.0, a, b)
}

func (v3 Vec3) ToSlice() []float32 {
	return v3.arr[:]
}

func (v3 Vec3) ToVec3() Vec3 {
	return v3
}

func (v3 Vec3) ToVec4() Vec4 {
	return Vec4{arr: toArray4(v3.arr[:])}
}

func (v3 Vec3) Add(other Vec) Vec {
	return Vec3{
		arr: toArray3(add(v3.arr[:], other.ToSlice())),
	}
}

func (v3 Vec3) Sub(other Vec) Vec {
	return Vec3{
		arr: toArray3(sub(v3.arr[:], other.ToSlice())),
	}
}

func (v3 Vec3) Mul(other Vec) Vec {
	return Vec3{
		arr: toArray3(mul(v3.ToSlice(), other.ToSlice())),
	}
}

func (v3 Vec3) Div(other Vec) Vec {
	return Vec3{
		arr: toArray3(div(v3.ToSlice(), other.ToSlice())),
	}
}

func (v3 Vec3) Cross(other Vec) Vec {
	return v3
}

func (v3 Vec3) Dot(other Vec) float32 {
	return dot(v3.ToSlice(), other.ToSlice())
}

func (v3 Vec3) ScalarAdd(c float32) Vec {
	return Vec3{
		arr: toArray3(
			scalarAdd(v3.ToSlice(), c)),
	}
}

func (v3 Vec3) ScalarSub(c float32) Vec {
	return Vec3{
		arr: toArray3(
			scalarSub(v3.ToSlice(), c)),
	}
}

func (v3 Vec3) ScalarMul(c float32) Vec {
	return Vec3{
		arr: toArray3(
			scalarMul(v3.ToSlice(), c)),
	}
}

func (v3 Vec3) ScalarDiv(c float32) Vec {
	return Vec3{
		arr: toArray3(
			scalarDiv(v3.ToSlice(), c)),
	}
}

func (v3 Vec3) X() float32 {
	return v3.arr[0]
}

func (v3 Vec3) Y() float32 {
	return v3.arr[1]
}

func (v3 Vec3) Z() float32 {
	return v3.arr[2]
}

func (v3 Vec3) Equals(other Vec) bool {
	return equals(v3.ToSlice(), other.ToSlice())
}

func (v4 Vec4) ToSlice() []float32 {
	return v4.arr[:]
}

func (v4 Vec4) ToVec3() Vec3 {
	return Vec3{
		arr: toArray3(v4.ToSlice()),
	}
}

func (v4 Vec4) ToVec4() Vec4 {
	return v4
}

func (v4 Vec4) Add(other Vec) Vec {
	return Vec4{
		arr: toArray4(add(v4.arr[:], other.ToSlice())),
	}
}

func (v4 Vec4) Sub(other Vec) Vec {
	return Vec4{
		arr: toArray4(sub(v4.arr[:], other.ToSlice())),
	}
}

func (v4 Vec4) Mul(other Vec) Vec {
	return Vec4{
		arr: toArray4(mul(v4.ToSlice(), other.ToSlice())),
	}
}

func (v4 Vec4) Div(other Vec) Vec {
	return Vec4{
		arr: toArray4(div(v4.ToSlice(), other.ToSlice())),
	}
}

func (v4 Vec4) Dot(other Vec) float32 {
	return dot(v4.ToSlice(), other.ToSlice())
}

func (v4 Vec4) Cross(other Vec) Vec {
	return v4
}

func (v4 Vec4) ScalarAdd(c float32) Vec {
	return Vec4{
		arr: toArray4(
			scalarAdd(v4.ToSlice(), c)),
	}
}

func (v4 Vec4) ScalarSub(c float32) Vec {
	return Vec4{
		arr: toArray4(
			scalarSub(v4.ToSlice(), c)),
	}
}

func (v4 Vec4) ScalarMul(c float32) Vec {
	return Vec4{
		arr: toArray4(
			scalarMul(v4.ToSlice(), c)),
	}
}

func (v4 Vec4) ScalarDiv(c float32) Vec {
	return Vec4{
		arr: toArray4(
			scalarDiv(v4.ToSlice(), c)),
	}
}

func (v4 Vec4) Equals(other Vec) bool {
	return equals(v4.ToSlice(), other.ToSlice())
}

func (v4 Vec4) X() float32 {
	return v4.arr[0]
}

func (v4 Vec4) Y() float32 {
	return v4.arr[1]
}

func (v4 Vec4) Z() float32 {
	return v4.arr[2]
}

func (v4 Vec4) W() float32 {
	return v4.arr[3]
}

func NewVec3(x, y, z float32) Vec3 {
	return Vec3{
		arr: [3]float32{
			x,
			y,
			z,
		},
	}
}

func NewVec4(x, y, z, w float32) Vec4 {
	return Vec4{
		arr: [4]float32{
			x,
			y,
			z,
			w,
		},
	}
}
