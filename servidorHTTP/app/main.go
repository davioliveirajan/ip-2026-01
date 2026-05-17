package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"servidorHTTP/servidorHTTP/app/handlers"
	"servidorHTTP/servidorHTTP/app/utils"
)

func main() {
	utils.ConnectToDB()
	if err := utils.EnsurePacientesTable(); err != nil {
		log.Fatalf("Erro ao preparar tabela de pacientes: %v", err)
	}

	staticDir := "servidorHTTP/static"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		staticDir = "static"
	}
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		staticDir = "../static"
	}

	fileserver := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fileserver)
	http.HandleFunc("/pacientes", handlers.PacientesHandler)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	var localIP string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			localIP = ipNet.IP.String()
			break
		}
	}

	port := "3000"

	if localIP == "" {
		localIP = "127.0.0.1"
	}

	fmt.Printf("Servidor rodando em: http://%s:%s/\n", localIP, port)

	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		log.Fatal(err)
	}
}
