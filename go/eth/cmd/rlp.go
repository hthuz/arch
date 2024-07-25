/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/spf13/cobra"
)

type Ah struct {
	Name string
	size int
}

// rlpCmd represents the rlp command
var rlpCmd = &cobra.Command{
	Use:   "rlp",
	Short: "Test RLP encode & decode",
	Run: func(cmd *cobra.Command, args []string) {

		ah := Ah{Name: "Cool", size: 30}
		encoded, err := rlp.EncodeToBytes(ah)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+X\n", encoded)
		fmt.Println(string(encoded))

		var decoded Ah
		rlp.DecodeBytes(encoded, &decoded)
		fmt.Println(decoded)

		url := "enode://2df6082696a4da1bace518305fdae47c4ca5342f83f1a3d33dc6836b461ec60240060180dd1710e0e66cdb5055733bbdcfe6a543de16bd2362777e6d7f8d2d6e@127.0.0.1:30307"
		node := enode.MustParse(url)
		fmt.Println(node.IP().String())
		fmt.Printf("%+v\n", node)

	},
}

func init() {
	rootCmd.AddCommand(rlpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rlpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rlpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
