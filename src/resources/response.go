package resources

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func JsonResponse[D any](message string, code int, status string, data D) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

type MetaPagination struct {
	Page 		int 		`json:"page"`
	Limit    	int    		`json:"limit"`
	Total  		int 		`json:"total"`
}

type StatusPagination struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ResponsePagination struct {
	Status 	StatusPagination 	`json:"status"`
	Meta 	MetaPagination 		`json:"meta"`
	Data 	any  		`json:"data"`
}

func JsonResponsePagination[D any](message string, code int, page int, limit int, total int, data D) ResponsePagination {
	status := StatusPagination{
		Message	: message,
		Code	: code,
	}

	meta := MetaPagination{
		Page	: 	page,
		Limit	:   limit,
		Total	:  	total,
	}

	jsonResponse := ResponsePagination{
		Status	: 	status,
		Meta	: 	meta,
		Data	: 	data,
	}

	return jsonResponse
}
