package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  `Trace the IP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
				showData()
			}
		} else {
			fmt.Println("Please provide an IP address")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type Ip struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
}

func showData() {
	url := "http://ipinfo.io/1.1.1.1/geo"
	responseByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Fatal("Unable to unmarshal the data")
	}
	fmt.Println("Data Found :")

	fmt.Printf("IP :%s\nCITY :%s\nREGION :%s\nCOUNTRY :%s\nLOCATION :%s\nTIMEZONE :%s\n", data.Ip, data.City, data.Region, data.Country, data.Loc, data.Timezone)
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read response")
	}

	return responseByte
}
