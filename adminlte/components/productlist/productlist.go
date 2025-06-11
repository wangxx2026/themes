package productlist

import (
	"html/template"

	adminTemplate "github.com/wangxx2026/go-admin/template"
)

type ProductList struct {
	*adminTemplate.BaseComponent

	Data []map[string]string
}

func New() ProductList {
	return ProductList{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "productlist",
			HTMLData: List["productlist"],
		},
	}
}

func (p ProductList) SetData(value []map[string]string) ProductList {
	p.Data = value
	return p
}

func (p ProductList) GetContent() template.HTML { return p.GetContentWithData(p) }
