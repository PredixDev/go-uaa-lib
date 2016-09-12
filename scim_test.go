package lib_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	lib "github.com/PredixDev/go-uaa-lib"
	. "github.com/onsi/gomega"
)

func TestScimRequestWithInvalidTarget(t *testing.T) {
	RegisterTestingT(t)

	scim := NewScim("Invalid URL", "", "")
	_, _, err := scim.Request("GET", "/ABC", nil, nil)
	Expect(err).ToNot(BeNil())
}

func TestScimRequestWithInvalidStatus(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/ABC"))

		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	_, _, err := scim.Request("GET", "/ABC", nil, nil)

	Expect(err).ToNot(BeNil())
	Expect(err.Error()).To(Equal("Invalid status response: 500"))
}

func TestScimRequestWithInvalidStatusAndError(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/ABC"))

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error_description": "Some error"}`))
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	_, _, err := scim.Request("GET", "/ABC", nil, nil)

	Expect(err).ToNot(BeNil())
	Expect(err.Error()).To(Equal("Invalid status response: 500. Some error"))
}

func TestScimRequestWithBody(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("POST"))
		Expect(r.Header.Get("Content-Type")).To(Equal("application/json"))
		requestBody, err := ioutil.ReadAll(r.Body)

		Expect(err).To(BeNil())
		Expect(requestBody).To(Equal([]byte(`{"field1":"Value1","field2":1234}`)))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	_, _, err := scim.Request("POST", "", SampleBody{
		Field1: "Value1",
		Field2: 1234,
	}, nil)

	Expect(err).To(BeNil())
}

func TestScimRequestWithHeaders(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("POST"))
		Expect(r.Header.Get("Header1")).To(Equal("Header-Value-1"))
		Expect(r.Header.Get("Header2")).To(Equal("Header-Value-2"))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	_, _, err := scim.Request("POST", "", nil, map[string]string{
		"Header1": "Header-Value-1",
		"Header2": "Header-Value-2",
	})

	Expect(err).To(BeNil())
}

func TestScimGetClientsWithNoStartAndCount(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/oauth/clients"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"resources":[{"client_id": "client1","scope": ["uaa.none"],"resource_ids": ["none"],"authorized_grant_types":["client_credentials"],"authorities":["clients.read","uaa.resource"]}]}`))
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	clients, err := scim.GetClients(-1, -1)

	Expect(err).To(BeNil())
	Expect(len(clients)).To(Equal(1))
	Expect(clients[0].ID).To(Equal("client1"))
	Expect(clients[0].Scopes).To(ConsistOf("uaa.none"))
	Expect(clients[0].ResourceIDs).To(ConsistOf("none"))
	Expect(clients[0].GrantTypes).To(ConsistOf("client_credentials"))
	Expect(clients[0].Authorities).To(ConsistOf("clients.read", "uaa.resource"))
}

func TestScimGetClientsWithStartOnly(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/oauth/clients?startIndex=10"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"resources":[{"client_id": "client1","scope": ["uaa.none"],"resource_ids": ["none"],"authorized_grant_types":["client_credentials"],"authorities":["clients.read","uaa.resource"]}]}`))
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	clients, err := scim.GetClients(10, -1)

	Expect(err).To(BeNil())
	Expect(len(clients)).To(Equal(1))
	Expect(clients[0].ID).To(Equal("client1"))
	Expect(clients[0].Scopes).To(ConsistOf("uaa.none"))
	Expect(clients[0].ResourceIDs).To(ConsistOf("none"))
	Expect(clients[0].GrantTypes).To(ConsistOf("client_credentials"))
	Expect(clients[0].Authorities).To(ConsistOf("clients.read", "uaa.resource"))
}

func TestScimGetClientsWithCountOnly(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/oauth/clients?count=5"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"resources":[{"client_id": "client1","scope": ["uaa.none"],"resource_ids": ["none"],"authorized_grant_types":["client_credentials"],"authorities":["clients.read","uaa.resource"]}]}`))
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	clients, err := scim.GetClients(-1, 5)

	Expect(err).To(BeNil())
	Expect(len(clients)).To(Equal(1))
	Expect(clients[0].ID).To(Equal("client1"))
	Expect(clients[0].Scopes).To(ConsistOf("uaa.none"))
	Expect(clients[0].ResourceIDs).To(ConsistOf("none"))
	Expect(clients[0].GrantTypes).To(ConsistOf("client_credentials"))
	Expect(clients[0].Authorities).To(ConsistOf("clients.read", "uaa.resource"))
}

