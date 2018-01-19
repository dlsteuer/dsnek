package swu

import (
	"fmt"
	"google.golang.org/grpc/naming"
	"net"
	"strings"
	"time"
)

type resolver struct {
	log *Logger
}

type watcher struct {
	target       string
	currentHosts []string
	log          *Logger
}

var watcherStruct watcher

func NewResolver(log *Logger) naming.Resolver {
	return &resolver{log: log}
}

func init() {
	watcherStruct = watcher{}
}

func (r *resolver) Resolve(target string) (naming.Watcher, error) {
	watcherStruct.target = target
	watcherStruct.currentHosts = []string{}
	watcherStruct.log = r.log
	return &watcherStruct, nil
}

// Follows the GRPC naming.Resolver spec:
//
// Next blocks until an update or error happens. It may return one or more
// updates. The first call should get the full set of the results. It should
// return an error if and only if Watcher cannot recover.
func (w *watcher) Next() ([]*naming.Update, error) {
	w.log.Info("Watcher.Next() called for updates, current hosts: %v\n", w.currentHosts)
	parts := strings.Split(w.target, ":")
	host := parts[0]
	port := parts[1]

	noChanges := true
	var ips []string
	for noChanges {
		var err error
		ips, err = net.LookupHost(host)
		if err != nil {
			w.log.Error("Error while looking up ip addresses: %v", err)
			ips = []string{}
		}
		for _, ip := range ips {
			if !inArray(ip, w.currentHosts) {
				noChanges = false
				break
			}
		}
		if noChanges {
			time.Sleep(5 * time.Second)
		}
	}

	updates := []*naming.Update{}

	// find additions
	for _, ip := range ips {
		if !inArray(ip, w.currentHosts) {
			updates = append(updates, &naming.Update{
				Op:   naming.Add,
				Addr: fmt.Sprintf("%v:%v", ip, port),
			})
			w.currentHosts = append(w.currentHosts, ip)
		}
	}

	// find deletions
	toRemove := []string{}
	for _, ip := range w.currentHosts {
		if !inArray(ip, ips) {
			updates = append(updates, &naming.Update{
				Op:   naming.Delete,
				Addr: fmt.Sprintf("%v:%v", ip, port),
			})
			toRemove = append(toRemove, ip)
		}
	}

	for _, ip := range toRemove {
		w.currentHosts = removeFromSlice(ip, w.currentHosts)
	}

	// display updates
	for _, update := range updates {
		w.log.Info("\tUpdate: %+v\n", *update)
	}
	return updates, nil
}

func (w *watcher) Close() {}

func inArray(item string, array []string) bool {
	for _, i := range array {
		if i == item {
			return true
		}
	}
	return false
}

func removeFromSlice(item string, items []string) []string {
	for index, i := range items {
		if i != item {
			continue
		}
		items = append(items[:index], items[index+1:]...)
		break
	}
	return items
}
