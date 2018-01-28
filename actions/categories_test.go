package actions

import (
	"fmt"

	"github.com/lenfree/pos/models"
	"github.com/satori/go.uuid"
)

func (as *ActionSuite) Test_CategoriesResource_List() {
	categories := models.Categories{
		{Name: "tires", Description: "all types of tires"},
		{Name: "bolts", Description: "all types of bolts"},
	}
	for _, t := range categories {
		err := as.DB.Create(&t)
		as.NoError(err)
	}

	res := as.JSON("/api/v1/categories").Get()
	body := res.Body.String()
	for _, t := range categories {
		as.Contains(body, fmt.Sprintf("%s", t.Name))
		as.Contains(body, fmt.Sprintf("%s", t.Description))
	}
	c, err := as.DB.Count(categories)
	as.NoError(err)
	as.Equal(2, c)
}

func (as *ActionSuite) Test_CategoriesResource_Show() {
	categories := models.Categories{
		{ID: uuid.UUID{2}, Name: "bolts", Description: "all types of bolts"},
	}
	for _, t := range categories {
		verrs, err := as.DB.ValidateAndCreate(&t)
		as.NoError(err)
		as.False(verrs.HasAny())
	}

	res := as.JSON("/api/v1/categories/%s", categories[0].ID).Get()
	as.Contains(res.Body.String(), fmt.Sprintf("%s", categories[0].ID))
	as.Contains(res.Body.String(), fmt.Sprintf("%s", categories[0].Name))
	as.Contains(res.Body.String(), fmt.Sprintf("%s", categories[0].Description))
}

func (as *ActionSuite) Test_CategoriesResource_New() {
	res := as.JSON("/api/v1/categories/new").Get()
	as.Equal(501, res.Code)
	as.Contains(res.Body.String(), "not implemented")
}

func (as *ActionSuite) Test_CategoriesResource_Create() {
	w := &models.Category{
		Name:        "bolts",
		Description: "all types of bolts",
	}
	res := as.JSON("/api/v1/categories").Post(w)
	err := as.DB.First(w)
	as.NoError(err)
	as.NotZero(w.ID)
	as.Equal("bolts", w.Name)
	as.Contains(res.Body.String(), fmt.Sprintf("%v", w.ID))
	as.Contains(res.Body.String(), fmt.Sprintf("%s", w.Name))
	as.Contains(res.Body.String(), fmt.Sprintf("%s", w.Description))
}

func (as *ActionSuite) Test_CategoriesResource_Edit() {
	res := as.JSON("/api/v1/categories/new").Get()
	as.Equal(501, res.Code)
	as.Contains(res.Body.String(), "not implemented")
}

func (as *ActionSuite) Test_CategoriesResource_Update() {
	category := &models.Category{
		Name:        "bolts",
		Description: "all helmets",
	}
	verrs, err := as.DB.ValidateAndCreate(category)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.JSON("/api/v1/categories/%s", category.ID).Put(&models.Category{
		ID:          category.ID,
		Name:        "Helmet",
		Description: "all helmets",
	})
	as.Equal(200, res.Code)

	err = as.DB.Reload(category)
	as.NoError(err)
	as.Equal("Helmet", category.Name)
	as.Equal("all helmets", category.Description)
}

func (as *ActionSuite) Test_CategoriesResource_Destroy() {
	uuid1, _ := uuid.NewV4()
	uuid2, _ := uuid.NewV4()
	categories := models.Categories{
		{ID: uuid1, Name: "tires", Description: "all types of tires"},
		{ID: uuid2, Name: "bolts", Description: "all types of bolts"},
	}
	for _, t := range categories {
		err := as.DB.Create(&t)
		as.NoError(err)
	}

	res := as.JSON("/api/v1/categories/%s", uuid1).Delete()
	body := res.Body.String()
	as.Contains(body, fmt.Sprintf("%s", categories[0].ID))
	as.Contains(body, fmt.Sprintf("%s", categories[0].Name))
	as.Contains(body, fmt.Sprintf("%s", categories[0].Description))
	c, err := as.DB.Count(categories)
	as.NoError(err)
	as.Equal(1, c)
}
