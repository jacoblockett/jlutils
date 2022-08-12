package jlutils

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(u string, p string) error {
	r, err := http.Get(u)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	o, err := os.Create(p)
	if err != nil {
		return err
	}
	defer o.Close()

	_, err = io.Copy(o, r.Body)
	if err != nil {
		return err
	}

	return nil
}
