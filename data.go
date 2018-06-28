package forms

import "net/url"

type Data interface {
	Get(key string) interface{}
	GetAll(key string) interface{}
}

type JSON map[string]interface{}

func (data JSON) Get(name string) interface{} {
	return data[name]
}

func (data JSON) GetAll(name string) interface{} {
	return data[name]
}

type FormData url.Values

func (data FormData) Get(name string) interface{} {
	if values, ok := data[name]; ok && values != nil {
		return values[0]
	} else {
		return nil
	}
}

func (data FormData) GetAll(name string) interface{} {
	return data[name]
}
