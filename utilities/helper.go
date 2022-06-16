package utilities

import(
	"github.com/spf13/cast"
)


func IsEmpty(s interface{}) bool {
	if s == nil || cast.ToString(s) =="" {
		return true;
	}
	return false;
}