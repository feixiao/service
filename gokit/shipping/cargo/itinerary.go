package cargo

import (
	"time"

	"github.com/feixiao/learning/gokit/shipping/location"
	"github.com/feixiao/learning/gokit/shipping/voyage"
)

// Leg  表示航行中两个地点之间的运输
type Leg struct {
	VoyageNumber   voyage.Number     `json:"voyage_number"`
	LoadLocation   location.UNLocode `json:"from"`
	UnloadLocation location.UNLocode `json:"to"`
	LoadTime       time.Time         `json:"load_time"`
	UnloadTime     time.Time         `json:"unload_time"`
}

// NewLeg 创建新的Leg
func NewLeg(voyageNumber voyage.Number, loadLocation, unloadLocation location.UNLocode, loadTime, unloadTime time.Time) Leg {
	return Leg{
		VoyageNumber:   voyageNumber,
		LoadLocation:   loadLocation,
		UnloadLocation: unloadLocation,
		LoadTime:       loadTime,
		UnloadTime:     unloadTime,
	}
}

// Itinerary 表示如果从出发点到达目的地
type Itinerary struct {
	Legs []Leg `json:"legs"`
}

// InitialDepartureLocation   返回起点的UNLocode
func (i Itinerary) InitialDepartureLocation() location.UNLocode {
	if i.IsEmpty() {
		return location.UNLocode("")
	}
	return i.Legs[0].LoadLocation
}

// FinalArrivalLocation 返回终点的UNLocode
func (i Itinerary) FinalArrivalLocation() location.UNLocode {
	if i.IsEmpty() {
		return location.UNLocode("")
	}
	return i.Legs[len(i.Legs)-1].UnloadLocation
}

// FinalArrivalTime 返回重点的到达时间
func (i Itinerary) FinalArrivalTime() time.Time {
	return i.Legs[len(i.Legs)-1].UnloadTime
}

// IsEmpty 检查路线是否为空
func (i Itinerary) IsEmpty() bool {
	return i.Legs == nil || len(i.Legs) == 0
}

// IsExpected checks if the given handling event is expected when executing this itinerary.
func (i Itinerary) IsExpected(event HandlingEvent) bool {
	if i.IsEmpty() {
		return true
	}

	switch event.Activity.Type {
	case Receive:
		return i.InitialDepartureLocation() == event.Activity.Location
	case Load:
		for _, l := range i.Legs {
			if l.LoadLocation == event.Activity.Location && l.VoyageNumber == event.Activity.VoyageNumber {
				return true
			}
		}
		return false
	case Unload:
		for _, l := range i.Legs {
			if l.UnloadLocation == event.Activity.Location && l.VoyageNumber == event.Activity.VoyageNumber {
				return true
			}
		}
		return false
	case Claim:
		return i.FinalArrivalLocation() == event.Activity.Location
	}

	return true
}
