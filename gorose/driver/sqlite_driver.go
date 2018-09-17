package driver

type SqliteDriver struct {
}

func (sql SqliteDriver) GetDsn(d string) (string, error) {
	return d + " driver", nil
}
