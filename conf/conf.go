package conf

type Server struct {
	Mode string
	Host string
	Port string
	Zap  ZapConfig
}
