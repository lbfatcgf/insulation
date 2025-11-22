package hashutil

import "github.com/tjfoc/gmsm/sm3"

func Sm3(data []byte) []byte {
	return sm3.Sm3Sum(data)
}

func Sm3String(data string) string {
	return string(sm3.Sm3Sum([]byte(data)))
}
