package cloud

import (
	"context"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

type CloudMeta struct {
	Cloud string `uri:"cloud_name" binding:"required"`
}

func (s *cloudRouter) listNamespace(c *restful.Request, res *restful.Response) {
	config, err := rest.InClusterConfig()
	if err != nil {
		// if running outside cluster
		kubeconfig := filepath.Join(os.Getenv("HOME"), "C:\\Users\\42245\\.kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("已成功加载 clientset")

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	res.WriteEntity(namespaces)
}
