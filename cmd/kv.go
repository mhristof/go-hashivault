package cmd

import (
	"fmt"

	"github.com/mhristof/go-vault/vault"
	"github.com/spf13/cobra"
)

var kvCmd = &cobra.Command{
	Use:   "kv",
	Short: "Interact with kv",
	Run: func(cmd *cobra.Command, args []string) {
		addr, err := cmd.Flags().GetString("addr")
		if err != nil {
			panic(err)
		}

		v := vault.New(&vault.Input{
			Addr: addr,
		})

		fmt.Println(fmt.Sprintf("v: %+v %T", v, v))
	},
}

func init() {
	rootCmd.AddCommand(kvCmd)
}
