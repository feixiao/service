package main

import(
	"errors"
	"strings"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"context"
	"net/http"
	"encoding/json"
	"log"
)

// sample 原地址：  https://github.com/go-kit/kit/tree/master/examples/stringsvc1
// 注解：http://blog.chunshengster.me/2016/05/28/go_kit_ru_men_er_di_yi_ge_go_kit_cheng_xu/

// 验证方法：
// curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
// curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count

// ErrEmpty 提示输入为空
var ErrEmpty = errors.New("empty string")


func main() {
	svc := stringService{}
	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeiCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase",uppercaseHandler)
	http.Handle("/count",countHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}


// 服务定义

// StringService 提供了针对string的操作
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int 
}

type stringService struct{}

func (stringService) Uppercase(s string )(string,error) {
	if s == "" {
		return "",ErrEmpty
	}
	return strings.ToUpper(s),nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// 请求和应答定义

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct{
	V string `json:"v"`
	Err string `json:"err,omitempty"`
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


func makeiCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context,request interface{}) (interface{},error) {
		req := request.(countRequest)
		v:= svc.Count(req.S)
		return countResponse{v}, nil
	}
}


