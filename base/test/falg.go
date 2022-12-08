package test

import "flag"

type Config struct {
	Host       string
	ListenPort int
	User       string
	Password   string
	Debug      bool
}

var GlobalConfig Config

func HH() {
	flag.StringVar(&GlobalConfig.Host, "h", "127.0.0.1", "监听的IP")
	flag.IntVar(&GlobalConfig.ListenPort, "p", 8080, "监听的端口")
	flag.StringVar(&GlobalConfig.User, "u", "root", "登陆的用户")
	flag.StringVar(&GlobalConfig.Password, "pwd", "123456", "用户密码")
	flag.BoolVar(&GlobalConfig.Debug, "d", false, "是否开启debug模式")
	flag.Parse()
}
