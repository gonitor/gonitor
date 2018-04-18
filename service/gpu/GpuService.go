package gpu

import (
	"os/exec"

	"github.com/gonitor/gonitor/util"
)

// ServiceGetInfo gets the GPU information.
func ServiceGetInfo() (string, error) {
	output, err := exec.Command("/bin/bash", "-c", "nvidia-smi -q -x").Output()
	return util.ConvertXMLToJSON(string(output)), err
}
