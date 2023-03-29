package main

import (
	"github.com/emicklei/go-restful"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"os"
	"path/filepath"
)

type DeploymentResource struct {
	clientset *kubernetes.Clientset
}

type StsResource struct {
	clientset *kubernetes.Clientset
}

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
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
	deploymentService := new(DeploymentResource)
	deploymentService.clientset = clientset

	stsService := new(StsResource)
	stsService.clientset = clientset

	deploymentWS := new(restful.WebService)
	deploymentWS.Path("/deployment").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	deploymentWS.Route(deploymentWS.GET("/{name}").To(deploymentService.getDeployment))
	deploymentWS.Route(deploymentWS.DELETE("/{name}").To(deploymentService.deleteDeployment))
}

func (r DeploymentResource) getDeployment(request *restful.Request, response *restful.Response) {
	deploymentName := request.PathParameter("name")
	deployment, err := r.clientset.AppsV1().Deployments(v1.NamespaceDefault).Get(deploymentName, metav1.)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
	} else {
		response.WriteEntity(deployment)
	}
}

func (r DeploymentResource) deleteDeployment(request *restful.Request, response *restful.Response) {
	deploymentName := request.PathParameter("name")
	err := r.clientset.AppsV1().Deployments(v1.NamespaceDefault).Delete(deploymentName, &metav1.DeleteOptions{})

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
	} else {
		response.WriteHeader(http.StatusNoContent)
	}
}
