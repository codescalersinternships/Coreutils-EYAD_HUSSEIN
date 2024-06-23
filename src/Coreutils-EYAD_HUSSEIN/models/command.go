package models


type Command struct {
	Name string
	Flags []string
	Description string
}



func CreateCommand(name string, flags []string, description string) *Command {
	return &Command{
		Name: name,
		Flags: flags,
		Description: description,
	}
}
