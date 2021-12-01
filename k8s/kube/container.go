package kube

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
)

func GenerateNginxContainer(name string, envMap map[string]string) []apiv1.Container {
	env := make([]apiv1.EnvVar, 0)

	for k, v := range envMap {
		fmt.Println(k, v)
		env = append(env, apiv1.EnvVar{Name: k, Value: v})
	}

	containers := []apiv1.Container{
		{
			Name:  name,
			Image: "nginx:1.12",
			Ports: []apiv1.ContainerPort{
				{
					Name:          "http",
					Protocol:      apiv1.ProtocolTCP,
					ContainerPort: 80,
				},
			},
			Env: env,
		},
	}

	return containers
}
