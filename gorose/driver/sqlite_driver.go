package driver

type SqliteDriver struct {
}
func (sql SqliteDriver) Drive(d string) (string, error) {
	return d+" driver",nil
}