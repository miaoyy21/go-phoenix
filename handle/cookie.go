package handle

import (
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, name string, value string, expire time.Time) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		Value:    value,
		HttpOnly: true,
		Expires:  expire,
	}

	http.SetCookie(w, cookie)
}
