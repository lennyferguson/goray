package src

import (
	"math"
	"testing"
)

func compareFloat32(a float32, b float32) bool {
	return math.Abs(float64(a-b)) < delta
}

func TestNewVec3(t *testing.T) {
	v1 := NewVec3(0.0, 0.0, 0.0)
	if !compareFloat32(v1.X(), 0.0) ||
		!compareFloat32(v1.Y(), 0.0) ||
		!compareFloat32(v1.Z(), 0.0) {
		t.Error("failed to correctly instantiate Vec 3")
	}

	if !v1.Equals(NewVec3(0.0, 0.0, 0.0)) {
		t.Error("Compare function does not work")
	}
}

func TestAddVec3(t *testing.T) {
	v1 := NewVec3(1.0, 1.0, 1.0)
	v2 := NewVec3(1.0, 2.0, 3.0)
	v3 := v1.Add(v2)
	expect := NewVec3(2.0, 3.0, 4.0)
	if !v3.Equals(expect) {
		t.Errorf("Expected %v does not match actual %v", expect, v3)
	}
}

func TestSubVec3(t *testing.T) {
	v1 := NewVec3(1.0, 1.0, 1.0)
	v2 := NewVec3(1.0, 2.0, 3.0)
	v3 := v2.Sub(v1)
	expect := NewVec3(0.0, 1.0, 2.0)
	if !v3.Equals(expect) {
		t.Errorf("Expected %v does not match actual %v", expect, v3)
	}
}

func TestMulVec3(t *testing.T) {
	v1 := NewVec3(2.0, 2.0, 2.0)
	v2 := NewVec3(1.0, 2.0, 3.0)
	v3 := v1.Mul(v2)
	expect := NewVec3(2.0, 4.0, 6.0)
	if !v3.Equals(expect) {
		t.Errorf("Expected %v does not match actual %v", expect, v3)
	}
}

func TestDivVec3(t *testing.T) {
	v1 := NewVec3(2.0, 2.0, 6.0)
	v2 := NewVec3(1.0, 2.0, 2.0)
	v3 := v1.Div(v2)
	expect := NewVec3(2.0, 1.0, 3.0)
	if !v3.Equals(expect) {
		t.Errorf("Expected %v does not match actual %v", expect, v3)
	}
}

func TestScalarMulVec3(t *testing.T) {
	v1 := NewVec3(1.0, 1.0, 1.0)
	v3 := v1.ScalarMul(3.0)
	expect := NewVec3(3.0, 3.0, 3.0)
	if !v3.Equals(expect) {
		t.Errorf("Expected %v does not match actual %v", expect, v3)
	}
}

func TestScalarDivVec3(t *testing.T) {
	v1 := NewVec3(4.0, 4.0, 4.0)
	v3 := v1.ScalarDiv(2.0)
	expect := NewVec3(2.0, 2.0, 2.0)
	if !v3.Equals(expect) {
		t.Errorf("Expected %v does not match actual %v", expect, v3)
	}
}

func TestScalarAddVec3(t *testing.T) {
	v1 := NewVec3(1.0, 1.0, 1.0)
	v3 := v1.ScalarAdd(1.0)
	expect := NewVec3(2.0, 2.0, 2.0)
	if !v3.Equals(expect) {
		t.Errorf("Expected %v does not match actual %v", expect, v3)
	}
}

func TestScalarSubVec3(t *testing.T) {
	v1 := NewVec3(2.0, 2.0, 2.0)
	v3 := v1.ScalarSub(1.0)
	expect := NewVec3(1.0, 1.0, 1.0)
	if !v3.Equals(expect) {
		t.Errorf("Expected %v does not match actual %v", expect, v3)
	}
}

func TestDotVec3(t *testing.T) {
	v1 := NewVec3(2.0, 2.0, 2.0)
	c := v1.Dot(v1)
	expect := float32(12.0)
	if !compareFloat32(c, expect) {
		t.Errorf("Expected %v does not match actual %v", expect, c)
	}
}

func TestNewVec4(t *testing.T) {
	v4 := NewVec4(1.0, 1.0, 1.0, 1.0)
	if !compareFloat32(v4.X(), 1.0) ||
		!compareFloat32(v4.Y(), 1.0) ||
		!compareFloat32(v4.Z(), 1.0) ||
		!compareFloat32(v4.W(), 1.0) {
		t.Errorf("Expected Vec4 [%v] to only contain 1.0", v4)
	}
	v3 := NewVec3(1.0, 1.0, 1.0)
	if v4.Equals(v3) {
		t.Errorf("%v should not equal %v", v4, v3)
	}
	if !v4.Equals(v4) {
		t.Errorf("%v should equal itself", v4)
	}
}

func TestAddVec4(t *testing.T) {
	u := NewVec4(1.0, 1.0, 1.0, 1.0)
	v := NewVec4(1.0, 2.0, 3.0, 4.0)
	w := u.Add(v)
	expect := NewVec4(2.0, 3.0, 4.0, 5.0)
	if !w.Equals(expect) {
		t.Errorf("actual value of %v + %v := %v does not match actual %v", u, v, expect, w)
	}
}

func TestSubVec4(t *testing.T) {
	u := NewVec4(1.0, 1.0, 1.0, 1.0)
	v := NewVec4(1.0, 2.0, 3.0, 4.0)
	w := v.Sub(u)
	expect := NewVec4(0.0, 1.0, 2.0, 3.0)
	if !w.Equals(expect) {
		t.Errorf("actual value of %v - %v := %v does not match actual %v", v, u, expect, w)
	}
}
