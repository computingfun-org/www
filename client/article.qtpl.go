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

//line article.qtpl:5
func (a ArticlePage) streamhead(qw422016 *qt422016.Writer) {
	//line article.qtpl:5
	qw422016.N().S(`
<title><%=s a.Title %> - Computing Fun</title>
<meta name="description" content="<%=s a.Details %>">
<meta name="author" content="">
<link rel="stylesheet" type="text/css" href="/client/pages/articles.css">
<link rel="stylesheet" type="text/css" href="/client/base/prettyprint.css">
<script src="https://cdn.rawgit.com/google/code-prettify/master/loader/run_prettify.js"></script>
`)
//line article.qtpl:12
}

//line article.qtpl:12
func (a ArticlePage) writehead(qq422016 qtio422016.Writer) {
	//line article.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line article.qtpl:12
	a.streamhead(qw422016)
	//line article.qtpl:12
	qt422016.ReleaseWriter(qw422016)
//line article.qtpl:12
}

//line article.qtpl:12
func (a ArticlePage) head() string {
	//line article.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
	//line article.qtpl:12
	a.writehead(qb422016)
	//line article.qtpl:12
	qs422016 := string(qb422016.B)
	//line article.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
	//line article.qtpl:12
	return qs422016
//line article.qtpl:12
}

//line article.qtpl:14
func (a ArticlePage) streambody(qw422016 *qt422016.Writer) {
	//line article.qtpl:14
	qw422016.N().S(`
<article class="article">
    <header class="title">
        <h1 class="title-main">
            `)
	//line article.qtpl:18
	qw422016.E().S(a.Title)
	//line article.qtpl:18
	qw422016.N().S(`
        </h1>
        <h2 class="title-sub">
            `)
	//line article.qtpl:21
	qw422016.E().S(a.Details)
	//line article.qtpl:21
	qw422016.N().S(`
        </h2>
        <div style="display: none;" class="info"><address><a rel="author" href="#"></a></address><time pubdate datetime=""></time></div>
    </header>
    <div class="article-content">
        `)
	//line article.qtpl:26
	qw422016.E().S(a.Content)
	//line article.qtpl:26
	qw422016.N().S(`
    </div>
</article>
`)
//line article.qtpl:29
}

//line article.qtpl:29
func (a ArticlePage) writebody(qq422016 qtio422016.Writer) {
	//line article.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line article.qtpl:29
	a.streambody(qw422016)
	//line article.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line article.qtpl:29
}

//line article.qtpl:29
func (a ArticlePage) body() string {
	//line article.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
	//line article.qtpl:29
	a.writebody(qb422016)
	//line article.qtpl:29
	qs422016 := string(qb422016.B)
	//line article.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
	//line article.qtpl:29
	return qs422016
//line article.qtpl:29
}

//line article.qtpl:31
func (a ArticlePage) streamnavLinks(qw422016 *qt422016.Writer) {
//line article.qtpl:31
}

//line article.qtpl:31
func (a ArticlePage) writenavLinks(qq422016 qtio422016.Writer) {
	//line article.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line article.qtpl:31
	a.streamnavLinks(qw422016)
	//line article.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line article.qtpl:31
}

//line article.qtpl:31
func (a ArticlePage) navLinks() string {
	//line article.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
	//line article.qtpl:31
	a.writenavLinks(qb422016)
	//line article.qtpl:31
	qs422016 := string(qb422016.B)
	//line article.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
	//line article.qtpl:31
	return qs422016
//line article.qtpl:31
}