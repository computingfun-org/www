// This file is automatically generated by qtc from "page.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line page.qtpl:2
package client

//line page.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line page.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line page.qtpl:2
type Page interface {
	//line page.qtpl:2
	head() string
	//line page.qtpl:2
	streamhead(qw422016 *qt422016.Writer)
	//line page.qtpl:2
	writehead(qq422016 qtio422016.Writer)
	//line page.qtpl:2
	navLinks() string
	//line page.qtpl:2
	streamnavLinks(qw422016 *qt422016.Writer)
	//line page.qtpl:2
	writenavLinks(qq422016 qtio422016.Writer)
	//line page.qtpl:2
	body() string
	//line page.qtpl:2
	streambody(qw422016 *qt422016.Writer)
	//line page.qtpl:2
	writebody(qq422016 qtio422016.Writer)
//line page.qtpl:2
}

//line page.qtpl:11
func StreamHTML(qw422016 *qt422016.Writer, p Page) {
	//line page.qtpl:11
	qw422016.N().S(`<!DOCTYPE html><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><link rel="shortcut icon" href="/client/ico/fav.ico"><link rel="icon" type="image/png" href="/client/ico/192.png" sizes="192x192"><link rel="apple-touch-icon" sizes="180x180" href="/client/ico/180.png"><link rel="stylesheet" type="text/css" href="/client/base/body.css"><link rel="stylesheet" type="text/css" href="/client/base/nav.css"><link rel="stylesheet" type="text/css" href="/client/base/nav-link.css"><link rel="stylesheet" type="text/css" href="https://use.fontawesome.com/releases/v5.1.0/css/all.css" integrity="sha384-lKuwvrZot6UHsBSfcMvOkWwlCMgc0TaWr+30HWe3a4ltaBwTZhyTEggF5tJv8tbt" crossorigin="anonymous"><script src="/client/base/dark-theme.js" async defer></script>`)
	//line page.qtpl:25
	p.streamhead(qw422016)
	//line page.qtpl:25
	qw422016.N().S(`</head><body onload="DarkThemeLoad();" class="dark-theme-tag"><nav class="dark-theme-tag"><div>Computing Fun</div><img src="/client/ico/50.png" alt="Computing Fun" height="50" width="50"><a id="nav-link-yt" href="https://www.youtube.com/channel/UCeZQbACMihORscFIwmydpzA" class="dark-theme-tag"><span>Channel</span><i class="fab fa-youtube"></i></a><a id="nav-link-patron" href="https://www.patreon.com/computingfun" class="dark-theme-tag"><span>Patron</span><i class="fab fa-patreon"></i></a><a id="nav-link-game" href="https://computingfun.org/games/" class="dark-theme-tag"><span>Games</span><i class="fas fa-gamepad"></i></a><a id="nav-link-article" href="https://computingfun.org/articles/" class="dark-theme-tag"><span>Articles</span><i class="far fa-newspaper"></i></a><a id="nav-link-twitch" href="#" class="dark-theme-tag"><span>Live</span><i class="fab fa-twitch"></i></a><a id="nav-link-mode" href="#" onclick="DarkThemeToggle();" class="dark-theme-tag"><span> Mode</span><i class="fas fa-palette"></i></a>`)
	//line page.qtpl:37
	p.streamnavLinks(qw422016)
	//line page.qtpl:37
	qw422016.N().S(`</nav><main>`)
	//line page.qtpl:40
	p.streambody(qw422016)
	//line page.qtpl:40
	qw422016.N().S(`</main></body></html>`)
//line page.qtpl:44
}

//line page.qtpl:44
func WriteHTML(qq422016 qtio422016.Writer, p Page) {
	//line page.qtpl:44
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line page.qtpl:44
	StreamHTML(qw422016, p)
	//line page.qtpl:44
	qt422016.ReleaseWriter(qw422016)
//line page.qtpl:44
}

//line page.qtpl:44
func HTML(p Page) string {
	//line page.qtpl:44
	qb422016 := qt422016.AcquireByteBuffer()
	//line page.qtpl:44
	WriteHTML(qb422016, p)
	//line page.qtpl:44
	qs422016 := string(qb422016.B)
	//line page.qtpl:44
	qt422016.ReleaseByteBuffer(qb422016)
	//line page.qtpl:44
	return qs422016
//line page.qtpl:44
}
