package xsys

import (
	"go-phoenix/handle"
	"net/http"
	"time"
)

func setCookie(ctx *handle.Context, name string, value string, expire time.Time) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		Value:    value,
		HttpOnly: false,
		Expires:  expire,
	}

	http.SetCookie(ctx.Writer, cookie)
}
