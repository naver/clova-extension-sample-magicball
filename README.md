### 개요 
'마법구슬'이라는 Clova extension의 REST API 서버 소스 코드입니다. '마법구슬' 익스텐션이 어떻게 작동하는지 보시려면, Clova 앱이나 Clova 스피커(WAVE, Friens)에서 '마법구슬 시작해줘'라고 말해보시길 바랍니다. 해당 익스텐션을 실행하면 마법구슬이 랜덤하게 20가지 대답 중 한 가지 값을 알려줍니다. 

### 사용환경
'마법구슬' Clova extension의 REST API 서버는 Go로 구현되어 있습니다.  Windows, MacOS, Linux 등 golang이 구동 가능한 OS면 실행 가능하며, 구체적인 목록들은 여기를 참고하셔서, 아래 가이드에 따라 Go를 먼저 설치하시길 바랍니다. https://github.com/golang-kr/golang-doc/wiki/%EC%84%A4%EC%B9%98-%EC%8B%9C%EC%9E%91%ED%95%98%EA%B8%B0 

### 설치방법
'마법구슬' REST API 서버 소스 코드 설치는 다음과 같이 해주시길 바랍니다.
1) Go 배포판 설치: https://github.com/golang-kr/golang-doc/wiki/%EC%84%A4%EC%B9%98-%EC%8B%9C%EC%9E%91%ED%95%98%EA%B8%B0 
2) 소스코드 다운로드:  # git clone https://github.com/naver/clova-extension-sample-magicball.git 
3) 소스코드 빌드: # make  
소스코드가 정상 빌드되면 bin 디렉토리 밑에 magicball이라는 실행파일이 생성됩니다. 

### 사용법 
'마법구슬' Clova extension의 REST API 서버는 Clova platform으로부터의 익스텐션 요청에 따라 20가지의 랜덤한 응답을 하도록 되어 있습니다. API 서버를 실행을 하더라도, Clova platform이 보내는 것과 동일한 API 요청을 해주셔야 정확하게 작동하는 점 참고 바랍니다. 실제 서비스를 위해서는 https 기반으로 외부에서 접근 가능한 도메인으로 해주셔야 합니다.
- API 서버 실행: bin/magicball 
- API 서버 테스팅: [Postman](https://www.getpostman.com/apps)에서 아래와 같이 json Request를 전송하고 json이 리턴되는지 테스트 해봅니다.
	- URL: http://localhost:10780/magicball 
	- 요청 방법: POST 
	- Body: raw ( JSON 선택 ) 
- 요청 예시)
```
{
    "version": "0.1.0",
    "session": {
        "sessionId": "867723c7-566a-45f9-8a0b-fc832e548bb8",
        "user": {
            "userId": "JOTur5fxSEuD4O5dbFO3Aw",
            "accessToken": "a510ccb5-72ad-42e7-934d-655139be1a69"
        },
        "new": true
    },
    "context": {
        "System": {
            "user": {
                "userId": "JOTur5fxSEuD4O5dbFO3Aw",
                "accessToken": "a510ccb5-72ad-42e7-934d-655139be1a69"
            },
            "device": {
                "deviceId": "0dcc90f8-4f3d-4ea6-ae71-9f990d43d4e5"
            }
        }
    },
    "request": {
        "type": "IntentRequest",
        "intent": {
            "name": "Clova.GuideIntent",
            "slots": null
        }
    }
}
```

![image.png](http://static.naver.net/clova/service/native_extensions/example/magicball.png)

### 라이선스
Naver & Line corp.

[LICENSE](https://github.com/naver/clova-extension-sample-magicball/blob/github-public/LICENSE)

```
Copyright 2018 NAVER Corp. & LINE Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```


