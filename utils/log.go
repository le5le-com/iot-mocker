package utils

import (
	"runtime"

	"github.com/rs/zerolog/log"
)

func PrintCallerLog(skip int) {
	callerFuncName := "unknown"

	pc, file, line, ok := runtime.Caller(skip)

	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			callerFuncName = fn.Name()
		}

		log.Debug().Msgf("caller: %s, file: %s, line: %d", callerFuncName, file, line)
	}
}
