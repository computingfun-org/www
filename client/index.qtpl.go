// This file is automatically generated by qtc from "index.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line index.qtpl:1
package client

//line index.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line index.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line index.qtpl:1
type IndexPage struct{}

//line index.qtpl:3
func (i IndexPage) streamhead(qw422016 *qt422016.Writer) {
	//line index.qtpl:3
	qw422016.N().S(`
<title>Computing Fun</title>
<meta name="description" content="We build, teach and deploy all types of software solutions. From teaching non tech to developers. Building software and web servers of small businesses.">
<link rel="stylesheet" type="text/css" href="/client/pages/index.css">
`)
//line index.qtpl:7
}

//line index.qtpl:7
func (i IndexPage) writehead(qq422016 qtio422016.Writer) {
	//line index.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line index.qtpl:7
	i.streamhead(qw422016)
	//line index.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line index.qtpl:7
}

//line index.qtpl:7
func (i IndexPage) head() string {
	//line index.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
	//line index.qtpl:7
	i.writehead(qb422016)
	//line index.qtpl:7
	qs422016 := string(qb422016.B)
	//line index.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
	//line index.qtpl:7
	return qs422016
//line index.qtpl:7
}

//line index.qtpl:9
func (i IndexPage) streambody(qw422016 *qt422016.Writer) {
	//line index.qtpl:9
	qw422016.N().S(`
<div class="nav-boxes">
    <a style="background-color: #ff0000;" href="https://www.youtube.com/channel/UCeZQbACMihORscFIwmydpzA"><i class="fab fa-youtube"></i></a>
    <a style="background-color: #f85944;" href="https://www.patreon.com/computingfun"><i class="fab fa-patreon"></i></a>
    <a style="background-color: #171A21;" href="https://computingfun.org/games/"><i class="fas fa-gamepad"></i></a>
    <a style="background-color: #77ab59;" href="https://computingfun.org/articles/"><i class="far fa-newspaper"></i></a>
    <a style="background-color: #3B5998;" href="https://www.facebook.com/ComputingFun/"><i class="fab fa-facebook-f"></i></a>
    <a style="background-color: #6441A4;" href="#"><i class="fab fa-twitch"></i></a>
</div>
`)
//line index.qtpl:18
}

//line index.qtpl:18
func (i IndexPage) writebody(qq422016 qtio422016.Writer) {
	//line index.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line index.qtpl:18
	i.streambody(qw422016)
	//line index.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line index.qtpl:18
}

//line index.qtpl:18
func (i IndexPage) body() string {
	//line index.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
	//line index.qtpl:18
	i.writebody(qb422016)
	//line index.qtpl:18
	qs422016 := string(qb422016.B)
	//line index.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
	//line index.qtpl:18
	return qs422016
//line index.qtpl:18
}

//line index.qtpl:20
func (i IndexPage) streamnavLinks(qw422016 *qt422016.Writer) {
//line index.qtpl:20
}

//line index.qtpl:20
func (i IndexPage) writenavLinks(qq422016 qtio422016.Writer) {
	//line index.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line index.qtpl:20
	i.streamnavLinks(qw422016)
	//line index.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line index.qtpl:20
}

//line index.qtpl:20
func (i IndexPage) navLinks() string {
	//line index.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
	//line index.qtpl:20
	i.writenavLinks(qb422016)
	//line index.qtpl:20
	qs422016 := string(qb422016.B)
	//line index.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
	//line index.qtpl:20
	return qs422016
//line index.qtpl:20
}
