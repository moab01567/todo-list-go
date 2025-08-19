package googleauth

type OauthIssuer string

const (
	GoogleAuth OauthIssuer = "Google"
)

type OauthUserInfo struct {
	Issuer  OauthIssuer
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
