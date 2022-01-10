package env

import (
	"fmt"
	"os"
	"path/filepath"
)

type conf struct {
	RunMode     string `yaml:"RUN_MODE"`
	Port        string `yaml:"PORT"`
	Project     string `yaml:"PROJECT"`
	Domain      string `yaml:"DOMAIN"`
	ProjectName string `yaml:"PROJECT_NAME"`

	JueJine struct {
		Cookie          string `yaml:"COOKIE"`
	} `yaml:"JUEJINE"`
}


//获取当前目录
func GetAppPath() string {
	pwdPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	lastDir := pwdPath
	const myUniqueRelativePath = "configs/.env.yaml"

	for {
		currentPath := fmt.Sprintf("%s/%s", lastDir, myUniqueRelativePath)

		fi, err := os.Stat(currentPath)
		if err == nil {
			switch mode := fi.Mode(); {
			case mode.IsRegular():
				//logrus.Infoln("config path: " + currentPath)
				return lastDir
			}
		}

		newDir := filepath.Dir(lastDir)
		if newDir == "/" || newDir == lastDir {
			panic("not find")
		}
		lastDir = newDir
	}
}
