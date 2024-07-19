/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// callCmd represents the call command
var callCmd = &cobra.Command{
	Use:   "call",
	Short: "Call Contract ",
	Run: func(cmd *cobra.Command, args []string) {
		call()
	},
}

var tokenContracts []string = []string{
	"https://bscscan.com/token/0x2170ed0880ac9a755fd29b2688956bd959f933f8", //1
	"https://bscscan.com/token/0x55d398326f99059ff775485246999027b3197955", //2
	"https://bscscan.com/token/0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", //3
	"https://bscscan.com/token/0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d", //4
	"https://bscscan.com/token/0x8965349fb649a33a30cbfda057d8ec2c48abe2a2", //5
	"https://bscscan.com/token/0x1d2f0da169ceb9fc7b3144628db156f3f6c60dbe", //6
	"https://bscscan.com/token/0x76A797A59Ba2C17726896976B7B3747BfD1d220f", //7
	"https://bscscan.com/token/0xba2ae424d960c26247dd6c32edc70b295c744c43", //8
	"https://bscscan.com/token/0x3ee2200efb3400fabb9aacf31297cbdd1d435d47", //9
	"https://bscscan.com/token/0xce7de646e7208a4ef112cb6ed5038fa6cc6b12e3", //10
	"https://bscscan.com/token/0x2859e4544c4bb03966803b044a93563bd2d0dd4d", //11
	"https://bscscan.com/token/0x1ce0c2827e2ef14d5c4f29a091d735a204794041", //12
	"https://bscscan.com/token/0x7083609fce4d1d8dc0c979aab8c869ea2c873402", //13
	"https://bscscan.com/token/0xf8a0bf9cf54bb92f17374d9e9a321e6a111a51bd", //14
	"https://bscscan.com/token/0x8ff795a6f4d97e7887c79bea79aba5cc76444adf", //15
	"https://bscscan.com/token/0x1fa4a73a3f0133f0025378af00236f3abdee5d63", //16
	"https://bscscan.com/token/0x4338665cbb7b2485a8855a139b75d5e34ab0db94", //17
	"https://bscscan.com/token/0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3", //18
	"https://bscscan.com/token/0xcc42724c6683b7e57334c4e856f4c9965ed682bd", //19
	"https://bscscan.com/token/0x25d887ce7a35172c62febfd67a1856f20faebb00", //20
	"https://bscscan.com/token/0xbf5140a22578168fd562dccf235e5d43a02ce9b1", //21
	"https://bscscan.com/token/0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c", //22
	"https://bscscan.com/token/0x031b41e504677879370e9dbcf937283a8691fa7f", //23
	"https://bscscan.com/token/0x3d6545b08693dae087e957cb1180ee38b9e3c25e", //24
	"https://bscscan.com/token/0x43c934a845205f0b514417d757d7235b8f53f1b9", //25
	"https://bscscan.com/token/0x87d23E1bcfB01a1d721627D2d9A8e5fD0C639999",
}

func call() {

	for i, url := range tokenContracts {
		contract := url[len(url)-42:]
		bytes := callContract(contract, "decimals()")
		fmt.Println(i, contract, int(bytes[31]))
	}

	// bytes := callContract("0x2170ed0880ac9a755fd29b2688956bd959f933f8", "decimals()")
	// fmt.Println(len(bytes))
	// fmt.Println(bytes)
	// decimals := int(bytes[31])
	// fmt.Println(decimals)
	// fmt.Println(common.BytesToHash(bytes))
	// fmt.Println(common.Bytes2Hex(bytes))

}

func callContract(to string, method string) []byte {
	client, err := ethclient.Dial("https://bsc-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	toAddr := common.HexToAddress(to)
	bytes, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &toAddr,
		Data: crypto.Keccak256([]byte(method)),
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func init() {
	rootCmd.AddCommand(callCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// callCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
