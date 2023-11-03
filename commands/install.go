package commands

import (
	"github.com/captainhook-go/captainhook/exec"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
)

func setupInstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install CaptainHook",
		Long:  "Install CaptainHook into your local .git/hooks directory",
		Run: func(cmd *cobra.Command, args []string) {
			appIO := io.NewDefaultIO(io.NORMAL, make(map[string]string))

			force, _ := cmd.Flags().GetBool("force")
			skip, _ := cmd.Flags().GetBool("skip-existing")

			conf, err := setUpConfig(cmd)
			if err != nil {
				DisplayCommandError(err)
			}

			repo, errRepo := git.NewRepository(conf.GitDirectory())
			if errRepo != nil {
				DisplayCommandError(errRepo)
			}

			installer := exec.NewInstaller(appIO, conf, repo)
			installer.SkipExisting(skip)
			installer.Force(force)
			instError := installer.Run()
			if instError != nil {
				DisplayCommandError(instError)
			}
		},
	}

	setUpFlags(cmd)
	configurationAware(cmd)
	repositoryAware(cmd)

	return cmd
}

func setUpFlags(cmd *cobra.Command) {
	var skip = false
	cmd.Flags().BoolP("skip-existing", "s", skip, "skip existing hooks")
	var force = false
	cmd.Flags().BoolP("force", "f", force, "force installation, overwrite existing hooks")
}
