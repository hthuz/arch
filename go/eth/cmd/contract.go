/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"eth/config"
	"eth/contract"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// contractCmd represents the contract command
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "Deploying and Interact with contract",
	Run: func(cmd *cobra.Command, args []string) {
		deploy()
	},
}

func deploy() {
	file, err := os.Open(config.Account1KeyStorePath + "/UTC--2024-07-02T07-08-41.129070472Z--6a62503be31643a98c5e5c7a584238e6fc815793")
	if err != nil {
		log.Fatal(err)
	}
	client, _ := ethclient.Dial(config.TestNode)
	chainid, _ := client.ChainID(context.Background())

	auth, err := bind.NewTransactorWithChainID(file, "123456", chainid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(auth)
	address, tx, instance, err := contract.DeploySimpleERC20(auth, client, big.NewInt(100000))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deployed addr ", address)
	fmt.Println("deploying tx ", tx)

	balance, err := instance.BalanceOf(&bind.CallOpts{Pending: true}, common.HexToAddress(config.TestAccount1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balanced ", balance)
}

func init() {
	rootCmd.AddCommand(contractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// contractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// contractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
