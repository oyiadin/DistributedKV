package internal

type ResponseItem struct {
	Value   string
	Success bool
}

type Response struct {
	Count int
	Items []ResponseItem
}

func MakeResponseOfFailure(message string) *Response {
	return &Response{
		Count: 1,
		Items: []ResponseItem{
			{
				Value:   message,
				Success: false,
			},
		},
	}
}

func MakeResponseWithOneResult(value string) *Response {
	return &Response{
		Count: 1,
		Items: []ResponseItem{
			{
				Value:   value,
				Success: true,
			},
		},
	}
}

func MakeResponseWithManyResults(values []string) *Response {
	items := make([]ResponseItem, len(values))
	for i, v := range values {
		items[i] = ResponseItem{
			Value:   v,
			Success: true,
		}
	}

	return &Response{
		Count: len(values),
		Items: items,
	}
}
