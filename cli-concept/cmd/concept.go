package cmd

import (
//	"fmt"
	"errors"
	"github.com/spf13/cobra"
)


// conceptCmd represents the concept command
var conceptCmd = &cobra.Command{
	Use:   "concept",
	Short: "concept command",
	Long: `concept allows to print different golang concepts`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("concept called")
	//},
	//RunE forces to execute a subcommand
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Provide item to the concept command")	
	},
}


//Verbose bool
var Verbose bool

func init() {
	rootCmd.AddCommand(conceptCmd)
	//Adding a persistent flag which will run at this level and any subcommand
	//This flag has been marked as required. It will always run even when the flag 
	//is not defined
	conceptCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Printing name\n")
	//conceptCmd.MarkFlagRequired("verbose")
}
