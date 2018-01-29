package cargo

import (
	"errors"
	"time"

	"github.com/feixiao/learning/gokit/shipping/location"
	"github.com/feixiao/learning/gokit/shipping/voyage"
)

// HandlingActivity represents how and where a cargo can be handled, and can
// be used to express predictions about what is expected to happen to a cargo
// in the future.
type HandlingActivity struct {
	Type         HandlingEventType
	Location     location.UNLocode
	VoyageNumber voyage.Number
}

// HandlingEvent 用于注册“何时”的事件，比如：货物在给定的时间点被卸货
// HandlingEvent（处理事件）是对Cargo采取的不同操作，如将它装上船或清关。这个类可以被细化为一个由不同种类的事件（如装货、卸货或由收货人提货）构成的层次结构。
type HandlingEvent struct {
	TrackingID TrackingID			
	Activity   HandlingActivity
}

// HandlingEventType 描述事件类型
type HandlingEventType int

// 枚举有效的事件类型
const (
	NotHandled HandlingEventType = iota
	Load
	Unload
	Receive
	Claim
	Customs
)

// 事件类型转换为字符串
func (t HandlingEventType) String() string {
	switch t {
	case NotHandled:
		return "Not Handled"
	case Load:
		return "Load"
	case Unload:
		return "Unload"
	case Receive:
		return "Receive"
	case Claim:
		return "Claim"
	case Customs:
		return "Customs"
	}

	return ""
}

// HandlingHistory 表示货运类型
type HandlingHistory struct {
	HandlingEvents []HandlingEvent
}

// MostRecentlyCompletedEvent 返回最近完成处理的事件
func (h HandlingHistory) MostRecentlyCompletedEvent() (HandlingEvent, error) {
	if len(h.HandlingEvents) == 0 {
		return HandlingEvent{}, errors.New("delivery history is empty")
	}

	return h.HandlingEvents[len(h.HandlingEvents)-1], nil
}

// HandlingEventRepository 提供事件访问接口.
type HandlingEventRepository interface {
	Store(e HandlingEvent)
	QueryHandlingHistory(TrackingID) HandlingHistory
}

// HandlingEventFactory 创建事件
type HandlingEventFactory struct {
	CargoRepository    Repository
	VoyageRepository   voyage.Repository
	LocationRepository location.Repository
}

// CreateHandlingEvent 创建有效的事件
func (f *HandlingEventFactory) CreateHandlingEvent(registered time.Time, completed time.Time, id TrackingID,
	voyageNumber voyage.Number, unLocode location.UNLocode, eventType HandlingEventType) (HandlingEvent, error) {

	if _, err := f.CargoRepository.Find(id); err != nil {
		return HandlingEvent{}, err
	}

	if _, err := f.VoyageRepository.Find(voyageNumber); err != nil {
		// TODO: This is pretty ugly, but when creating a Receive event, the voyage number is not known.
		if len(voyageNumber) > 0 {
			return HandlingEvent{}, err
		}
	}

	if _, err := f.LocationRepository.Find(unLocode); err != nil {
		return HandlingEvent{}, err
	}

	return HandlingEvent{
		TrackingID: id,
		Activity: HandlingActivity{
			Type:         eventType,
			Location:     unLocode,
			VoyageNumber: voyageNumber,
		},
	}, nil
}
