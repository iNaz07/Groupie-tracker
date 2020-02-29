package locations

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		res, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations" + r.URL.Query().Get("id"))
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		var locations Locations
		if err := json.Unmarshal(body, &locations); err != nil {
			panic(err)
		}

		respBytes := new(bytes.Buffer)
		json.NewEncoder(respBytes).Encode(locations)
		w.Write(respBytes.Bytes())
	}
}
