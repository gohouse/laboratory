package driver

type IDriver interface {
	GetDsn(d string) (string, error)
}
