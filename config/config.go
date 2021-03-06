package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type Env struct {
	DEBUG    bool   `env:"DEBUG"`
	ENV      string `env:"ENV"`
	DATAPATH string `env:"DATAPATH"`
}

const projectDirName = "p2pdb-server" // change to relevant project name

func init() {
	// projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	// currentWorkDirectory, _ := os.Getwd()
	// rootPath := projectName.Find([]byte(currentWorkDirectory))
	// fmt.Printf("%+v\n", rootPath)
	// err := godotenv.Load(string(rootPath) + `/.env`)
	// //err := godotenv.Load(`../.env`)
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
}

func IsDebug() bool {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
	return conf.DEBUG
}

func GetEnv() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
	return conf.ENV
}

func GetDataPath() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
	return conf.DATAPATH
}
