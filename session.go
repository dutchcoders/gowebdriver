package webdriver

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

func (wds *Session) Close() (*Session, error) {
	req, err := wds.wd.NewRequest("DELETE", fmt.Sprintf("/wd/hub/session/%s", wds.SessionId), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := wds.wd.Do(req, &wds); err != nil {
		return nil, err
	}

	return wds, nil

}

func (wds *Session) Screenshot() (io.Reader, error) {
	req, err := wds.wd.NewRequest("GET", fmt.Sprintf("/wd/hub/session/%s/screenshot", wds.SessionId), nil)
	if err != nil {
		return nil, err
	}

	image := &bytes.Buffer{}

	if err := wds.wd.Do(req, image); err != nil {
		panic(err)
		return nil, err
	}

	r := base64.NewDecoder(base64.StdEncoding, image)
	return r, nil
}

func (wds *Session) WindowSize(width, height int) (*Session, error) {
	body := struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}{
		Width:  width,
		Height: height,
	}

	req, err := wds.wd.NewRequest("POST", fmt.Sprintf("/wd/hub/session/%s/window/current/size", wds.SessionId), body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := wds.wd.Do(req, &wds); err != nil {
		return nil, err
	}

	return wds, nil

}
func (wds *Session) Url(url string) (*Session, error) {
	body := struct {
		Url string `json:"url"`
	}{
		Url: url,
	}

	req, err := wds.wd.NewRequest("POST", fmt.Sprintf("/wd/hub/session/%s/url", wds.SessionId), body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := wds.wd.Do(req, &wds); err != nil {
		return nil, err
	}

	return wds, nil

}