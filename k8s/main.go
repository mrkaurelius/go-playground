package main

import (
	"flag"
	"path/filepath"

	"example.com/m/v2/k8s-examples/simple"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func createClientSet() *kubernetes.Clientset {
	var kubeconfig *string
	// kubeconfig = "/home/mrk/.kube/config"
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
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

	ingressClient := simple.CreateIngressClient(clientset)
	ingress := simple.BuildIngress()
	simple.DeployIngress(ingress, ingressClient)

	// TODO burada kaldim
	serviceClient := simple.CreateServiceClient(clientset)
	service := simple.GenerateService()
	simple.DeployService(serviceClient, service)

	podClient := simple.CreatePodClient(clientset)
	pod := simple.BuildPod()
	simple.DeployPod(podClient, pod)
}
