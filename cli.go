package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "inncfg",
	Short: "Configuration management tool",
	Long:  "A tool for managing configuration files.",
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		exist := configExists()

		if !exist {
			err := viper.SafeWriteConfig()
			if err != nil {
				fmt.Printf("Error creating config file: %s\n", err)
				return
			}
		}

		fmt.Printf("Config found! Creating .env files...\n")

		var secrets map[string][]string

		if err := viper.UnmarshalKey("secrets", &secrets); err != nil {
			fmt.Printf("Error unmarshalling secrets: %s\n", err)
		}

		for key, value := range secrets {
			newFile := fmt.Sprintf("%s.env", key)

			fmt.Printf("Creating %s...\n", newFile)

			file, err := os.Create(newFile)

			if err != nil {
				fmt.Printf("Error creating file: %s\n", err)
				return
			}

			for _, v := range value {
				fmt.Printf("Enter value for %s: ", v)
				text, _ := reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)

				if text == "" {
					// assign text a secure string
					text, err = generateSecureString(64)

					if err != nil {
						fmt.Printf("Error generating secure string: %s\n", err)
						return
					}
				}

				file.WriteString(fmt.Sprintf("%s=%s\n", v, text))
			}

			file.Close()
		}
	},
}

var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Manage secrets",
}

var secretsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new secret",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		exist := configExists()

		if !exist {
			fmt.Printf("No config found! Run 'inncfg init' to create one.\n")
			return
		}

		var secrets = make(map[string][]string)

		if err := viper.UnmarshalKey("secrets", &secrets); err != nil {
			fmt.Printf("Error unmarshalling secrets: %s\n", err)
		}

		if len(args) < 2 {
			fmt.Printf("Usage: inncfg secrets add <group> <key>\n")
			return
		}

		key := args[0]
		value := args[1]

		secrets[key] = append(secrets[key], value)

		viper.Set("secrets", secrets)

		err := viper.WriteConfig()

		if err != nil {
			fmt.Printf("Error writing config file: %s\n", err)
		}

		fmt.Printf("Enter value for %s: ", value)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "" {
			// assign text a secure string
			text, err = generateSecureString(64)

			if err != nil {
				fmt.Printf("Error generating secure string: %s\n", err)
				return
			}
		}

		newFile := fmt.Sprintf("%s.env", key)

		//check if file exists
		if _, err := os.Stat(newFile); os.IsNotExist(err) {
			file, err := os.Create(newFile)

			if err != nil {
				fmt.Printf("Error creating file: %s\n", err)
			}

			file.WriteString(fmt.Sprintf("%s=%s\n", value, text))

			fmt.Printf("Successfully added secret %s to group %s\n", value, key)

			return
		}

		file, err := os.OpenFile(newFile, os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
		}

		file.WriteString(fmt.Sprintf("%s=%s\n", value, text))

		fmt.Printf("Successfully added secret %s to group %s\n", value, key)
	},
}

var secretsRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a secret",
}

var secretsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all secrets",
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(secretsCmd)
	secretsCmd.AddCommand(secretsAddCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
