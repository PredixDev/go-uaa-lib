// This file was generated by counterfeiter
package libfakes

import (
	"sync"

	"github.com/PredixDev/go-uaa-lib"
)

type FakeScim struct {
	GetClientsStub        func(start int, count int) ([]lib.Client, error)
	getClientsMutex       sync.RWMutex
	getClientsArgsForCall []struct {
		start int
		count int
	}
	getClientsReturns struct {
		result1 []lib.Client
		result2 error
	}
	GetClientStub        func(clientID string) (*lib.Client, error)
	getClientMutex       sync.RWMutex
	getClientArgsForCall []struct {
		clientID string
	}
	getClientReturns struct {
		result1 *lib.Client
		result2 error
	}
	CreateClientStub        func(client *lib.Client) error
	createClientMutex       sync.RWMutex
	createClientArgsForCall []struct {
		client *lib.Client
	}
	createClientReturns struct {
		result1 error
	}
	PutClientStub        func(client *lib.Client) error
	putClientMutex       sync.RWMutex
	putClientArgsForCall []struct {
		client *lib.Client
	}
	putClientReturns struct {
		result1 error
	}
	DeleteClientStub        func(clientID string) error
	deleteClientMutex       sync.RWMutex
	deleteClientArgsForCall []struct {
		clientID string
	}
	deleteClientReturns struct {
		result1 error
	}
	GetUsersStub        func(start int, count int) ([]lib.User, error)
	getUsersMutex       sync.RWMutex
	getUsersArgsForCall []struct {
		start int
		count int
	}
	getUsersReturns struct {
		result1 []lib.User
		result2 error
	}
	GetUserStub        func(userName string) (*lib.User, error)
	getUserMutex       sync.RWMutex
	getUserArgsForCall []struct {
		userName string
	}
	getUserReturns struct {
		result1 *lib.User
		result2 error
	}
	CreateUserStub        func(user *lib.User) error
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		user *lib.User
	}
	createUserReturns struct {
		result1 error
	}
	PutUserStub        func(user *lib.User) error
	putUserMutex       sync.RWMutex
	putUserArgsForCall []struct {
		user *lib.User
	}
	putUserReturns struct {
		result1 error
	}
	DeleteUserStub        func(userID string) error
	deleteUserMutex       sync.RWMutex
	deleteUserArgsForCall []struct {
		userID string
	}
	deleteUserReturns struct {
		result1 error
	}
	RequestStub        func(method, path string, body interface{}, headers map[string]string) (status int, responseBody []byte, err error)
	requestMutex       sync.RWMutex
	requestArgsForCall []struct {
		method  string
		path    string
		body    interface{}
		headers map[string]string
	}
	requestReturns struct {
		result1 int
		result2 []byte
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScim) GetClients(start int, count int) ([]lib.Client, error) {
	fake.getClientsMutex.Lock()
	fake.getClientsArgsForCall = append(fake.getClientsArgsForCall, struct {
		start int
		count int
	}{start, count})
	fake.recordInvocation("GetClients", []interface{}{start, count})
	fake.getClientsMutex.Unlock()
	if fake.GetClientsStub != nil {
		return fake.GetClientsStub(start, count)
	} else {
		return fake.getClientsReturns.result1, fake.getClientsReturns.result2
	}
}

func (fake *FakeScim) GetClientsCallCount() int {
	fake.getClientsMutex.RLock()
	defer fake.getClientsMutex.RUnlock()
	return len(fake.getClientsArgsForCall)
}

func (fake *FakeScim) GetClientsArgsForCall(i int) (int, int) {
	fake.getClientsMutex.RLock()
	defer fake.getClientsMutex.RUnlock()
	return fake.getClientsArgsForCall[i].start, fake.getClientsArgsForCall[i].count
}

