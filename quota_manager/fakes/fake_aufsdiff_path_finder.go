// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/quota_manager"
)

type FakeAUFSDiffPathFinder struct {
	GetDiffLayerPathStub        func(rootFSPath string) string
	getDiffLayerPathMutex       sync.RWMutex
	getDiffLayerPathArgsForCall []struct {
		rootFSPath string
	}
	getDiffLayerPathReturns struct {
		result1 string
	}
}

func (fake *FakeAUFSDiffPathFinder) GetDiffLayerPath(rootFSPath string) string {
	fake.getDiffLayerPathMutex.Lock()
	fake.getDiffLayerPathArgsForCall = append(fake.getDiffLayerPathArgsForCall, struct {
		rootFSPath string
	}{rootFSPath})
	fake.getDiffLayerPathMutex.Unlock()
	if fake.GetDiffLayerPathStub != nil {
		return fake.GetDiffLayerPathStub(rootFSPath)
	} else {
		return fake.getDiffLayerPathReturns.result1
	}
}

func (fake *FakeAUFSDiffPathFinder) GetDiffLayerPathCallCount() int {
	fake.getDiffLayerPathMutex.RLock()
	defer fake.getDiffLayerPathMutex.RUnlock()
	return len(fake.getDiffLayerPathArgsForCall)
}

func (fake *FakeAUFSDiffPathFinder) GetDiffLayerPathArgsForCall(i int) string {
	fake.getDiffLayerPathMutex.RLock()
	defer fake.getDiffLayerPathMutex.RUnlock()
	return fake.getDiffLayerPathArgsForCall[i].rootFSPath
}

func (fake *FakeAUFSDiffPathFinder) GetDiffLayerPathReturns(result1 string) {
	fake.GetDiffLayerPathStub = nil
	fake.getDiffLayerPathReturns = struct {
		result1 string
	}{result1}
}

var _ quota_manager.AUFSDiffPathFinder = new(FakeAUFSDiffPathFinder)
