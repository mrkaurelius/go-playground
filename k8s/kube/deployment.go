package kube

import (
	"context"
	"encoding/json"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func CreateDeploymentClient(clientset *kubernetes.Clientset) v1.DeploymentInterface {
	// clientset.NetworkingV1().Ingresses() // bele
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	return deploymentsClient
}

func GenerateDeploymentObject(name string, containers []apiv1.Container) *appsv1.Deployment {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"deneme1": "deneme",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo", // project'e gore label olmali
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":    "demo",
						"deneme": "biriki", // project'e gore label olmali
					},
				},
				Spec: apiv1.PodSpec{
					Containers: containers},
			},
		},
	}

	deploymentBytes, _ := json.MarshalIndent(deployment, "", "  ")
	fmt.Println(string(deploymentBytes))

	return deployment
}

func CreateDeployment(deploymentsClient v1.DeploymentInterface, deploymentObject *appsv1.Deployment) {
	result, err := deploymentsClient.Create(context.TODO(), deploymentObject, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

func int32Ptr(i int32) *int32 { return &i }
