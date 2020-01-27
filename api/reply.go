package api

import (
	"encoding/json"
	"net/http"
)

type replyRequest struct {
	Sentence string `json:"sentence"`
}
type replyResponse struct {
	Reply string `json:"reply"`
}

func (api *API) handleReply(w http.ResponseWriter, r *http.Request) {
	var data replyRequest

	reqDecodeErr := json.NewDecoder(r.Body).Decode(&data)
	if reqDecodeErr != nil {
		http.Error(w, reqDecodeErr.Error(), http.StatusBadRequest)
		return
	}

	reply := replyResponse{
		Reply: api.geralt.Reply(r.Context(), data.Sentence),
	}

	jsonResponse, resEncodeErr := json.Marshal(&reply)
	if resEncodeErr != nil {
		http.Error(w, resEncodeErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
