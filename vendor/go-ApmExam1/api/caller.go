package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go-ApmCommon/model"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
)

var client = apmhttp.WrapClient(http.DefaultClient)
var config model.TomlConfig

func Init() {
	config = *model.GetConfig()
}

func ApiGetUserInfo(ctx context.Context, id string) ([]byte, error) {
	span, ctx := apm.StartSpan(ctx, "UpdateRequestCount", "custom")
	defer span.End()
	url := fmt.Sprintf("http://%s%s/userInfo/%s", config.Servers["ApmExam2"].IP, config.Servers["ApmExam2"].PORT, id)
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
