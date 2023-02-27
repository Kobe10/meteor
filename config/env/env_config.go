package env

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 环境常量配置
const (
	ZEnvLocal   string = "local"
	ZEnvDev     string = "dev"
	ZEnvFeature string = "feature"
	ZEnvPreprod string = "preprod"
	ZEnvProd    string = "prod"
)

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
