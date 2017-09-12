package config

import (
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"log"
)

var AppConfig = struct {
	Server struct {
		LogLevel       string `default:"debug" yaml:"logLevel"` // 因为变量使用的是驼峰命名，需要使用 yaml 标签指定它的对应关系
		Port           int    `default:"8080"`                  // 属性名只有首字母大小，配置文件纯小写的情况则可以省略 yaml 标签
		LogModelEnable bool                                     // 此字段用于控制大部分第三方模块的日志是否开启
	}

	MySQL struct {
		Url       string `required:"true"`
		Username  string `required:"true"`
		Password  string `required:"true"`
		Database  string `required:"true"`
		MaxIdle   int    `default:"10" yaml:"maxIdle"`
		MaxActive int    `default:"50" yaml:"maxActive"`
	} `yaml:"mysql"`
}{}

func init() {
	err := configor.Load(&AppConfig, "config.yml")
	if err != nil {
		log.Fatalf("AppConfig init failed！because %q", err)
	}

	logLevel, err := logrus.ParseLevel(AppConfig.Server.LogLevel)
	if err != nil {
		log.Println(err)
	}
	if logLevel >= 5 {
		AppConfig.Server.LogModelEnable = true
	}

	logrus.SetLevel(logLevel)
}
