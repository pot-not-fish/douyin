package dao

// mysql连接配置
var MySQLConfig = struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}{
	Username: "",
	Password: "",
	Host:     "",
	Port:     3306,
	DBName:   "",
}

// redis连接配置
var RedisConfig = struct {
	Address  string
	Password string
}{
	Address:  "",
	Password: "",
}
