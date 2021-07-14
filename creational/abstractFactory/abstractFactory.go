package abstractFactory

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//抽象工厂模式，和工厂模式的区别就是返回的不是实例而是接口

// We define a Doer interface, that has the method signature
// of the `http.Client` structs `Do` method
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// This gives us a regular HTTP client from the `net/http` package
func NewHTTPClient() Doer {
	return &http.Client{}
}

type mockHTTPClient struct{}

func (*mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// The `NewRecorder` method of the httptest package gives us
	// a new mock request generator
	res := httptest.NewRecorder()

	// calling the `Result` method gives us
	// the default empty *http.Response object
	return res.Result(), nil
}

// This gives us a mock HTTP client, which returns
// an empty response for any request sent to it
func NewMockHTTPClient() Doer {
	return &mockHTTPClient{}
}

//假如你想测试这段代码
func QueryUser(doer Doer) error {
	req, err := http.NewRequest("Get", "http://iam.api.marmotedu.com:8080/v1/secrets", nil)
	if err != nil {
		return err
	}

	_, err = doer.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func TestQueryUser(t *testing.T) {
	doer := NewMockHTTPClient()
	if err := QueryUser(doer); err != nil {
		t.Errorf("QueryUser failed, err: %v", err)
	}
}
