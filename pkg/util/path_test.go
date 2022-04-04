package util

import (
	"fmt"
	"path"
	"testing"
)

func TestGetBasePath(t *testing.T) {
	p := GetBasePath()
	fmt.Println(p)
	res := path.Join(p, "gorm.db")
	fmt.Println(res)
}
