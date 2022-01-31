package main

import (
	"fmt"
)

func createKubeConfig(kc* kubeConfig, cluster_name string, cad string, server_name string) (retValue string) {

	fmt.Println("Create KubeConfig")

	// Format cluster data
	cluster_ := Cluster{certificate_authority_data: cad, server: server_name}
	cluster_object := Cluster_obj{cluster: cluster_, name: cluster_name}
	cluster_list := []Cluster_obj{cluster_object}

	// Format user data
	arg_list := []string{"--region", "us-east-1", "eks", "get-token", "--cluster-name", cluster_name}
	env_ := Env{name: "AWS_PROFILE", value: "nonprod"}
	exec_ := Exec{args: arg_list, command: "aws", env: env_, interactiveMode: "ifAvailable", provideClusterInfo: false}

	// Format context info
	context_ := Context{cluster: cluster_name, user: cluster_name}
	context_obj := Context_obj{context: context_, name: cluster_name}
	context_list := []Context_obj{context_obj}

	//Format user info
	user := User{name: cluster_name, exec: exec_} 
	user_list := []User{user}

	kc.clusters = cluster_list
	kc.contexts = context_list
	kc.users = user_list

	if retVal := generateJson(kc); retVal != "Success" {
		fmt.Println(retVal)
		return retVal
	}

	return "Success"

}