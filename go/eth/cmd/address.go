/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// addressCmd represents the address command
var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Run test about hex and address",
	Run: func(cmd *cobra.Command, args []string) {
		address()
	},
}

func address() {
	var addr common.Address
	addr = common.HexToAddress("0x37e9627A91DD13e453246856D58797Ad6583D762")
	fmt.Println(addr)
	fmt.Println(addr.Hex())
	fmt.Println(addr.String())
	fmt.Println(addr.Bytes())
	fmt.Println(common.BytesToAddress(addr.Bytes()))
	fmt.Println(common.Bytes2Hex(addr.Bytes()))
	fmt.Println(common.BytesToHash(addr.Bytes()))

	fmt.Println("===========")
	var hash common.Hash
	hash = common.HexToHash("0x7cbd47e6dd4e5ea2a68389f4454545a1fb60dc3dfdb30b3fecdb9c3862a63278")
	fmt.Println(hash)
	fmt.Println(hash.Hex())
	fmt.Println(hash.String())
	fmt.Println(hash.Bytes())
}

func init() {
	rootCmd.AddCommand(addressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
