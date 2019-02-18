package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// conceptHelloCmd represents the conceptHello command
var conceptHelloCmd = &cobra.Command{
	Use:   "conceptHello",
	Short: "Prints name",
	Long: `Function to print names
	Hello World by default.
	Hello Name if --name|-n [name] is passed
	`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		_, ok := cmd.Flags().GetString("verbose")
		if name == "" {
			fmt.Println("Hello World")
		}else {
			fmt.Printf("Hello %s\n", name)	
		}
		if ok != nil { 
			fmt.Println("verbosing")	
		}
	},
}

func init() {
	conceptCmd.AddCommand(conceptHelloCmd)
	//Adding local flags(at the subcommand level)
	conceptHelloCmd.Flags().StringP("name", "n", "", "Set your name")
}
