package chatp2p

type Group struct {
	Name      string
	Owner     int64
}

func NewGroup(name string) *Group {
	return &Group{Name: name}
}

func (g *Group) Create(user *User) (bool, error) {
	res := Serialize(*user)
	return rds.HSetNX(keyGroup, g.Name, res).Result()
}

type GroupUser struct {
	GroupName string
	Users     []string // username
}
