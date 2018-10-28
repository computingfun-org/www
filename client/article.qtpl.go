// This file is automatically generated by qtc from "article.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line article.qtpl:1
package client

//line article.qtpl:1
import "gitlab.com/computingfun/www/articles"

//line article.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line article.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line article.qtpl:3
type ArticlePage articles.Article

//line article.qtpl:7
func (a ArticlePage) streamhead(qw422016 *qt422016.Writer) {
	//line article.qtpl:7
	qw422016.N().S(`<title>`)
	//line article.qtpl:8
	qw422016.E().S(a.Title)
	//line article.qtpl:8
	qw422016.N().S(`- Computing Fun</title><meta name="description" content="<%=s a.Details %>"><meta name="author" content=""><link rel="stylesheet" type="text/css" href="/client/pages/articles.css"><link rel="stylesheet" type="text/css" href="/client/base/prettyprint.css"><script src="https://cdn.rawgit.com/google/code-prettify/master/loader/run_prettify.js"></script>`)
//line article.qtpl:14
}

//line article.qtpl:14
func (a ArticlePage) writehead(qq422016 qtio422016.Writer) {
	//line article.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line article.qtpl:14
	a.streamhead(qw422016)
	//line article.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line article.qtpl:14
}

//line article.qtpl:14
func (a ArticlePage) head() string {
	//line article.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
	//line article.qtpl:14
	a.writehead(qb422016)
	//line article.qtpl:14
	qs422016 := string(qb422016.B)
	//line article.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
	//line article.qtpl:14
	return qs422016
//line article.qtpl:14
}

//line article.qtpl:16
func (a ArticlePage) streambody(qw422016 *qt422016.Writer) {
	//line article.qtpl:16
	qw422016.N().S(`<article class="article"><header class="title"><h1 class="title-main">`)
	//line article.qtpl:20
	qw422016.E().S(a.Title)
	//line article.qtpl:20
	qw422016.N().S(`</h1><h2 class="title-sub">`)
	//line article.qtpl:23
	qw422016.E().S(a.Details)
	//line article.qtpl:23
	qw422016.N().S(`</h2><div style="display: none;" class="info"><address><a rel="author" href="#"></a></address><time pubdate datetime=""></time></div></header><div class="article-content">`)
	//line article.qtpl:28
	qw422016.N().S(a.Content)
	//line article.qtpl:28
	qw422016.N().S(`</div></article>`)
//line article.qtpl:31
}

//line article.qtpl:31
func (a ArticlePage) writebody(qq422016 qtio422016.Writer) {
	//line article.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line article.qtpl:31
	a.streambody(qw422016)
	//line article.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line article.qtpl:31
}

//line article.qtpl:31
func (a ArticlePage) body() string {
	//line article.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
	//line article.qtpl:31
	a.writebody(qb422016)
	//line article.qtpl:31
	qs422016 := string(qb422016.B)
	//line article.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
	//line article.qtpl:31
	return qs422016
//line article.qtpl:31
}

//line article.qtpl:33
func (a ArticlePage) streamnavLinks(qw422016 *qt422016.Writer) {
//line article.qtpl:33
}

//line article.qtpl:33
func (a ArticlePage) writenavLinks(qq422016 qtio422016.Writer) {
	//line article.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line article.qtpl:33
	a.streamnavLinks(qw422016)
	//line article.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line article.qtpl:33
}

//line article.qtpl:33
func (a ArticlePage) navLinks() string {
	//line article.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
	//line article.qtpl:33
	a.writenavLinks(qb422016)
	//line article.qtpl:33
	qs422016 := string(qb422016.B)
	//line article.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
	//line article.qtpl:33
	return qs422016
//line article.qtpl:33
}
