package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/pure"
	"github.com/parnurzeal/gorequest"
	//mw "github.com/go-playground/pure/examples/middleware/logging-recovery"
)

func main() {

	p := pure.New()
	//p.Use(mw.LoggingAndRecovery(true))

	p.Get("/", stt_models)

	http.ListenAndServe(":3007", p.Serve())
}

func stt_models(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello World"))

	//https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779

	/*
	  "url": "https://stream.watsonplatform.net/speech-to-text/api",
	  "password": "provided by ibm bluemix",
	  "username": "provided by ibm bluemix"
	*/

	request := gorequest.New().Timeout(5 * time.Second)
	request.SetBasicAuth("username", "password")

	url := "https://stream.watsonplatform.net/speech-to-text/api/v1/models"

	resp, body, err := request.Get(url).End()

	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
		fmt.Fprintf(w, body)
	}
}

/*
Output body

{
   "models": [
      {
         "name": "fr-FR_BroadbandModel",
         "language": "fr-FR",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/fr-FR_BroadbandModel",
         "rate": 16000,
         "supported_features": {
            "custom_language_model": false,
            "speaker_labels": false
         },
         "description": "French broadband model."
      },
      {
         "name": "en-US_NarrowbandModel",
         "language": "en-US",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/en-US_NarrowbandModel",
         "rate": 8000,
         "supported_features": {
            "custom_language_model": true,
            "speaker_labels": true
         },
         "description": "US English narrowband model."
      },
      {
         "name": "pt-BR_BroadbandModel",
         "language": "pt-BR",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/pt-BR_BroadbandModel",
         "rate": 16000,
         "supported_features": {
            "custom_language_model": false,
            "speaker_labels": false
         },
         "description": "Brazilian Portuguese broadband model."
      },
      {
         "name": "ja-JP_NarrowbandModel",
         "language": "ja-JP",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/ja-JP_NarrowbandModel",
         "rate": 8000,
         "supported_features": {
            "custom_language_model": true,
            "speaker_labels": true
         },
         "description": "Japanese narrowband model."
      },
      {
         "name": "zh-CN_BroadbandModel",
         "language": "zh-CN",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/zh-CN_BroadbandModel",
         "rate": 16000,
         "supported_features": {
            "custom_language_model": false,
            "speaker_labels": false
         },
         "description": "Mandarin broadband model."
      },
      {
         "name": "ja-JP_BroadbandModel",
         "language": "ja-JP",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/ja-JP_BroadbandModel",
         "rate": 16000,
         "supported_features": {
            "custom_language_model": true,
            "speaker_labels": false
         },
         "description": "Japanese broadband model."
      },
      {
         "name": "pt-BR_NarrowbandModel",
         "language": "pt-BR",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/pt-BR_NarrowbandModel",
         "rate": 8000,
         "supported_features": {
            "custom_language_model": false,
            "speaker_labels": false
         },
         "description": "Brazilian Portuguese narrowband model."
      },
      {
         "name": "es-ES_BroadbandModel",
         "language": "es-ES",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/es-ES_BroadbandModel",
         "rate": 16000,
         "supported_features": {
            "custom_language_model": false,
            "speaker_labels": false
         },
         "description": "Spanish broadband model."
      },
      {
         "name": "ar-AR_BroadbandModel",
         "language": "ar-AR",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/ar-AR_BroadbandModel",
         "rate": 16000,
         "supported_features": {
            "custom_language_model": false,
            "speaker_labels": false
         },
         "description": "Modern Standard Arabic broadband model."
      },
      {
         "name": "zh-CN_NarrowbandModel",
         "language": "zh-CN",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/zh-CN_NarrowbandModel",
         "rate": 8000,
         "supported_features": {
            "custom_language_model": false,
            "speaker_labels": false
         },
         "description": "Mandarin narrowband model."
      },
      {
         "name": "en-UK_BroadbandModel",
         "language": "en-UK",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/en-UK_BroadbandModel",
         "rate": 16000,
         "supported_features": {
            "custom_language_model": false,
            "speaker_labels": false
         },
         "description": "UK English broadband model."
      },
      {
         "name": "es-ES_NarrowbandModel",
         "language": "es-ES",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/es-ES_NarrowbandModel",
         "rate": 8000,
         "supported_features": {
            "custom_language_model": false,
            "speaker_labels": true
         },
         "description": "Spanish narrowband model."
      },
      {
         "name": "en-US_BroadbandModel",
         "language": "en-US",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/en-US_BroadbandModel",
         "rate": 16000,
         "supported_features": {
            "custom_language_model": true,
            "speaker_labels": false
         },
         "description": "US English broadband model."
      },
      {
         "name": "en-UK_NarrowbandModel",
         "language": "en-UK",
         "url": "https://stream.watsonplatform.net/speech-to-text/api/v1/models/en-UK_NarrowbandModel",
         "rate": 8000,
         "supported_features": {
            "custom_language_model": false,
            "speaker_labels": false
         },
         "description": "UK English narrowband model."
      }
   ]
}
*/
