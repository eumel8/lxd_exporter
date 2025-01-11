package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/eumel8/lxd_exporter/pkg/metrics"
	lxd "github.com/lxc/incus/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	version = "staging-UNVERSIONED"

	port = kingpin.Flag(
		"port", "Provide the port to listen on").Default("9472").Int16()
	listen = kingpin.Flag(
		"listen", "Provide the interface to listen on").Default("0.0.0.0").String()
	socket = kingpin.Flag(
		"socket", "Provide the socket to listen on").Default("/var/snap/lxd/common/lxd/unix.socket").String()
)

func main() {
	logger := log.New(os.Stderr, "lxd_exporter: ", log.LstdFlags)

	kingpin.Version(version)
	kingpin.Parse()

	socketFile := *socket
	server, err := lxd.ConnectIncusUnix(socketFile, &lxd.ConnectionArgs{InsecureSkipVerify: true, SkipGetServer: true})
	if err != nil {
		logger.Fatalf("Unable to contact LXD server: %s", err)
		return
	}

	prometheus.MustRegister(metrics.NewCollector(logger, server))
	http.Handle("/metrics", promhttp.Handler())

	serveAddress := fmt.Sprintf("%s:%d", *listen, *port)
	log.Print("Server listening on ", serveAddress)
	log.Fatal(http.ListenAndServe(serveAddress, nil))
}
