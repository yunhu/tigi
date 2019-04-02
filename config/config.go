package config

type DB struct {
	Dev    Conf
	Online Conf
}

type Conf struct {
	Redis  redis
	Mysql  mysql
	Cache  redis
	Server server
}

type redis struct {
	IP   string `toml:"ip"`
	DB   int    `toml:"db"`
	Pass string `toml:"pass"`
}

type mysql struct {
	IP   string `toml:"ip"`
	DB   string `toml:"db"`
	Pass string `toml:"pass"`
	User string `toml:"user"`
}

type server struct {
	Port int `toml:"port"`
}
