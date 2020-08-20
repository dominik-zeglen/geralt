package api

import (
	"encoding/json"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type ReplyRequest struct {
	Sentence string `json:"sentence"`
}
type ReplyResponse struct {
	Reply string `json:"reply"`
}

func (api *API) handleReply(w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(
		r.Context(),
		"handler-reply",
	)
	defer span.Finish()

	var data ReplyRequest

	reqDecodeErr := json.NewDecoder(r.Body).Decode(&data)
	if reqDecodeErr != nil {
		http.Error(w, reqDecodeErr.Error(), http.StatusBadRequest)
		return
	}

	span.LogFields(
		log.String("sentence", data.Sentence),
	)

	reply := ReplyResponse{
		Reply: api.geralt.Reply(ctx, data.Sentence),
	}

	jsonResponse, resEncodeErr := json.Marshal(&reply)
	if resEncodeErr != nil {
		http.Error(w, resEncodeErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
