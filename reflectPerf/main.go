package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Name    string `json:"server-name"`
	IP      string `json:"server-ip"`
	URL     string `json:"server-url"`
	Timeout string `json:"timeout"`
}

//先从配置文件读取，然后再检查环境变量中是否进行了配置，
//如果环境变量中进行了设置，则以其为准；
func readConfig(path string) *Config {
	var conf Config
	// 1. 读取配置文件
	if len(path) != 0 {
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		json.NewDecoder(f).Decode(&conf)
	}

	// 2. 检查环境变量 环境变量的设置形如 CONFIG_xxx_xxx xxx为json tag
	typ := reflect.TypeOf(conf)
	value := reflect.Indirect(reflect.ValueOf(&conf))
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		// 2.1 找到带有 json Tag 的属性
		if v, ok := field.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exist := os.LookupEnv(key); exist {
				value.FieldByName(field.Name).Set(reflect.ValueOf(env))
			}
		}
	}
	return &conf
}

func main() {
	os.Setenv("CONFIG_SERVER_IP", "124.222.182.143")
	path := flag.String("path", "config.json", "the config filename")
	conf := readConfig(*path)
	fmt.Printf("conf: %v\n", *conf)
}
