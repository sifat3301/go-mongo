package serializers

//type
type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func (serializer *UserResponse) Response() UserResponse {
	return UserResponse{
		Status:  200,
		Message: serializer.Message,
		Data:    serializer.Data,
	}
}
