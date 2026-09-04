package cmd

import (
	"errors"
	"fmt"
	"os"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/util"
	"strings"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User management",
}

var (
	isAdmin      bool
	passwordFile string
)

var userAddCmd = &cobra.Command{
	Use:   "add username [password]",
	Short: "Add a user",
	Args: func(cmd *cobra.Command, args []string) error {
		switch {
		case passwordFile != "" && len(args) == 1:
			return nil
		case passwordFile == "" && len(args) == 2:
			return nil
		default:
			return errors.New("provide either a password argument or --password-file")
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		password := ""
		if passwordFile != "" {
			bytes, err := os.ReadFile(passwordFile)
			if err != nil {
				return fmt.Errorf("read password file: %w", err)
			}
			password = strings.TrimSuffix(strings.TrimSuffix(string(bytes), "\n"), "\r")
		} else {
			password = args[1]
		}
		if password == "" {
			return errors.New("password must not be empty")
		}

		passwordHash, err := util.Password(password)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}
		user := &model.User{
			Username:        args[0],
			Password:        passwordHash,
			Name:            args[0],
			LicensedDevices: 0,
			LoginVerify:     model.LOGIN_ACCESS_TOKEN,
			IsAdmin:         isAdmin,
			Status:          1,
		}
		cfg := config.GetServerConfig()
		engine, err := db.NewEngine(cfg.Db)
		if err != nil {
			return fmt.Errorf("create database engine: %w", err)
		}
		defer engine.Close()

		if _, err := engine.Insert(user); err != nil {
			return fmt.Errorf("add user: %w", err)
		}
		fmt.Println("Add success")
		return nil
	},
}

func init() {
	userAddCmd.Flags().BoolVarP(&isAdmin, "admin", "a", false, "Set user admin")
	userAddCmd.Flags().StringVar(&passwordFile, "password-file", "", "Read the password from a file")
	userCmd.AddCommand(userAddCmd)
	RootCmd.AddCommand(userCmd)
}
