package climatic_webapp

import (
	"appengine"
	"encoding/json"
)

type Deliverables struct {
	Id      int64 `datastore:",noindex"`
	Modules []Module
	OrgName string
}

type Module struct {
	Name  string
	Tiers []Tier
}

type Tier struct {
	Name  string
	Tasks []Task
}

type Task struct {
	Name    string
	BlobKey appengine.BlobKey
}

type UploadUrl struct {
	url string
}

type DeliverablesAsString struct {
	S string
}

func (d *Deliverables) FromString(c appengine.Context, s string) error {
	return json.Unmarshal([]byte(s), d)
}

func (d *Deliverables) ToString(c appengine.Context) (string, error) {
	b, err := json.Marshal(d)
	return string(b), err
}
