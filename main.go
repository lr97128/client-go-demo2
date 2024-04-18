package main

import (
	"context"
	"flag"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "config", "(optional) absolute path to the kubeconfig file")
	flag.Parse()
	//get config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	deploy, err := clientSet.AppsV1().Deployments("kube-system").Get(context.TODO(), "coredns", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(deploy.Name)
	deploys, err := clientSet.AppsV1().Deployments("kube-system").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, deploy1 := range deploys.Items {
		fmt.Println(deploy1.Name)
	}
	pod, err := clientSet.CoreV1().Pods("kube-system").Get(context.TODO(), "kube-apiserver-k8s-node01", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(pod.Name)
}
