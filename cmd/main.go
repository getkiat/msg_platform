package main

import (
	"msg_platform/server/api"
	"msg_platform/server/config"
)

func main() {

	loadedConfig := config.LoadConfig()

	basicsAPI := api.NewAPI(loadedConfig)

	basicsAPI.Run()

	
	// handler := server.NewRouter()
	// srv := &http.Server{
	// 	Addr:    fmt.Sprintf(":%v", defaultHTTPPort),
	// 	Handler: handler,
	// }
	// go func() {
	// 	srv.
	// 	srv.ListenAndServe()
	// }()

	// // Wait for an interrupt
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// <-c

	// // Attempt a graceful shutdown
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// srv.Shutdown(ctx)
}
