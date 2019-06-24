package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nodias/go-ApmCommon/model"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
)

var client = apmhttp.WrapClient(http.DefaultClient)
var config model.TomlConfig

func init() {
	config.New("config.toml")
}

func ApiGetUserInfo(ctx context.Context, id string) ([]byte, error) {
	span, ctx := apm.StartSpan(ctx, "UpdateRequestCount", "custom")
	defer span.End()
	url := fmt.Sprintf("http://%s%s/userInfo/%s", config.Servers["local3"].IP, config.Servers["local3"].PORT, id)
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
