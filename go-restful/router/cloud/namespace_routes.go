package cloud

import (
	"github.com/emicklei/go-restful/v3"
)

type CloudMeta struct {
	Cloud string `uri:"cloud_name" binding:"required"`
}

func (s *cloudRouter) listNamespace(c *restful.Request, response *restful.Response) {

}
