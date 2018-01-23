package location

import "errors"

//  UNLocode 美国的位置编码，唯一的标示特定的位置。
// http://www.unece.org/cefact/locode/
// http://www.unece.org/cefact/locode/DocColumnDescription.htm#LOCODE
type UNLocode string 

// Location 在我们的模型中表示位置(起点，终点，移动位置)
type Location struct {
	UNLocode UNLocode
	Name  string
}

// ErrUnknown 用于表示当找不到某个位置的时候
var ErrUnknown = errors.New("unknown location")

// Repository 提供查询地理库的接口
type Repository interface {
	Find(locode UNLocode)(*Location, error)
	FindAll() []*Location
}

