package main

import (
	"errors"
	"log"
	"net/http"

	redigo "github.com/garyburd/redigo/redis"
)

func main() {

	_, err := RedisConnect()
	if err != nil {
		log.Println("fail ...")
	} else {
		log.Println("success ...")
	}

	_, err = Url()
	if err != nil {
		log.Println("url failed ...")
		log.Println(err)

	} else {
		log.Println("url success ...")
	}

}

func RedisConnect() (string, error) {

	conn, err := redigo.Dial("tcp", ":6379")

	if err != nil {
		return "", err
	}

	defer conn.Close()

	_, err = redigo.String(conn.Do("ping"))

	if err != nil {
		return "", err
	}
	return "success", nil
}

func Url() (string, error) {
	client := http.Client{}
	resp, err := client.Head("http://baidu.com")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {

		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("not 200")
	}

	return "success", nil
}
