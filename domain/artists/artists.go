package artists

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")

		res, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		var artists []Artist
		if err := json.Unmarshal(body, &artists); err != nil {
			panic(err)
		}

		if id != "" {

			for _, artist := range artists {
				if strconv.Itoa(artist.Id) == id {

					respBytes := new(bytes.Buffer)
					json.NewEncoder(respBytes).Encode(artist)
					w.Write(respBytes.Bytes())

					break
				}
			}

		} else {

			respBytes := new(bytes.Buffer)
			json.NewEncoder(respBytes).Encode(artists)
			w.Write(respBytes.Bytes())
		}
	}
}
