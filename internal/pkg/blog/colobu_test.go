package blog

import (
	"fmt"
	"testing"
)

func TestNewColobuInfo(t *testing.T) {
	data := NewColobuInfo()
	data.SetWebPage()
	fmt.Println(data.GetData())
}
