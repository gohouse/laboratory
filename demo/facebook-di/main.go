package main

import "github.com/facebookgo/inject"
type Conf struct {

}
type DB struct {

}
type UserController struct {
	UserService *UserService `inject:""`
	Conf        *Conf        `inject:""`
}

type PostController struct {
	UserService *UserService `inject:""`
	PostService *PostService `inject:""`
	Conf        *Conf        `inject:""`
}

type UserService struct {
	Db   *DB   `inject:""`
	Conf *Conf `inject:""`
}

type PostService struct {
	Db *DB `inject:""`
}

type Server struct {
	UserApi *UserController `inject:""`
	PostApi *PostController `inject:""`
}

func (*Server) Run()  {

}

func main() {
	conf := &Conf{} // *Conf
	db := &DB{} // *DB

	server := Server{}

	graph := inject.Graph{}

	if err := graph.Provide(
		&inject.Object{
			Value: &server,
		},
		&inject.Object{
			Value: conf,
		},
		&inject.Object{
			Value: db,
		},
	); err != nil {
		panic(err)
	}

	if err := graph.Populate(); err != nil {
		panic(err)
	}

	server.Run()
}
