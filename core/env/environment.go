// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-07-30 12:43
// version: 1.0.0
// desc   :

package env

import (
	"os"
	"strings"

	"github.com/yhyzgn/winter/core/prop"
)

// Environment ...
type Environment interface {
	Get(key string) interface{}

	GetWithDefault(key string, def interface{}) interface{}
}

// SystemEnvironment ...
type SystemEnvironment struct {
	p *prop.Properties
}

// NewSystemEnvironment ...
func NewSystemEnvironment() *SystemEnvironment {
	return &SystemEnvironment{p: prop.New()}
}

func (se *SystemEnvironment) loadSystemEnv() {
	for _, env := range os.Environ() {
		kv := strings.SplitN(env, "=", 2)
		if len(kv) == 1 {
			continue
		}

		k, v := kv[0], kv[1]
		if k == "" || v == "" {
			continue
		}

		// 加载系统环境变量
		se.p.Set(k, v)
	}
}

func (se *SystemEnvironment) loadCmdArgs() {
	for i := 0; i < len(os.Args); i++ {

		s := os.Args[i]
		if !strings.HasPrefix(s, "-") {
			continue
		}

		k, v := s[1:], ""
		if i >= len(os.Args)-1 {
			se.p.Set(k, v)
			break
		}

		if !strings.HasPrefix(os.Args[i+1], "-") {
			v = os.Args[i+1]
			i++
		}
		se.p.Set(k, v)
	}
}

// Prepare ...
func (se *SystemEnvironment) Prepare() {
	se.loadSystemEnv()
	se.loadCmdArgs()
}

// Get ...
func (se *SystemEnvironment) Get(key string) interface{} {
	return se.p.Get(key)
}

// GetWithDefault ...
func (se *SystemEnvironment) GetWithDefault(key string, def interface{}) interface{} {
	return se.p.GetWithDefault(key, def)
}
