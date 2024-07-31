/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"eth/config"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Get Latest Block",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial(config.PublicNode)
		if err != nil {
			log.Fatal("dail error", err)
		}

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			var s string
			for s != "exit" {
				fmt.Print(">")
				fmt.Scanln(&s)
				if s == "s" {
					ctx_timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)
					no, err := client.BlockNumber(ctx_timeout)
					if err != nil {
						log.Println("get block number error", err)
					}
					fmt.Println("no:", no)
				}
			}
			cancel()
		}()
		<-ctx.Done()
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
