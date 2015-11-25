// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/docker_drivers/aufs"
	"github.com/docker/docker/pkg/archive"
)

type FakeGraphDriver struct {
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct{}
	stringReturns struct {
		result1 string
	}
	CreateStub        func(id, parent string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		id     string
		parent string
	}
	createReturns struct {
		result1 error
	}
	RemoveStub        func(id string) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		id string
	}
	removeReturns struct {
		result1 error
	}
	GetStub        func(id, mountLabel string) (dir string, err error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		id         string
		mountLabel string
	}
	getReturns struct {
		result1 string
		result2 error
	}
	PutStub        func(id string) error
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		id string
	}
	putReturns struct {
		result1 error
	}
	ExistsStub        func(id string) bool
	existsMutex       sync.RWMutex
	existsArgsForCall []struct {
		id string
	}
	existsReturns struct {
		result1 bool
	}
	StatusStub        func() [][2]string
	statusMutex       sync.RWMutex
	statusArgsForCall []struct{}
	statusReturns struct {
		result1 [][2]string
	}
	GetMetadataStub        func(id string) (map[string]string, error)
	getMetadataMutex       sync.RWMutex
	getMetadataArgsForCall []struct {
		id string
	}
	getMetadataReturns struct {
		result1 map[string]string
		result2 error
	}
	CleanupStub        func() error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct{}
	cleanupReturns struct {
		result1 error
	}
	DiffStub        func(id, parent string) (archive.Archive, error)
	diffMutex       sync.RWMutex
	diffArgsForCall []struct {
		id     string
		parent string
	}
	diffReturns struct {
		result1 archive.Archive
		result2 error
	}
	ChangesStub        func(id, parent string) ([]archive.Change, error)
	changesMutex       sync.RWMutex
	changesArgsForCall []struct {
		id     string
		parent string
	}
	changesReturns struct {
		result1 []archive.Change
		result2 error
	}
	ApplyDiffStub        func(id, parent string, diff archive.ArchiveReader) (size int64, err error)
	applyDiffMutex       sync.RWMutex
	applyDiffArgsForCall []struct {
		id     string
		parent string
		diff   archive.ArchiveReader
	}
	applyDiffReturns struct {
		result1 int64
		result2 error
	}
	DiffSizeStub        func(id, parent string) (size int64, err error)
	diffSizeMutex       sync.RWMutex
	diffSizeArgsForCall []struct {
		id     string
		parent string
	}
	diffSizeReturns struct {
		result1 int64
		result2 error
	}
}

func (fake *FakeGraphDriver) String() string {
	fake.stringMutex.Lock()
	fake.stringArgsForCall = append(fake.stringArgsForCall, struct{}{})
	fake.stringMutex.Unlock()
	if fake.StringStub != nil {
		return fake.StringStub()
	} else {
		return fake.stringReturns.result1
	}
}

func (fake *FakeGraphDriver) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *FakeGraphDriver) StringReturns(result1 string) {
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeGraphDriver) Create(id string, parent string) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		id     string
		parent string
	}{id, parent})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(id, parent)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeGraphDriver) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeGraphDriver) CreateArgsForCall(i int) (string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].id, fake.createArgsForCall[i].parent
}

func (fake *FakeGraphDriver) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGraphDriver) Remove(id string) error {
	fake.removeMutex.Lock()
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		id string
	}{id})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(id)
	} else {
		return fake.removeReturns.result1
	}
}

func (fake *FakeGraphDriver) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeGraphDriver) RemoveArgsForCall(i int) string {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].id
}

func (fake *FakeGraphDriver) RemoveReturns(result1 error) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGraphDriver) Get(id string, mountLabel string) (dir string, err error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		id         string
		mountLabel string
	}{id, mountLabel})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(id, mountLabel)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeGraphDriver) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeGraphDriver) GetArgsForCall(i int) (string, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].id, fake.getArgsForCall[i].mountLabel
}

