package simple

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

func DeployIngress(ingress networkingv1.Ingress, ingressClient v1.IngressInterface) {
	result, err := ingressClient.Create(context.TODO(), &ingress, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deployed ingress  %s.\n", result.GetObjectMeta().GetName())

}

func CreateIngressClient(clientset *kubernetes.Clientset) v1.IngressInterface {
	ingressClient := clientset.NetworkingV1().Ingresses(corev1.NamespaceDefault)
	return ingressClient
}

func BuildIngress() networkingv1.Ingress {
	fmt.Println("Generating ingress")

	// Sadece
	// apiVerison vb gibi alanlar cluter'a ekledikten sonra olusur herhalde ?
	// Benim olusturdugum object ile default object arasinda farklar var mi acaba ?

	ingress := networkingv1.Ingress{}
	ingress.Name = "simple-ingress-generated"
	annotations := make(map[string]string)
	annotations["nginx.ingress.kubernetes.io/rewrite-target"] = "/$1"

	ingress.Annotations = annotations
	var pathType networkingv1.PathType = "Prefix"
	rule := networkingv1.IngressRule{Host: "local.mrk1debian",
		IngressRuleValue: networkingv1.IngressRuleValue{
			HTTP: &networkingv1.HTTPIngressRuleValue{
				Paths: []networkingv1.HTTPIngressPath{
					networkingv1.HTTPIngressPath{
						Path:     "/simplegenerated",
						PathType: &pathType,
						Backend: networkingv1.IngressBackend{
							Service: &networkingv1.IngressServiceBackend{
								Name: "simple-service-generated",
								Port: networkingv1.ServiceBackendPort{
									Number: 80,
								},
							},
						},
					},
				},
			},
		}}

	ingress.Spec.Rules = make([]networkingv1.IngressRule, 1)
	ingress.Spec.Rules = append(ingress.Spec.Rules, rule)

	return ingress
}
