// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/quota_manager"
	"github.com/pivotal-golang/lager"
)

type FakeBaseSizer struct {
	BaseSizeStub        func(logger lager.Logger, rootfsPath string) (uint64, error)
	baseSizeMutex       sync.RWMutex
	baseSizeArgsForCall []struct {
		logger     lager.Logger
		rootfsPath string
	}
	baseSizeReturns struct {
		result1 uint64
		result2 error
	}
}

func (fake *FakeBaseSizer) BaseSize(logger lager.Logger, rootfsPath string) (uint64, error) {
	fake.baseSizeMutex.Lock()
	fake.baseSizeArgsForCall = append(fake.baseSizeArgsForCall, struct {
		logger     lager.Logger
		rootfsPath string
	}{logger, rootfsPath})
	fake.baseSizeMutex.Unlock()
	if fake.BaseSizeStub != nil {
		return fake.BaseSizeStub(logger, rootfsPath)
	} else {
		return fake.baseSizeReturns.result1, fake.baseSizeReturns.result2
	}
}

func (fake *FakeBaseSizer) BaseSizeCallCount() int {
	fake.baseSizeMutex.RLock()
	defer fake.baseSizeMutex.RUnlock()
	return len(fake.baseSizeArgsForCall)
}

func (fake *FakeBaseSizer) BaseSizeArgsForCall(i int) (lager.Logger, string) {
	fake.baseSizeMutex.RLock()
	defer fake.baseSizeMutex.RUnlock()
	return fake.baseSizeArgsForCall[i].logger, fake.baseSizeArgsForCall[i].rootfsPath
}

func (fake *FakeBaseSizer) BaseSizeReturns(result1 uint64, result2 error) {
	fake.BaseSizeStub = nil
	fake.baseSizeReturns = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

var _ quota_manager.BaseSizer = new(FakeBaseSizer)
