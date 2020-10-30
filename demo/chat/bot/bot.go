package bot

type Command struct {
	Command  string
	Argument string
}

func IsBot(text string) bool {
	return true
}

func Handle(text string) {

}
