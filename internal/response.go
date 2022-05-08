package internal

type ResponseItem struct {
	Value   string
	Success bool
}

type Response struct {
	Count int
	Items []ResponseItem
}

func (r Response) ToMap() (*H, error) {
	results := H{}
	results["count"] = r.Count
	items := make([]H, r.Count)

	for index, item := range r.Items {
		itemMap := H{}
		itemMap["value"] = item.Value
		itemMap["success"] = item.Success
		items[index] = itemMap
	}
	results["items"] = items

	return &results, nil
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