func (fake *FakeScim) GetClientsReturns(result1 []lib.Client, result2 error) {
	fake.GetClientsStub = nil
	fake.getClientsReturns = struct {
		result1 []lib.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeScim) GetClient(clientID string) (*lib.Client, error) {
	fake.getClientMutex.Lock()
	fake.getClientArgsForCall = append(fake.getClientArgsForCall, struct {
		clientID string
	}{clientID})
	fake.recordInvocation("GetClient", []interface{}{clientID})
	fake.getClientMutex.Unlock()
	if fake.GetClientStub != nil {
		return fake.GetClientStub(clientID)
	} else {
		return fake.getClientReturns.result1, fake.getClientReturns.result2
	}
}

func (fake *FakeScim) GetClientCallCount() int {
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	return len(fake.getClientArgsForCall)
}

func (fake *FakeScim) GetClientArgsForCall(i int) string {
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	return fake.getClientArgsForCall[i].clientID
}

func (fake *FakeScim) GetClientReturns(result1 *lib.Client, result2 error) {
	fake.GetClientStub = nil
	fake.getClientReturns = struct {
		result1 *lib.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeScim) CreateClient(client *lib.Client) error {
	fake.createClientMutex.Lock()
	fake.createClientArgsForCall = append(fake.createClientArgsForCall, struct {
		client *lib.Client
	}{client})
	fake.recordInvocation("CreateClient", []interface{}{client})
	fake.createClientMutex.Unlock()
	if fake.CreateClientStub != nil {
		return fake.CreateClientStub(client)
	} else {
		return fake.createClientReturns.result1
	}
}

func (fake *FakeScim) CreateClientCallCount() int {
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	return len(fake.createClientArgsForCall)
}

func (fake *FakeScim) CreateClientArgsForCall(i int) *lib.Client {
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	return fake.createClientArgsForCall[i].client
}

func (fake *FakeScim) CreateClientReturns(result1 error) {
	fake.CreateClientStub = nil
	fake.createClientReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScim) PutClient(client *lib.Client) error {
	fake.putClientMutex.Lock()
	fake.putClientArgsForCall = append(fake.putClientArgsForCall, struct {
		client *lib.Client
	}{client})
	fake.recordInvocation("PutClient", []interface{}{client})
	fake.putClientMutex.Unlock()
	if fake.PutClientStub != nil {
		return fake.PutClientStub(client)
	} else {
		return fake.putClientReturns.result1
	}
}

func (fake *FakeScim) PutClientCallCount() int {
	fake.putClientMutex.RLock()
	defer fake.putClientMutex.RUnlock()
	return len(fake.putClientArgsForCall)
}

func (fake *FakeScim) PutClientArgsForCall(i int) *lib.Client {
	fake.putClientMutex.RLock()
	defer fake.putClientMutex.RUnlock()
	return fake.putClientArgsForCall[i].client
}

func (fake *FakeScim) PutClientReturns(result1 error) {
	fake.PutClientStub = nil
	fake.putClientReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScim) DeleteClient(clientID string) error {
	fake.deleteClientMutex.Lock()
	fake.deleteClientArgsForCall = append(fake.deleteClientArgsForCall, struct {
		clientID string
	}{clientID})
	fake.recordInvocation("DeleteClient", []interface{}{clientID})
	fake.deleteClientMutex.Unlock()
	if fake.DeleteClientStub != nil {
		return fake.DeleteClientStub(clientID)
	} else {
		return fake.deleteClientReturns.result1
	}
}

func (fake *FakeScim) DeleteClientCallCount() int {
	fake.deleteClientMutex.RLock()
	defer fake.deleteClientMutex.RUnlock()
	return len(fake.deleteClientArgsForCall)
}

func (fake *FakeScim) DeleteClientArgsForCall(i int) string {
	fake.deleteClientMutex.RLock()
	defer fake.deleteClientMutex.RUnlock()
	return fake.deleteClientArgsForCall[i].clientID
}

func (fake *FakeScim) DeleteClientReturns(result1 error) {
	fake.DeleteClientStub = nil
	fake.deleteClientReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScim) GetUsers(start int, count int) ([]lib.User, error) {
	fake.getUsersMutex.Lock()
	fake.getUsersArgsForCall = append(fake.getUsersArgsForCall, struct {
		start int
		count int
	}{start, count})
	fake.recordInvocation("GetUsers", []interface{}{start, count})
	fake.getUsersMutex.Unlock()
	if fake.GetUsersStub != nil {
		return fake.GetUsersStub(start, count)
	} else {
		return fake.getUsersReturns.result1, fake.getUsersReturns.result2
	}
}

func (fake *FakeScim) GetUsersCallCount() int {
	fake.getUsersMutex.RLock()
	defer fake.getUsersMutex.RUnlock()
	return len(fake.getUsersArgsForCall)
}

func (fake *FakeScim) GetUsersArgsForCall(i int) (int, int) {
	fake.getUsersMutex.RLock()
	defer fake.getUsersMutex.RUnlock()
	return fake.getUsersArgsForCall[i].start, fake.getUsersArgsForCall[i].count
}

func (fake *FakeScim) GetUsersReturns(result1 []lib.User, result2 error) {
	fake.GetUsersStub = nil
	fake.getUsersReturns = struct {
		result1 []lib.User
		result2 error
	}{result1, result2}
}

func (fake *FakeScim) GetUser(userName string) (*lib.User, error) {
	fake.getUserMutex.Lock()
	fake.getUserArgsForCall = append(fake.getUserArgsForCall, struct {
		userName string
	}{userName})
	fake.recordInvocation("GetUser", []interface{}{userName})
	fake.getUserMutex.Unlock()
	if fake.GetUserStub != nil {
		return fake.GetUserStub(userName)
	} else {
		return fake.getUserReturns.result1, fake.getUserReturns.result2
	}
}

func (fake *FakeScim) GetUserCallCount() int {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	return len(fake.getUserArgsForCall)
}

func (fake *FakeScim) GetUserArgsForCall(i int) string {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	return fake.getUserArgsForCall[i].userName
}

func (fake *FakeScim) GetUserReturns(result1 *lib.User, result2 error) {
	fake.GetUserStub = nil
	fake.getUserReturns = struct {
		result1 *lib.User
		result2 error
	}{result1, result2}
}

func (fake *FakeScim) CreateUser(user *lib.User) error {
	fake.createUserMutex.Lock()
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		user *lib.User
	}{user})
	fake.recordInvocation("CreateUser", []interface{}{user})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub(user)
	} else {
		return fake.createUserReturns.result1
	}
}

