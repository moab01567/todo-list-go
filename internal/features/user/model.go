package user

// this user is used as table in database
type Account struct {
	AccountId string `gorm:"primaryKey"`
	Sub       string
	Issuer    string
	FirstName string
	LastName  string
	Email     string
}
