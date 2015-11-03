package climatic_webapp

import (
	"encoding/json"
	"net/http"

	"appengine"
	"appengine/datastore"
)

func init() {
	API{"/admin/api/deliverables", Resource(DeliverablesRes{})}.Register()
}

type DeliverablesRes struct {
	PutNotSupported
	DeleteNotSupported
}

// orgKey returns the key used for all entries for this organization.
func orgKey(c appengine.Context, orgName string) *datastore.Key {
	// The string orgName here could be varied to have multiple Organizations.
	return datastore.NewKey(c, "OrgName", orgName, 0, nil)
}

// deliverables_key returns the key used for deliverables entry.
func deliverablesKey(c appengine.Context, orgName string) *datastore.Key {
	return datastore.NewKey(c, "DeliverablesAsString", "deliverables_keyID", 0, orgKey(c, orgName))
}

func GetDeliverables(c appengine.Context, orgName string) (*Deliverables, error) {

	das := new(DeliverablesAsString)
	k := deliverablesKey(c, orgName)
	if err := datastore.Get(c, k, das); err != nil {
		if err == datastore.ErrNoSuchEntity {
			das.S, err = new(Deliverables).ToString(c) //Pretend as if you have fetched an empty entity.
		} else {
			return nil, err
		}
	}

	d := new(Deliverables)
	if err := d.FromString(c, das.S); err != nil {
		return nil, err
	}
	return d, nil
}

func (m DeliverablesRes) Get(c appengine.Context, r *http.Request) (interface{}, error) {
	c.Infof("inside DeliverablesRes get")
	orgName := "Org1" //TODO: Handle with sessions?

	return GetDeliverables(c, orgName)

}

func (m DeliverablesRes) Post(c appengine.Context, r *http.Request) (interface{}, error) {
	c.Infof("inside DeliverablesRes post")
	orgName := "Org1" //TODO: Handle with sessions?

	d := new(Deliverables)
	if err := json.NewDecoder(r.Body).Decode(d); err != nil {
		return nil, err
	}
	k := deliverablesKey(c, orgName)
	das := new(DeliverablesAsString)
	var err error
	das.S, err = d.ToString(c)
	if err != nil {
		return nil, err
	}

	if _, err := datastore.Put(c, k, das); err != nil {
		return nil, err
	}
	return nil, nil
}
