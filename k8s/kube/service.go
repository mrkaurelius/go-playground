package kube

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"

	apiv1 "k8s.io/api/core/v1"
)

func DeployService(service apiv1.Service, serviceClient v1.ServiceInterface) {
	result, err := serviceClient.Create(context.TODO(), &service, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deployed service  %s.\n", result.GetObjectMeta().GetName())
}

func CreateServiceClient(clientset *kubernetes.Clientset) v1.ServiceInterface {
	serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	return serviceClient
}

func GenerateService() apiv1.Service {
	service := apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "myservice",
			Labels: map[string]string{
				"app": "myapp",
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports:     nil,
			Selector:  nil,
			ClusterIP: "",
		}}

	return service
}
