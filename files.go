package climatic_webapp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"appengine"
	"appengine/blobstore"
)

const UPLOADURL_API = "/admin/api/uploadurl"
const REMOVE_TASKFILE_API = "/admin/api/removetaskfile"

func init() {
	API{UPLOADURL_API, Resource(UploadRes{})}.Register()
	API{REMOVE_TASKFILE_API, Resource(RemoveTaskFileRes{})}.Register()
}

type UploadRes struct {
	PutNotSupported
	DeleteNotSupported
}

type RemoveTaskFileRes struct {
	GetNotSupported
	PutNotSupported
	DeleteNotSupported
}

func (x UploadRes) Get(c appengine.Context, r *http.Request) (interface{}, error) {
	return blobstore.UploadURL(c, UPLOADURL_API, nil)
}

func (x UploadRes) Post(c appengine.Context, r *http.Request) (interface{}, error) {
	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {
		return nil, err
	}
	file := blobs["file"]
	if len(file) == 0 {
		return nil, fmt.Errorf("No file uploaded")
	}
	bk := file[0].BlobKey
	PrintInBox(c, "File saved with blobkey", bk)
	return bk, nil
}

func (x RemoveTaskFileRes) Post(c appengine.Context, r *http.Request) (interface{}, error) {
	t := new(Task)
	if err := json.NewDecoder(r.Body).Decode(t); err != nil {
		return nil, err
	}
	PrintInBox(c, "Trying to remove", t.BlobKey)
	if err := blobstore.Delete(c, t.BlobKey); err != nil {
		return nil, err
	}
	return nil, blobstore.Delete(c, t.BlobKey)
}
