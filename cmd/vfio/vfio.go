/*
 * Copyright (c) 2017 Martin Polednik
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

import (
	"flag"
	"os"

	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
	"github.com/wbpascal/kubevirt-device-plugins/pkg/pci"
)

func main() {
	flag.Parse()

	// Let's start by making sure vfio_pci module is loaded. Without that, binds/unbinds will fail.
	// Small caveat: the loaded module is called vfio_pci, but when it's being probed the name to use is vfio-pci!
	if !pci.IsModuleLoaded("vfio_pci") {
		err := pci.LoadModule("vfio-pci")
		// If we were not able to load the module, we're out of luck.
		if err != nil {
			os.Exit(1)
		}
	}

	manager := dpm.NewManager(pci.PCILister{})
	manager.Run()
}
