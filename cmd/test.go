/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/anatolyjudov/job/app/model"
	"github.com/anatolyjudov/job/app/service"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Quick test run command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")

		var service = service.SqliteStorageServiceFactory()

		//service.Init()

		testUser := model.User{}
		testUser.SetName("Toli").SetAvatarFromString("AY")

		service.SaveUser(testUser)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
