package account

const (
	GetUserDataAPI       = "/user/cloud/userData.html"
	GetUserInfoById      = "/lan/user/info"
	GetUserAuthStatusAPI = "/sweep/auth"
)

type Account struct {
	UserId        uint
	UserName      string
	CookieStr     string
	APIUriPrefix  string
	APISignSecret string
}

// setAPIUrlPrefix 设置api schema和host，默认是官网st环境
func (a *Account) setAPIUrlPrefix() {
	if a.APIUriPrefix == "" {
		a.APIUriPrefix = "https://www.st.landui.cn"
	}
}
