package builder

type MysqlBuilder struct {
}

func (sql MysqlBuilder) BuildQuery() (string, error) {
	return "MysqlBuilder BuildQuery", nil
}

func (sql MysqlBuilder) BuildExecute() (string, error) {
	return "MysqlBuilder BuildExecute", nil
}
