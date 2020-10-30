module github.com/gohouse/demo/gin

go 1.14

require (
	github.com/gin-gonic/autotls v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/gohouse/gopass v0.0.0-20200115061301-96fd8f65e09d
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
)

replace github.com/gohouse/gopass => ./../../../gopass
