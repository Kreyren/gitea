// DNM-LEGAL(Krey): License
// DNM-DOCS(Krey): Needs docs

package org

import (
	"fmt"

	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
)

const (
	tplDiscussionsHome base.TplName = "org/discussion/home"
	tplDiscussionsNew  base.TplName = "org/discussion/new"
)

func DiscussionHome(ctx *context.Context) {
	ctx.HTML(200, tplDiscussionsHome)
}

// DNM(Krey): Terrible name
func NewDiscussion(ctx *context.Context) {
	ctx.HTML(200, tplDiscussionsNew)
}

// POST: Creates new discussion
func CreateNewDiscussion(ctx *context.Context) {
	fmt.Println("WHEEE")
	return
}
