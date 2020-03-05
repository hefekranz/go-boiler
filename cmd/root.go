package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var (
    Version = "master"
    Date    = "undefined"
    Commit  = "undefined"
)


var RootCmd = &cobra.Command{
    Use:   "boiler",
    Short: "this is a boilerplate application",
    // Uncomment the following line if your bare application
    // has an action associated with it:
    //	Run: func(cmd *cobra.Command, args []string) {
    //	    fmt.Println("Let's get started!")
    //    },
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}