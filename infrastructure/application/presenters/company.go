package presenters

import (
	"encoding/json"
	"net/http"
	"quiz-1/business/company"
	"quiz-1/infrastructure/application/common/constants"
	"quiz-1/infrastructure/application/common/serializers"
	"quiz-1/infrastructure/application/dto"
	"strconv"
)

type companyPresenter struct {
	service company.IService
}

func NewCompanyPresenter(service *company.IService) *companyPresenter {
	return &companyPresenter{service: *service}
}

func (p *companyPresenter) ListCompany(w http.ResponseWriter, req *http.Request) {
	res, err := p.service.FindAll()
	if err != nil {
		serializers.BuildResponse(w, "Internal Server Error", err, http.StatusInternalServerError)
	} else {
		serializers.BuildResponse(w, "Success retrieving data", res, http.StatusOK)
	}
}

func (p *companyPresenter) StoreCompany(w http.ResponseWriter, req *http.Request) {
	//storeCompanyDto := serializers.RequestBody(req, &dto.StoreCompany{})
	storeCompanyDto := dto.StoreCompany{}
	json.NewDecoder(req.Body).Decode(&storeCompanyDto)
	_, validationError := storeCompanyDto.Validate()

	if validationError != nil {
		serializers.BuildResponse(w, constants.VALIDATION_ERROR, validationError, http.StatusUnprocessableEntity)
	} else {
		var _, err = p.service.Store(
			company.IStoreCompany{
				Name:        storeCompanyDto.Name,
				Description: storeCompanyDto.Description,
				Problem:     storeCompanyDto.Problem,
			},
		)

		if err != nil {
			serializers.BuildResponse(w, constants.SERVER_ERROR, err, http.StatusInternalServerError)
		} else {
			serializers.BuildResponse(w, constants.SUCCESS_CREATING_DATA, true, http.StatusCreated)
		}
	}
}

func (p *companyPresenter) GetCompanyById(w http.ResponseWriter, req *http.Request) {
	params := serializers.RequestParams(req)
	parsedId, _ := strconv.Atoi(params["id"])
	result, err := p.service.FindOne(parsedId)
	if err != nil {
		serializers.BuildResponse(w, constants.SERVER_ERROR, err, http.StatusInternalServerError)
	} else {
		serializers.BuildResponse(w, constants.SUCCESS_RETRIEVING_DATA, result, http.StatusOK)
	}
}

func (p *companyPresenter) UpdateCompany(w http.ResponseWriter, req *http.Request) {
	params := serializers.RequestParams(req)
	id, _ := strconv.Atoi(params["id"])

	updateCompanyDto := dto.UpdateCompany{}
	json.NewDecoder(req.Body).Decode(&updateCompanyDto)

	_, validationError := updateCompanyDto.Validate()
	if validationError != nil {
		serializers.BuildResponse(w, constants.VALIDATION_ERROR, validationError, http.StatusUnprocessableEntity)
	}

	var res, err = p.service.Update(id, company.IUpdateCompany{
		Name:        updateCompanyDto.Name,
		Description: updateCompanyDto.Description,
		Problem:     updateCompanyDto.Problem,
	})

	if err != nil {
		serializers.BuildResponse(w, constants.SERVER_ERROR, err, http.StatusInternalServerError)
	}

	serializers.BuildResponse(w, constants.SUCCESS_UPDATING_DATA, res, http.StatusOK)
}

func (p *companyPresenter) DeleteCompany(w http.ResponseWriter, req *http.Request) {
	params := serializers.RequestParams(req)
	id, _ := strconv.Atoi(params["id"])

	err := p.service.Delete(id)
	if err != nil {
		serializers.BuildResponse(w, constants.SERVER_ERROR, err, http.StatusInternalServerError)
	}

	serializers.BuildResponse(w, constants.SUCCESS_DELETING_DATA, err, http.StatusOK)
}
