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
