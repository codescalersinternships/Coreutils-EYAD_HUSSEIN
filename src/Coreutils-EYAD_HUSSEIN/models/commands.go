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
		CreateCommand("head", []string{"-n"}, "Description for head"),
		CreateCommand("tail", []string{"-n"}, "Description for tail"),
		CreateCommand("wc", []string{"-n", "-w", "-c"}, "Description for wc"),
		CreateCommand("cat", []string{"-n"}, "Description for cat"),
		CreateCommand("echo", []string{"-n"}, "Description for echo"),
		CreateCommand("true", []string{}, "Description for true"),
		CreateCommand("false", []string{}, "Description for false"),
		CreateCommand("env", []string{}, "Description for env"),
	}


	return commands
}