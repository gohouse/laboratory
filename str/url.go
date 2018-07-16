package main

import (
	"net/url"
	"fmt"
)

func main() {
	urlstr := "http://abc.com/a/b?k=d&e=f&m=n"
	//urlstr := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(urlstr)
	if err != nil {
		panic(err)
	}

	//fmt.Println(u.Scheme)
	//fmt.Println(u.User)
	//fmt.Println(u.User.Username())
	//p, _ := u.User.Password()
	//fmt.Println(p)
	//
	fmt.Println(u.Host)
	//host, port, _ := net.SplitHostPort(u.Host)
	//fmt.Println(host)
	//fmt.Println(port)
	//
	fmt.Println(u.Path)
	//fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
