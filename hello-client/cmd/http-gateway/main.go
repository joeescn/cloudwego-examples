package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joeescn/hello-svc/kitex_gen/api"
)

func main() {

	call("")
	call("http-gateway")

}

func call(message string) {
	url := "http://127.0.0.1:8080/gateway/hello-svc/echo"

	req := &api.Request{Message: message}
	payload := new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(req)
	if err != nil {
		klog.Error(err)
		return
	}

	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		klog.Error(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		klog.Error(err)
		return
	}
	klog.Info(string(body))
}