func (fake *FakeGraphDriver) GetReturns(result1 string, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGraphDriver) Put(id string) error {
	fake.putMutex.Lock()
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		id string
	}{id})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(id)
	} else {
		return fake.putReturns.result1
	}
}

func (fake *FakeGraphDriver) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeGraphDriver) PutArgsForCall(i int) string {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].id
}

func (fake *FakeGraphDriver) PutReturns(result1 error) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGraphDriver) Exists(id string) bool {
	fake.existsMutex.Lock()
	fake.existsArgsForCall = append(fake.existsArgsForCall, struct {
		id string
	}{id})
	fake.existsMutex.Unlock()
	if fake.ExistsStub != nil {
		return fake.ExistsStub(id)
	} else {
		return fake.existsReturns.result1
	}
}

func (fake *FakeGraphDriver) ExistsCallCount() int {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return len(fake.existsArgsForCall)
}

func (fake *FakeGraphDriver) ExistsArgsForCall(i int) string {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return fake.existsArgsForCall[i].id
}

func (fake *FakeGraphDriver) ExistsReturns(result1 bool) {
	fake.ExistsStub = nil
	fake.existsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGraphDriver) Status() [][2]string {
	fake.statusMutex.Lock()
	fake.statusArgsForCall = append(fake.statusArgsForCall, struct{}{})
	fake.statusMutex.Unlock()
	if fake.StatusStub != nil {
		return fake.StatusStub()
	} else {
		return fake.statusReturns.result1
	}
}

func (fake *FakeGraphDriver) StatusCallCount() int {
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	return len(fake.statusArgsForCall)
}

func (fake *FakeGraphDriver) StatusReturns(result1 [][2]string) {
	fake.StatusStub = nil
	fake.statusReturns = struct {
		result1 [][2]string
	}{result1}
}

func (fake *FakeGraphDriver) GetMetadata(id string) (map[string]string, error) {
	fake.getMetadataMutex.Lock()
	fake.getMetadataArgsForCall = append(fake.getMetadataArgsForCall, struct {
		id string
	}{id})
	fake.getMetadataMutex.Unlock()
	if fake.GetMetadataStub != nil {
		return fake.GetMetadataStub(id)
	} else {
		return fake.getMetadataReturns.result1, fake.getMetadataReturns.result2
	}
}

func (fake *FakeGraphDriver) GetMetadataCallCount() int {
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	return len(fake.getMetadataArgsForCall)
}

func (fake *FakeGraphDriver) GetMetadataArgsForCall(i int) string {
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	return fake.getMetadataArgsForCall[i].id
}

func (fake *FakeGraphDriver) GetMetadataReturns(result1 map[string]string, result2 error) {
	fake.GetMetadataStub = nil
	fake.getMetadataReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeGraphDriver) Cleanup() error {
	fake.cleanupMutex.Lock()
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct{}{})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub()
	} else {
		return fake.cleanupReturns.result1
	}
}

func (fake *FakeGraphDriver) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeGraphDriver) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGraphDriver) Diff(id string, parent string) (archive.Archive, error) {
	fake.diffMutex.Lock()
	fake.diffArgsForCall = append(fake.diffArgsForCall, struct {
		id     string
		parent string
	}{id, parent})
	fake.diffMutex.Unlock()
	if fake.DiffStub != nil {
		return fake.DiffStub(id, parent)
	} else {
		return fake.diffReturns.result1, fake.diffReturns.result2
	}
}

func (fake *FakeGraphDriver) DiffCallCount() int {
	fake.diffMutex.RLock()
	defer fake.diffMutex.RUnlock()
	return len(fake.diffArgsForCall)
}

func (fake *FakeGraphDriver) DiffArgsForCall(i int) (string, string) {
	fake.diffMutex.RLock()
	defer fake.diffMutex.RUnlock()
	return fake.diffArgsForCall[i].id, fake.diffArgsForCall[i].parent
}

func (fake *FakeGraphDriver) DiffReturns(result1 archive.Archive, result2 error) {
	fake.DiffStub = nil
	fake.diffReturns = struct {
		result1 archive.Archive
		result2 error
	}{result1, result2}
}

