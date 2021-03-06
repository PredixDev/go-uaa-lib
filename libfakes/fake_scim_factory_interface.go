// This file was generated by counterfeiter
package libfakes

import (
	"sync"

	"github.com/PredixDev/go-uaa-lib"
)

type FakeScimFactoryInterface struct {
	NewStub        func(target lib.Target, context lib.Context) lib.Scim
	newMutex       sync.RWMutex
	newArgsForCall []struct {
		target  lib.Target
		context lib.Context
	}
	newReturns struct {
		result1 lib.Scim
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScimFactoryInterface) New(target lib.Target, context lib.Context) lib.Scim {
	fake.newMutex.Lock()
	fake.newArgsForCall = append(fake.newArgsForCall, struct {
		target  lib.Target
		context lib.Context
	}{target, context})
	fake.recordInvocation("New", []interface{}{target, context})
	fake.newMutex.Unlock()
	if fake.NewStub != nil {
		return fake.NewStub(target, context)
	} else {
		return fake.newReturns.result1
	}
}

func (fake *FakeScimFactoryInterface) NewCallCount() int {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return len(fake.newArgsForCall)
}

func (fake *FakeScimFactoryInterface) NewArgsForCall(i int) (lib.Target, lib.Context) {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return fake.newArgsForCall[i].target, fake.newArgsForCall[i].context
}

func (fake *FakeScimFactoryInterface) NewReturns(result1 lib.Scim) {
	fake.NewStub = nil
	fake.newReturns = struct {
		result1 lib.Scim
	}{result1}
}

func (fake *FakeScimFactoryInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeScimFactoryInterface) recordInvocation(key string, args []interface{}) {
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

var _ lib.ScimFactoryInterface = new(FakeScimFactoryInterface)
