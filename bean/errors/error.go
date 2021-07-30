// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-07-30 10:10
// version: 1.0.0
// desc   :

package errors

// BeanError bean 错误
type BeanError struct {
	msg string
}

// Error ...
func (b *BeanError) Error() string {
	return b.msg
}

// NewBeanError ...
func NewBeanError(msg string) *BeanError {
	return &BeanError{msg: msg}
}

// NoSuchBeanDefinitionError bean 定义错误
type NoSuchBeanDefinitionError struct {
	beanName string
	BeanError
}

// Error ...
func (n *NoSuchBeanDefinitionError) Error() string {
	panic("implement me")
}

// NewNoSuchBeanDefinitionError ...
func NewNoSuchBeanDefinitionError(beanName string) *NoSuchBeanDefinitionError {
	return &NoSuchBeanDefinitionError{beanName: beanName}
}
