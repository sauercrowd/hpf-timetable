package web

import (
	"log"
	"net/http"
)

type WebContext struct {
}

type TimetableHandler struct {
	*WebContext
	Handler func(*WebContext, http.ResponseWriter, *http.Request) (int, error)
}

//Custom Handler in order to provide context
func (th TimetableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := th.Handler(th.WebContext, w, r)
	if err == nil {
		log.Printf("[OK] [HTTP %d]: %s", status, r.RequestURI)
	}
	if err != nil {
		log.Printf("[ERROR] [HTTP %d]: %s %q", status, r.RequestURI, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}

func BulkTimetableInsert(ctx *WebContext, w http.ResponseWriter, r *http.Request) (int, error) {
	return 200, nil
}

func RegisterRoutes() {
	http.Handle("/api/v1/timetable/bulk", TimetableHandler{nil, BulkTimetableInsert})
}
