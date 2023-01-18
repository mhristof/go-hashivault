package vault

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
)

type Input struct {
	Addr  string
	Token string
}

type client struct {
	addr string
	get  *http.Request
}

func New(input *Input) *client {
	var ret client
	if input.Addr != "" {
		ret.addr = input.Addr
	}

	if input.Token == "" {
		input.Token = getToken()
	}

	req, err := http.NewRequest("GET", ret.addr, nil)
	if err != nil {
		panic("panic!")
	}

	ret.get = req
	ret.get.Header.Set("X-Vault-Token", input.Token)

	health := ret.Health()

	return &ret
}

func getToken() string {
	token := os.Getenv("VAULT_TOKEN")
	if token != "" {
		return token
	}

	tokenPath, err := homedir.Expand("~/.vault-token")
	if err == nil {
		data, err := os.ReadFile(tokenPath)
		if err == nil {
			return strings.TrimSpace(string(data))
		}
	}

	return ""
}

func (v *client) Health() *HealthResponse {
	body := v.Get("v1/sys/health")

	var health HealthResponse
	err := json.Unmarshal(body, &health)
	if err != nil {
		panic(err)
	}

	return &health
}

func (v *client) Get(urlPath string) []byte {
	client := &http.Client{}
	addr, err := url.Parse(v.addr + "/" + urlPath)
	if err != nil {
		panic(err)
	}

	v.get.URL = addr

	resp, err := client.Do(v.get)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}
