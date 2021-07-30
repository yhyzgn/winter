// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-07-30 11:25
// version: 1.0.0
// desc   :

package definition

import (
	"math"
	"reflect"
)

// Singleton ...
const (
	Singleton uint = iota // 单例，默认
	Prototype             // 原型
)

// HighestOrder ...
const (
	HighestOrder = math.MinInt32
	LowestOrder  = math.MaxInt32
)

type beanStatus int

// Default ...
const (
	Default   = beanStatus(0) // 默认状态
	Resolving = beanStatus(1) // 正在决议
	Resolved  = beanStatus(2) // 已决议
	Wiring    = beanStatus(3) // 正在注入
	Wired     = beanStatus(4) // 注入完成
	Deleted   = beanStatus(5) // 已删除
)

// BeanDefinition bean 信息定义
type BeanDefinition struct {
	typeName string
	name     string
	typ      reflect.Type
	val      reflect.Value
	scope    uint
	status   beanStatus // 状态
	primary  bool       // 是否为主版本
	order    int        // 收集时的顺序
}

// New ...
func New(name string, typ reflect.Type) *BeanDefinition {
	return &BeanDefinition{
		name:    name,
		typ:     typ,
		scope:   Singleton,
		status:  Default,
		primary: true,
		order:   LowestOrder,
	}
}
