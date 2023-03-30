package menu

import (
	"github.com/emicklei/go-restful/v3"
)

type menuRouter struct{}

func NewRouter(restfulEngine *restful.Container) {
	u := &menuRouter{}
	u.initRoutes(restfulEngine)
}

func (m *menuRouter) initRoutes(restfulEngine *restful.Container) {

}
