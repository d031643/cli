// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeUserProvidedServiceInstanceRepository struct {
	CreateStub        func(string, string, string, map[string]interface{}, []string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 map[string]interface{}
		arg5 []string
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	GetSummariesStub        func() (models.UserProvidedServiceSummary, error)
	getSummariesMutex       sync.RWMutex
	getSummariesArgsForCall []struct {
	}
	getSummariesReturns struct {
		result1 models.UserProvidedServiceSummary
		result2 error
	}
	getSummariesReturnsOnCall map[int]struct {
		result1 models.UserProvidedServiceSummary
		result2 error
	}
	UpdateStub        func(models.ServiceInstanceFields) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 models.ServiceInstanceFields
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserProvidedServiceInstanceRepository) Create(arg1 string, arg2 string, arg3 string, arg4 map[string]interface{}, arg5 []string) error {
	var arg5Copy []string
	if arg5 != nil {
		arg5Copy = make([]string, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 map[string]interface{}
		arg5 []string
	}{arg1, arg2, arg3, arg4, arg5Copy})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4, arg5Copy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *FakeUserProvidedServiceInstanceRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeUserProvidedServiceInstanceRepository) CreateCalls(stub func(string, string, string, map[string]interface{}, []string) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeUserProvidedServiceInstanceRepository) CreateArgsForCall(i int) (string, string, string, map[string]interface{}, []string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeUserProvidedServiceInstanceRepository) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserProvidedServiceInstanceRepository) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserProvidedServiceInstanceRepository) GetSummaries() (models.UserProvidedServiceSummary, error) {
	fake.getSummariesMutex.Lock()
	ret, specificReturn := fake.getSummariesReturnsOnCall[len(fake.getSummariesArgsForCall)]
	fake.getSummariesArgsForCall = append(fake.getSummariesArgsForCall, struct {
	}{})
	fake.recordInvocation("GetSummaries", []interface{}{})
	fake.getSummariesMutex.Unlock()
	if fake.GetSummariesStub != nil {
		return fake.GetSummariesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSummariesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserProvidedServiceInstanceRepository) GetSummariesCallCount() int {
	fake.getSummariesMutex.RLock()
	defer fake.getSummariesMutex.RUnlock()
	return len(fake.getSummariesArgsForCall)
}

func (fake *FakeUserProvidedServiceInstanceRepository) GetSummariesCalls(stub func() (models.UserProvidedServiceSummary, error)) {
	fake.getSummariesMutex.Lock()
	defer fake.getSummariesMutex.Unlock()
	fake.GetSummariesStub = stub
}

func (fake *FakeUserProvidedServiceInstanceRepository) GetSummariesReturns(result1 models.UserProvidedServiceSummary, result2 error) {
	fake.getSummariesMutex.Lock()
	defer fake.getSummariesMutex.Unlock()
	fake.GetSummariesStub = nil
	fake.getSummariesReturns = struct {
		result1 models.UserProvidedServiceSummary
		result2 error
	}{result1, result2}
}

func (fake *FakeUserProvidedServiceInstanceRepository) GetSummariesReturnsOnCall(i int, result1 models.UserProvidedServiceSummary, result2 error) {
	fake.getSummariesMutex.Lock()
	defer fake.getSummariesMutex.Unlock()
	fake.GetSummariesStub = nil
	if fake.getSummariesReturnsOnCall == nil {
		fake.getSummariesReturnsOnCall = make(map[int]struct {
			result1 models.UserProvidedServiceSummary
			result2 error
		})
	}
	fake.getSummariesReturnsOnCall[i] = struct {
		result1 models.UserProvidedServiceSummary
		result2 error
	}{result1, result2}
}

func (fake *FakeUserProvidedServiceInstanceRepository) Update(arg1 models.ServiceInstanceFields) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 models.ServiceInstanceFields
	}{arg1})
	fake.recordInvocation("Update", []interface{}{arg1})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1
}

func (fake *FakeUserProvidedServiceInstanceRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeUserProvidedServiceInstanceRepository) UpdateCalls(stub func(models.ServiceInstanceFields) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeUserProvidedServiceInstanceRepository) UpdateArgsForCall(i int) models.ServiceInstanceFields {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUserProvidedServiceInstanceRepository) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserProvidedServiceInstanceRepository) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserProvidedServiceInstanceRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.getSummariesMutex.RLock()
	defer fake.getSummariesMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserProvidedServiceInstanceRepository) recordInvocation(key string, args []interface{}) {
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

var _ api.UserProvidedServiceInstanceRepository = new(FakeUserProvidedServiceInstanceRepository)
