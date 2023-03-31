package cloud

import (
	"context"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"os"
	"path/filepath"

	"github.com/emicklei/go-restful/v3"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *cloudRouter) createDeployment(request *restful.Request, response *restful.Response) {

}

func (s *cloudRouter) getDeployment(req *restful.Request, res *restful.Response) {
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
	deploymentname := req.PathParameter("name")
	deployment, err := clientset.AppsV1().Deployments(deploymentname).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		res.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	res.WriteEntity(deployment)
}

func (r *cloudRouter) DeleteDeployment(req *restful.Request, res *restful.Response) {
	config, err := rest.InClusterConfig()
	if err != nil {
		// if running outside cluster
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deploymentName := req.PathParameter("name")
	err = clientset.AppsV1().Deployments(v1.NamespaceDefault).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})

	if err != nil {
		res.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}
