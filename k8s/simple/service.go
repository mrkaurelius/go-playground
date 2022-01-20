package simple

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"

	apiv1 "k8s.io/api/core/v1"
)

func DeployService(serviceClient v1.ServiceInterface, service apiv1.Service) {
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
			Name: "simple-service-generated",
			Labels: map[string]string{
				"app": "myapp",
			},
		},
		Spec: apiv1.ServiceSpec{
			Type: apiv1.ServiceTypeNodePort,
			Selector: map[string]string{
				"app": "simple-generated",
			},
			Ports: []apiv1.ServicePort{
				apiv1.ServicePort{
					Name:       "http",
					Port:       80,
					TargetPort: intstr.FromInt(5678), // Neden direkt int degilde ayri bir struct ?
				},
			},
		}}

	return service
}
