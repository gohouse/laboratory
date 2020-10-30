package chat

type MsgType int

const (
	MTbot MsgType = iota
	MTmsg
)

type Message struct {
	Username string
	Content  string
	Type     MsgType
}
