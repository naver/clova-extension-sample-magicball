package intent

import (
	"magicball/protocol"
	"math/rand"
	"time"
)

var answers = []string{
	"그거 확실하게 될것같아", "그건 분명히 되겠어", "의심할 여지가 없어", "맞아! 확실해", "믿어도 좋겠어",
	"내가 보기엔 맞는것 같아", "거의 확실해보여", "전망이 좋아", "응, 그래", "괜찮아 보이네",
	"뭔가 흐릿흐릿하게 보이네", "지금은 잘 안보여, 나중에 다시 물어봐줘", "음...지금은 말하지 않는편이 나을것 같아", "지금은 잘 모르겠어", "정신을 집중하고 다시 물어봐줘",
	"꿈도 꾸지말게나", "아니라고 말해주겠네", "내가 가진 정보로는 별로야", "전망이 그리 좋진 않아", "뭔가 좀 의심스러워",
}

func GetAnswer() (protocol.CEKResponsePayload, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	randomIndex := rand.Intn(len(answers))

	responsePayload := protocol.CEKResponsePayload{
		OutputSpeech: protocol.MakeOutputSpeechList(
			protocol.Value{
				Lang:  "",
				Value: "https://ssl.pstatic.net/static/clova/service/native_extensions/magicball/magic_ball_sound.mp3",
				Type:  "URL",
			},
			protocol.Value{
				Lang:  "ko",
				Value: "마법구슬이 \"" + answers[randomIndex] + "\"라고 말합니다.",
				Type:  "PlainText",
			},
		),
		ShouldEndSession: true,
	}

	return responsePayload, nil
}
