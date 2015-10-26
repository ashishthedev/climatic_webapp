package climatic_webapp

import (
	"appengine"
	"net/url"
)

func init() {
	m := API{"/admin/api/module/{moduleName}", Resource(ModuleRes{})}
	m.Register()
	t := API{"/test1/", Resource(Test{})}
	t.Register()
}

type ModuleRes struct {
	PutNotSupported
	DeleteNotSupported
}

type Test struct {
	PutNotSupported
	PostNotSupported
	DeleteNotSupported
}

func (t Test) Get(c appengine.Context, values url.Values) (interface{}, error) {
	c.Infof("inside test resource get")
	return "Found", nil
}

func (m ModuleRes) Get(c appengine.Context, values url.Values) (interface{}, error) {
	c.Infof("inside module resource get")
	return "Found", nil
}
func (m ModuleRes) Post(c appengine.Context, values url.Values) (interface{}, error) {
	c.Infof("inside module resource post")
	return "Found", nil
}
