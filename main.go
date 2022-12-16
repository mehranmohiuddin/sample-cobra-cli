package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var cmdRun = &cobra.Command{
		Use:   "run",
		Short: "Run the app",
		Long: `Run the app on the specified port(s).

The language flag is optional. Allowed values for language are: english, spanish, french.`,
		Example: `./app run -p 8080,8081 --language english
./app run --port 8080 --language spanish
./app run -p 8080 -l french`,
		Run: func(cmd *cobra.Command, args []string) {
			// get the value of the "port" flag
			port, err := cmd.Flags().GetStringSlice("port")
			if err != nil {
				fmt.Println(err)
				return
			}

			// get the value of the "language" flag
			language, err := cmd.Flags().GetString("language")
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Running app on port(s):")
			for _, p := range port {
				fmt.Println(p)
			}

			// convert the value of the "language" flag to lowercase
			language = strings.ToLower(language)
			if language != "" {
				if language != "english" && language != "spanish" && language != "french" {
					fmt.Println("Unsupported language")
				} else {
					fmt.Printf("Language: %s\n", language)
				}
			}
		},
	}

	cmdRun.Flags().StringSliceP("port", "p", []string{}, "port to run the app on")
	cmdRun.MarkFlagRequired("port")
	cmdRun.Flags().StringP("language", "l", "", "language to use")

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdRun)
	rootCmd.Execute()
}
