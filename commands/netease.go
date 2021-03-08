package commands

import (
	tea "github.com/anhoder/bubbletea"
	"github.com/anhoder/go-musicfox/constants"
	"github.com/anhoder/go-musicfox/ui"
	"github.com/gookit/gcli/v2"
)

func NewPlayerCommand() *gcli.Command {
	return &gcli.Command{
		Name: "netease",
		// allow color tag and {$cmd} will be replace to 'demo'
		UseFor: "Command line player for Netease Cloud Music",
		Func: runPlayer,
	}
}

func runPlayer(cmd *gcli.Command, args []string) error{
	neteaseModel := ui.NeteaseModel{TotalDuration: constants.StartupLoadingDuration}
	program := tea.NewProgram(neteaseModel)
	program.EnterAltScreen()
	defer program.ExitAltScreen()
	if err := program.Start(); err != nil {
		return err
	}

	return nil
}