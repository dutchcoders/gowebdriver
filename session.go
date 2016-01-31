package webdriver

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

func (wds *Session) Close() (*Session, error) {
	req, err := wds.wd.NewRequest("DELETE", fmt.Sprintf("/wd/hub/session/%s", wds.SessionId), nil)
	if err != nil {
		return nil, err
	}

	if err := wds.wd.Do(req, &wds); err != nil {
		return nil, err
	}

	return wds, nil

}

func (wds *Session) Source() (io.Reader, error) {
	req, err := wds.wd.NewRequest("GET", fmt.Sprintf("/wd/hub/session/%s/source", wds.SessionId), nil)
	if err != nil {
		return nil, err
	}

	source := &bytes.Buffer{}
	if err := wds.wd.Do(req, source); err != nil {
		return nil, err
	}

	return source, nil
}

func (wds *Session) Screenshot() (io.Reader, error) {
	req, err := wds.wd.NewRequest("GET", fmt.Sprintf("/wd/hub/session/%s/screenshot", wds.SessionId), nil)
	if err != nil {
		return nil, err
	}

	image := &bytes.Buffer{}

	if err := wds.wd.Do(req, image); err != nil {
		return nil, err
	}

	r := base64.NewDecoder(base64.StdEncoding, image)
	return r, nil
}

func (wds *Session) SetWindowSize(width, height int) (*Session, error) {
	body := struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}{
		Width:  width,
		Height: height,
	}

	req, err := wds.wd.NewRequest("POST", fmt.Sprintf("/wd/hub/session/%s/window/current/size", wds.SessionId), body)
	if err != nil {
		return nil, err
	}

	if err := wds.wd.Do(req, &wds); err != nil {
		return nil, err
	}

	return wds, nil

}

func (wds *Session) Back() error {
	req, err := wds.wd.NewRequest("POST", fmt.Sprintf("/wd/hub/session/%s/back", wds.SessionId), nil)
	if err != nil {
		return err
	}

	if err := wds.wd.Do(req, nil); err != nil {
		return err
	}

	return nil
}

func (wds *Session) Forward() error {
	req, err := wds.wd.NewRequest("POST", fmt.Sprintf("/wd/hub/session/%s/forward", wds.SessionId), nil)
	if err != nil {
		return err
	}

	if err := wds.wd.Do(req, &wds); err != nil {
		return err
	}

	return nil
}

func (wds *Session) Refresh() error {
	req, err := wds.wd.NewRequest("POST", fmt.Sprintf("/wd/hub/session/%s/refresh", wds.SessionId), nil)
	if err != nil {
		return err
	}

	if err := wds.wd.Do(req, &wds); err != nil {
		return err
	}

	return nil
}

func (wds *Session) Title() (string, error) {
	req, err := wds.wd.NewRequest("GET", fmt.Sprintf("/wd/hub/session/%s/title", wds.SessionId), nil)
	if err != nil {
		return "", err
	}

	u := ""
	if err := wds.wd.Do(req, &u); err != nil {
		return "", err
	}

	return u, nil
}

func (wds *Session) Url() (string, error) {
	req, err := wds.wd.NewRequest("GET", fmt.Sprintf("/wd/hub/session/%s/url", wds.SessionId), nil)
	if err != nil {
		return "", err
	}

	u := ""
	if err := wds.wd.Do(req, &u); err != nil {
		return "", err
	}

	return u, nil
}

func (wds *Session) SetUrl(url string) error {
	body := struct {
		Url string `json:"url"`
	}{
		Url: url,
	}

	req, err := wds.wd.NewRequest("POST", fmt.Sprintf("/wd/hub/session/%s/url", wds.SessionId), body)
	if err != nil {
		return err
	}

	if err := wds.wd.Do(req, &wds); err != nil {
		return err
	}

	return nil
}
