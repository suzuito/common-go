package cgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/cgcp"
)

const ctxGCPResource = "gcp_ctx_resource"

func GetGCPResource(ctx *gin.Context) (*cgcp.GCPContextResource, error) {
	v, err := getCtxVariable(ctx, ctxGCPResource, &cgcp.GCPContextResource{})
	return v.(*cgcp.GCPContextResource), err
}

func SetGCPResource(ctx *gin.Context, v *cgcp.GCPContextResource) {
	ctx.Set(ctxGCPResource, v)
}
