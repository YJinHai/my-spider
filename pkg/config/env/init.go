package env

import (
	"fmt"
	"io/ioutil"
	"my-spider/pkg/log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type conf struct {
	RunMode     string `yaml:"RUN_MODE"`
	Port        string `yaml:"PORT"`
	Project     string `yaml:"PROJECT"`
	Domain      string `yaml:"DOMAIN"`
	ProjectName string `yaml:"PROJECT_NAME"`

	JueJine struct {
		Cookie string `yaml:"COOKIE"`
	} `yaml:"JUEJINE"`
}

var Conf *conf

func init() {
	getConf()
}

//获取env配置
func getConf() {
	envPath := GetAppPath() + "/configs/.env.yaml"

	yamlFile, err := ioutil.ReadFile(envPath)
	if err != nil {
		log.Logger.Log(log.LevelDebug, "yamlFile.get err ", err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		log.Logger.Log(log.LevelDebug, "unmarchar", err)
	}
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
				//logrus.Infoln("configs path: " + currentPath)
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
