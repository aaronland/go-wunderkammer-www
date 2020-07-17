package www

import (
	"encoding/json"
	"encoding/xml"
	_ "errors"
	"github.com/aaronland/go-wunderkammer/oembed"
	_ "log"
	"net/http"
)

func NewOEmbedHandler(db oembed.OEmbedDatabase) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		q := req.URL.Query()

		url := q.Get("url")

		if url == "" {
			http.Error(rsp, "Missing url", http.StatusBadRequest)
			return
		}

		format := q.Get("format")

		switch format {
		case "", "json":
			// pass
		case "xml":
			http.Error(rsp, "Not implemented", http.StatusBadRequest)
		default:
			http.Error(rsp, "Invalid format", http.StatusBadRequest)
			return
		}

		ph, err := db.GetOEmbedWithURL(ctx, url)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		switch format {
		case "", "json":

			rsp.Header().Set("Content-type", "application/json")

			enc := json.NewEncoder(rsp)
			err = enc.Encode(ph)

			if err != nil {
				http.Error(rsp, err.Error(), http.StatusInternalServerError)
				return
			}

		case "xml":

			rsp.Header().Set("Content-type", "text/xml")

			enc := xml.NewEncoder(rsp)
			err = enc.Encode(ph)

			if err != nil {
				http.Error(rsp, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			// pass
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
