package climatic_webapp

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	API{"/admin/api/module/{moduleName}", Resource(ModuleResource{})}.Register()
	API{"/admin/api/module/{moduleName}/deliverable/{deliverableName}", Resource(DeliverableResource{})}.Register()
}

type ModuleResource struct {
	PutNotSupported
	DeleteNotSupported
}

func (m ModuleResource) Get(c appengine.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	c.Infof("inside module resource get", "Got the resource", vars["moduleName"])
	return "Found", nil
}
func (m ModuleResource) Post(c appengine.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	c.Infof("inside module resource post", "Got the resource", vars["moduleName"])
	return "Found", nil
}

type DeliverableResource struct {
	PutNotSupported
	DeleteNotSupported
}

func (m DeliverableResource) Get(c appengine.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	c.Infof("inside DeliverableResource get", "Got the resource", vars["moduleName"])
	return "Found", nil
}

func (m DeliverableResource) Post(c appengine.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	c.Infof("inside deliverable resource post", "Got the resource", vars["moduleName"], vars["deliverableName"])
	return "Found", nil
}
