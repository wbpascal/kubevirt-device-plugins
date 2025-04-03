package main

import (
	"flag"

	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
	"github.com/wbpascal/kubevirt-device-plugins/pkg/kvm"
)

func main() {
	flag.Parse()

	manager := dpm.NewManager(kvm.KVMLister{})
	manager.Run()
}
