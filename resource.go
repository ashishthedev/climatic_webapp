package climatic_webapp

import (
	"fmt"
	"net/http"

	"appengine"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Resource interface {
	Get(c appengine.Context, r *http.Request) (interface{}, error)
	Post(c appengine.Context, r *http.Request) (interface{}, error)
	Put(c appengine.Context, r *http.Request) (interface{}, error)
	Delete(c appengine.Context, r *http.Request) (interface{}, error)
}

type ExampleOfGetResource struct {
	PutNotSupported
	PostNotSupported
	DeleteNotSupported
}

type (
	GetNotSupported    struct{}
	PostNotSupported   struct{}
	PutNotSupported    struct{}
	DeleteNotSupported struct{}
)

func (GetNotSupported) Get(c appengine.Context, r *http.Request) (interface{}, error) {
	return nil, fmt.Errorf("GET not implemented")
}
func (PostNotSupported) Post(c appengine.Context, r *http.Request) (interface{}, error) {
	return nil, fmt.Errorf("POST not implemented")
}
func (PutNotSupported) Put(c appengine.Context, r *http.Request) (interface{}, error) {
	return nil, fmt.Errorf("PUT not implemented")
}
func (DeleteNotSupported) Delete(c appengine.Context, r *http.Request) (interface{}, error) {
	return nil, fmt.Errorf("DELETE not implemented")
}

type API struct {
	path     string
	resource Resource
}

func (api API) Abort(c appengine.Context, w http.ResponseWriter, err error, httpStatusCode int) {
	c.Errorf("%v", err.Error())
	http.Error(w, err.Error(), httpStatusCode)
}

func (api API) Register() API {
	router.Handle(api.path, apiHandler(api))
	return api
}

//Convenience wrapper function
func apiHandler(api API) gaeHandler {
	h := func(c appengine.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {

		var data interface{}
		var err error

		// r.ParseForm()
		// values := r.Form

		switch r.Method {
		case GET:
			data, err = api.resource.Get(c, r)
		case POST:
			data, err = api.resource.Post(c, r)
		case PUT:
			data, err = api.resource.Put(c, r)
		case DELETE:
			data, err = api.resource.Delete(c, r)
		default:
			return nil, fmt.Errorf(r.Method + " not implemented")
		}

		if err = WriteJson(w, data); err != nil {
			return nil, err
		}
		return nil, nil
	}
	return gaeHandler(h)
}
