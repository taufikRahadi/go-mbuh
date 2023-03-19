package presenters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quiz-1/business/book"
	"quiz-1/infrastructure/application/common/constants"
	"quiz-1/infrastructure/application/common/serializers"
	"quiz-1/infrastructure/application/dto"
	"strconv"
)

type bookPresenter struct {
	service book.IService
}

func NewBookPresenter(service *book.IService) *bookPresenter {
	return &bookPresenter{
		service: *service,
	}
}

func (bp *bookPresenter) ListBooks(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	var search string = ""
	if len(query["q"]) > 0 {
		search = query["q"][0]
	}
	res, err := bp.service.FindAll(search)

	if err != nil {
		fmt.Println(err)
		serializers.BuildResponse(w, constants.SERVER_ERROR, err, http.StatusInternalServerError)
	}
	serializers.BuildResponse(w, constants.SUCCESS_RETRIEVING_DATA, res, http.StatusOK)
}

func (bp *bookPresenter) StoreBook(w http.ResponseWriter, req *http.Request) {
	storeBookDto := dto.StoreBook{}
	json.NewDecoder(req.Body).Decode(&storeBookDto)

	_, validationError := storeBookDto.Validate()
	if validationError != nil {
		serializers.BuildResponse(w, constants.VALIDATION_ERROR, validationError, http.StatusUnprocessableEntity)
	}

	result, err := bp.service.Store(book.IStoreBook{
		Name:        storeBookDto.Name,
		Description: storeBookDto.Description,
	})

	if err != nil {
		serializers.BuildResponse(w, constants.SERVER_ERROR, err, http.StatusInternalServerError)
	}

	serializers.BuildResponse(w, constants.SUCCESS_CREATING_DATA, result, http.StatusCreated)
}

func (bp *bookPresenter) GetBookById(w http.ResponseWriter, req *http.Request) {
	params := serializers.RequestParams(req)
	parsedId, _ := strconv.Atoi(params["id"])
	result, err := bp.service.FindOne(parsedId)
	if err != nil {
		serializers.BuildResponse(w, constants.VALIDATION_ERROR, err, http.StatusInternalServerError)
	}
	serializers.BuildResponse(w, constants.SUCCESS_RETRIEVING_DATA, result, http.StatusOK)
}

func (bp *bookPresenter) UpdateBook(w http.ResponseWriter, req *http.Request) {
	params := serializers.RequestParams(req)
	id, _ := strconv.Atoi(params["id"])

	updateBookDto := dto.UpdateBook{
		Name:        req.FormValue("name"),
		Description: req.FormValue("description"),
	}
	json.NewDecoder(req.Body).Decode(&updateBookDto)

	_, validationError := updateBookDto.Validate()

	if validationError != nil {
		serializers.BuildResponse(w, constants.VALIDATION_ERROR, validationError, http.StatusUnprocessableEntity)
	}

	res, err := bp.service.Update(id, book.IUpdateBook{
		Name:        updateBookDto.Name,
		Description: updateBookDto.Description,
	})

	if err != nil {
		serializers.BuildResponse(w, constants.SERVER_ERROR, err, http.StatusInternalServerError)
	}

	serializers.BuildResponse(w, constants.SUCCESS_UPDATING_DATA, res, http.StatusOK)
}

func (bp *bookPresenter) DeleteBook(w http.ResponseWriter, req *http.Request) {
	params := serializers.RequestParams(req)
	id, _ := strconv.Atoi(params["id"])

	err := bp.service.Delete(id)
	if err != nil {
		serializers.BuildResponse(w, constants.SERVER_ERROR, err, http.StatusInternalServerError)
	}

	serializers.BuildResponse(w, constants.SUCCESS_DELETING_DATA, true, http.StatusOK)
}
