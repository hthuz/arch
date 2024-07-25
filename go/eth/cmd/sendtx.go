/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"eth/config"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/cobra"
)

// sendtxCmd represents the sendtx command
var sendtxCmd = &cobra.Command{
	Use:   "sendtx",
	Short: "eth send tx",
	Run: func(cmd *cobra.Command, args []string) {

		// Get keystore
		ks := keystore.NewKeyStore(config.Account1KeyStorePath, keystore.StandardScryptN, keystore.StandardScryptP)
		// manager := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
		account1 := ks.Accounts()[0]

		client, err := ethclient.Dial(config.TestNode)
		if err != nil {
			log.Fatal(err)
		}

		chainid, err := client.ChainID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(config.TestAccount1))
		if err != nil {
			log.Fatal(err)
		}

		tipCap, _ := client.SuggestGasTipCap(context.Background())
		feeCap, _ := client.SuggestGasPrice(context.Background())

		var (
			to       = common.HexToAddress(config.TestAccount2)
			value    = new(big.Int).Mul(big.NewInt(2), big.NewInt(params.Wei))
			gasLimit = uint64(21111) // Minimum gasUsed for a ETH transfer is 21000, tx will fail if gasLimit < 21000
		)
		// Create a new transaction
		tx := types.NewTx(
			&types.DynamicFeeTx{
				ChainID:   chainid,
				Nonce:     nonce,
				GasTipCap: tipCap,
				GasFeeCap: feeCap,
				Gas:       gasLimit,
				To:        &to,
				Value:     value,
				Data:      nil,
			})

		// signedTX, err := types.SignTx(tx, types.NewLondonSigner(chainid), sk)
		signedTx, err := ks.SignTxWithPassphrase(account1, "123456", tx, chainid)
		if err != nil {
			log.Fatal(err)
		}

		// send transaction
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Waiting")
		time.Sleep(6 * time.Second)
		// get transaction receipt
		receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
		if err != nil {
			log.Println(signedTx.Hash(), " pending")
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", receipt)

	},
}

func init() {
	rootCmd.AddCommand(sendtxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendtxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendtxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
