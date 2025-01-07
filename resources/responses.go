package resources

// ErrorResponse formats an error message
func ErrorResponse(message string, err interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "error",
		"message": message,
		"error":   err,
	}
}

// SuccessResponse formats a success message
func SuccessResponse(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}
