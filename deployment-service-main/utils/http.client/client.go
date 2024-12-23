package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type Credentials struct {
	XClientID     string
	XClientSecret string
}

type HttpClient struct {
	Endpoint    string
	Client      *http.Client
	RetryClient *http.Client
	MaxRetry    int
	Credentials Credentials
}

type IHttpClient interface {
	Get(queries map[string]string) (*HttpResponse, error)
	Post(data map[string]interface{}, queries map[string]string) (*HttpResponse, error)
	PostWithCustomHeaders(data map[string]interface{}, multipartPayload *bytes.Buffer, queries map[string]string, headers map[string]string) (*HttpResponse, error)
	GetWithCustomHeaders(queries map[string]string, headers map[string]string) (*HttpResponse, error)
	GetWithRetry(queries map[string]string, headers map[string]string) (*HttpResponse, error)
	PostWithRetry(data map[string]interface{}, multipartPayload *bytes.Buffer, queries map[string]string, headers map[string]string) (*HttpResponse, error)
}

type HttpResponse struct {
	StatusCode int
	Body       string
}

func NewHttpClient(endpoint string, args ...string) IHttpClient {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 3
	retryClient.RetryWaitMax = time.Second * 2
	var creds Credentials
	if len(args) > 0 {
		creds = Credentials{XClientID: args[0], XClientSecret: args[1]}
	}
	c := &HttpClient{
		Endpoint:    endpoint,
		Client:      &http.Client{Timeout: 2 * time.Minute},
		Credentials: creds,
		RetryClient: retryClient.StandardClient(),
	}
	return c
}

func (h *HttpClient) Get(queries map[string]string) (*HttpResponse, error) {
	req, err := http.NewRequest(http.MethodGet, h.Endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Client-Id", h.Credentials.XClientID)
	req.Header.Add("Client-Secret", h.Credentials.XClientSecret)
	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()
	for k, v := range queries {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()

	resp, err := h.Client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return nil, err
	}

	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := HttpResponse{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}
	return &response, nil
}

func (h *HttpClient) Post(data map[string]interface{}, queries map[string]string) (*HttpResponse, error) {

	marshalled, err := json.Marshal(data)
	req, err := http.NewRequest(http.MethodPost, h.Endpoint, bytes.NewBuffer(marshalled))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Client-Id", h.Credentials.XClientID)
	req.Header.Add("Client-Secret", h.Credentials.XClientSecret)
	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()
	for k, v := range queries {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()

	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := HttpResponse{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

func (h *HttpClient) PostWithCustomHeaders(data map[string]interface{}, multipartPayload *bytes.Buffer, queries map[string]string, headers map[string]string) (*HttpResponse, error) {
	var marshalled []byte
	var err error
	var req *http.Request
	if multipartPayload != nil {
		req, err = http.NewRequest(http.MethodPost, h.Endpoint, multipartPayload)
	} else {
		marshalled, err = json.Marshal(data)
		req, err = http.NewRequest(http.MethodPost, h.Endpoint, bytes.NewBuffer(marshalled))
	}
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := HttpResponse{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

func (h *HttpClient) GetWithCustomHeaders(queries map[string]string, headers map[string]string) (*HttpResponse, error) {

	req, err := http.NewRequest(http.MethodGet, h.Endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()
	for k, v := range queries {
		query.Add(k, v)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	req.URL.RawQuery = query.Encode()
	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := HttpResponse{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

func (h *HttpClient) GetWithRetry(queries map[string]string, headers map[string]string) (*HttpResponse, error) {
	req, err := http.NewRequest(http.MethodGet, h.Endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()
	for k, v := range queries {
		query.Add(k, v)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	req.URL.RawQuery = query.Encode()
	resp, err := h.RetryClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := HttpResponse{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

func (h *HttpClient) PostWithRetry(data map[string]interface{}, multipartPayload *bytes.Buffer, queries map[string]string, headers map[string]string) (*HttpResponse, error) {
	var marshalled []byte
	var err error
	var req *http.Request
	if multipartPayload != nil {
		req, err = http.NewRequest(http.MethodPost, h.Endpoint, multipartPayload)
	} else {
		marshalled, err = json.Marshal(data)
		req, err = http.NewRequest(http.MethodPost, h.Endpoint, bytes.NewBuffer(marshalled))
	}
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := h.RetryClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := HttpResponse{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}
