package actions

import (
	"fmt"

	"github.com/lenfree/pos/models"
	uuid "github.com/satori/go.uuid"
)

func (as *ActionSuite) Test_ItemsResource_List() {
	uuid1, _ := uuid.NewV4()
	uuid2, _ := uuid.NewV4()
	categories := models.Categories{
		{Name: "tires", Description: "all types of tires", ID: uuid1},
		{Name: "bolts", Description: "all types of bolts", ID: uuid2},
	}
	for _, t := range categories {
		err := as.DB.Create(&t)
		as.NoError(err)
	}
	items := models.Items{
		{
			Name:        "kumho 2x2",
			Description: "kumho 2x2",
			CategoryID:  categories[0].ID,
			Price:       12.5,
		},
		{
			Name:        "full face - Shoei",
			Description: "full face Shoei",
			CategoryID:  categories[1].ID,
			Price:       7.25,
		},
	}
	for _, t := range items {
		err := as.DB.Create(&t)
		as.NoError(err)
	}

	res := as.JSON("/api/v1/items").Get()
	body := res.Body.String()
	for _, t := range items {
		as.Contains(body, fmt.Sprintf("%s", t.Name))
		as.Contains(body, fmt.Sprintf("%s", t.Description))
		as.Contains(body, fmt.Sprintf("%s", t.CategoryID))
		as.Contains(body, fmt.Sprintf("%g", t.Price))
	}
	c, err := as.DB.Count(items)
	as.NoError(err)
	as.Equal(2, c)
}

func (as *ActionSuite) Test_ItemsResource_Show() {
	uuid1, _ := uuid.NewV4()
	categories := models.Categories{
		{ID: uuid1, Name: "bolts", Description: "all types of bolts"},
	}
	for _, t := range categories {
		verrs, err := as.DB.ValidateAndCreate(&t)
		as.NoError(err)
		as.False(verrs.HasAny())
	}

	items := models.Items{
		{
			ID:          uuid.UUID{1},
			Name:        "kumho 2x2",
			Description: "kumho 2x2",
			CategoryID:  uuid1,
			Price:       12.5,
		},
	}
	for _, t := range items {
		verrs, err := as.DB.ValidateAndCreate(&t)
		as.NoError(err)
		as.False(verrs.HasAny())
	}
	res := as.JSON("/api/v1/items/%s", items[0].ID).Get()
	as.Contains(res.Body.String(), fmt.Sprintf("%s", items[0].ID))
	as.Contains(res.Body.String(), fmt.Sprintf("%s", items[0].Name))
	as.Contains(res.Body.String(), fmt.Sprintf("%s", items[0].Description))
	as.Contains(res.Body.String(), fmt.Sprintf("%g", items[0].Price))
	as.Contains(res.Body.String(), fmt.Sprintf("%s", items[0].CategoryID))
}

func (as *ActionSuite) Test_ItemsResource_New() {
	res := as.JSON("/api/v1/categories/new").Get()
	as.Equal(501, res.Code)
	as.Contains(res.Body.String(), "not implemented")
}

func (as *ActionSuite) Test_ItemsResource_Create() {
	uuid1, _ := uuid.NewV4()
	item := &models.Item{
		Name:        "kumho 2x2",
		Description: "kumho 2x2",
		CategoryID:  uuid1,
		Price:       12.5,
	}
	res := as.JSON("/api/v1/items").Post(item)
	err := as.DB.First(item)
	as.NoError(err)
	as.NotZero(item.ID)
	as.Contains(res.Body.String(), fmt.Sprintf("%s", item.ID))
	as.Contains(res.Body.String(), fmt.Sprintf("%s", item.Name))
	as.Contains(res.Body.String(), fmt.Sprintf("%s", item.Description))
	as.Contains(res.Body.String(), fmt.Sprintf("%g", item.Price))
}

func (as *ActionSuite) Test_ItemsResource_Edit() {
	res := as.JSON("/api/v1/categories/new").Get()
	as.Equal(501, res.Code)
	as.Contains(res.Body.String(), "not implemented")
}

func (as *ActionSuite) Test_ItemsResource_Update() {
	uuid1, _ := uuid.NewV4()
	uuid2, _ := uuid.NewV4()
	item := &models.Item{
		Name:        "kumho 2x2",
		Description: "kumho 2x2",
		CategoryID:  uuid1,
		Price:       12.5,
	}
	verrs, err := as.DB.ValidateAndCreate(item)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.JSON("/api/v1/items/%s", item.ID).Put(&models.Item{
		ID:          item.ID,
		Name:        "Pirelli",
		Description: "Pirelli tyres",
		Price:       19,
		CategoryID:  uuid2,
	})
	as.Equal(200, res.Code)

	err = as.DB.Reload(item)
	as.NoError(err)
	as.Equal("Pirelli", item.Name)
	as.Equal("Pirelli tyres", item.Description)
	as.Equal(float32(19), item.Price)
	as.Equal(uuid2, item.CategoryID)
}

func (as *ActionSuite) Test_ItemsResource_Destroy() {
	uuid1, _ := uuid.NewV4()
	uuid2, _ := uuid.NewV4()
	items := models.Items{
		{
			Name:        "kumho 2x2",
			Description: "kumho 2x2",
			ID:          uuid1,
			Price:       12.5,
		},
		{
			Name:        "full face - Shoei",
			Description: "full face Shoei",
			ID:          uuid2,
			Price:       7.25,
		},
	}
	for _, t := range items {
		err := as.DB.Create(&t)
		as.NoError(err)
	}

	res := as.JSON("/api/v1/items/%s", uuid1).Delete()
	body := res.Body.String()
	as.Contains(body, fmt.Sprintf("%s", items[0].ID))
	as.Contains(body, fmt.Sprintf("%s", items[0].Name))
	as.Contains(body, fmt.Sprintf("%s", items[0].Description))
	c, err := as.DB.Count(items)
	as.NoError(err)
	as.Equal(1, c)
}
