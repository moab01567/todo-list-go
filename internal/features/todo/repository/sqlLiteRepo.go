package repository

type SqlRepo struct {
	dbFile string
}

func NewSqlRepo(dbFile string) SqlRepo {
	return SqlRepo{dbFile: dbFile}

}
