package helper

import (
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func ParamInt(params httprouter.Params, key string) int {
	var (
		id  int
		err error
	)

	paramId := params.ByName(key)
	if id, err = strconv.Atoi(paramId); err != nil {
		panic(err)
	}

	return id
}
