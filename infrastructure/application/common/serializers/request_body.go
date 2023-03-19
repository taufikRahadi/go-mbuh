package serializers

import (
	"encoding/json"
	"net/http"
	"quiz-1/business/core"
)

func RequestBody(req *http.Request, dto core.IDto) any {
	requestDto := dto
	json.NewDecoder(req.Body).Decode(&requestDto)

	return requestDto
}
