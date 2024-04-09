package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/guid"
	"gitlab.landui.cn/gomod/logs"
	"sort"
	"strings"
	"time"
)

type ReqUser struct {
	Status string      `json:"status"`
	Data   ReqUserData `json:"data"`
}
type ReqUserData struct {
	UserId   uint   `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// GetUserInfo 调用用户中心接口获取用户信息
func GetUserInfo(cookieStr, apiUrl string) (ReqUser, int, error) {
	sid, err := getSid(cookieStr)
	if err != nil {
		return ReqUser{}, 0, err
	}
	header := map[string]string{"Cookie": fmt.Sprintf("PHPSESSID=%s;", sid)}
	var req ReqUser
	httpClient := resty.New()
	resp, err := httpClient.R().SetHeaders(header).Get(apiUrl)
	if resp.StatusCode() != 200 {
		return req, resp.StatusCode(), errors.New("响应状态码错误")
	}
	if strings.Contains(resp.String(), "系统发生错误") {
		return req, resp.StatusCode(), errors.New("系统发生错误")
	}
	err = json.Unmarshal(resp.Body(), &req)
	if err != nil {
		return req, resp.StatusCode(), errors.New("参数解析失败")
	}

	if req.Status == "n" {
		return req, resp.StatusCode(), errors.New("查询失败")
	}
	return req, resp.StatusCode(), err
}

func sorts(text string) string {
	var array []string
	for _, v := range text {
		array = append(array, string(v))
	}
	sort.Strings(array)
	newText := ""
	for _, v := range array {
		newText += v
	}
	return newText
}

// RealNameAuthentication 判断用户是否实名
func RealNameAuthentication(userId uint, apiUrl, signSecret string) error {
	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, signSecret)
	newText := sorts(text)
	ciphertext := gmd5.MustEncryptString(newText)

	body := map[string]interface{}{
		"time_stamp": fmt.Sprintf("%d", times),
		"nonce_str":  randStr,
		"sign":       ciphertext,
		"userid":     userId,
	}
	httpClient := resty.New()
	var res map[string]interface{}
	resp, err := httpClient.R().SetBody(body).SetResult(&res).Post(apiUrl)
	if err != nil {
		log := logs.New()
		log.Error("请求失败", err)
		return errors.New("请求失败")
	}
	if fmt.Sprint(res["status"]) != "y" {
		fmt.Println("times=>", times)
		fmt.Println("------------")
		fmt.Println(resp.String())
		fmt.Println(times)
		log := logs.New()
		log.Error("用户未实名错误"+resp.String(), errors.New(""))
		return errors.New("用户未实名认证，请先实名认证")
	}
	return nil
}
