// Code generated by counterfeiter. DO NOT EDIT.
package requirementsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/models"
	"code.cloudfoundry.org/cli/cf/requirements"
)

type FakeApplicationRequirement struct {
	ExecuteStub        func() error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
	}
	executeReturns struct {
		result1 error
	}
	executeReturnsOnCall map[int]struct {
		result1 error
	}
	GetApplicationStub        func() models.Application
	getApplicationMutex       sync.RWMutex
	getApplicationArgsForCall []struct {
	}
	getApplicationReturns struct {
		result1 models.Application
	}
	getApplicationReturnsOnCall map[int]struct {
		result1 models.Application
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeApplicationRequirement) Execute() error {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
	}{})
	fake.recordInvocation("Execute", []interface{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.executeReturns
	return fakeReturns.result1
}

func (fake *FakeApplicationRequirement) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeApplicationRequirement) ExecuteCalls(stub func() error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = stub
}

func (fake *FakeApplicationRequirement) ExecuteReturns(result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApplicationRequirement) ExecuteReturnsOnCall(i int, result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApplicationRequirement) GetApplication() models.Application {
	fake.getApplicationMutex.Lock()
	ret, specificReturn := fake.getApplicationReturnsOnCall[len(fake.getApplicationArgsForCall)]
	fake.getApplicationArgsForCall = append(fake.getApplicationArgsForCall, struct {
	}{})
	fake.recordInvocation("GetApplication", []interface{}{})
	fake.getApplicationMutex.Unlock()
	if fake.GetApplicationStub != nil {
		return fake.GetApplicationStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getApplicationReturns
	return fakeReturns.result1
}

func (fake *FakeApplicationRequirement) GetApplicationCallCount() int {
	fake.getApplicationMutex.RLock()
	defer fake.getApplicationMutex.RUnlock()
	return len(fake.getApplicationArgsForCall)
}

func (fake *FakeApplicationRequirement) GetApplicationCalls(stub func() models.Application) {
	fake.getApplicationMutex.Lock()
	defer fake.getApplicationMutex.Unlock()
	fake.GetApplicationStub = stub
}

func (fake *FakeApplicationRequirement) GetApplicationReturns(result1 models.Application) {
	fake.getApplicationMutex.Lock()
	defer fake.getApplicationMutex.Unlock()
	fake.GetApplicationStub = nil
	fake.getApplicationReturns = struct {
		result1 models.Application
	}{result1}
}

func (fake *FakeApplicationRequirement) GetApplicationReturnsOnCall(i int, result1 models.Application) {
	fake.getApplicationMutex.Lock()
	defer fake.getApplicationMutex.Unlock()
	fake.GetApplicationStub = nil
	if fake.getApplicationReturnsOnCall == nil {
		fake.getApplicationReturnsOnCall = make(map[int]struct {
			result1 models.Application
		})
	}
	fake.getApplicationReturnsOnCall[i] = struct {
		result1 models.Application
	}{result1}
}

func (fake *FakeApplicationRequirement) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.getApplicationMutex.RLock()
	defer fake.getApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeApplicationRequirement) recordInvocation(key string, args []interface{}) {
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

var _ requirements.ApplicationRequirement = new(FakeApplicationRequirement)
