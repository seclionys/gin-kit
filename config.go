package gin_kit

type Config struct {
	Mode string
	Log  *LogConfig
}

type LogConfig struct {
	Level      string
	MaxAge     int
	MaxSize    int
	MaxBackups int
}
