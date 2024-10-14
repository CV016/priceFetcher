package main

import (
	"flag"
	// "context"
	// "fmt"
	// "log"

	// "github.com/CV016/priceFetcher/client"
)

func main() {
	// client := client.New("http://localhost:3000")

	// price, err := client.FetchPrice(context.Background(), "ETH")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)


	listenAddr := flag.String("listenaddr", ":3000", "The service is running")
	flag.Parse() // Make sure to parse flags before accessing them

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
