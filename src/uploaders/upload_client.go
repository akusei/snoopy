package uploaders

import (
	"io"
	"net/http"
)

type UploadClient interface {
	Upload(client *http.Client, reader io.Reader) (string, error)
}
