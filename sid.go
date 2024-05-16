package account

import (
	"encoding/base64"
	"errors"
	"net/url"
	"strings"
)

func GetSid(cookieStr string) (string, error) {
	if strings.TrimSpace(cookieStr) == "" {
		return "", errors.New("CookieStr不能为空")
	}

	sids, err := url.QueryUnescape(cookieStr)
	if err != nil {
		return "", errors.New("url解码失败")
	}
	// 进行 Base64 解码
	decodedBytes, err := base64.StdEncoding.DecodeString(sids)
	if err != nil {
		return "", errors.New("base64解码失败")
	}
	newSids := ""
	for k, v := range decodedBytes {
		newSids += string(rune(int(v) + k%8))
	}
	//sid := "7h16o9bfdj7kd56ed205civ4t1"
	//return sid, nil
	return newSids, nil
}

func parseCookieString(cookieStr string) map[string]string {
	// 解析 Cookie 字符串为 Cookie 的 Map
	cookies := make(map[string]string)

	// 根据分号和空格拆分字符串获取每个 Cookie
	cookieList := strings.Split(cookieStr, "; ")
	for _, cookie := range cookieList {
		// 根据等号拆分键值对
		pair := strings.Split(cookie, "=")
		if len(pair) == 2 {
			cookies[pair[0]] = pair[1]
		}
	}

	return cookies
}
