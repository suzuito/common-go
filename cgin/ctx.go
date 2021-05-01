package cgin

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

var errNotFoundCtxVariable = fmt.Errorf("Not found ctx variable")
var errCannotAssignableCtxVariable = fmt.Errorf("Cannot assignable ctx variable")

func getCtxVariable(ctx *gin.Context, key string, targetType interface{}) (interface{}, error) {
	v, exists := ctx.Get(key)
	if !exists {
		return nil, xerrors.Errorf("key '%s' : %w", key, errNotFoundCtxVariable)
	}
	t := reflect.TypeOf(targetType)
	vv := reflect.ValueOf(v)
	if !vv.Type().AssignableTo(t) {
		return nil, xerrors.Errorf("Cannot var '%s' assign type '%s' to '%s' : %w", key, vv.Type().Name(), t.Name(), errCannotAssignableCtxVariable)
	}
	return v, nil
}
