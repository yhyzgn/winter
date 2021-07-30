// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-07-30 10:03
// version: 1.0.0
// desc   :

package factory

import (
	"reflect"

	"github.com/yhyzgn/winter/bean/errors"
)

// BeanFactory
//
// bean 工厂
type BeanFactory interface {
	// GetByName 根据 name 获取 bean
	GetByName(name string) (value interface{}, err errors.BeanError)

	// GetByNameAndType 根据 name 和类型获取 bean
	GetByNameAndType(name string, typ reflect.Type) (value interface{}, err errors.BeanError)

	// GetByNameAndArgs 根据 name 和构造函数参数列表获取 bean
	GetByNameAndArgs(name string, args ...interface{}) (value interface{}, err errors.BeanError)

	// GetByType 根据类型获取 bean
	GetByType(typ reflect.Type) (value interface{}, err errors.BeanError)

	// GetByTypeAndArgs 根据类型和构造函数参数列表获取 bean
	GetByTypeAndArgs(typ reflect.Type, args ...interface{}) (value interface{}, err errors.BeanError)

	// ContainsBeanByName 根据 name 判断 bean 是否存在
	ContainsBeanByName(name string) (contains bool)

	// ContainsBeanByType 根据类型判断 bean 是否存在
	ContainsBeanByType(typ reflect.Type) (contains bool)

	// GetBeanProviderByType 根据类型获取实例提供者
	GetBeanProviderByType(typ reflect.Type) (provider ObjectProvider)

	// IsSingletonByName 根据 name 判断是否是单例 bean
	IsSingletonByName(name string) (is bool, err errors.NoSuchBeanDefinitionError)

	// IsSingletonByType 根据类型判断是否是单例 bean
	IsSingletonByType(typ reflect.Type) (is bool, err errors.NoSuchBeanDefinitionError)

	// IsPrototypeByName 根据 name 判断是否是原型 bean
	IsPrototypeByName(name string) (is bool, err errors.NoSuchBeanDefinitionError)

	// IsPrototypeByType 根据类型判断是否是原型 bean
	IsPrototypeByType(typ reflect.Type) (is bool, err errors.NoSuchBeanDefinitionError)

	// IsTypeMatched 类型是否匹配
	IsTypeMatched(typ reflect.Type) (matched bool, err errors.NoSuchBeanDefinitionError)

	// GetType 根据 name 获取 bean 的类型
	GetType(name string) (typ reflect.Type, err errors.NoSuchBeanDefinitionError)
}
