package config

// Config 全局配置结构
type Config struct {
	App   *Application `toml:"app" yaml:"app" json:"app"`
	MySQL *MySQL       `toml:"mysql" yaml:"mysql" json:"mysql"`
	Redis *Redis       `toml:"redis" yaml:"redis" json:"redis"`
	Log   *Log         `toml:"log" yaml:"log" json:"log"`
}

// Application 应用配置
type Application struct {
	Name string `toml:"name" yaml:"name" json:"name" env:"APP_NAME"`
	Host string `toml:"host" yaml:"host" json:"host" env:"APP_HOST"`
	Port int    `toml:"port" yaml:"port" json:"port" env:"APP_PORT"`
	Mode string `toml:"mode" yaml:"mode" json:"mode" env:"APP_MODE"` // debug, release, test
}

// MySQL 数据库配置
type MySQL struct {
	Host     string `toml:"host" yaml:"host" json:"host" env:"MYSQL_HOST"`
	Port     int    `toml:"port" yaml:"port" json:"port" env:"MYSQL_PORT"`
	User     string `toml:"user" yaml:"user" json:"user" env:"MYSQL_USER"`
	Password string `toml:"password" yaml:"password" json:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" yaml:"database" json:"database" env:"MYSQL_DATABASE"`
	Charset  string `toml:"charset" yaml:"charset" json:"charset" env:"MYSQL_CHARSET"`
}

// Redis 缓存配置
type Redis struct {
	Host     string `toml:"host" yaml:"host" json:"host" env:"REDIS_HOST"`
	Port     int    `toml:"port" yaml:"port" json:"port" env:"REDIS_PORT"`
	Password string `toml:"password" yaml:"password" json:"password" env:"REDIS_PASSWORD"`
	DB       int    `toml:"db" yaml:"db" json:"db" env:"REDIS_DB"`
}

// Log 日志配置
type Log struct {
	Level  string `toml:"level" yaml:"level" json:"level" env:"LOG_LEVEL"`     // debug, info, warn, error
	Format string `toml:"format" yaml:"format" json:"format" env:"LOG_FORMAT"` // json, text
	Output string `toml:"output" yaml:"output" json:"output" env:"LOG_OUTPUT"` // stdout, file
	Path   string `toml:"path" yaml:"path" json:"path" env:"LOG_PATH"`
}

func Default() *Config {
	return &Config{
		App: &Application{
			Name: "bibi-cms",
			Host: "0.0.0.0",
			Port: 8080,
			Mode: "debug",
		},
		MySQL: &MySQL{
			Host:    "127.0.0.1",
			Port:    3306,
			Charset: "utf8mb4",
		},
		Redis: &Redis{
			Host: "127.0.0.1",
			Port: 6379,
			DB:   0,
		},
		Log: &Log{
			Level:  "info",
			Format: "json",
			Output: "stdout",
		},
	}
}
