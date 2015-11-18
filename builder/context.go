package builder

import (
	"github.com/go-xiaohei/pugo-static/model"
	"github.com/go-xiaohei/pugo-static/render"
)

// build context, maintain parse data, posts, pages or others
type Context struct {
	Theme           *render.Theme
	DstDir          string
	Version         builderVersion
	isCopyAllAssets bool
	isSuffixed      bool // generate suffixed url

	Posts         []*model.Post
	PostPageCount int
	Pages         []*model.Page
	indexPosts    []*model.Post // temp posts for index page
	indexPager    *model.Pager

	Tags     map[string]*model.Tag
	tagPosts map[string][]*model.Post

	Navs    model.Navs
	Meta    *model.Meta
	Comment *model.Comment
}

// return global view data for template compilation
func (ctx *Context) ViewData() map[string]interface{} {
	m := map[string]interface{}{
		"Version": ctx.Version,
		"Nav":     ctx.Navs,
		"Meta":    ctx.Meta,
		"Title":   ctx.Meta.Title + " - " + ctx.Meta.Subtitle,
		"Desc":    ctx.Meta.Desc,
		"Comment": ctx.Comment,
	}
	return m
}