func TestScimGetClientsWithStartAndCount(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/oauth/clients?startIndex=20&count=3"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"resources":[{"client_id": "client1","scope": ["uaa.none"],"resource_ids": ["none"],"authorized_grant_types":["client_credentials"],"authorities":["clients.read","uaa.resource"]}]}`))
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	clients, err := scim.GetClients(20, 3)

	Expect(err).To(BeNil())
	Expect(len(clients)).To(Equal(1))
	Expect(clients[0].ID).To(Equal("client1"))
	Expect(clients[0].Scopes).To(ConsistOf("uaa.none"))
	Expect(clients[0].ResourceIDs).To(ConsistOf("none"))
	Expect(clients[0].GrantTypes).To(ConsistOf("client_credentials"))
	Expect(clients[0].Authorities).To(ConsistOf("clients.read", "uaa.resource"))
}

func TestScimGetClientsWithError(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	clients, err := scim.GetClients(1, 2)

	Expect(err).ToNot(BeNil())
	Expect(clients).To(BeNil())
}

func TestScimGetClientWhereClientNotFound(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/oauth/clients/does-not-exist-client"))

		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	client, err := scim.GetClient("does-not-exist-client")

	Expect(err).To(BeNil())
	Expect(client).To(BeNil())
}

func TestScimGetClientWhereClientIsFound(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/oauth/clients/some-client"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"client_id": "some-client","scope": ["uaa.none", "openid"],"resource_ids": ["openid"],"authorized_grant_types":["client_credentials","password"],"authorities":["clients.write"]}`))
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	client, err := scim.GetClient("some-client")

	Expect(err).To(BeNil())
	Expect(client).ToNot(BeNil())
	Expect(client.ID).To(Equal("some-client"))
	Expect(client.Scopes).To(ConsistOf("uaa.none", "openid"))
	Expect(client.ResourceIDs).To(ConsistOf("openid"))
	Expect(client.GrantTypes).To(ConsistOf("client_credentials", "password"))
	Expect(client.Authorities).To(ConsistOf("clients.write"))
}

func TestScimGetClientWithError(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	client, err := scim.GetClient("some-client")

	Expect(err).ToNot(BeNil())
	Expect(client).To(BeNil())
}

func TestScimCreateClient(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("POST"))
		Expect(r.URL.RequestURI()).To(Equal("/oauth/clients"))

		requestBody, err := ioutil.ReadAll(r.Body)

		Expect(err).To(BeNil())
		Expect(requestBody).To(Equal([]byte(`{"client_id":"client1"}`)))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	err := scim.CreateClient(&lib.Client{ID: "client1"})

	Expect(err).To(BeNil())
}

func TestScimPutClient(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("PUT"))
		Expect(r.URL.RequestURI()).To(Equal("/oauth/clients/client2"))

		requestBody, err := ioutil.ReadAll(r.Body)

		Expect(err).To(BeNil())
		Expect(requestBody).To(Equal([]byte(`{"client_id":"client2","name":"Some Client"}`)))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	err := scim.PutClient(&lib.Client{ID: "client2", Name: "Some Client"})

	Expect(err).To(BeNil())
}

func TestScimDeleteClient(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("DELETE"))
		Expect(r.URL.RequestURI()).To(Equal("/oauth/clients/client3"))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	err := scim.DeleteClient("client3")

	Expect(err).To(BeNil())
}

func TestScimGetUsersWithNoStartAndCount(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/Users"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"resources":[{"id": "ABCD","username": "User1"},{"id": "EFGH","username": "User2"}]}`))
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	users, err := scim.GetUsers(-1, -1)

	Expect(err).To(BeNil())
	Expect(len(users)).To(Equal(2))
	Expect(users[0].ID).To(Equal("ABCD"))
	Expect(users[0].UserName).To(Equal("User1"))
	Expect(users[1].ID).To(Equal("EFGH"))
	Expect(users[1].UserName).To(Equal("User2"))
}

func TestScimGetUsersWithStartOnly(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/Users?startIndex=10"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"resources":[{"id": "1234","meta": {"created": "2016-01-01T00:00:00.000Z","lastModified": "2016-01-02T00:01:10.123Z"},"username": "Some-User"}]}`))
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	users, err := scim.GetUsers(10, -1)

	Expect(err).To(BeNil())
	Expect(len(users)).To(Equal(1))
	Expect(users[0].ID).To(Equal("1234"))
	Expect(users[0].Meta.Created).To(Equal("2016-01-01T00:00:00.000Z"))
	Expect(users[0].Meta.LastModified).To(Equal("2016-01-02T00:01:10.123Z"))
	Expect(users[0].UserName).To(Equal("Some-User"))
}

