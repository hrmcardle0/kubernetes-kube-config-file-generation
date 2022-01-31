package main

import (
	"fmt"
	"bufio"
	"os"
)

func readFromCli(kc *kubeConfig) (retVal string) {

	// Read details from user
	reader := bufio.NewReader(os.Stdin)

	// cluster name
	fmt.Print("Cluster Name:")
	cluster_name, _ := reader.ReadString('\n')

	// certificate authority data
	fmt.Print("Certificate Authority Data:")
	cad, _ := reader.ReadString('\n')
	
	// server
	fmt.Print("Server Endpoint API Url:")
	server, _ := reader.ReadString('\n')


	return createKubeConfig(kc, cluster_name, cad, server)

}