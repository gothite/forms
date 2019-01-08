package forms

import (
	"mime/multipart"
	"net/url"
)

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
	}

	return nil
}

func (data FormData) GetAll(name string) interface{} {
	return data[name]
}

type MultipartFormData struct{ form *multipart.Form }

func (data MultipartFormData) Get(name string) interface{} {
	var all = data.GetAll(name)

	if values, ok := all.([]string); ok && values != nil {
		return values[0]
	} else if files, ok := all.([]*multipart.FileHeader); ok && files != nil {
		return files[0]
	}

	return nil
}

func (data MultipartFormData) GetAll(name string) interface{} {
	if values, ok := data.form.Value[name]; ok && values != nil {
		return values
	} else if files, ok := data.form.File[name]; ok && files != nil {
		return files
	}

	return nil
}

type Value struct {
	Value interface{}
}

func (value Value) Get(name string) interface{} {
	return value.Value
}

func (value Value) GetAll(name string) interface{} {
	return value.Value
}
