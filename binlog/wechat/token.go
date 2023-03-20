package wechat

import (
	"fmt"
	"github.com/qinxiaogit/k8sToGo/binlog/util"
	"log"
	"github.com/tidwall/gjson"
)

type Wechat struct {
}

const (
	APP_ID           = "wxc3b4fdc2c724a3c3"
	GRANT_TYPE       = "client_credential"
	APP_SECRET       = "14a03c6ccf007677e584212b59ef842e"
	WECHAT_GET_TOKEN = "https://api.weixin.qq.com/cgi-bin/token"

	WECHAT_DD_MESSAGE = "https://api.weixin.qq.com/cgi-bin/message/template/send"
)

func (w *Wechat) GetToke() string {
	httpClient := &util.Http{}

	//body, err := httpClient.Get(WECHAT_GET_TOKEN, map[string]string{"grant_type": GRANT_TYPE, "appid": APP_ID, "secret": APP_SECRET})
	//
	//if err != nil {
	//	log.Fatal(err)
	//	return ""
	//}
	body,err := httpClient.Get("http://www.jubaohuizhong.com/access_token.json", map[string]string{})

	if err != nil{
		log.Fatal(err)
		return ""
	}
	fmt.Println(string(body))
	return gjson.Parse(body).Get("access_token").String()
}

func (w *Wechat) SendSubscribeMsg() {
	httpClient := &util.Http{}

	body, err := httpClient.PostJson(WECHAT_DD_MESSAGE+"?access_token="+w.GetToke(), map[string]interface{}{
		"touser":            "o2D7SwSW1kuPENZBH960eCcmk8aI",
		"template_id":       "K0SH8PoI8vy9Nx__hrrH7DkhomTHc3K6HxowJARo3Jk",
	//	"page":              "index",
	//	"miniprogram_state": "developer",
		"lang":              "zh_CN",
		"data": map[string]interface{}{
			"first":         map[string]string{"value": "您有新的商城订单"},
			"keyword1":         map[string]string{"value": "1"},
			"keyword2":        map[string]string{"value": "2"},
			"keyword3": map[string]string{"value": "3"},
			"keyword4": map[string]string{"value": "4"},
			"keyword5": map[string]string{"value": "5"},
			"remark": map[string]string{"value": "remark"},
		},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(body)
}
