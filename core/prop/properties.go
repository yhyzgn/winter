// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-07-30 12:40
// version: 1.0.0
// desc   :

package prop

import (
	"fmt"
	"reflect"

	"github.com/yhyzgn/winter/core/util/cast"
)

// Properties ...
type Properties struct {
	mp map[string]string
}

// New 返回一个空的属性列表。
func New() *Properties {
	return &Properties{mp: make(map[string]string)}
}

// Set 设置属性
func (p *Properties) Set(key string, val interface{}) {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Map:
		for _, k := range v.MapKeys() {
			mapKey := cast.ToString(k.Interface())
			mapValue := v.MapIndex(k).Interface()
			p.Set(key+"."+mapKey, mapValue)
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			subKey := fmt.Sprintf("%s[%d]", key, i)
			subValue := v.Index(i).Interface()
			p.Set(subKey, subValue)
		}
	default:
		p.mp[key] = cast.ToString(val)
	}
}

// Get ...
func (p *Properties) Get(key string) interface{} {
	return p.GetWithDefault(key, nil)
}

// GetWithDefault ...
func (p *Properties) GetWithDefault(key string, def interface{}) interface{} {
	if val, ok := p.mp[key]; ok {
		return val
	}

	if nil != def {
		return cast.ToString(def)
	}
	return nil
}
