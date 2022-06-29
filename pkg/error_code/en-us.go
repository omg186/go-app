package error_code

var enUSText = map[string]string{
	SUCCESS:            "OK",
	FAILURE:            "FAIL",
	NotFound:           "resources not found",
	ServerError:        "Internal server error",
	TooManyRequests:    "Too many requests",
	ParamBindError:     "Parameter error",
	AuthorizationError: "Authorization error",
	RBACError:          "No access",
}
