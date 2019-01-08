package forms

import (
	"mime/multipart"
	"net/url"
	"testing"
)

func TestJSON(test *testing.T) {
	var data = map[string]interface{}{"id": 1, "city": "Boston"}
	var json = JSON(data)

	if id := json.Get("id"); id != data["id"] {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", data["id"])
		test.Fatalf("Actual: %s", id)
	}

	if city := json.GetAll("city"); city != data["city"] {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", data["city"])
		test.Fatalf("Actual: %s", city)
	}
}

func TestFormData(test *testing.T) {
	var data = url.Values{"id": {"1"}, "city": {"Boston"}}
	var form = FormData(data)

	if id := form.Get("id"); id != data["id"][0] {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", data["id"][0])
		test.Fatalf("Actual: %s", id)
	}

	if value := form.Get("undefined"); value != nil {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: nil")
		test.Fatalf("Actual: %s", value)
	}

	if city := form.GetAll("city").([]string); len(city) == len(data["city"]) && city[0] != data["city"][0] {
		test.Errorf("Incorrect values!")
		test.Errorf("Expected: %s", data["city"])
		test.Fatalf("Actual: %s", city)
	}
}

func TestMultipartFormData(test *testing.T) {
	var data = &multipart.Form{
		Value: map[string][]string{"id": {"1"}},
		File:  map[string][]*multipart.FileHeader{"file": {&multipart.FileHeader{}}},
	}
	var form = MultipartFormData{data}

	if id := form.Get("id"); id != data.Value["id"][0] {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", data.Value["id"][0])
		test.Fatalf("Actual: %s", id)
	}

	if value := form.Get("undefined"); value != nil {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: nil")
		test.Fatalf("Actual: %s", value)
	}

	if file := form.Get("file"); file != data.File["file"][0] {
		test.Errorf("Incorrect values!")
		test.Errorf("Expected: %v", data.File["file"][0])
		test.Fatalf("Actual: %v", file)
	}
}

func TestValue(test *testing.T) {
	var value = "hello"
	var data = Value{value}

	if v := data.Get("id"); v != value {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", value)
		test.Fatalf("Actual: %s", v)
	}

	if v := data.GetAll("undefined"); v != value {
		test.Errorf("Incorrect value!")
		test.Errorf("Expected: %s", value)
		test.Fatalf("Actual: %s", v)
	}
}
