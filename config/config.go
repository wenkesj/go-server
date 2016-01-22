package config;

import (
  "log"
  "encoding/json";
  "io/ioutil";
  "path/filepath";
  "github.com/wenkesj/go-server/httpserver";
);

type RedisConfig struct {
  Host string `json:"host"`;
  Port string `json:"port"`;
};

type MongoConfig struct {
  URL string `json:"url"`;
  Port string `json:"port"`;
  Database string `json:"database"`;
};

type Config struct {
  Server httpserver.HTTPServer `json:"server"`;
  Mongo *MongoConfig `json:"mongo"`;
  Redis *RedisConfig `json:"redis"`;
};

func Server(env string) (*Config, error) {
  log.Printf("Config loading ... \n");
  config := &Config{};
  var (
    data []byte;
    err  error;
  );
  path := filepath.Join("config", "config." + env + ".json");
  if data, err = ioutil.ReadFile(path); err != nil {
    return nil, err;
  }
  log.Printf("Config opened. \n");
  return config, json.Unmarshal(data, config);
}
