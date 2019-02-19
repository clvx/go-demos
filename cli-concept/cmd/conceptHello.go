package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// conceptHelloCmd represents the conceptHello command
var conceptHelloCmd = &cobra.Command{
	Use:   "conceptHello",
	Short: "Prints name",
	//Adding arguments
	Args: cobra.ExactArgs(1),
	Long: `Function to print names
	cli-concept concept conceptHello [string] <-l|--last string> <-v>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("last")
		ok, _ := cmd.Flags().GetBool("verbose")
		if name == "" {
			fmt.Printf("Hello %s\n", args[0])
		}else {
			fmt.Printf("Hello %s %s\n", args[0], name)	
		}
		if ok == true { 
			fmt.Println("verbosing")	
		}
	},
}

func init() {
	conceptCmd.AddCommand(conceptHelloCmd)
	//Adding local flags(at the subcommand level)
	conceptHelloCmd.Flags().StringP("last", "l", "", "Set your last name")
}
