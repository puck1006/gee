package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

//trace 打印出函数调用栈
func traic(message string) string {
	var pcs [32]uintptr

	n := runtime.Callers(3, pcs[:])

	var str strings.Builder

	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t %s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {

				message := fmt.Sprintf("%s", err)

				log.Printf("%s\n\n", traic(message))
				c.Fail(http.StatusInternalServerError, "Internal Serve Error")

			}
		}()

		c.Next()
	}
}
