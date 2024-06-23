package utils

import (
	"Coreutils-EYAD_HUSSEIN/models"
)


func CreateCommands() []*models.Command {
	Commands := []*models.Command{
		models.CreateCommand("head", []string{"-n"}, "Description for head"),
		models.CreateCommand("tail", []string{"-n"}, "Description for tail"),
		models.CreateCommand("wc", []string{"-n", "-w", "-c"}, "Description for wc"),
		models.CreateCommand("cat", []string{"-n"}, "Description for cat"),
		models.CreateCommand("echo", []string{"-n"}, "Description for echo"),
	}


	return Commands
}


