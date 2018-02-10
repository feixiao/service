package voyage

import (
	"errors"
	"time"

	"github.com/feixiao/learning/gokit/shipping/location"
)

// Number 唯一标识一个特定的航行
type Number string

// Voyage 表示一系列的货物移动
type Voyage struct {
	Number   Number
	Schedule Schedule
}

// New creates a voyage with a voyage number and a provided schedule.
func New(n Number, s Schedule) *Voyage {
	return &Voyage{Number: n, Schedule: s}
}

// Schedule 表示 航线规划
type Schedule struct {
	CarrierMovements []CarrierMovement //  由一系列满足Specification（规格） 的CarrierMovement（运输动作） 来完成运送目标。
}

// CarrierMovement 表示船只从一个地方航行到另一个地方的。
type CarrierMovement struct {
	DepartureLocation location.UNLocode // 离开位置
	ArrivalLocation   location.UNLocode // 到达位置
	DepartureTime     time.Time         // 离开时间
	ArrivalTime       time.Time         // 到达时间
}

// ErrUnknown 表示找不到对于的航线
var ErrUnknown = errors.New("unknown voyage")

// Repository 提供访问航线接口。
type Repository interface {
	Find(Number) (*Voyage, error)
}
