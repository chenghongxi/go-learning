package main

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sharedInformerd := informers.NewSharedInformerFactory(clientset, 0)

	gvrs := []schema.GroupVersionResource{
		{Group: "app", Version: "v1", Resource: "deployments"},
		{Group: "", Version: "v1", Resource: "pods"},
	}
	for _, gvr := range gvrs {
		if _, err = sharedInformerd.ForResource(gvr); err != nil {
			panic(err)
		}
	}

	// Start all informers
	sharedInformerd.Start(ctx.Done())
}
