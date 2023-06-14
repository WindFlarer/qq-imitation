package config

type ConfStruct struct {
	Mysql Mysql `yaml:"mysql"`
}

var Conf ConfStruct
