package config

type Conf struct {
	Redis []redis
	Mysql []mysql
	Port  int
}

type redis struct {
	IP      string `toml:"ip"`
	DbIndex int64  `toml:"dbIndex"`
	Pass    string `toml:"pass"`
	DbName  string `toml:"dbName"`
}

type mysql struct {
	IP      string `toml:"ip"`
	DbName  string `toml:"dbName"`
	Pass    string `toml:"pass"`
	User    string `toml:"user"`
	Timeout int64  `toml:"timeout"`
}
