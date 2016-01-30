package webdriver

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type RequestOrigins struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Version string `json:"version"`
}

type LoggingPrefs struct {
	Browser string `json:"browser"`
	Driver  string `json:"driver"`
}

func BrowserName(val string) func(dc map[string]interface{}) {
	return func(dc map[string]interface{}) {
		dc["browserName"] = val
	}
}

var (
	PageLoadingStrategyEager  = PageLoadingStrategy("eager")
	PageLoadingStrategyNormal = PageLoadingStrategy("normal")
)

func PageLoadingStrategy(val string) func(dc map[string]interface{}) {
	return func(dc map[string]interface{}) {
		dc["pageLoadingStrategy"] = val
	}
}

func Platform(val string) func(dc map[string]interface{}) {
	return func(dc map[string]interface{}) {
		dc["platform"] = val
	}
}

func Version(val string) func(dc map[string]interface{}) {
	return func(dc map[string]interface{}) {
		dc["version"] = val
	}
}

func HandlesAlerts(val bool) func(dc map[string]interface{}) {
	return func(dc map[string]interface{}) {
		dc["handlesAlerts"] = val
	}
}

func JavascriptEnabled(val bool) func(dc map[string]interface{}) {
	return func(dc map[string]interface{}) {
		dc["javascriptEnabled"] = val
	}
}

func LocationContextEnabled(val bool) func(dc map[string]interface{}) {
	return func(dc map[string]interface{}) {
		dc["locationContextEnabled"] = val
	}
}

func Rotatable(val bool) func(dc map[string]interface{}) {
	return func(dc map[string]interface{}) {
		dc["rotatable"] = val
	}
}

func AcceptSslCerts(val bool) func(dc map[string]interface{}) {
	return func(dc map[string]interface{}) {
		dc["acceptSslCerts"] = val
	}
}

func CustomCapability(key, val string) func(dc map[string]interface{}) {
	return func(dc map[string]interface{}) {
		dc[key] = val
	}
}

type Settings struct {
	UserAgent string `json:"userAgent"`
}

type Page struct {
	CustomHeaders map[string]string `json:"customHeaders"`
	Settings      Settings          `json:"settings"`
}

type Proxy struct {
	Type  string `json:"proxyType"`
	Socks string `json:"socksProxy"`
	HTTP  string `json:"httpProxy"`
	SSL   string `json:"sslProxy"`
	FTP   string `json:"ftpProxy"`
}

type Driver struct {
	client  *http.Client
	dc      map[string]interface{}
	BaseURL *url.URL
}

type Session struct {
	wd *Driver

	SessionId                string `json:"sessionId"`
	AcceptSslCerts           bool   `json:"acceptSslCerts"`
	ApplicationCacheEnabled  bool   `json:"applicationCacheEnabled"`
	BrowserConnectionEnabled bool   `json:"browserConnectionEnabled"`
	BrowserName              string `json:"browserName"`
	CssSelectorsEnabled      bool   `json:"cssSelectorsEnabled"`
	DatabaseEnabled          bool   `json:"databaseEnabled"`
	DriverName               string `json:"driverName"`
	DriverVersion            string `json:"driverVersion"`
	HandlesAlerts            bool   `json:"handlesAlerts"`
	JavascriptEnabled        bool   `json:"javascriptEnabled"`
	LocationContextEnabled   bool   `json:"locationContextEnabled"`
	NativeEvents             bool   `json:"nativeEvents"`
	Platform                 string `json:"platform"`
	Proxy                    struct {
		ProxyType string `json:"proxyType"`
	} `json:"proxy"`
	Rotatable         bool   `json:"rotatable"`
	TakesScreenshot   bool   `json:"takesScreenshot"`
	Version           string `json:"version"`
	WebStorageEnabled bool   `json:"webStorageEnabled"`
}

func (c *Driver) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "text/json; charset=UTF-8")
	req.Header.Add("Accept", "text/json")
	return req, nil
}

func New(fns ...func(map[string]interface{})) *Driver {
	dc := map[string]interface{}{}

	for _, fn := range fns {
		fn(dc)
	}

	return &Driver{
		client: http.DefaultClient,
		dc:     dc,
	}
}

type driverResponse struct {
	SessionId string           `json:"sessionId"`
	State     string           `json:"state"`
	Status    int64            `json:"status"`
	Value     *json.RawMessage `json:"value"`
}

func (wd *Driver) do(req *http.Request) (*driverResponse, error) {
	resp, err := wd.client.Do(req)
	if err != nil {
		return nil, err
	}

	dr := driverResponse{}

	r := resp.Body
	defer r.Close()

	err = json.NewDecoder(r).Decode(&dr)
	if err != nil {
		return nil, err
	}

	if dr.Status != 0 {
		return &dr, &Error{
			Status: dr.Status,
			State:  dr.State,
		}
	}

	return &dr, nil
}

func (wd *Driver) Do(req *http.Request, v interface{}) error {
	dr, err := wd.do(req)
	if err != nil {
		return err
	}

	if dr.Value == nil {
		return nil
	}

	switch v := v.(type) {
	case io.Writer:
		value := ""
		if err = json.Unmarshal(*dr.Value, &value); err != nil {
			return err
		}

		v.Write([]byte(value))
	case interface{}:
		return json.Unmarshal(*dr.Value, &v)
	}

	return nil
}

func (wd *Driver) Connect(u string) (*Session, error) {
	baseURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	wd.BaseURL = baseURL

	body := struct {
		DesiredCapabilities map[string]interface{} `json:"desiredCapabilities"`
	}{
		DesiredCapabilities: wd.dc,
	}

	wds := Session{
		wd: wd,
	}

	req, err := wd.NewRequest("POST", "/wd/hub/session", body)
	if err != nil {
		return nil, err
	}

	dr, err := wd.do(req)
	if err != nil {
		return nil, err
	}

	wds.SessionId = dr.SessionId

	err = json.Unmarshal(*dr.Value, &wds)
	if err != nil {
		return nil, err
	}

	return &wds, nil

}
