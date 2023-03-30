package main

import (
	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/pkg/apis/apps"
	"net/http"
	"os"
	"path/filepath"
)

type Resource struct {
	clientset *kubernetes.Clientset
}

func main() {
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

	wsDeployment := new(restful.WebService)
	wsDeployment.
		Path("/deployment").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	wsStsfulset := new(restful.WebService)
	wsStsfulset.
		Path("/stsfulset").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	resourceDeployment := Resource{clientset: clientset}
	resourceStsfulset := Resource{clientset: clientset}

	wsDeployment.Route(wsDeployment.GET("/{name}").To(resourceDeployment.GetDeployment))
	wsDeployment.Route(wsDeployment.DELETE("/{name}").To(resourceDeployment.DeleteDeployment))
	wsDeployment.Route(wsDeployment.PUT("").To(resourceDeployment.CreateDeployment))
	wsDeployment.Route(wsDeployment.POST("").To(resourceDeployment.UpdateDeployment))
	wsDeployment.Route(wsDeployment.GET("").To(resourceDeployment.ListDeployments))

	wsStsfulset.Route(wsStsfulset.GET("/{name}").To(resourceStsfulset.GetStsfulset))
	wsStsfulset.Route(wsStsfulset.DELETE("/{name}").To(resourceStsfulset.DeleteStsfulset))
	wsStsfulset.Route(wsStsfulset.PUT("").To(resourceStsfulset.CreateStsfulset))
	wsStsfulset.Route(wsStsfulset.POST("").To(resourceStsfulset.UpdateStsfulset))
	wsStsfulset.Route(wsStsfulset.GET("").To(resourceStsfulset.ListStsfulsets))

	restful.DefaultContainer.Add(wsDeployment)
	restful.DefaultContainer.Add(wsStsfulset)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (r *Resource) GetDeployment(request *restful.Request, response *restful.Response) {
	deploymentName := request.PathParameter("name")
	deployment, err := r.clientset.AppsV1().Deployments(v1.NamespaceDefault).Get(deploymentName, metav1.GetOptions{})

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(deployment)
}

func (r *Resource) ListDeployments(request *restful.Request, response *restful.Response) {
	l, err := r.clientset.AppsV1().Deployments(v1.NamespaceDefault).List(metav1.ListOptions{})

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(l)
}

func (r *Resource) CreateDeployment(request *restful.Request, response *restful.Response) {
	deployment := new(apps.Deployment)
	err := request.ReadEntity(deployment)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	result, err := r.clientset.AppsV1().Deployments(v1.NamespaceDefault).Create(deployment)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteHeaderAndEntity(http.StatusCreated, result)
}

func (r *Resource) UpdateDeployment(request *restful.Request, response *restful.Response) {
	deployment := new(apps.Deployment)
	err := request.ReadEntity(deployment)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	if deployment.ObjectMeta.Name == "" {
		response.WriteHeaderAndEntity(http.StatusBadRequest, "Name parameter missing from request.")
		return
	}

	_, err = r.clientset.AppsV1().Deployments(deployment.ObjectMeta.Namespace).Get(deployment.ObjectMeta.Name, metav1.GetOptions{})

	if err != nil {
		statusErr, ok := err.(*errors.StatusError)
		if ok && statusErr.Status().Reason == metav1.StatusReasonNotFound {
			response.WriteHeaderAndEntity(http.StatusNotFound, "Deployment not found.")
			return
		}

		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	result, err := r.clientset.AppsV1().Deployments(deployment.ObjectMeta.Namespace).Update(deployment)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(result)
}

func (r *Resource) DeleteDeployment(request *restful.Request, response *restful.Response) {
	deploymentName := request.PathParameter("name")
	err := r.clientset.AppsV1().Deployments(v1.NamespaceDefault).Delete(deploymentName, &metav1.DeleteOptions{})

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}

func (r *Resource) GetStsfulset(request *restful.Request, response *restful.Response) {
	stsfulsetName := request.PathParameter("name")
	stsfulset, err := r.clientset.AppsV1().StatefulSets(v1.NamespaceDefault).Get(stsfulsetName, metav1.GetOptions{})

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(stsfulset)
}

func (r *Resource) ListStsfulsets(request *restful.Request, response *restful.Response) {
	l, err := r.clientset.AppsV1().StatefulSets(v1.NamespaceDefault).List(metav1.ListOptions{})

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(l)
}

func (r *Resource) CreateStsfulset(request *restful.Request, response *restful.Response) {
	stsfulset := new(apps.StatefulSet)
	err := request.ReadEntity(stsfulset)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	result, err := r.clientset.AppsV1().StatefulSets(v1.NamespaceDefault).Create(stsfulset)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteHeaderAndEntity(http.StatusCreated, result)
}

func (r *Resource) UpdateStsfulset(request *restful.Request, response *restful.Response) {
	stsfulset := new(apps.StatefulSet)
	err := request.ReadEntity(stsfulset)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	if stsfulset.ObjectMeta.Name == "" {
		response.WriteHeaderAndEntity(http.StatusBadRequest, "Name parameter missing from request.")
		return
	}

	_, err = r.clientset.AppsV1().StatefulSets(stsfulset.ObjectMeta.Namespace).Get(stsfulset.ObjectMeta.Name, metav1.GetOptions{})

	if err != nil {
		statusErr, ok := err.(*errors.StatusError)
		if ok && statusErr.Status().Reason == metav1.StatusReasonNotFound {
			response.WriteHeaderAndEntity(http.StatusNotFound, "StatefulSet not found.")
			return
		}

		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	result, err := r.clientset.AppsV1().StatefulSets(stsfulset.ObjectMeta.Namespace).Update(stsfulset)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(result)
}

func (r *Resource) DeleteStsfulset(request *restful.Request, response *restful.Response) {
	stsfulsetName := request.PathParameter("name")
	err := r.clientset.AppsV1().StatefulSets(v1.NamespaceDefault).Delete(stsfulsetName, &metav1.DeleteOptions{})

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}
