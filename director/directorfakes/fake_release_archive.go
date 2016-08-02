// This file was generated by counterfeiter
package directorfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-init/director"
)

type FakeReleaseArchive struct {
	InfoStub        func() (string, string, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct{}
	infoReturns     struct {
		result1 string
		result2 string
		result3 error
	}
	FileStub        func() (director.UploadFile, error)
	fileMutex       sync.RWMutex
	fileArgsForCall []struct{}
	fileReturns     struct {
		result1 director.UploadFile
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseArchive) Info() (string, string, error) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct{}{})
	fake.recordInvocation("Info", []interface{}{})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub()
	} else {
		return fake.infoReturns.result1, fake.infoReturns.result2, fake.infoReturns.result3
	}
}

func (fake *FakeReleaseArchive) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeReleaseArchive) InfoReturns(result1 string, result2 string, result3 error) {
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeReleaseArchive) File() (director.UploadFile, error) {
	fake.fileMutex.Lock()
	fake.fileArgsForCall = append(fake.fileArgsForCall, struct{}{})
	fake.recordInvocation("File", []interface{}{})
	fake.fileMutex.Unlock()
	if fake.FileStub != nil {
		return fake.FileStub()
	} else {
		return fake.fileReturns.result1, fake.fileReturns.result2
	}
}

func (fake *FakeReleaseArchive) FileCallCount() int {
	fake.fileMutex.RLock()
	defer fake.fileMutex.RUnlock()
	return len(fake.fileArgsForCall)
}

func (fake *FakeReleaseArchive) FileReturns(result1 director.UploadFile, result2 error) {
	fake.FileStub = nil
	fake.fileReturns = struct {
		result1 director.UploadFile
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseArchive) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.fileMutex.RLock()
	defer fake.fileMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeReleaseArchive) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ director.ReleaseArchive = new(FakeReleaseArchive)