func (fake *FakeGraphDriver) Changes(id string, parent string) ([]archive.Change, error) {
	fake.changesMutex.Lock()
	fake.changesArgsForCall = append(fake.changesArgsForCall, struct {
		id     string
		parent string
	}{id, parent})
	fake.changesMutex.Unlock()
	if fake.ChangesStub != nil {
		return fake.ChangesStub(id, parent)
	} else {
		return fake.changesReturns.result1, fake.changesReturns.result2
	}
}

func (fake *FakeGraphDriver) ChangesCallCount() int {
	fake.changesMutex.RLock()
	defer fake.changesMutex.RUnlock()
	return len(fake.changesArgsForCall)
}

func (fake *FakeGraphDriver) ChangesArgsForCall(i int) (string, string) {
	fake.changesMutex.RLock()
	defer fake.changesMutex.RUnlock()
	return fake.changesArgsForCall[i].id, fake.changesArgsForCall[i].parent
}

func (fake *FakeGraphDriver) ChangesReturns(result1 []archive.Change, result2 error) {
	fake.ChangesStub = nil
	fake.changesReturns = struct {
		result1 []archive.Change
		result2 error
	}{result1, result2}
}

func (fake *FakeGraphDriver) ApplyDiff(id string, parent string, diff archive.ArchiveReader) (size int64, err error) {
	fake.applyDiffMutex.Lock()
	fake.applyDiffArgsForCall = append(fake.applyDiffArgsForCall, struct {
		id     string
		parent string
		diff   archive.ArchiveReader
	}{id, parent, diff})
	fake.applyDiffMutex.Unlock()
	if fake.ApplyDiffStub != nil {
		return fake.ApplyDiffStub(id, parent, diff)
	} else {
		return fake.applyDiffReturns.result1, fake.applyDiffReturns.result2
	}
}

func (fake *FakeGraphDriver) ApplyDiffCallCount() int {
	fake.applyDiffMutex.RLock()
	defer fake.applyDiffMutex.RUnlock()
	return len(fake.applyDiffArgsForCall)
}

func (fake *FakeGraphDriver) ApplyDiffArgsForCall(i int) (string, string, archive.ArchiveReader) {
	fake.applyDiffMutex.RLock()
	defer fake.applyDiffMutex.RUnlock()
	return fake.applyDiffArgsForCall[i].id, fake.applyDiffArgsForCall[i].parent, fake.applyDiffArgsForCall[i].diff
}

func (fake *FakeGraphDriver) ApplyDiffReturns(result1 int64, result2 error) {
	fake.ApplyDiffStub = nil
	fake.applyDiffReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeGraphDriver) DiffSize(id string, parent string) (size int64, err error) {
	fake.diffSizeMutex.Lock()
	fake.diffSizeArgsForCall = append(fake.diffSizeArgsForCall, struct {
		id     string
		parent string
	}{id, parent})
	fake.diffSizeMutex.Unlock()
	if fake.DiffSizeStub != nil {
		return fake.DiffSizeStub(id, parent)
	} else {
		return fake.diffSizeReturns.result1, fake.diffSizeReturns.result2
	}
}

func (fake *FakeGraphDriver) DiffSizeCallCount() int {
	fake.diffSizeMutex.RLock()
	defer fake.diffSizeMutex.RUnlock()
	return len(fake.diffSizeArgsForCall)
}

func (fake *FakeGraphDriver) DiffSizeArgsForCall(i int) (string, string) {
	fake.diffSizeMutex.RLock()
	defer fake.diffSizeMutex.RUnlock()
	return fake.diffSizeArgsForCall[i].id, fake.diffSizeArgsForCall[i].parent
}

func (fake *FakeGraphDriver) DiffSizeReturns(result1 int64, result2 error) {
	fake.DiffSizeStub = nil
	fake.diffSizeReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

var _ aufs.GraphDriver = new(FakeGraphDriver)
