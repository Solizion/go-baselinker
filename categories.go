package baselinker

import (
	"encoding/json"
	"net/http"
)

type CategoryStorageId struct {
	StorageId string `json:"storage_id" validate:"required,max=30"`
}

type GetCategoriesListResponse struct {
	*BaseResponse
	Categories []Category `json:"categories"`
}

type GetCategoriesListParameters struct {
	*CategoryStorageId
}

type AddEditCategoryResponse struct {
	*BaseResponse
	CategoryId int `json:"category_id"`
}

type EditCategoryParameters struct {
	*CategoryStorageId
	*EditCategory
}

type AddCategoryParameters struct {
	*CategoryStorageId
	*AddCategory
}

type addCategoryParameters struct {
	*CategoryStorageId
	*Category
}

// Documentation: https://api.baselinker.com/index.php?method=getCategories
func (baseLinker *BaseLinker) GetCategories(parameters GetCategoriesListParameters) ([]Category, Error) {
	var (
		response GetCategoriesListResponse
	)

	err := baseLinker.validator.Struct(parameters)
	if nil != err {
		return response.Categories, NewSimpleError(err)
	}

	requestForm := baseLinker.createRequestForm("getCategories", parameters)
	resp, err := http.PostForm(baseLinker.Url, requestForm)

	if nil != err {
		return response.Categories, NewSimpleError(err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)

	if nil != err {
		return response.Categories, NewSimpleError(err)
	}

	if !response.IsSuccess() {
		return response.Categories, response
	}

	return response.Categories, nil
}

func (baseLinker *BaseLinker) AddCategory(parameters AddCategoryParameters) (int, Error) {
	err := baseLinker.validator.Struct(parameters)
	if nil != err {
		return 0, NewSimpleError(err)
	}

	addCategoryParameters := addCategoryParameters{
		Category: &Category{
			BaseCategory: parameters.AddCategory.BaseCategory,
		},
		CategoryStorageId: parameters.CategoryStorageId,
	}

	return baseLinker.addCategory(addCategoryParameters)
}

func (baseLinker *BaseLinker) EditCategory(parameters EditCategoryParameters) (int, Error) {
	err := baseLinker.validator.Struct(parameters)
	if nil != err {
		return 0, NewSimpleError(err)
	}

	addCategoryParameters := addCategoryParameters{
		Category: &Category{
			Id:           parameters.EditCategory.Id,
			BaseCategory: parameters.EditCategory.BaseCategory,
		},
		CategoryStorageId: parameters.CategoryStorageId,
	}

	return baseLinker.addCategory(addCategoryParameters)
}

// Documentation: https://api.baselinker.com/index.php?method=addCategory
func (baseLinker *BaseLinker) addCategory(parameters addCategoryParameters) (int, Error) {
	var (
		response AddEditCategoryResponse
	)

	requestForm := baseLinker.createRequestForm("addCategory", parameters)
	resp, err := http.PostForm(baseLinker.Url, requestForm)

	if nil != err {
		return response.CategoryId, NewSimpleError(err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)

	if nil != err {
		return response.CategoryId, NewSimpleError(err)
	}

	if !response.IsSuccess() {
		return response.CategoryId, response
	}

	return response.CategoryId, nil
}
