// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/docker_drivers/aufs"
)

type FakeRetrier struct {
	RunStub        func(work func() error) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		work func() error
	}
	runReturns struct {
		result1 error
	}
}

func (fake *FakeRetrier) Run(work func() error) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		work func() error
	}{work})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(work)
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeRetrier) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeRetrier) RunArgsForCall(i int) func() error {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].work
}

func (fake *FakeRetrier) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

var _ aufs.Retrier = new(FakeRetrier)
