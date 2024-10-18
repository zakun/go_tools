package config

import (
	"gopkg.in/ini.v1"
)

var Conf *ini.File

func init() {
	if Conf == nil {
		var err error
		Conf, err = ini.Load("my.ini")
		if err != nil {
			panic(err)
		}
	}
}

func Get(key string, section ...any) *ini.Key {
	var sec string
	if len(section) > 0 {
		sec = section[0].(string)
	} else {
		sec = ""
	}

	return Conf.Section(sec).Key(key)
}
