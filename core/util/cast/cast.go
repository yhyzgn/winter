// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-07-30 14:47
// version: 1.0.0
// desc   :

package cast

import (
	"time"

	"github.com/spf13/cast"
)

// ToBool casts an interface{} to a bool.
func ToBool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}

// ToBoolE casts an interface{} to a bool.
func ToBoolE(i interface{}) (bool, error) {
	return cast.ToBoolE(i)
}

// ToInt casts an interface{} to an int.
func ToInt(i interface{}) int {
	v, _ := ToInt64E(i)
	return int(v)
}

// ToInt64 casts an interface{} to an int64.
func ToInt64(i interface{}) int64 {
	v, _ := ToInt64E(i)
	return v
}

// ToInt64E casts an interface{} to an int64.
func ToInt64E(i interface{}) (int64, error) {
	return cast.ToInt64E(i)
}

// ToUint64 casts an interface{} to a uint64.
func ToUint64(i interface{}) uint64 {
	v, _ := ToUint64E(i)
	return v
}

// ToUint64E casts an interface{} to a uint64.
func ToUint64E(i interface{}) (uint64, error) {
	return cast.ToUint64E(i)
}

// ToFloat64 casts an interface{} to a float64.
func ToFloat64(i interface{}) float64 {
	v, _ := ToFloat64E(i)
	return v
}

// ToFloat64E casts an interface{} to a float64.
func ToFloat64E(i interface{}) (float64, error) {
	return cast.ToFloat64E(i)
}

// ToString casts an interface{} to a string.
func ToString(i interface{}) string {
	v, err := ToStringE(i)
	if err != nil {
		return err.Error()
	}
	return v
}

// ToStringE casts an interface{} to a string.
func ToStringE(i interface{}) (string, error) {
	return cast.ToStringE(i)
}

// ToDuration casts an interface{} to a time.Duration.
func ToDuration(i interface{}) time.Duration {
	v, _ := ToDurationE(i)
	return v
}

// ToDurationE casts an interface{} to a time.Duration.
func ToDurationE(i interface{}) (time.Duration, error) {
	return cast.ToDurationE(i)
}

// ToTime casts an interface{} to a time.Time.
func ToTime(i interface{}) time.Time {
	v, _ := ToTimeE(i)
	return v
}

// ToTimeE casts an interface{} to a time.Time.
func ToTimeE(i interface{}) (time.Time, error) {
	return cast.ToTimeE(i)
}
