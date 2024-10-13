package server

import "github.com/GlebSr/app/internal/app/storage/simple"

func Start() error {
	storage := simple.CreateStorage()
	serv := NewServer(&storage)
	return serv.Listen(":8080")
}
