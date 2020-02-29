package relations

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		res, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + r.URL.Query().Get("id"))
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		var relations Relation
		if err := json.Unmarshal(body, &relations); err != nil {
			panic(err)
		}

		respBytes := new(bytes.Buffer)
		json.NewEncoder(respBytes).Encode(relations)
		w.Write(respBytes.Bytes())
	}
}
