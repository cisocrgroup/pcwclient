package main // import "github.com/finkf/pcwclient"
import (
	"github.com/spf13/cobra"
)

var mainCommand = &cobra.Command{
	Use:   "pcwclient",
	Short: "Command line client for pocoweb",
	Long: `
Command line client for pocoweb. You can use it to automate or test
the pocoweb post-correction.

In order to use the command line client, you should use the
POCOWEB_URL and POCOWEB_AUTH environment varibales to set the url and
the authentification token respectively or use the appropriate --url
and --auth parameters accordingly.`,
}

func init() {
	mainCommand.AddCommand(&listCommand)
	mainCommand.AddCommand(&newCommand)
	mainCommand.AddCommand(&loginCommand)
	mainCommand.AddCommand(&logoutCommand)
	mainCommand.AddCommand(&printCommand)
	mainCommand.AddCommand(&versionCommand)
	mainCommand.AddCommand(&searchCommand)
	mainCommand.AddCommand(&correctCommand)
	mainCommand.AddCommand(&downloadCommand)
	mainCommand.AddCommand(&pkgCommand)
	downloadCommand.AddCommand(&downloadBookCommand)
	downloadCommand.AddCommand(&downloadPoolCommand)
	pkgCommand.AddCommand(&pkgAssignCommand)
	pkgCommand.AddCommand(&pkgReassignCommand)
	pkgCommand.AddCommand(&pkgSplitCommand)
	mainCommand.AddCommand(&deleteCommand)
	mainCommand.AddCommand(&startCommand)
	listCommand.AddCommand(&listBooksCommand)
	listCommand.AddCommand(&listUsersCommand)
	listCommand.AddCommand(&listPatternsCommand)
	listCommand.AddCommand(&listSuggestionsCommand)
	listCommand.AddCommand(&listSuspiciousCommand)
	listCommand.AddCommand(&listAdaptiveCommand)
	listCommand.AddCommand(&listELCommand)
	listCommand.AddCommand(&listRRDMCommand)
	listCommand.AddCommand(&listCharsCommand)
	newCommand.AddCommand(&newUserCommand)
	newCommand.AddCommand(&newBookCommand)
	startCommand.AddCommand(&startProfileCommand)
	startCommand.AddCommand(&startELCommand)
	startCommand.AddCommand(&startRRDMCommand)
	deleteCommand.AddCommand(&deleteBooksCommand)
	deleteCommand.AddCommand(&deleteUsersCommand)

	mainCommand.SilenceUsage = true
	mainCommand.SilenceErrors = true
	mainCommand.PersistentFlags().BoolVarP(&opts.format.json, "json", "J", false,
		"output raw json")
	mainCommand.PersistentFlags().BoolVarP(&opts.skipVerify,
		"skip-verify", "S", false, "ignore invalid ssl certificates")
	mainCommand.PersistentFlags().BoolVarP(&opts.debug, "debug", "D", false,
		"enable debug output")
	mainCommand.PersistentFlags().StringVarP(&opts.pocowebURL, "url", "U",
		getURL(), "set pocoweb url")
	mainCommand.PersistentFlags().StringVarP(&opts.format.template, "format", "F",
		"", "set output format")
	mainCommand.PersistentFlags().StringVarP(&opts.authToken, "auth", "A",
		getAuth(), "set auth token")
}

func main() {
	chk(mainCommand.Execute())
}
