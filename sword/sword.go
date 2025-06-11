package sword

import (
	"strings"

	"github.com/GoAdminGroup/themes/common"
	"github.com/GoAdminGroup/themes/sword/resource"
	adminTemplate "github.com/wangxx2026/go-admin/template"
	"github.com/wangxx2026/go-admin/template/components"
	"github.com/wangxx2026/go-admin/template/types"
)

type Theme struct {
	ThemeName string
	components.Base
	*common.BaseTheme
}

var Sword = Theme{
	ThemeName: "sword",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: TemplateList,
		},
	},
	BaseTheme: &common.BaseTheme{
		AssetPaths:   resource.AssetPaths,
		TemplateList: TemplateList,
	},
}

func init() {
	adminTemplate.Add("sword", &Sword)
}

func Get() *Theme {
	return &Sword
}

func (t *Theme) Name() string {
	return t.ThemeName
}

func (t *Theme) GetTmplList() map[string]string {
	return TemplateList
}

func (t *Theme) GetAssetList() []string {
	return resource.AssetsList
}

func (t *Theme) GetAsset(path string) ([]byte, error) {
	path = strings.Replace(path, "/assets/dist", "assets/dist", -1)
	return resource.AssetFS.ReadFile(path)
}
