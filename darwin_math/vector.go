package darwin_math

import (
	"math"

	proto "github.com/anirul/go_darwin_server/darwin_proto"
)

func CreateVector2(x float64, y float64) proto.Vector2 {
	return proto.Vector2{X: x, Y: y}
}

func CreateVector3(x float64, y float64, z float64) proto.Vector3 {
	return proto.Vector3{X: x, Y: y, Z: z}
}

func CreateVector4(x float64, y float64, z float64, w float64) proto.Vector4 {
	return proto.Vector4{X: x, Y: y, Z: z, W: w}
}

func Dot(l *proto.Vector3, r *proto.Vector3) float64 {
	return l.X*r.X + l.Y*r.Y + l.Z*r.Z
}

func Length(vec3 *proto.Vector3) float64 {
	return math.Sqrt(Dot(vec3, vec3))
}

func Minus(l *proto.Vector3, r *proto.Vector3) *proto.Vector3 {
	return &proto.Vector3{X: l.X - r.X, Y: l.Y - r.Y, Z: l.Z - r.Z}
}

func Add(l *proto.Vector3, r *proto.Vector3) *proto.Vector3 {
	return &proto.Vector3{X: l.X + r.X, Y: l.Y + r.Y, Z: l.Z + r.Z}
}

func Distance(l *proto.Vector3, r *proto.Vector3) float64 {
	return Length(Minus(r, l))
}

func Cross(l *proto.Vector3, r *proto.Vector3) *proto.Vector3 {
	return &proto.Vector3{
		X: l.Y*r.Z - l.Z*r.Y,
		Y: l.Z*r.X - l.X*r.Z,
		Z: l.X*r.Y - l.Y*r.X,
	}
}

func Normalize(vec3 *proto.Vector3) *proto.Vector3 {
	length := Length(vec3)
	return &proto.Vector3{
		X: vec3.X / length,
		Y: vec3.Y / length,
		Z: vec3.Z / length,
	}
}

func Times(vec3 *proto.Vector3, val float64) *proto.Vector3 {
	return &proto.Vector3{
		X: vec3.X * val,
		Y: vec3.Y * val,
		Z: vec3.Z * val,
	}
}

func ProjectOnPlane(vec3 *proto.Vector3, normal *proto.Vector3) *proto.Vector3 {
	dotVN := Dot(vec3, normal)
	normalSqrd := Dot(normal, normal)
	projVonN := Times(normal, dotVN/normalSqrd)
	vPlane := Minus(vec3, projVonN)
	return Times(Normalize(vPlane), Length(vec3))
}
