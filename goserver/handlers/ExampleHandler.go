package handlers

import (
	"encoding/json"
	"net/http"
	e "goserver/domain"
)

func GetExamples(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		results :=  e.GetExamples()
		err := json.NewEncoder(w).Encode(mapToSlice(results))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// https://web3.coach/golang-how-to-convert-map-to-slice-three-gotchas
func mapToSlice(elementMap map[int]e.Example) []e.Example {

	result := make([]e.Example, 0, len(elementMap))

	for _, el := range elementMap {
		result = append(result, el)
	}

	return result
}