package climatic_webapp

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func init() {
	API{"/admin/api/deliverables", Resource(DeliverablesRes{})}.Register()
}

type DeliverablesRes struct {
	PutNotSupported
	DeleteNotSupported
}

func (m DeliverablesRes) Get(c appengine.Context, r *http.Request) (interface{}, error) {
	//vars := mux.Vars(r)
	c.Infof("inside DeliverablesRes get")
	d := Deliverables{
		OrgName: "CWAOrg",
		Modules: []Module{
			{
				Name: "Orientation",
				Tiers: []Tier{
					{
						Name: "Tier 1",
						Tasks: []Task{
							{
								Name:    "Foo Task",
								Date:    time.Now(),
								BlobKey: appengine.BlobKey(""),
							},
						},
					},
				},
			},
			{
				Name: "Governance",
				Tiers: []Tier{
					{
						Name: "Tier 1",
						Tasks: []Task{
							{
								Name:    "Foo Task",
								Date:    time.Now(),
								BlobKey: appengine.BlobKey(""),
							},
						},
					},
				},
			},
			{
				Name: "Security",
				Tiers: []Tier{
					{
						Name:  "Tier 1",
						Tasks: []Task{},
					},
				},
			},
		},
	}
	return d, nil
}
func (m DeliverablesRes) Post(c appengine.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	c.Infof("inside module resource post", "Got the resource", vars["moduleName"])
	return "Found", nil
}
