// Copyright UCloud. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package ipamd

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/ucloud/uk8s-cni-vpc/pkg/deviceplugin"

	"github.com/ucloud/ucloud-sdk-go/ucloud/metadata"
)

// UNI limit equals to vCPU number
func getNodeUNILimits() int {
	//	http://100.80.80.80/meta-data/v1/uhost/cpu
	mdcli := metadata.NewClient()
	cpu, err := mdcli.GetMetadata("/uhost/cpu")
	if err != nil {
		return runtime.NumCPU()
	} else {
		cpuNo, err := strconv.Atoi(cpu)
		if err != nil {
			return runtime.NumCPU()
		} else {
			return cpuNo
		}
	}
}

func startDevicePlugin() error {
	// Init deviceplugin daemon for UNI
	s := deviceplugin.NewUNIDevicePlugin(getNodeUNILimits())
	err := s.Serve(deviceplugin.ResourceName)
	if err != nil {
		return fmt.Errorf("failed to set deviceplugin on node, %v", err)
	}
	return nil
}
