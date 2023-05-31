package db

type Config struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Address      string `json:"address"`
	Port         int    `json:"port"`
	DbName       string `json:"db_name"`
	DialTimeout  string `json:"dial_timeout"`
	ReadTimeout  string `json:"read_timeout"`
	WriteTimeout string `json:"write_timeout"`
}
