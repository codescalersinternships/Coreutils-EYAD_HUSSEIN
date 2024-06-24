package models


type Command struct {
	Name string
	Flags []string
}



func CreateCommand(name string, flags []string) *Command {
	return &Command{
		Name: name,
		Flags: flags,
	}
}
