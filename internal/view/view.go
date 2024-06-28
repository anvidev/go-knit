package view

import (
	"context"
	"go-starter/internal/model"

)

// FIX: Remove hardcoded user key
func GetSessionAttributes(ctx context.Context) model.SessionAttributes {
	if sessAttr, ok := ctx.Value("user").(model.SessionAttributes); ok {
		return sessAttr
	}
	return model.SessionAttributes{}
}

