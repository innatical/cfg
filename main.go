package main

func main() {
	loadConfig()

	if err := Execute(); err != nil {
		panic(err)
	}
}
