package handles

import (
	"log"
	"reflect"
)

type (
	GetBeanRequest      struct{}
	GetAppConfigRequest struct{}
)

func Push(reqType reflect.Type) {
	log.Printf("log type: %s", reqType.Name())
	switch reqType.Name() {
	case "GetBeanRequest":
		go func() { sendBeanChannel <- 1 }()
	}
}
