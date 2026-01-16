package option

import (
	"api-service-template/pkg/database"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type (
	// Options 配置
	Options struct {
		viper *viper.Viper

		db *gorm.DB

		ConfigFile string `yaml:"-"`

		ListenPort int             `yaml:"listenPort"`
		Database   database.Option `yaml:"database"`
	}
)

// New 构造函数
func New(file string) *Options {
	opt := &Options{
		viper:      viper.New(),
		ConfigFile: file,
	}

	return opt
}

// Parse 解析参数
func (opt *Options) Parse() error {
	if opt.ConfigFile != "" {
		opt.viper.SetConfigFile(opt.ConfigFile)

		if err := opt.viper.ReadInConfig(); err != nil {
			return err
		}
	}

	opt.viper.Unmarshal(opt, func(c *mapstructure.DecoderConfig) {
		c.TagName = "yaml"
	})

	return nil
}

// Prepare 初始化
func (opt *Options) Prepare() error {
	db, err := database.NewDB(opt.Database)
	if err != nil {
		return err
	}

	opt.db = db
	return nil
}

// GetDB 获取数据库连接
func (opt *Options) GetDB() *gorm.DB {
	return opt.db
}
