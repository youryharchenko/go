package wg_respose

import (
	"fmt"
)

func R200(data, text string) []byte {
	var d = "{}"
	var t = "OK"
	if data != "" {
		d = data
	}
	if text != "" {
		t = text
	}
	json := fmt.Sprintf(`{"status": 200, "statusText": "%s", "data": %s}`, t, d)
	fmt.Println("R200 json: ", json)
	return []byte(json)
}

func R201(data, text string) []byte {
	var d = "{}"
	var t = "Created"
	if data != "" {
		d = data
	}
	if text != "" {
		t = text
	}
	json := fmt.Sprintf(`{"status": 201, "statusText": "%s", "data": "%s"}`, t, d)
	return []byte(json)
}

func R202(data, text string) []byte {
	var d = "{}"
	var t = "Accepted"
	if data != "" {
		d = data
	}
	if text != "" {
		t = text
	}
	json := fmt.Sprintf(`{"status": 202, "statusText": "%s", "data": "%s"}`, t, d)
	return []byte(json)
}

func R204(data, text string) []byte {
	var d = "{}"
	var t = "No Content"
	if data != "" {
		d = data
	}
	if text != "" {
		t = text
	}
	json := fmt.Sprintf(`{"status": 204, "statusText": "%s", "data": "%s"}`, t, d)
	return []byte(json)
}

func R400(data, text string) []byte {
	var d = "{}"
	var t = "Bad Request"
	if data != "" {
		d = data
	}
	if text != "" {
		t = text
	}
	json := fmt.Sprintf(`{"status": 400, "statusText": "%s", "data": "%s"}`, t, d)
	return []byte(json)
}

func R401(data, text string) []byte {
	var d = "{}"
	var t = "Unauthorized"
	if data != "" {
		d = data
	}
	if text != "" {
		t = text
	}
	json := fmt.Sprintf(`{"status": 401, "statusText": "%s", "data": "%s"}`, t, d)
	return []byte(json)
}

func R403(data, text string) []byte {
	var d = "{}"
	var t = "Forbidden"
	if data != "" {
		d = data
	}
	if text != "" {
		t = text
	}
	json := fmt.Sprintf(`{"status": 403, "statusText": "%s", "data": "%s"}`, t, d)
	return []byte(json)
}

func R404(data, text string) []byte {
	var d = "{}"
	var t = "Not Found"
	if data != "" {
		d = data
	}
	if text != "" {
		t = text
	}
	json := fmt.Sprintf(`{"status": 404, "statusText": "%s", "data": "%s"}`, t, d)
	return []byte(json)
}
