package main

import (
	"os"
	"os/signal"
	"syscall"

	many_many "codeflow/template/producerConsumer/many-many"
	many_one "codeflow/template/producerConsumer/many-one"
	one_many "codeflow/template/producerConsumer/one-many"
	one_one "codeflow/template/producerConsumer/one-one"
	"codeflow/template/producerConsumer/out"
)

func main() {
	o := out.NewOut()
	go o.OutPut()

	one_one.Exec()
	one_many.Exec()
	many_one.Exec()
	many_many.Exec()

	// out.Println("temp")

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
