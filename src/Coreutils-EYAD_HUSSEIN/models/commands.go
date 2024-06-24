package models

var Commands []*Command
var CommandsMap map[string]*Command

func init() {
	Commands = createCommands()
	CommandsMap = make(map[string]*Command)
	for _, command := range Commands {
		CommandsMap[command.Name] = command
	}
}

func createCommands() []*Command {
	commands := []*Command{
		CreateCommand("head", []string{"-n"}),
		CreateCommand("tail", []string{"-n"}),
		CreateCommand("wc", []string{"-l", "-w", "-c"}),
		CreateCommand("cat", []string{"-n"}),
		CreateCommand("echo", []string{"-n"}),
		CreateCommand("true", []string{}),
		CreateCommand("false", []string{}),
		CreateCommand("env", []string{}),
		CreateCommand("yes", []string{}),
		CreateCommand("tree", []string{"-L"}),
	}

	return commands
}
