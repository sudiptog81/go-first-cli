/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var gopherName string

		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
		} else {
			gopherName = "gopher"
		}

		URL := "https://raw.githubusercontent.com/scraly/gophers/main/" + gopherName + ".png"

		fmt.Println("Try to get '" + gopherName + "' Gopher...")

		// Get the data
		resp, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			out, err := os.Create(gopherName + ".png")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer out.Close()

			_, err = io.Copy(out, resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Gopher '" + gopherName + "' downloaded successfully at " + out.Name() + " 🙌")
		} else {
			fmt.Println("Gopher '" + gopherName + "' not found 👀")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
