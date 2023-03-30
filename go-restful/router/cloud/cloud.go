package cloud

import (
	"github.com/emicklei/go-restful/v3"
	"k8s.io/client-go/kubernetes"
)

// cloudRouter is a router to talk with the cloud controller
type cloudRouter struct {
	clientset *kubernetes.Clientset
}

// NewRouter initializes a new cloud router
func NewRouter(restfulEngine *restful.Container) {
	s := &cloudRouter{}
	s.initRoutes(restfulEngine)
}

func (s *cloudRouter) initRoutes(restfulEngine *restful.Container) {
	cloudRoute := new(restful.WebService)
	cloudRoute.
		Path("/clouds").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	// Namespaces API
	cloudRoute.Route(cloudRoute.GET("/v1/:cloud_name/namespaces").To(s.listNamespace))

	//  Deployments API
	cloudRoute.Route(cloudRoute.GET("/apps/v1/:cloud_name/namespaces/:namespace/deployments").To(s.createDeployment))
	cloudRoute.Route(cloudRoute.GET("/apps/v1/:cloud_name/namespaces/:namespace/deployments").To(s.getDeployment))

}
