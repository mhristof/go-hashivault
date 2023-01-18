package cmd

import (
	"fmt"

	"github.com/mhristof/go-vault/vault"
	"github.com/spf13/cobra"
)

var kvDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump KV storage",
	Run: func(cmd *cobra.Command, args []string) {
		addr, err := cmd.Flags().GetString("addr")
		if err != nil {
			panic(err)
		}
		v := vault.New(&vault.Input{
			Addr: addr,
		})

		fmt.Println(fmt.Sprintf("v: %+v %T", v, v))
		// do something
	},
}

func init() {
	kvCmd.AddCommand(kvDumpCmd)
}
