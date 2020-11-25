package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	// Cfg is
	Cfg *ini.File

	//RunMode is
	RunMode string

	//HTTPPort is
	HTTPPort int
	//ReadTimeout is
	ReadTimeout time.Duration
	//WriteTimeout is
	WriteTimeout time.Duration

	//PageSize is
	PageSize int
	//JwtSecret is
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

// LoadBase is
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// LoadServer is
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// LoadApp is
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JwtSecret").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PageSize").MustInt(10)
}
