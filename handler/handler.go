package handler

import (
	"encoding/json"
	"log"
	"magicball/intent"
	"magicball/protocol"
	"net/http"
)

// ServeHTTP handles CEK requests
func Dispatch(w http.ResponseWriter, r *http.Request) {

	var req protocol.CEKRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("JSON decoding failed")
		respondError(w, "서버와의 연결이 원활하지 않네요")
		return
	}

	reqType := req.Request.Type

	var response protocol.CEKResponse

	switch reqType {
	case "LaunchRequest":
		response = protocol.MakeCEKResponse(handleLaunchRequest())
	case "SessionEndedRequest":
		response = protocol.MakeCEKResponse(handleEndRequest())

	case "IntentRequest":
		if result, err := intent.GetAnswer(); err == nil {
			response = protocol.MakeCEKResponse(result)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(&response)
	w.Write(b)
}

func handleLaunchRequest() protocol.CEKResponsePayload {
	return protocol.CEKResponsePayload{
		OutputSpeech:     protocol.MakeSimpleOutputSpeech("안녕하세요? 마법구슬이에요. 무엇이든 저에게 물어보세요."),
		ShouldEndSession: false,
	}
}

func handleEndRequest() protocol.CEKResponsePayload {
	return protocol.CEKResponsePayload{
		OutputSpeech:     protocol.MakeSimpleOutputSpeech(""),
		ShouldEndSession: true,
	}
}

func respondError(w http.ResponseWriter, msg string) {
	response := protocol.MakeCEKResponse(
		protocol.CEKResponsePayload{
			OutputSpeech: protocol.MakeSimpleOutputSpeech(msg),
		})

	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(&response)
	w.Write(b)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {}
