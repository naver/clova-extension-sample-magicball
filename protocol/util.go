package protocol

// MakeCEKResponse creates CEKResponse instance with given params
func MakeCEKResponse(responsePayload CEKResponsePayload) CEKResponse {
	response := CEKResponse{
		Version:  "0.1.0",
		Response: responsePayload,
	}

	return response
}

// MakeOutputSpeech creates OutputSpeech instance with given params
func MakeSimpleOutputSpeech(msg string) OutputSpeech {
	return OutputSpeech{
		Type: "SimpleSpeech",
		Values: Value{
			Lang:  "ko",
			Value: msg,
			Type:  "PlainText",
		},
	}
}

// MakeOutputSpeech creates OutputSpeech instance with given params
func MakeOutputSpeechList(value ...Value) OutputSpeech {
	return OutputSpeech{
		Type:   "SpeechList",
		Values: value,
	}
}
