package main

import (
	"fmt"

	"github.com/finkf/pcwgo/api"
	"github.com/spf13/cobra"
)

var loginCommand = cobra.Command{
	Use:   "login [EMAIL PASSWORD]",
	Short: "login to pocoweb or get logged in user",
	RunE:  runLogin,
	Args:  exactArgs(0, 2),
}

func runLogin(cmd *cobra.Command, args []string) error {
	if len(args) == 2 {
		user, password := args[0], args[1]
		return login(user, password)
	}
	return getLogin()
}

func login(user, password string) error {
	// if args.debug {
	// 	log.SetLevel(log.DebugLevel)
	// }
	url := getURL()
	if url == "" {
		return fmt.Errorf("login: missing url: use --url or POCOWEB_URL")
	}
	c, err := api.Login(url, user, password, opts.skipVerify)
	if err != nil {
		return fmt.Errorf("login: %v", err)
	}
	format(c.Session)
	return nil
}

func getLogin() error {
	c := api.Authenticate(getURL(), getAuth(), opts.skipVerify)
	var session api.Session
	if err := get(c, c.URL("login"), &session); err != nil {
		return fmt.Errorf("get login: %v", err)
	}
	format(session)
	return nil
}

var logoutCommand = cobra.Command{
	Use:   "logout",
	Short: "logout from pocoweb",
	RunE:  runLogout,
	Args:  cobra.NoArgs,
}

func runLogout(_ *cobra.Command, args []string) error {
	c := api.Authenticate(getURL(), getAuth(), opts.skipVerify)
	if err := get(c, c.URL("login"), nil); err != nil {
		return fmt.Errorf("logout: %v", err)
	}
	return nil
}
