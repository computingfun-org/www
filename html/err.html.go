// Code generated by hero.
// source: C:\Users\zac22\go\src\gitlab.com\computingfun\computingfun.org\html\err.html
// DO NOT EDIT!
package html

import (
	"io"

	"github.com/shiyanhui/hero"
)

func Err(e *ErrMessage, w io.Writer) (int, error) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" name="viewport" content="width=device-width, initial-scale=1">
        <title>`)
	hero.EscapeHTML(e.Title, _buffer)
	_buffer.WriteString(` - Computing Fun`)

	_buffer.WriteString(`</title>
        <link rel="icon" type="image/svg+xml" href="/client/icon/main.svg">
        `)
	_buffer.WriteString(`
        <link rel="stylesheet" type="text/css" href="/client/base/body.css">
        <link rel="stylesheet" type="text/css" href="/client/base/navbar.css">
        <link rel="stylesheet" type="text/css" href="https://use.fontawesome.com/releases/v5.1.0/css/all.css" integrity="sha384-lKuwvrZot6UHsBSfcMvOkWwlCMgc0TaWr+30HWe3a4ltaBwTZhyTEggF5tJv8tbt" crossorigin="anonymous">
        `)
	_buffer.WriteString(`
    <link rel="stylesheet" type="text/css" href="/client/base/error.css">
`)

	_buffer.WriteString(`
    </head>
    <body>
        <nav class="navbar">
            <a class="navbar-title" href="/">Computing Fun</a>
            <img src="/client/icon/main.svg" alt="Computing Fun" height="50" width="50" class="navbar-icon">
            <a class="navbar-link" style="background-color: orange;" href="#">Find more articles.<i class="fa fa-newspaper"></i></a>
            <a class="navbar-link" style="background-color: grey;" href="https://www.patreon.com/computingfun">Become a Patron.<i class="fab fa-patreon"></i></a>
            <a class="navbar-link" style="background-color: red;" href="https://www.youtube.com/channel/UCeZQbACMihORscFIwmydpzA">Watch for free on YouTube.<i class="fab fa-youtube"></i></a>
            <a class="navbar-link" style="background-color: purple;" href="#">See us live on Twitch.<i class="fab fa-twitch"></i></a>
            <a class="navbar-link" style="background-color: orangered;" href="http://git.computingfun.org">Check out our GitLab.<i class="fab fa-gitlab"></i></a>
        </nav>
        <main class="content">`)
	_buffer.WriteString(`
    <img class="icon" src="/client/icon/err.svg" alt="Computing Fun Error" height="250" width="250">
    <div class="message">`)
	hero.EscapeHTML(e.Message, _buffer)
	_buffer.WriteString(`</div>
    <div class="code">`)
	hero.EscapeHTML(e.Code, _buffer)
	_buffer.WriteString(`</div>
`)

	_buffer.WriteString(`</main>
    </body>
</html>`)
	return w.Write(_buffer.Bytes())

}
