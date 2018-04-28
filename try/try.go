package try

type exception struct {
	error interface{}
	handle func()
}

func Try(actor func()) *exception {
	return &exception{handle:actor}
}

func (ex *exception) Catch(actor func(e interface{})) *exception {

	defer func() {
		if err := recover(); err!=nil {
			actor(err)
		}
	}()

	ex.handle()

	return ex
}
