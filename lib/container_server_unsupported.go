// +build !linux

package lib

import (
	"github.com/kubernetes-sigs/cri-o/lib/sandbox"
	"github.com/kubernetes-sigs/cri-o/oci"
	"github.com/pkg/errors"
)

func (c *ContainerServer) addSandboxPlatform(sb *sandbox.Sandbox) {
	// nothin' doin'
}

func (c *ContainerServer) removeSandboxPlatform(sb *sandbox.Sandbox) {
	// nothin' doin'
}

func (c *ContainerServer) getContainerStats(ctr *oci.Container, previousStats *ContainerStats) (*ContainerStats, error) {
	// nothin' doin'
	return nil, errors.New("container stats not supported")
}
