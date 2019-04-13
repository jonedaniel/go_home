package main

import (
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"net/http"
	"strings"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, e := r.Cookie("auth")
	if e == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	h.next.ServeHTTP(w, r)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to provider %s:%s", provider, err), http.StatusBadRequest)
			return
		}
		loginUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to GetBeginAuthURL for %s:%s", provider, err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", loginUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		p, e := gomniauth.Provider(provider)
		if e != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get provider %s : %s", p, e), http.StatusBadRequest)
			return
		}
		s := string(r.URL.RawQuery)
		maps, e := objx.FromURLQuery(s)
		credentials, e := p.CompleteAuth(maps)
		if e != nil {
			http.Error(w, fmt.Sprintf("Error when trying to complete auth for %s:%s", p, e), http.StatusInternalServerError)
			return
		}
		user, e := p.GetUser(credentials)
		if e != nil {
			http.Error(w, fmt.Sprintf("error when trying to get user info %s :%s", p, e), http.StatusInternalServerError)
			return
		}
		authCookieVal := objx.New(map[string]interface{}{
			"name":       user.Name(),
			"avatar_url": user.AvatarURL(),
		}).MustBase64()
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieVal,
			Path:  "/"})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
