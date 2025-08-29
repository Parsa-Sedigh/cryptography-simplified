package main

func main() {
	sig, digest := send()
	receive(sig, digest)
}
