package driver

type IDriver interface {
	Drive(d string) (string, error)
}
