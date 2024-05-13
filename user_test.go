package account

import (
	"fmt"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	account := Account{
		CookieStr:    "ZXFkcTRxaDBwa2hpX2ZpZ3AvYzBjXi5bNDY%3D",
		APIUriPrefix: "https://www.st.landui.cn",
	}
	res, code, err := account.GetUserInfo()

	fmt.Println(res)
	fmt.Println(code)
	fmt.Println(err)
}

func TestGetUserInfoById(t *testing.T) {
	account := Account{
		UserId:       14610,
		UserName:     "86326328",
		APIUriPrefix: "https://www.st.landui.cn",
	}
	res, err := account.GetUserInfoById()

	fmt.Println(res)
	fmt.Println(err)
}
