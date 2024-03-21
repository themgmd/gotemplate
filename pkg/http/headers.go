package http

func PrepareHeaders(inputHeaders ...string) map[string]string {
	headers := make(map[string]string, 2)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	for i := 0; i < len(inputHeaders)-1; i += 2 {
		headers[inputHeaders[i]] = inputHeaders[i+1]
	}

	return headers
}
