package climatic_webapp

import (
	"appengine"
	"github.com/gorilla/mux"
	"net/http"
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
								Name:    "Task 1",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 2",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 3",
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
								Name:    "Task 1",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 2",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 3",
								BlobKey: appengine.BlobKey(""),
							},
						},
					},
				},
			},
			{
				Name: "Scope",
				Tiers: []Tier{
					{
						Name: "Tier 1",
						Tasks: []Task{
							{
								Name:    "Task 1",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 2",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 3",
								BlobKey: appengine.BlobKey(""),
							},
						},
					},
				},
			},
			{
				Name: "Risk Analysis",
				Tiers: []Tier{
					{
						Name: "Tier 1",
						Tasks: []Task{
							{
								Name:    "Task 1",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 2",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 3",
								BlobKey: appengine.BlobKey(""),
							},
						},
					},
				},
			},
			{
				Name: "Resilence Plan",
				Tiers: []Tier{
					{
						Name: "Tier 1",
						Tasks: []Task{
							{
								Name:    "Task 1",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 2",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 3",
								BlobKey: appengine.BlobKey(""),
							},
						},
					},
				},
			},
			{
				Name: "Business Case",
				Tiers: []Tier{
					{
						Name: "Tier 1",
						Tasks: []Task{
							{
								Name:    "Task 1",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 2",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 3",
								BlobKey: appengine.BlobKey(""),
							},
						},
					},
				},
			},
			{
				Name: "Implementation",
				Tiers: []Tier{
					{
						Name: "Tier 1",
						Tasks: []Task{
							{
								Name:    "Task 1",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 2",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 3",
								BlobKey: appengine.BlobKey(""),
							},
						},
					},
				},
			},
			{
				Name: "Evolution",
				Tiers: []Tier{
					{
						Name: "Tier 1",
						Tasks: []Task{
							{
								Name:    "Task 1",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 2",
								BlobKey: appengine.BlobKey(""),
							},
							{
								Name:    "Task 3",
								BlobKey: appengine.BlobKey(""),
							},
						},
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
