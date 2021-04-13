/**
 * @Author: yutaoluo@tencent.com
 * @Description: 调用python服务相关page service
 * @File: page_service_predict
 * @Date: 2021/4/7 15:08
 */

package pageservicepredict

import (
	"log"
	"net/http"
	"net/url"
	"predictProcessorServer/conf"
	"strings"
)

func Call(dagStr string) *http.Response {
	data := url.Values{
		"send1": {"test"},
		"send2": {"param"},
	}

	resp, err := http.Post(conf.PythonURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("[Call] new request failed, err:[%v]", err)
	}
	return resp
}