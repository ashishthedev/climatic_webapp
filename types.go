package climatic_webapp

import (
	"time"

	"appengine"
)

type Deliverables struct {
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
	Date    time.Time
	BlobKey appengine.BlobKey
}

type UploadUrl struct {
	url string
}
