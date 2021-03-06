// This file was generated by counterfeiter
package libfakes

import (
	"sync"

	"github.com/PredixDev/go-uaa-lib"
)

type FakeTokenIssuerFactoryInterface struct {
	NewStub        func(target, clientID, clientSecret string, skipSslValidation bool, caCertFile string) lib.TokenIssuer
	newMutex       sync.RWMutex
	newArgsForCall []struct {
		target            string
		clientID          string
		clientSecret      string
		skipSslValidation bool
		caCertFile        string
	}
	newReturns struct {
		result1 lib.TokenIssuer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenIssuerFactoryInterface) New(target string, clientID string, clientSecret string, skipSslValidation bool, caCertFile string) lib.TokenIssuer {
	fake.newMutex.Lock()
	fake.newArgsForCall = append(fake.newArgsForCall, struct {
		target            string
		clientID          string
		clientSecret      string
		skipSslValidation bool
		caCertFile        string
	}{target, clientID, clientSecret, skipSslValidation, caCertFile})
	fake.recordInvocation("New", []interface{}{target, clientID, clientSecret, skipSslValidation, caCertFile})
	fake.newMutex.Unlock()
	if fake.NewStub != nil {
		return fake.NewStub(target, clientID, clientSecret, skipSslValidation, caCertFile)
	} else {
		return fake.newReturns.result1
	}
}

func (fake *FakeTokenIssuerFactoryInterface) NewCallCount() int {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return len(fake.newArgsForCall)
}

func (fake *FakeTokenIssuerFactoryInterface) NewArgsForCall(i int) (string, string, string, bool, string) {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return fake.newArgsForCall[i].target, fake.newArgsForCall[i].clientID, fake.newArgsForCall[i].clientSecret, fake.newArgsForCall[i].skipSslValidation, fake.newArgsForCall[i].caCertFile
}

func (fake *FakeTokenIssuerFactoryInterface) NewReturns(result1 lib.TokenIssuer) {
	fake.NewStub = nil
	fake.newReturns = struct {
		result1 lib.TokenIssuer
	}{result1}
}

func (fake *FakeTokenIssuerFactoryInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTokenIssuerFactoryInterface) recordInvocation(key string, args []interface{}) {
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

var _ lib.TokenIssuerFactoryInterface = new(FakeTokenIssuerFactoryInterface)
