package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var currentWorkDir, _ = os.Getwd()
var licenseFile = filepath.Join(currentWorkDir, "README.md")

func TestMarkdown(t *testing.T) {
	file, _ := os.Open(licenseFile)
	content, _ := ioutil.ReadAll(file)

	res := MdMatchStringByStartAndEnd(string(content), "## 项目代码规范", "## 注意")
	fmt.Println(res)
}
