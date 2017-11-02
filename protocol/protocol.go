package protocol

type CEKIntent struct {
	IntentType string             `json:"intent"` //RULE 'json"intent' would be deprecated(keep this for compatibility for a while). instead of, use 'json:"name"'.
	Name       string             `json:"name"`   //RULE should be set same as IntentType value.
	Slots      map[string]CEKSlot `json:"slots"`
}

type CEKSlot struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type RequestCommon struct {
	Type      string `json:"type"`
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Locale    string `json:"locale"`
}

type CEKSession struct {
	New       bool   `json:"new"`
	SessionId string `json:"sessionId`
	User      struct {
		AccessToken string `json:"accessToken"`
		UserId      string `json:"userId`
	}
	SessionAttributes interface{} `json:"sessionAttributes"`
}

type CEKRequestPayload struct {
	RequestCommon
	Intent CEKIntent   `json:"intent,omitempty"`
	Event  interface{} `json:"event,omitempty"`
}

type CEKRequest struct {
	Version  string                 `json:"version"`
	Session  CEKSession             `json:"session"`
	Contexts map[string]interface{} `json:"context"`
	Request  CEKRequestPayload      `json:"request"`
}

type Value struct {
	Lang  string `json:"lang"`
	Type  string `json:"type"` // Will be deprecated
	Value string `json:"value"`
}

type OutputSpeech struct {
	Type   string      `json:"type"`
	Values interface{} `json:"values"`
}

type CEKResponsePayload struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	Card             interface{}  `json:"card,omitempty"`
	Directives       interface{}  `json:"directives"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

type CEKResponse struct {
	Version           string             `json:"version"`
	SessionAttributes interface{}        `json:"sessionAttributes"`
	Response          CEKResponsePayload `json:"response"`
}

type Card struct {
	ActionList    []CardValue `json:"actionList"`
	BgUrl         CardValue   `json:"bgUrl`
	HighlightText CardValue   `json:"highlightText"`
	MainText      CardValue   `json:"mainText"`
	ParagraphText CardValue   `json:"paragraphText"`
	ReferenceText CardValue   `json:"referenceText"`
	ReferenceUrl  CardValue   `json:"referenceUrl"`
	SentenceText  CardValue   `json:"sentenceText"`
	SubText       CardValue   `json:"subText"`
	TableList     []CardValue `json:"tableList"`
	emotionCode   CardValue   `json:"emotionCode"`
	motionCode    CardValue   `json:"motionCode"`
	Type          string      `json:"type"`
}

type CardValue struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
