package main

import (
	"os"
	"os/signal"
	"syscall"

	many_many "code_for_review/goroutine/producer-consumer/many-many"
	"code_for_review/goroutine/producer-consumer/out"
)

func main() {
	o := out.NewOut()
	go o.OutPut()

	// one_one.Exec()
	// one_many.Exec()
	// many_one.Exec()
	many_many.Exec()

	// out.Println("temp")

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
