package crypto

import "net/http"

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type FileWriter interface {
	Write(p []byte) (n int, err error)
	Close() error
}

type FileCreator interface {
	Create(name string) (FileWriter, error)
}
