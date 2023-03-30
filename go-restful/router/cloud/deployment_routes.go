package cloud

import (
	"context"
	"net/http"

	"github.com/emicklei/go-restful/v3"
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
