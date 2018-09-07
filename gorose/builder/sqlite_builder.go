package builder

type SqliteBuilder struct {
}

func (sql SqliteBuilder) BuildQuery() (string, error) {
	return "SqliteBuilder BuildQuery",nil
}

func (sql SqliteBuilder) BuildExecute() (string, error) {
	return "SqliteBuilder BuildExecute",nil
}
