package googleauth

type OauthUserInfo struct {
	Sub     string
	Name    string
	Email   string
	Picture string
}
type OauthTokenInfo struct {
	Token_type   string
	Access_token string
	Scope        string
	Id_token     string
}
