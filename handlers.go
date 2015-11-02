package climatic_webapp

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"appengine"
	"appengine/blobstore"
)

type gaeHandler func(c appengine.Context, w http.ResponseWriter, r *http.Request) (interface{}, error)

func (h gaeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	val, err := h(c, w, r)
	if err == nil {
		err = json.NewEncoder(w).Encode(val)
	}
	if err != nil {
		c.Errorf("%v", err.Error())
		CLWAReportErrorThroughMail(r, "Error happended while serving"+r.URL.Path, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type urlPack struct {
	urlPath            string
	handler            http.HandlerFunc
	data               interface{}
	dependentTemplates []string
}

type TEMPLATE_AND_DATA struct {
	t    *template.Template
	data interface{}
}

var BASE_TEMPLATE = "templates/base.html"
var ADMIN_BASE = "templates/adminbase.html"
var PRE_BUILT_TEMPLATES_WITH_DATA = make(map[string]TEMPLATE_AND_DATA)
var PAGE_NOT_FOUND_TEMPLATE = template.Must(template.ParseFiles("templates/pageNotFound.html", BASE_TEMPLATE))

var router = new(mux.Router)

func initStaticHTMLUrlMaps() {
	urlPacks := []urlPack{
		{"/", pageHandler, nil, []string{"templates/home.html", BASE_TEMPLATE}},
		{"/admin", pageHandler, nil, []string{"templates/admin_tasks.html", BASE_TEMPLATE, ADMIN_BASE}},
		{"/admin/tasks", pageHandler, nil, []string{"templates/admin_tasks.html", BASE_TEMPLATE, ADMIN_BASE}},
		{"/admin/modules", pageHandler, nil, []string{"templates/admin_modules.html", BASE_TEMPLATE, ADMIN_BASE}},
		{"/admin/tiers", pageHandler, nil, []string{"templates/admin_tiers.html", BASE_TEMPLATE, ADMIN_BASE}},
	}

	for _, x := range urlPacks {
		PRE_BUILT_TEMPLATES_WITH_DATA[x.urlPath] = TEMPLATE_AND_DATA{template.Must(template.ParseFiles(x.dependentTemplates...)), x.data}
		router.HandleFunc(x.urlPath, x.handler)
	}
}

func init() {
	initStaticHTMLUrlMaps()
	http.Handle("/", router)
	http.HandleFunc("/serve", ServeFile)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	t := PRE_BUILT_TEMPLATES_WITH_DATA[r.URL.Path].t
	data := PRE_BUILT_TEMPLATES_WITH_DATA[r.URL.Path].data
	if t == nil {
		t = PAGE_NOT_FOUND_TEMPLATE
		data = nil
	}

	if err := t.ExecuteTemplate(w, "base", data); err != nil {
		c := appengine.NewContext(r)
		c.Errorf("%v", err.Error())
		CLWAReportErrorThroughMail(r, "Error happended while serving"+r.URL.Path, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ServeFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bk := r.FormValue("BlobKey")
	PrintInBox(appengine.NewContext(r), "Got the BlobKey as ", bk)
	blobstore.Send(w, appengine.BlobKey(bk))
}
