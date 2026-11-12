package graph

import (
	"errors"
	"fmt"
	"math"
)

const eps = 1e-5

type Vec2f struct {
	x, y float32
}

func NewVec2f(x, y float32) Vec2f {
	return Vec2f{x, y}
}

func (v *Vec2f) GetX() float32 {
	return v.x
}

func (v *Vec2f) GetY() float32 {
	return v.y
}

func (v *Vec2f) SetX(x float32) {
	v.x = x
}

func (v *Vec2f) SetY(y float32) {
	v.y = y
}

func (v1 Vec2f) GetLengthSquared() float32 {
	return v1.x*v1.x + v1.y*v1.y
}

func (v1 Vec2f) GetLength() float32 {
	return float32(math.Sqrt(float64(v1.GetLengthSquared())))
}

func (v1 Vec2f) EqualTo(v2 Vec2f) bool {
	return v1.x == v2.x && v1.y == v2.y
}

func (v1 Vec2f) AlmostEqualTo(v2 Vec2f) bool {
	return math.Abs(float64(v1.x-v2.x)) <= eps && math.Abs(float64(v1.y-v2.y)) <= eps
}

func (v1 Vec2f) EqualZero() bool {
	return v1.x == 0 && v1.y == 0
}

func (v1 Vec2f) AlmostEqualZero() bool {
	return math.Abs(float64(v1.x)) <= eps && math.Abs(float64(v1.y)) <= eps
}

func (v1 Vec2f) Add(v2 Vec2f) Vec2f {
	return NewVec2f(v1.x+v2.x, v1.y+v2.y)
}

func (v1 Vec2f) Sub(v2 Vec2f) Vec2f {
	return NewVec2f(v1.x-v2.x, v1.y-v2.y)
}

func (v1 Vec2f) Dot(v2 Vec2f) float32 {
	return v1.x*v2.x + v1.y*v2.y
}

func (v1 Vec2f) Cross(v2 Vec2f) float32 {
	return v1.x*v2.y - v1.y*v2.x
}

func (v1 Vec2f) Scale(k_ float32) Vec2f {
	return NewVec2f(k_*v1.x, k_*v1.y)
}

func (v1 Vec2f) GetDistanceSquaredWith(v2 Vec2f) float32 {
	return v1.Sub(v2).GetLengthSquared()
}

func (v1 Vec2f) GetDistanceWith(v2 Vec2f) float32 {
	return v1.Sub(v2).GetLength()
}

func (v1 Vec2f) GetAngleWith(v2 Vec2f) float32 {
	cross := v1.Cross(v2)
	dot := v1.Dot(v2)
	return float32(math.Atan2(float64(cross), float64(dot)))
}

func (v1 Vec2f) GetNormalized() (Vec2f, error) {
	if v1.AlmostEqualZero() {
		return NewVec2f(0.0, 0.0), errors.New("vector is too close to a zero vector")
	}

	oneOverLength := 1.0 / v1.GetLength()
	return NewVec2f(oneOverLength*v1.x, oneOverLength*v1.y), nil
}

func (v *Vec2f) Print() {
	fmt.Println("\tx:", v.x)
	fmt.Println("\ty:", v.y)
}
