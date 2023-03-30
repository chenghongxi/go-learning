package main

import (
	"go-learning/go-restful/options"
	"go-learning/go-restful/router"
)

func main() {
	//config, err := rest.InClusterConfig()
	//if err != nil {
	//	// if running outside cluster
	//	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	//	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//}

	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err.Error())
	//}
	opts := &options.Options{}

	router.InstallRoutes(opts)

}
