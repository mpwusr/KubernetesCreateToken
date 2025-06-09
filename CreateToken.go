package main

import (
	"context"
	"fmt"
	authv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	config, _ := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	clientset, _ := kubernetes.NewForConfig(config)
	saClient := clientset.CoreV1().ServiceAccounts("default")

	req := &authv1.TokenRequest{
		Spec: authv1.TokenRequestSpec{
			Audiences:         []string{"https://kubernetes.default.svc"},
			ExpirationSeconds: func(i int64) *int64 { return &i }(3600),
		},
	}

	token, err := saClient.CreateToken(context.TODO(), "podcast", req, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Token:", token.Status.Token)
}
