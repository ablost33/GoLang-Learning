package main

import (
	api_authentication "github.com/ablost33/GoLang-Learning/random_exercises/api-authentication"
)

func main() {
	in := api_authentication.GetCatRequest{}
	api_authentication.GetCatImages(in)
}
