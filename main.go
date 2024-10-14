package main

import (
	"flag"
)

func main() {

	// var svc PriceFetcher

	listenAddr := flag.String("listenaddr",":3000","The service is running")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr,svc)
	server.Run()

	
}