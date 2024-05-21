package config

import (
	"github.com/BurntSushi/toml"
)

func LoadConfig(P string) *Config {
	if P == "" {
		P = "config.toml"
	}
	var conf Config

	// TOML 파일 로드
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		panic(err)
	}

	// 설정 내용을 출력
	// /fmt.Printf("%+v\n", conf)
	return &conf
}
