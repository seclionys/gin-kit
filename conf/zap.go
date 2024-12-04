package conf

type ZapConfig struct {
	Level      string
	MaxAge     int
	MaxSize    int
	MaxBackups int
	LogPath    string
}
