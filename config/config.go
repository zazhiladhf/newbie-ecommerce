package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App        App        `yaml:"app"`
	DB         DB         `yaml:"db"`
	Redis      Redis      `yaml:"redis"`
	Meili      Meili      `yaml:"meili"`
	JWT        JWT        `yaml:"jwt"`
	Cloudinary Cloudinary `yaml:"cloudinary"`
	Payment    Payment    `yaml:"payment"`
	MongoDB    MongoDB    `yaml:"mongo"`
}

type App struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	LifeTime int    `yaml:"lifeTime"`
}

type Meili struct {
	Host   string `yaml:"host"`
	ApiKey string `yaml:"key"`
}

type JWT struct {
	SecretKey string `yaml:"key"`
}

type Cloudinary struct {
	Name      string `yaml:"name"`
	ApiKey    string `yaml:"apiKey"`
	ApiSecret string `yaml:"apiSecret"`
}

type Payment struct {
	SecretKey       string         `yaml:"secretKey"`
	Redirect        redirectConfig `yaml:"redirectUrl"`
	InvoiceDuration int            `yaml:"invoiceDuration"`
}

type redirectConfig struct {
	Success string `yaml:"success"`
	Failure string `yaml:"failure"`
}

type MongoDB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Driver   string `yaml:"driver"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

var Cfg *Config

func LoadConfig(filename string) (err error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	cfg := Config{}

	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return
	}

	Cfg = &cfg
	return
}
