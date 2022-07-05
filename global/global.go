package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"scry/config"
)

var (
	Mysql  *gorm.DB
	Viper  *viper.Viper
	Config config.Config
)
