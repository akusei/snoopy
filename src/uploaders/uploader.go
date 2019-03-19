package uploaders

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
	"time"
)

type Uploader struct {
	uploader  UploadClient
	client    *http.Client
	useragent string
}

func New(uploader UploadClient, options ...UploadOption) *Uploader {
	x := &Uploader{
		uploader:  uploader,
		useragent: "Mozilla/5.0 (Windows; U; MSIE 7.0; Windows NT 6.0; en-US)",
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: time.Second * 30,
		},
	}

	for _, option := range options {
		option(x)
	}

	return x
}

func (up *Uploader) UploadFile(filename string) (link string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	link, err = up.uploader.Upload(up.client, file)

	file.Close()

	return
}

func (up *Uploader) Upload(reader io.Reader) (string, error) {
	return up.uploader.Upload(up.client, reader)
}
