package driver

type MysqlDriver struct {
}

func (sql MysqlDriver) GetDsn(d string) (string, error) {
	return d + " driver", nil
}
