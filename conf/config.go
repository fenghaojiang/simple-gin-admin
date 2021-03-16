package conf

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	IsDev     bool         `toml:"IsDev"`
	Port      string       `toml:"Port"`
	MySQLConf *MySQLConfig `toml:"MySQL"`
}

type MySQLConfig struct {
	Addr     string `toml:"addr"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbName"`
}

var GlobalConfig = new(Config)

func Parse() error {
	if _, err := toml.DecodeFile("./config.toml", &GlobalConfig); err != nil {
		return err
	}
	return nil
}
