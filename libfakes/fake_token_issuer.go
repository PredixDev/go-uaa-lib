// This file was generated by counterfeiter
package libfakes

import (
	"net/url"
	"sync"

	"github.com/PredixDev/go-uaa-lib"
)

type FakeTokenIssuer struct {
	ClientCredentialsGrantStub        func(scopes []string) (*lib.TokenResponse, error)
	clientCredentialsGrantMutex       sync.RWMutex
	clientCredentialsGrantArgsForCall []struct {
		scopes []string
	}
	clientCredentialsGrantReturns struct {
		result1 *lib.TokenResponse
		result2 error
	}
	PasswordGrantStub        func(username, password string, scopes []string) (*lib.TokenResponse, error)
	passwordGrantMutex       sync.RWMutex
	passwordGrantArgsForCall []struct {
		username string
		password string
		scopes   []string
	}
	passwordGrantReturns struct {
		result1 *lib.TokenResponse
		result2 error
	}
	RequestTokenStub        func(params url.Values, scopes []string) (*lib.TokenResponse, error)
	requestTokenMutex       sync.RWMutex
	requestTokenArgsForCall []struct {
		params url.Values
		scopes []string
	}
	requestTokenReturns struct {
		result1 *lib.TokenResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenIssuer) ClientCredentialsGrant(scopes []string) (*lib.TokenResponse, error) {
	var scopesCopy []string
	if scopes != nil {
		scopesCopy = make([]string, len(scopes))
		copy(scopesCopy, scopes)
	}
	fake.clientCredentialsGrantMutex.Lock()
	fake.clientCredentialsGrantArgsForCall = append(fake.clientCredentialsGrantArgsForCall, struct {
		scopes []string
	}{scopesCopy})
	fake.recordInvocation("ClientCredentialsGrant", []interface{}{scopesCopy})
	fake.clientCredentialsGrantMutex.Unlock()
	if fake.ClientCredentialsGrantStub != nil {
		return fake.ClientCredentialsGrantStub(scopes)
	} else {
		return fake.clientCredentialsGrantReturns.result1, fake.clientCredentialsGrantReturns.result2
	}
}

func (fake *FakeTokenIssuer) ClientCredentialsGrantCallCount() int {
	fake.clientCredentialsGrantMutex.RLock()
	defer fake.clientCredentialsGrantMutex.RUnlock()
	return len(fake.clientCredentialsGrantArgsForCall)
}

func (fake *FakeTokenIssuer) ClientCredentialsGrantArgsForCall(i int) []string {
	fake.clientCredentialsGrantMutex.RLock()
	defer fake.clientCredentialsGrantMutex.RUnlock()
	return fake.clientCredentialsGrantArgsForCall[i].scopes
}

func (fake *FakeTokenIssuer) ClientCredentialsGrantReturns(result1 *lib.TokenResponse, result2 error) {
	fake.ClientCredentialsGrantStub = nil
	fake.clientCredentialsGrantReturns = struct {
		result1 *lib.TokenResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeTokenIssuer) PasswordGrant(username string, password string, scopes []string) (*lib.TokenResponse, error) {
	var scopesCopy []string
	if scopes != nil {
		scopesCopy = make([]string, len(scopes))
		copy(scopesCopy, scopes)
	}
	fake.passwordGrantMutex.Lock()
	fake.passwordGrantArgsForCall = append(fake.passwordGrantArgsForCall, struct {
		username string
		password string
		scopes   []string
	}{username, password, scopesCopy})
	fake.recordInvocation("PasswordGrant", []interface{}{username, password, scopesCopy})
	fake.passwordGrantMutex.Unlock()
	if fake.PasswordGrantStub != nil {
		return fake.PasswordGrantStub(username, password, scopes)
	} else {
		return fake.passwordGrantReturns.result1, fake.passwordGrantReturns.result2
	}
}

func (fake *FakeTokenIssuer) PasswordGrantCallCount() int {
	fake.passwordGrantMutex.RLock()
	defer fake.passwordGrantMutex.RUnlock()
	return len(fake.passwordGrantArgsForCall)
}

func (fake *FakeTokenIssuer) PasswordGrantArgsForCall(i int) (string, string, []string) {
	fake.passwordGrantMutex.RLock()
	defer fake.passwordGrantMutex.RUnlock()
	return fake.passwordGrantArgsForCall[i].username, fake.passwordGrantArgsForCall[i].password, fake.passwordGrantArgsForCall[i].scopes
}

func (fake *FakeTokenIssuer) PasswordGrantReturns(result1 *lib.TokenResponse, result2 error) {
	fake.PasswordGrantStub = nil
	fake.passwordGrantReturns = struct {
		result1 *lib.TokenResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeTokenIssuer) RequestToken(params url.Values, scopes []string) (*lib.TokenResponse, error) {
	var scopesCopy []string
	if scopes != nil {
		scopesCopy = make([]string, len(scopes))
		copy(scopesCopy, scopes)
	}
	fake.requestTokenMutex.Lock()
	fake.requestTokenArgsForCall = append(fake.requestTokenArgsForCall, struct {
		params url.Values
		scopes []string
	}{params, scopesCopy})
	fake.recordInvocation("RequestToken", []interface{}{params, scopesCopy})
	fake.requestTokenMutex.Unlock()
	if fake.RequestTokenStub != nil {
		return fake.RequestTokenStub(params, scopes)
	} else {
		return fake.requestTokenReturns.result1, fake.requestTokenReturns.result2
	}
}

func (fake *FakeTokenIssuer) RequestTokenCallCount() int {
	fake.requestTokenMutex.RLock()
	defer fake.requestTokenMutex.RUnlock()
	return len(fake.requestTokenArgsForCall)
}

func (fake *FakeTokenIssuer) RequestTokenArgsForCall(i int) (url.Values, []string) {
	fake.requestTokenMutex.RLock()
	defer fake.requestTokenMutex.RUnlock()
	return fake.requestTokenArgsForCall[i].params, fake.requestTokenArgsForCall[i].scopes
}

func (fake *FakeTokenIssuer) RequestTokenReturns(result1 *lib.TokenResponse, result2 error) {
	fake.RequestTokenStub = nil
	fake.requestTokenReturns = struct {
		result1 *lib.TokenResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeTokenIssuer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clientCredentialsGrantMutex.RLock()
	defer fake.clientCredentialsGrantMutex.RUnlock()
	fake.passwordGrantMutex.RLock()
	defer fake.passwordGrantMutex.RUnlock()
	fake.requestTokenMutex.RLock()
	defer fake.requestTokenMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTokenIssuer) recordInvocation(key string, args []interface{}) {
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

var _ lib.TokenIssuer = new(FakeTokenIssuer)
