package handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type requiredParams struct {
	Method     string `form:"method,required" json:"method"`
	MerchantId string `form:"merchant_id,required" json:"merchant_id"`
	BizParams  string `form:"biz_params,required" json:"biz_params"`
}

var SvcMap = make(map[string]genericclient.Client)

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {
	svcName := c.Param("svc")
	svcMethod := c.Param("method")
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, errors.New("ErrBadRequest"))
		return
	}
	// 这里不能直接传入 c.RequestBodyStream() ， body 会被吞掉，还没看源码，可能是因为原始的 body 不支持重复读吧
	req, err := http.NewRequest(http.MethodPost, "", bytes.NewReader(c.Request.BodyBytes()))
	if err != nil {
		hlog.Warnf("new http request failed: %v", err)
		c.JSON(http.StatusOK, errors.New("ErrRequestServerFail"))
		return
	}
	req.URL.Path = fmt.Sprintf("/%s/%s", svcName, svcMethod)

	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		hlog.Errorf("convert request failed: %v", err)
		c.JSON(http.StatusOK, errors.New("ErrServerHandleFail"))
		return
	}
	resp, err := cli.GenericCall(ctx, "", customReq)
	respMap := make(map[string]interface{})
	if err != nil {
		hlog.Errorf("GenericCall err:%v", err)
		bizErr, ok := kerrors.FromBizStatusError(err)
		if !ok {
			c.JSON(http.StatusOK, errors.New("ErrServerHandleFail"))
			return
		}
		respMap["ErrCode"] = bizErr.BizStatusCode()
		respMap["ErrMessage"] = bizErr.BizMessage()
		c.JSON(http.StatusOK, respMap)
		return
	}
	realResp, ok := resp.(*generic.HTTPResponse)
	if !ok {
		c.JSON(http.StatusOK, errors.New("ErrServerHandleFail"))
		return
	}
	realResp.Body["ErrCode"] = 0
	realResp.Body["ErrMessage"] = "ok"
	c.JSON(http.StatusOK, realResp.Body)
}
