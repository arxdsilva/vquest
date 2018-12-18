package actions

// Action is some kind of user interaction with the client
// Can be a movement, spell or anything else
type Action struct {
	Type     int
	Target   string
	Accepted bool
}

const (
	Connection = iota
	Movement
)
