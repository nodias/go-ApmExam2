package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
)

var client = apmhttp.WrapClient(http.DefaultClient)

func ApiGetUserInfo(ctx context.Context, id string) ([]byte, error) {
	span, ctx := apm.StartSpan(ctx, "UpdateRequestCount", "custom")
	defer span.End()
	url := fmt.Sprintf("http://localhost:7003/userInfo/%s", id)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("%s\n", string(data))
	return data, nil
}

func UpdateRequestCount(ctx context.Context, name int) ([]byte, error) {
	span, ctx := apm.StartSpan(ctx, "updateRequestCount", "custom")
	defer span.End()
	req, _ := http.NewRequest("GET", "http://localhost:7003/hello/world", nil)

	client := apmhttp.WrapClient(http.DefaultClient)
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("%s\n", string(data))
	return data, nil
}

func ApiGetUsers(ctx context.Context) ([]byte, error) {
	span, ctx := apm.StartSpan(ctx, "ApiGetUsers", "custom")
	defer span.End()
	req, _ := http.NewRequest("GET", "http://localhost:7003/users", nil)
	resp, err := client.Do(req.WithContext(ctx))
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

func ApiGetUser(ctx context.Context, uid int) ([]byte, error) {
	span, ctx := apm.StartSpan(ctx, "ApiGetUser", "custom")
	defer span.End()
	url := fmt.Sprintf("http://localhost:7003/user/%d", uid)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req.WithContext(ctx))
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
