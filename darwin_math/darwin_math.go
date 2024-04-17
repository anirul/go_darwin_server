package darwin_math

import (
	"math"
	"math/rand"
	"time"

	proto "github.com/anirul/go_darwin_server/darwin_proto"
)

func RandomVec3() *proto.Vector3 {
	return &proto.Vector3{
		X: 2*rand.Float64() - 1,
		Y: 2*rand.Float64() - 1,
		Z: 2*rand.Float64() - 1,
	}
}

func RadiusFromVolume(volume float64) float64 {
	return math.Cbrt((3 * volume) / (4 * math.Pi))
}

func IsIntersecting(p1 *proto.Physic, p2 *proto.Physic) bool {
	distance := Distance(p1.Position, p2.Position)
	sumRadius := p1.Radius + p2.Radius
	return distance < sumRadius
}

func IsAlmostIntersecting(p1 *proto.Physic, p2 *proto.Physic) bool {
	dot := Dot(Normalize(p1.Position), Normalize(p2.Position))
	return dot > AlmostIntersect
}

func TimeSecondNow() float64 {
	return float64(time.Now().UnixNano()) / 1e9
}
