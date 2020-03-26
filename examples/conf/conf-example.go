package main

import (
	"fmt"
	"time"

	"github.com/xelasov/golib/pkg/conf"
)

// DBConfig stuct of db config info
type DBConfig struct {
	Host     string
	Port     int
	Timeout  time.Duration
	Password string
	User     string
}

// Config - app config type
type Config struct {
	FromDb    DBConfig
	ToDb      DBConfig
	LogLevel  string
	Operation string
	Tables    []string
}

// define your config structure
// NOTE: all types that can be marshaled to/from JSON are supported
var cfg = Config{
	FromDb: DBConfig{
		Host:    "defHostValue",
		Port:    1111,
		Timeout: 1111,
	},
	ToDb: DBConfig{
		Host:    "defHostValue",
		Port:    1111,
		Timeout: 1111,
	},
	LogLevel:  "INFO",
	Operation: "pipe-data",
	Tables:    []string{"tbl1", "tbl2"},
}

func main() {

	// E.g. of a command line: env CONF_FILES="/tmp/doesnotexist.json,./target/config.json" CONF_JSON='{"Operation":"op-one", "LogLevel":"DEBUG"}' ./target/conf-example -conf.files=/tmp/notthere.json,./target/config.jso -conf.print -conf.json='{"Operation":"YYY", "LogLevel":"XXX"}'
	loader := conf.NewJSONConfig("/etc/conf/config.json", "~/config.json", conf.GetDefaultFilePath("config.json"))
	loader.LoadInto(&cfg)

	// use config here...

	fmt.Printf("\n\n=========================\n\n")
	fmt.Printf("Running the app\n")
	fmt.Printf("Using config: %v\n", cfg)
	fmt.Printf("\n\n=========================\n\n")
}
