package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ApiGetUsers() ([]byte, error) {
	resp, err := http.Get("http://localhost:7003/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("%s\n", string(data))
	return data, nil
}

func ApiGetUser(uid int) ([]byte, error) {
	url := fmt.Sprintf("http://localhost:7003/user/%d", uid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("%s\n", string(data))
	return data, nil
}
