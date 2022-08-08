package context

import "net/http"

type Context interface {
	Request() *http.Request
	Path() string
	Bind(i interface{}) error

	Params
	Cookies
	Response
}

type Params interface {
	PathParam(name string) string
	QueryParam(name string) string
	FormParam(name string) string
}

type Cookies interface {
	Cookie(name string) (*http.Cookie, error)
	SetCookie(cookie *http.Cookie)
}

type Response interface {
	Render(code int, name string, data interface{}) error

	HTML(code int, html string) error
	String(code int, s string) error
	JSON(code int, i interface{}) error
	JSONPretty(code int, i interface{}, indent string) error
	XML(code int, i interface{}) error
	XMLPretty(code int, i interface{}, indent string) error

	File(file string) error
	Attachment(file string, name string) error

	Redirect(code int, url string) error
}
