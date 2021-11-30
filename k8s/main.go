package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"example.com/m/v2/k8s-examples/kube"
)

func createClientSet() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	fmt.Println(config)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func main() {
	clientset := createClientSet()
	deploymentClient := kube.CreateDeploymentClient(clientset)

	envVars := make(map[string]string)
	envVars["deneme"] = "biriki"
	envVars["asdf"] = "ucdort"

	nginxContainer := kube.GenerateNginxContainer("deneme1", envVars)
	deploymentObject := kube.GenerateDeploymentObject("ahmetmehmet12345", nginxContainer)

	fmt.Println(deploymentClient, deploymentObject)
	kube.CreateDeployment(deploymentClient, deploymentObject)
}
