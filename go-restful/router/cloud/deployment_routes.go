package cloud

import (
	"context"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *cloudRouter) createDeployment(request *restful.Request, response *restful.Response) {

}

func (s *cloudRouter) getDeployment(request *restful.Request, response *restful.Response) {
	deploymentname := request.PathParameter("name")
	deployment, err := s.clientset.AppsV1().Deployments("default").Get(context.TODO(), deploymentname, metav1.GetOptions{})
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(deployment)
}

func (r *cloudRouter) DeleteDeployment(request *restful.Request, response *restful.Response) {
	deploymentName := request.PathParameter("name")
	err := r.clientset.AppsV1().Deployments(v1.NamespaceDefault).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}