func (fake *FakeScim) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeScim) CreateUserArgsForCall(i int) *lib.User {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return fake.createUserArgsForCall[i].user
}

func (fake *FakeScim) CreateUserReturns(result1 error) {
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScim) PutUser(user *lib.User) error {
	fake.putUserMutex.Lock()
	fake.putUserArgsForCall = append(fake.putUserArgsForCall, struct {
		user *lib.User
	}{user})
	fake.recordInvocation("PutUser", []interface{}{user})
	fake.putUserMutex.Unlock()
	if fake.PutUserStub != nil {
		return fake.PutUserStub(user)
	} else {
		return fake.putUserReturns.result1
	}
}

func (fake *FakeScim) PutUserCallCount() int {
	fake.putUserMutex.RLock()
	defer fake.putUserMutex.RUnlock()
	return len(fake.putUserArgsForCall)
}

func (fake *FakeScim) PutUserArgsForCall(i int) *lib.User {
	fake.putUserMutex.RLock()
	defer fake.putUserMutex.RUnlock()
	return fake.putUserArgsForCall[i].user
}

func (fake *FakeScim) PutUserReturns(result1 error) {
	fake.PutUserStub = nil
	fake.putUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScim) DeleteUser(userID string) error {
	fake.deleteUserMutex.Lock()
	fake.deleteUserArgsForCall = append(fake.deleteUserArgsForCall, struct {
		userID string
	}{userID})
	fake.recordInvocation("DeleteUser", []interface{}{userID})
	fake.deleteUserMutex.Unlock()
	if fake.DeleteUserStub != nil {
		return fake.DeleteUserStub(userID)
	} else {
		return fake.deleteUserReturns.result1
	}
}

func (fake *FakeScim) DeleteUserCallCount() int {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return len(fake.deleteUserArgsForCall)
}

func (fake *FakeScim) DeleteUserArgsForCall(i int) string {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return fake.deleteUserArgsForCall[i].userID
}

func (fake *FakeScim) DeleteUserReturns(result1 error) {
	fake.DeleteUserStub = nil
	fake.deleteUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScim) Request(method string, path string, body interface{}, headers map[string]string) (status int, responseBody []byte, err error) {
	fake.requestMutex.Lock()
	fake.requestArgsForCall = append(fake.requestArgsForCall, struct {
		method  string
		path    string
		body    interface{}
		headers map[string]string
	}{method, path, body, headers})
	fake.recordInvocation("Request", []interface{}{method, path, body, headers})
	fake.requestMutex.Unlock()
	if fake.RequestStub != nil {
		return fake.RequestStub(method, path, body, headers)
	} else {
		return fake.requestReturns.result1, fake.requestReturns.result2, fake.requestReturns.result3
	}
}

func (fake *FakeScim) RequestCallCount() int {
	fake.requestMutex.RLock()
	defer fake.requestMutex.RUnlock()
	return len(fake.requestArgsForCall)
}

func (fake *FakeScim) RequestArgsForCall(i int) (string, string, interface{}, map[string]string) {
	fake.requestMutex.RLock()
	defer fake.requestMutex.RUnlock()
	return fake.requestArgsForCall[i].method, fake.requestArgsForCall[i].path, fake.requestArgsForCall[i].body, fake.requestArgsForCall[i].headers
}

func (fake *FakeScim) RequestReturns(result1 int, result2 []byte, result3 error) {
	fake.RequestStub = nil
	fake.requestReturns = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeScim) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getClientsMutex.RLock()
	defer fake.getClientsMutex.RUnlock()
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	fake.createClientMutex.RLock()
	defer fake.createClientMutex.RUnlock()
	fake.putClientMutex.RLock()
	defer fake.putClientMutex.RUnlock()
	fake.deleteClientMutex.RLock()
	defer fake.deleteClientMutex.RUnlock()
	fake.getUsersMutex.RLock()
	defer fake.getUsersMutex.RUnlock()
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	fake.putUserMutex.RLock()
	defer fake.putUserMutex.RUnlock()
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	fake.requestMutex.RLock()
	defer fake.requestMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeScim) recordInvocation(key string, args []interface{}) {
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

var _ lib.Scim = new(FakeScim)