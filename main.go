package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/joho/godotenv"
)

func main() {
	// 处理参数
	motion := os.Args[1]
	env := os.Args[2]
	environment := firstToUp(env)
	godotenv.Load("config/env-" + env)

	// 连接 Etcd
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{
			"0.0.0.0:2379",
		},
		DialTimeout: 5 * time.Second,
	})
	kv := clientv3.NewKV(cli)

	if motion == "put" {
		// Mysql
		kv.Put(context.TODO(), "/"+environment+"/Mysql/DB_HOST", os.Getenv("DB_HOST"))
		kv.Put(context.TODO(), "/"+environment+"/Mysql/DB_PORT", os.Getenv("DB_PORT"))
		kv.Put(context.TODO(), "/"+environment+"/Mysql/DB_DATABASE", os.Getenv("DB_DATABASE"))
		kv.Put(context.TODO(), "/"+environment+"/Mysql/DB_USERNAME", os.Getenv("DB_USERNAME"))
		kv.Put(context.TODO(), "/"+environment+"/Mysql/DB_PASSWORD", os.Getenv("DB_PASSWORD"))
		kv.Put(context.TODO(), "/"+environment+"/Mysql/DB_CHARSET", os.Getenv("DB_CHARSET"))
		kv.Put(context.TODO(), "/"+environment+"/Mysql/DB_MAX_CONNECTIONS", os.Getenv("DB_MAX_CONNECTIONS"))
		kv.Put(context.TODO(), "/"+environment+"/Mysql/DB_MAX_OPEN_CONNECTIONS", os.Getenv("DB_MAX_OPEN_CONNECTIONS"))

		// Redis
		kv.Put(context.TODO(), "/"+environment+"/Redis/CACHE_HOST", os.Getenv("CACHE_HOST"))
		kv.Put(context.TODO(), "/"+environment+"/Redis/CACHE_PORT", os.Getenv("CACHE_PORT"))
		kv.Put(context.TODO(), "/"+environment+"/Redis/CACHE_PASS", os.Getenv("CACHE_PASS"))
	} else if motion == "get" {
		var data = make(map[string]interface{})

		// Mysql
		DB_HOST, _ := kv.Get(context.TODO(), "/"+environment+"/Mysql/DB_HOST")
		DB_PORT, _ := kv.Get(context.TODO(), "/"+environment+"/Mysql/DB_PORT")
		DB_DATABASE, _ := kv.Get(context.TODO(), "/"+environment+"/Mysql/DB_DATABASE")
		DB_USERNAME, _ := kv.Get(context.TODO(), "/"+environment+"/Mysql/DB_USERNAME")
		DB_PASSWORD, _ := kv.Get(context.TODO(), "/"+environment+"/Mysql/DB_PASSWORD")
		DB_CHARSET, _ := kv.Get(context.TODO(), "/"+environment+"/Mysql/DB_CHARSET")
		DB_MAX_CONNECTIONS, _ := kv.Get(context.TODO(), "/"+environment+"/Mysql/DB_MAX_CONNECTIONS")
		DB_MAX_OPEN_CONNECTIONS, _ := kv.Get(context.TODO(), "/"+environment+"/Mysql/DB_MAX_OPEN_CONNECTIONS")
		data["DB_HOST"] = string(DB_HOST.Kvs[0].Value)
		data["DB_PORT"] = string(DB_PORT.Kvs[0].Value)
		data["DB_DATABASE"] = string(DB_DATABASE.Kvs[0].Value)
		data["DB_USERNAME"] = string(DB_USERNAME.Kvs[0].Value)
		data["DB_PASSWORD"] = string(DB_PASSWORD.Kvs[0].Value)
		data["DB_CHARSET"] = string(DB_CHARSET.Kvs[0].Value)
		data["DB_MAX_CONNECTIONS"] = string(DB_MAX_CONNECTIONS.Kvs[0].Value)
		data["DB_MAX_OPEN_CONNECTIONS"] = string(DB_MAX_OPEN_CONNECTIONS.Kvs[0].Value)

		// Redis
		CACHE_HOST, _ := kv.Get(context.TODO(), "/"+environment+"/Redis/CACHE_HOST")
		CACHE_PORT, _ := kv.Get(context.TODO(), "/"+environment+"/Redis/CACHE_PORT")
		CACHE_PASS, _ := kv.Get(context.TODO(), "/"+environment+"/Redis/CACHE_PASS")
		data["CACHE_HOST"] = string(CACHE_HOST.Kvs[0].Value)
		data["CACHE_PORT"] = string(CACHE_PORT.Kvs[0].Value)
		data["CACHE_PASS"] = string(CACHE_PASS.Kvs[0].Value)

		// Print
		for k, v := range data {
			fmt.Printf("{%v} => {%v}\n", k, v)
		}
	}
}

// 首字母大写
func firstToUp(in string) (ret string) {
	first := strings.ToUpper(string([]rune(in)[:1]))
	last := string([]rune(in)[1:])
	return first + last
}
