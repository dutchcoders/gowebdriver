gowebdriver
===========

Golang Webdriver library 

## Sample
```
package main

import (
        "io"
        "os"

        webdriver "github.com/dutchcoders/gowebdriver"
)

func main() {
        wd := webdriver.New(
                webdriver.BrowserName("phantomjs"),
                webdriver.PageLoadingStrategyEager,
                webdriver.AcceptSslCerts(true),
                webdriver.Platform("ANY"),
                webdriver.Version(""),
                webdriver.LocationContextEnabled(true),
                webdriver.JavascriptEnabled(true),
                webdriver.HandlesAlerts(true),
                webdriver.Rotatable(true),
                webdriver.CustomCapability("phantomjs.page.settings.userAgent", "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36"),
                webdriver.CustomCapability("phantomjs.page.customHeaders.Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"),
                webdriver.CustomCapability("phantomjs.page.customHeaders.Accept-Language", "ru-RU"),
        )

        session, err := wd.Connect("http://127.0.0.1:4444")
        if err != nil {
                panic(err)
        }

        defer session.Close()

        if _, err = session.WindowSize(2048, 1680); err != nil {
                panic(err)
        }
        if _, err = session.Url("http://httpbin.org/headers"); err != nil {
                panic(err)
        }

        if r, err := session.Screenshot(); err != nil {
                panic(err)
        } else {
                w, _ := os.Create("screenshot.png")
                defer w.Close()

                io.Copy(w, r)
        }
}
```

## Testing
```
go test
```

## References
* https://github.com/stuart/elixir-webdriver/blob/master/lib/webdriver/session.ex
* https://w3c.github.io/webdriver/webdriver-spec.html
* https://selenium.googlecode.com/git/docs/api/java/constant-values.html#org.openqa.selenium.remote.CapabilityType.PAGE_LOADING_STRATEGY


## Contributors

**Remco Verhoef**
- <https://twitter.com/remco_verhoef>
- <https://twitter.com/dutchcoders>

## Copyright and license

Code and documentation copyright 2011-2016 Remco Verhoef.

Code released under [the MIT license](LICENSE).
~

