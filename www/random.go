package www

import (
	"encoding/base64"
	"fmt"
	"github.com/aaronland/go-wunderkammer/oembed"
	_ "log"
	"net/http"
	"regexp"
)

func NewRandomObjectHandler(db oembed.OEmbedDatabase) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		ph, err := db.GetRandomOEmbed(ctx)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		// START OF please put this in a function...

		object_uri := ph.ObjectURI
		redir_url := fmt.Sprintf("/object?url=%s", object_uri)

		http.Redirect(rsp, req, redir_url, http.StatusSeeOther)
		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}

func NewRandomImageHandler(db oembed.OEmbedDatabase) (http.Handler, error) {

	re, err := regexp.Compile(`^data:(image/(?:[a-z]+));base64,(.*)$`)

	if err != nil {
		return nil, err
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		ph, err := db.GetRandomOEmbed(ctx)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		// START OF please put this in a function...

		data_url := ph.DataURL

		if data_url == "" {
			http.Redirect(rsp, req, ph.URL, http.StatusSeeOther)
			return
		}

		m := re.FindStringSubmatch(data_url)

		if len(m) != 3 {
			http.Error(rsp, "Invalid data URL", http.StatusInternalServerError)
			return
		}

		content_type := m[1]
		b64_data := m[2]

		raw, err := base64.StdEncoding.DecodeString(b64_data)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-type", content_type)
		rsp.Write(raw)

		// END OF please put this in a function...
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
