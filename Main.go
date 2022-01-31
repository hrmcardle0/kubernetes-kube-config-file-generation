package main

import (
	"fmt"
)

/*
* Contains information regarding the context portion of the config file
*/
type Context struct {
	cluster string
	user string
}
type Context_obj struct {
	context Context
	name string
}

/*
* Contains information regarding the exec portion of the configuration, nested from User
*/
type Env struct {
	name string
	value string
}
type Exec struct {
	apiVersion string `default:"client.authentication.k8s.io/v1alpha1"`
	args []string
	command string
	env Env
	interactiveMode string
	provideClusterInfo bool
}

/*
* Stores user information
*/
type User struct {
	name string
	exec Exec
}

/*
* The following 2 structs are necessary because of the way the config file is formatted
*/ 
type Cluster struct {
	certificate_authority_data string
	server string
}
type Cluster_obj struct {
	cluster Cluster
	name string
}

/*
* The overall kubeConfig yaml structure
*/
type kubeConfig struct {
	apiVersion string `default:"v1" json:"apiVersion"`
	clusters []Cluster_obj `json:"clusters"`
	contexts []Context_obj `json:"contexts"`
	kind string `default:"Config" json:"kind"`
	preferences string `default:"" json:"preferences"`
	users []User `json:"users"`
}

/*
* Entry point 
*/
func main() {

	// Init kubeConfig, set defaults in-case of annotation failure
	kconfig := &kubeConfig{apiVersion: "v1", kind: "Config", preferences: ""}
	fmt.Println("Formatting Kubectl Config...")

	// Get CLI arguments
	if retVal := readFromCli(kconfig); retVal != "Success" {
		panic("Something went wrong. Killing")
	}

}