func TestScimGetUsersWithCountOnly(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/Users?count=5"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"resources":[{"id": "ABCD","username": "Some-User","name":{"givenName":"John","familyName":"Doe"},"emails":[{"value":"user@name.com"}]}]}`))
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	users, err := scim.GetUsers(-1, 5)

	Expect(err).To(BeNil())
	Expect(len(users)).To(Equal(1))
	Expect(users[0].ID).To(Equal("ABCD"))
	Expect(users[0].UserName).To(Equal("Some-User"))
	Expect(users[0].Name.GivenName).To(Equal("John"))
	Expect(users[0].Name.FamilyName).To(Equal("Doe"))
	Expect(len(users[0].Emails)).To(Equal(1))
	Expect(users[0].Emails[0].Value).To(Equal("user@name.com"))
	Expect(users[0].Emails[0].Primary).To(BeFalse())
}

func TestScimGetUsersWithStartAndCount(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/Users?startIndex=20&count=3"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"resources":[{"id": "XYZ","username": "User2","groups":["Group1","Group2"]}]}`))
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	users, err := scim.GetUsers(20, 3)

	Expect(err).To(BeNil())
	Expect(len(users)).To(Equal(1))
	Expect(users[0].ID).To(Equal("XYZ"))
	Expect(users[0].UserName).To(Equal("User2"))
	Expect(users[0].Groups).To(ConsistOf("Group1", "Group2"))
}

func TestScimGetUsersWithError(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	users, err := scim.GetUsers(1, 2)

	Expect(err).ToNot(BeNil())
	Expect(users).To(BeNil())
}

func TestScimGetUserWhereUserNotFound(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/Users?filter=userName+eq+%22does-not-exist-user%22"))

		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	user, err := scim.GetUser("does-not-exist-user")

	Expect(err).To(BeNil())
	Expect(user).To(BeNil())
}

func TestScimGetUserWhereUserIsFound(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("GET"))
		Expect(r.URL.RequestURI()).To(Equal("/Users?filter=userName+eq+%22some-user%22"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"resources":[{"id": "12345","username": "some-user"}]}`))
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	user, err := scim.GetUser("some-user")

	Expect(err).To(BeNil())
	Expect(user).ToNot(BeNil())
	Expect(user.ID).To(Equal("12345"))
	Expect(user.UserName).To(Equal("some-user"))
}

func TestScimGetUserWithError(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	user, err := scim.GetUser("some-user")

	Expect(err).ToNot(BeNil())
	Expect(user).To(BeNil())
}

func TestScimCreateUser(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("POST"))
		Expect(r.URL.RequestURI()).To(Equal("/Users"))

		requestBody, err := ioutil.ReadAll(r.Body)

		Expect(err).To(BeNil())
		Expect(requestBody).To(Equal([]byte(`{"id":"user1","meta":{},"userName":"User1","name":{},"emails":[{"value":"user1@domain.com","primary":true}]}`)))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	err := scim.CreateUser(&lib.User{ID: "user1", UserName: "User1", Emails: []lib.Value{
		lib.Value{
			Value:   "user1@domain.com",
			Primary: true,
		},
	}})

	Expect(err).To(BeNil())
}

func TestScimPutUser(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("PUT"))
		Expect(r.URL.RequestURI()).To(Equal("/Users/user2"))
		Expect(r.Header.Get("If-Match")).To(Equal("2"))

		requestBody, err := ioutil.ReadAll(r.Body)

		Expect(err).To(BeNil())
		Expect(requestBody).To(Equal([]byte(`{"id":"user2","meta":{"version":2},"userName":"Some-User","name":{}}`)))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	err := scim.PutUser(&lib.User{ID: "user2", UserName: "Some-User", Meta: lib.Meta{Version: 2}})

	Expect(err).To(BeNil())
}

func TestScimDeleteUser(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.Method).To(Equal("DELETE"))
		Expect(r.URL.RequestURI()).To(Equal("/Users/user3"))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	scim := NewScim(server.URL, "", "")
	err := scim.DeleteUser("user3")

	Expect(err).To(BeNil())
}

func NewScim(url, tokenType, accessToken string) lib.Scim {
	return lib.NewScim(FakeTarget{
		url: url,
	}, FakeContext{
		tokenType:   tokenType,
		accessToken: accessToken,
	})
}

type FakeTarget struct {
	url           string
	skipSslVerify bool
	caCertFile    string
}

func (t FakeTarget) URL() string {
	return t.url
}

func (t FakeTarget) SkipSslVerify() bool {
	return t.skipSslVerify
}

func (t FakeTarget) CaCertFile() string {
	return t.caCertFile
}

type FakeContext struct {
	tokenType   string
	accessToken string
}

func (c FakeContext) TokenType() string {
	return c.tokenType
}

func (c FakeContext) AccessToken() string {
	return c.accessToken
}

type SampleBody struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}
