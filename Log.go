// Copyright (c) 2020 Digital Transaction Limited
// All Rights Reserved.
//

package dtl_util

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var logIndent = 0
var logMutex sync.RWMutex

func indentText() string {
	logMutex.RLock()
	level := logIndent
	logMutex.RUnlock()

	return strings.Repeat("  ", level)
}

func LogInit(path string) {
	if path != "" {
		if f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666); err == nil {
			log.SetOutput(f)
		}
	}

	logMutex.Lock()
	logIndent = 0
	logMutex.Unlock()
}

func LogBegin(function string) {
	log.Printf("%s%s BEGIN", indentText(), function)
	LogIndent()
}

func LogEnd(function string, result ...interface{}) {
	LogUnindent()
	s := fmt.Sprintf("%s%s END", indentText(), function)
	if len(result) != 0 {
		s += fmt.Sprintf(" => (")
		for i, item := range result {
			if i != 0 {
				s += fmt.Sprintf(", ")
			}

			if text, ok := item.(string); ok {
				s += fmt.Sprintf("%q", text)
			} else if err, ok := item.(error); ok {
				s += fmt.Sprintf("error{%v}", err)
			} else {
				s += fmt.Sprintf("%v", item)
			}
		}
		s += fmt.Sprintf(")")
	}
	log.Print(s)
}

func Log(format string, a ...interface{}) {
	log.Printf(indentText()+format, a...)
}

func LogIndent() {
	logMutex.Lock()
	logIndent += 1
	logMutex.Unlock()
}

func LogUnindent() {
	logMutex.Lock()
	if logIndent > 0 {
		logIndent -= 1
	}
	logMutex.Unlock()
}
