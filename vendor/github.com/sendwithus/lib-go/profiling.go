package swu

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/e-dard/netbug"
	"github.com/sendwithus/lib-go/profiling"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func EnableProfiling(appName string) {
	profilingUrl := GetEnvVariable("SWU_SERVICE_PROFILING_URL", false)
	if len(profilingUrl) == 0 {
		internalLogger.Info("Unable to initialize profiling, please set SWU_SERVICE_PROFILING_URL")
		return
	}

	// check for required env vars
	req := GetEnvVariable("GO_INSTALL_TOOLS_IN_IMAGE", false)
	if len(req) == 0 {
		internalLogger.Info("Unable to initialize profiling, please set GO_INSTALL_TOOLS_IN_IMAGE")
		return
	}
	req = GetEnvVariable("GO_SETUP_GOPATH_IN_IMAGE", false)
	if len(req) == 0 {
		internalLogger.Info("Unable to initialize profiling, please set GO_SETUP_GOPATH_IN_IMAGE")
		return
	}
	go func() {
		RunCommand("go", "get", "github.com/uber/go-torch")
		internalLogger.Info("Downloaded go-torch")
	}()

	go func() {
		internalLogger.Info("Download flamegraph scripts...")
		if ok, _ := exists("Flamegraph"); !ok {
			RunCommand("git", "clone", "https://github.com/brendangregg/FlameGraph")
		}
		goPath := GetEnvVariable("GOPATH", true)
		RunCommand("cp", "FlameGraph/flamegraph.pl", fmt.Sprintf("%v/bin/", goPath))
		internalLogger.Info("Downloaded flamegraph scripts")
	}()
	internalLogger.Info("Registering with profiling server")
	dyno := GetEnvVariable("DYNO", true)
	addr, err := net.ResolveTCPAddr("tcp4", ":0")
	if err != nil {
		internalLogger.ErrorWithError(err, "Unable to resolve local ip address")
		return
	}

	ip := GetEnvVariable("HEROKU_PRIVATE_IP", false)
	if len(ip) == 0 {
		internalLogger.Error("Missing heroku private ip env variable")
		return
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		internalLogger.ErrorWithError(err, "Unable to listen on addr: %v", addr.String())
		return
	}

	port := l.Addr().(*net.TCPAddr).Port
	url := fmt.Sprintf("%v:%v", ip, port)

	go notifyServiceProfiling(appName, dyno, url)
	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for range ticker.C {
			go notifyServiceProfiling(appName, dyno, url)
		}
	}()
	internalLogger.Info("Setting up listeners for profiling requests")
	r := http.NewServeMux()
	r.HandleFunc("/flame", func(w http.ResponseWriter, r *http.Request) {
		_, errorBuffer := RunCommand("go", "tool", "pprof", "-raw", "-seconds=25", fmt.Sprintf("http://%v:%v/profiling/profile", ip, port))
		// error buffer contains where the file was saved
		lines := strings.Split(errorBuffer.String(), "\n")
		line := lines[2]
		words := strings.Split(line, " ")
		path := words[3]
		RunCommand("go-torch", path)
		f, err := os.Open("torch.svg")
		if err != nil {
			internalLogger.ErrorWithError(err, "Error while opening flamegraph to return")
		}
		io.Copy(w, f)
	})
	netbug.RegisterHandler("/profiling/", r)

	internalLogger.Info("TCP Listener created, listening on port %v", port)
	if err := http.Serve(l, r); err != nil {
		internalLogger.ErrorWithError(err, "Unable to listen for profiling information")
	}
}

func notifyServiceProfiling(app, dyno, url string) {
	profilingUrl := GetEnvVariable("SWU_SERVICE_PROFILING_URL", false)
	conn, err := grpc.Dial(profilingUrl, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		internalLogger.ErrorWithError(err, "Error while notifying service profiling")
	}
	client := profiling.NewServiceProfilingClient(conn)
	_, err = client.Register(context.Background(), &profiling.RegisterRequest{
		Name:    dyno,
		DataUrl: url,
		App:     app,
	})
	if err != nil {
		internalLogger.ErrorWithError(err, "Error while registering with service profiling")
	}
}

func RunCommand(name string, args ...string) (*bytes.Buffer, *bytes.Buffer) {
	cmd := exec.Command(name, args...)
	output := bytes.Buffer{}
	errorBuffer := bytes.Buffer{}
	cmd.Stdout = &output
	cmd.Stderr = &errorBuffer
	err := cmd.Run()
	if err != nil {
		internalLogger.ErrorWithError(err, "Error while running command %v: %v", name, errorBuffer.String())
	}
	return &output, &errorBuffer
}

func ExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
