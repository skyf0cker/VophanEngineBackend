package main

import (
	Router "VophanEngineBackend/router/v1"
	"net/http"
)

func main()  {

	router := Router.InitRouter()

	s := &http.Server{
		//Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Addr:           ":3000",
		Handler:        router,
		//ReadTimeout:    setting.ServerSetting.ReadTimeout,
		//WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
