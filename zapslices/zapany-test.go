package zapslices

import "github.com/pydio/cells-linter/zapslices/zap"

func shouldPassZapSlice() {
	var i interface{}
	zap.Any("pass", i)
}

func shouldFailZapSlice() {
	var ss []string
	zap.Any("fail", ss)
}
