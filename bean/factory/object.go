// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-07-30 10:20
// version: 1.0.0
// desc   :

package factory

import "github.com/yhyzgn/winter/bean/errors"

// ObjectFactory 对象工厂
type ObjectFactory interface {
	// GetObject 创建对象
	GetObject() (value interface{}, err errors.BeanError)
}

// ObjectProvider 对象提供者
type ObjectProvider interface {
	// GetObjectByArgs 用参数列表创建对象
	GetObjectByArgs(args ...interface{}) (value interface{}, err errors.BeanError)
	ObjectFactory
}
