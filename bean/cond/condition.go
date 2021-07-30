// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-07-30 11:49
// version: 1.0.0
// desc   :

package cond

import "github.com/yhyzgn/winter/bean/factory"

// Condition ...
type Condition interface {
	// Matches 条件匹配
	Matches(fct factory.BeanFactory) (matches bool, err error)
}

// not 对一个条件进行取反的 Condition 实现。
type not struct {
	c Condition
}

// Matches ...
func (n *not) Matches(fct factory.BeanFactory) (matches bool, err error) {
	matches, err = n.c.Matches(fct)
	return !matches, err
}

// onProperty 基于属性值匹配的 Condition 实现。
type onProperty struct {
	name           string
	havingValue    string
	matchIfMissing bool
}

// Matches ...
func (op *onProperty) Matches(fct factory.BeanFactory) (matches bool, err error) {
	return
}
