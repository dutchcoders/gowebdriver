package webdriver

import (
	"io"
	"os"
	"testing"
)

func TestScreenshot(t *testing.T) {
	wd := New(
		BrowserName("phantomjs"),
		PageLoadingStrategyEager,
		AcceptSslCerts(true),
		Platform("ANY"),
		Version(""),
		LocationContextEnabled(true),
		JavascriptEnabled(true),
		HandlesAlerts(true),
		Rotatable(true),
		CustomCapability("phantomjs.page.settings.userAgent", "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36"),
		CustomCapability("phantomjs.page.customHeaders.Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"),
		CustomCapability("phantomjs.page.customHeaders.Accept-Language", "ru-RU"),
	)

	session, err := wd.Connect("http://127.0.0.1:4444")
	if err != nil {
		t.Error(err)
	}

	defer session.Close()

	if _, err = session.WindowSize(2048, 1680); err != nil {
		t.Error(err)
	}

	if _, err = session.Url("http://httpbin.org/headers"); err != nil {
		t.Error(err)
	}

	if _, err := session.Source(); err != nil {
		t.Error(err)
	}

	if r, err := session.Screenshot(); err != nil {
		t.Error(err)
	} else {
		w, _ := os.Create("screenshot.png")
		defer w.Close()

		io.Copy(w, r)
	}
}
