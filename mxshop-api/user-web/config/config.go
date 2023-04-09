package config

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type LogConfig struct {
	Level    string `mapstructure:"level"`
	FilePath string `mapstructure:"file_path"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	Log         LogConfig     `mapstructure:"log"`
}
