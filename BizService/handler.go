package main

import (
	"context"
	"encoding/json"
	"fmt"
	"genhttpserv/kitex_gen/http"
	"log"
)

// BizServiceImpl implements the last service interface defined in the IDL.
type BizServiceImpl struct{}

func (s *BizServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	fmt.Print(request)

	// use jsoniter or other json parse sdk to assert request
	// m := request.(string)

	//this is a hashmap mapping method names to the functions
	funcMap := map[string]interface{}{
		"BizMethod1": BizMethod1,
		"BizMethod2": BizMethod2,
		"BizMethod3": BizMethod3,
	}

	if f, ok := funcMap[method]; ok {
		fmt.Printf("Recv: ok")
		if fn, ok := f.(func(context.Context, interface{}) (interface{}, error)); ok {
			return fn(ctx, request)
		}
	} else {
		return "{\"message\": \"error finding method\"", nil
	}
	return
}

// BizMethod1 implements the BizServiceImpl interface.
func BizMethod1(ctx context.Context, request interface{}) (response interface{}, err error) {

	m := request.(string)
	var j *http.BizRequest
	json.Unmarshal([]byte(m), &j)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("Recv: %v\n", j.Text)

	fmt.Printf("Recv: %v\n", m)
	return "{\"text\":" + "\"" + "BizMethod1 called" + " is your message\"}", nil
}

// BizMethod2 implements the BizServiceImpl interface.
func BizMethod2(ctx context.Context, request interface{}) (response interface{}, err error) {

	m := request.(string)
	var j *http.BizRequest
	json.Unmarshal([]byte(m), &j)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("Recv: %v\n", j.Text)

	fmt.Printf("Recv: %v\n", m)
	return "{\"text\":" + "\"" + "BizMethod2 called" + " is your message\"}", nil
}

// BizMethod3 implements the BizServiceImpl interface.
func BizMethod3(ctx context.Context, request interface{}) (response interface{}, err error) {

	m := request.(string)
	var j *http.BizRequest
	json.Unmarshal([]byte(m), &j)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("Recv: %v\n", j.Text)

	fmt.Printf("Recv: %v\n", m)
	return "{\"text\":" + "\"" + "BizMethod3 called" + " is your message\"}", nil
}
