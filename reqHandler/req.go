package reqHandler

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
)

func WxReq() {
	all := GetAll()
	for _, model := range all {
		fmt.Println(model)
	}

	req := requests.Requests()

	targetUrl := "https://xxcapp.xidian.edu.cn/uc/wap/login/check"

	for _, v := range all {
		payload := requests.Datas{
			"username": v.UserName,
			"password": v.Password,
		}

		var msg Check

		rsp, err := req.Post(targetUrl, payload)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(rsp.Content(), &msg)
		fmt.Println(msg)

		if msg.E != 0 {
			panic("登录失败")
		}

		form := requests.Datas{
			"sfzx":         "1",
			"tw":           "1",
			"area":         "陕西省 西安市 长安区",
			"city":         "西安市",
			"province":     "陕西省",
			"address":      "陕西省西安市长安区兴隆街道210国道",
			"geo_api_info": `{"type":"complete","position":{"Q":34.126833224827,"R":108.84419433593803,"lng":108.844194,"lat":34.126833},"location_type":"html5","message":"Get ipLocation failed.Get geolocation success.Convert Success.Get address success.","accuracy":74,"isConverted":true,"status":1,"addressComponent":{"citycode":"029","adcode":"610116","businessAreas":[],"neighborhoodType":"","neighborhood":"","building":"","buildingType":"","street":"雷甘路","streetNumber":"230号","country":"中国","province":"陕西省","city":"西安市","district":"长安区","township":"兴隆街道"},"formattedAddress":"陕西省西安市长安区兴隆街道210国道","roads":[],"crosses":[],"pois":[],"info":"SUCCESS"}`,
			"sfcyglq":      "0",
			"sfyzz":        "0",
			"qtqk":         "",
			"ymtys":        "0",
		}
		rsp2, err := req.Post("https://xxcapp.xidian.edu.cn/xisuncov/wap/open-report/save", form)
		fmt.Println(string(rsp2.Content()))
	}
}

type Check struct {
	E int    `json:"e"`
	M string `json:"m"`
}
