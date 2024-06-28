package view

import (
	"context"
	"fmt"
	"go-starter/internal/model"

)

func GetSessionAttributes(ctx context.Context) model.SessionAttributes {
	if sessAttr, ok := ctx.Value("user").(model.SessionAttributes); ok {
    fmt.Printf("THIS IS THE SESS ATTR: %v", sessAttr)
		return sessAttr
	}
	return model.SessionAttributes{}
}

