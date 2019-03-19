package uploaders

import (
	"net/http"
)

type UploadOption func(*Uploader)

func WithClient(client *http.Client) UploadOption {
	return func(uploader *Uploader) {
		uploader.client = client
	}
}

func WithUseragent(useragent string) UploadOption {
	return func(uploader *Uploader) {
		uploader.useragent = useragent
	}
}
