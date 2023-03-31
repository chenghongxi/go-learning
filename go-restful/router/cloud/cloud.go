package cloud

import (
	"github.com/emicklei/go-restful/v3"
	"net/http"
)

// cloudRouter is a router to talk with the cloud controller
type cloudRouter struct {
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
	cloudRoute.Route(cloudRoute.GET("/v1/namespaces").To(s.listNamespace))

	//  Deployments API
	cloudRoute.Route(cloudRoute.GET("/apps/v1/namespaces/:namespace/deployments").To(s.createDeployment))
	cloudRoute.Route(cloudRoute.GET("/apps/v1/:cloud_name/namespaces/:namespace/deployments").To(s.getDeployment))

	restful.Add(cloudRoute)
}

func Load() error {
	http.ListenAndServe(":8080", nil)
	return nil
}
