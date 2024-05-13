package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/guid"
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
func (a *Account) GetUserInfo() (ReqUser, int, error) {
	sid, err := getSid(a.CookieStr)
	if err != nil {
		return ReqUser{}, 0, err
	}
	header := map[string]string{"Cookie": fmt.Sprintf("PHPSESSID=%s;", sid)}
	var req ReqUser
	httpClient := resty.New()
	resp, err := httpClient.R().SetHeaders(header).Get(a.APIUriPrefix + GetUserDataAPI)
	if err != nil {
		return req, 0, fmt.Errorf("请求失败: %s", err)
	}
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

// RealNameAuthentication 判断用户是否实名
func (a *Account) RealNameAuthentication() error {
	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, a.APISignSecret)
	newText := sorts(text)
	ciphertext := gmd5.MustEncryptString(newText)

	body := map[string]interface{}{
		"time_stamp": fmt.Sprintf("%d", times),
		"nonce_str":  randStr,
		"sign":       ciphertext,
		"userid":     a.UserId,
	}
	httpClient := resty.New()
	var res map[string]interface{}
	resp, err := httpClient.R().SetBody(body).SetResult(&res).Post(a.APIUriPrefix + GetUserAuthStatusAPI)
	if err != nil {
		return fmt.Errorf("请求失败: %s", err)
	}
	if fmt.Sprint(res["status"]) != "y" {
		return fmt.Errorf("用户未实名认证，请先实名认证 %s", resp.String())
	}
	return nil
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
