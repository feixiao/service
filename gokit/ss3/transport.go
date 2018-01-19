package main 

import (
	"encoding/json"
	"context"
	"net/http"
	"github.com/go-kit/kit/endpoint"
	"bytes"
	"io/ioutil"
)
// 请求和应答定义

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct{
	V string `json:"v"`
	Err string `json:"err,omitempty"`
}

func decodeUppercaseResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response uppercaseResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeUppercaseRequest(_ context.Context,r *http.Request)(interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}


type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func decodeCountRequest(_ context.Context,r *http.Request)(interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// 端点,一个端点表示一个RPC服务


func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context,request interface{})(interface{},error) {
		req:= request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v,err.Error()},nil
		}
		return uppercaseResponse{v,""},nil
	}
}


func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context,request interface{}) (interface{},error) {
		req := request.(countRequest)
		v:= svc.Count(req.S)
		return countResponse{v}, nil
	}
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}