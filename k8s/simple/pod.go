package simple

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreatePodClient(clientset *kubernetes.Clientset) v1.PodInterface {
	podClinet := clientset.CoreV1().Pods(corev1.NamespaceDefault)
	return podClinet
}

func DeployPod(podClient v1.PodInterface, pod corev1.Pod) {
	result, err := podClient.Create(context.TODO(), &pod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Deployed pod %s.\n", result.GetObjectMeta().GetName())
}

func BuildPod() corev1.Pod {
	pod := corev1.Pod{ObjectMeta: metav1.ObjectMeta{
		Name: "simple-pod-generated",
		Labels: map[string]string{
			"app": "simple-generated",
		},
	},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "echo-simple",
					Image: "hashicorp/http-echo",
					Args:  []string{"-text=simplegenerated"},
					Ports: []corev1.ContainerPort{
						corev1.ContainerPort{
							ContainerPort: 5687,
						},
					},
				},
			},
		},
	}

	// TODO pritn pod json
	// fmt.Println()
	return pod
}
