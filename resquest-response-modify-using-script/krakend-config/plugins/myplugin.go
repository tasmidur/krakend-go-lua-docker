package main

import (
    "fmt"
    "errors"
)

func main() {}

func init() {
    fmt.Println("Martian plugin loaded!")
}

var ModifierRegisterer = registerer("martian-plugin")

type registerer string

func (r registerer) RegisterModifiers(f func(name string, factoryFunc func(map[string]interface{}) func(interface{}) (interface{}, error), appliesToRequest bool, appliesToResponse bool)) {
    f(string(r), r.modifyHeaders, true, false)
    fmt.Println(string(r), " registered-piyal!")
}

func (r registerer) modifyHeaders(cfg map[string]interface{}) func(interface{}) (interface{}, error) {
    return func(input interface{}) (interface{}, error) {
        resp, ok := input.(ResponseWrapper)
        if !ok {
            return nil, errors.New("unknown request type")
        }
        // Modify headers here
        resp.Headers()["X-Custom-Header"] = []string{"MyValue"}
        return resp, nil
    }
}

type ResponseWrapper interface {
    Headers() map[string][]string
}