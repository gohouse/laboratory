package driver

type MysqlDriver struct {
}
func (sql MysqlDriver) Drive(d string) (string, error) {
	return d+" driver",nil
}