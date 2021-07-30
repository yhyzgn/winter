// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-07-30 14:14
// version: 1.0.0
// desc   :

package env

import "testing"

func TestLoadSystemEnv(t *testing.T) {
	se := NewSystemEnvironment()

	se.Prepare()

	goRoot := se.Get("GOROOT")
	t.Log(goRoot)
}
