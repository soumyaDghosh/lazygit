package custom_commands

import (
	"runtime"

	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var BasicCmdFromConfig = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Using a custom command to create a new file",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("blah")
	},
	SetupConfig: func(cfg *config.AppConfig) {
		cmd := "touch myfile"
		if runtime.GOOS == "windows" {
			cmd = "copy NUL myfile"
		}
		cfg.UserConfig.CustomCommands = []config.CustomCommand{
			{
				Key:     "a",
				Context: "files",
				Command: cmd,
			},
		}
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Files().
			IsEmpty().
			IsFocused().
			Press("a").
			Lines(
				Contains("myfile"),
			)
	},
})
