package k8s

import (
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

func CreateService(name string, port int32) (*v1.Service, error) {
	client, err := OpenClient()
	if err != nil {
		return nil, err
	}
	client.Services(NAMESPACE).Delete(name, &v1.DeleteOptions{})
	return client.Services(NAMESPACE).Create(&v1.Service{
		ObjectMeta: v1.ObjectMeta{
			Name: name,
		},
		Spec: v1.ServiceSpec{
			Type:     v1.ServiceTypeLoadBalancer,
			Selector: map[string]string{"labels": name},
			Ports: []v1.ServicePort{
				{Port: port},
			},
		},
	})
}

func CreateDeployment(name string, image string) (*v1beta1.Deployment, error) {
	client, err := OpenClient()
	if err != nil {
		return nil, err
	}
	return client.Deployments(NAMESPACE).Create(&v1beta1.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name: name,
		},
		Spec: v1beta1.DeploymentSpec{
			Selector: &unversioned.LabelSelector{
				MatchLabels:      map[string]string{"labels": name},
				MatchExpressions: []unversioned.LabelSelectorRequirement{},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Name:   name,
					Labels: map[string]string{"labels": name},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  name,
							Image: image,
							Args:  []string{"server", "/export"},
						},
					},
				},
			},
		},
	})
}
