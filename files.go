package climatic_webapp

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/blobstore"
)

const UPLOAD_API = "/admin/api/upload"

func init() {
	API{UPLOAD_API, Resource(UploadRes{})}.Register()
}

type UploadRes struct {
	PutNotSupported
	DeleteNotSupported
}

func (x UploadRes) Get(c appengine.Context, r *http.Request) (interface{}, error) {
	//return blobstore.UploadURL(c, UPLOAD_API, nil)
	url, err := blobstore.UploadURL(c, UPLOAD_API, nil)
	PrintInBox(c, "Generated upload URL is : ", url)
	return url, err
}

func (x UploadRes) Post(c appengine.Context, r *http.Request) (interface{}, error) {
	PrintInBox(c, "Got the request to save file at path:", r.URL.Path)
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
