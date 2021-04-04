package inject

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"
)

type Injector struct {
	singletons map[string]interface{}
	factorys   map[string]func() interface{}
}

var injector = Injector{
	singletons: make(map[string]interface{}),
	factorys:   make(map[string]func() interface{}),
}

var injectorLock = make(chan bool, 1)

// Avoid concurrent uses of Injector. Used by unit test
func LockInjector() {
	injectorLock <- true
}

func ReleaseInjector() {
	<-injectorLock
}

func In(name string) interface{} {
	if object, ok := injector.singletons[name]; ok {
		return object
	}

	if handler, ok := injector.factorys[name]; ok {
		return handler()
	}

	panic(errors.New(fmt.Sprintf("Could not resolve injection named %s", name)))
}

func RegisterSingleton(name string, instance interface{}) {
	log.Printf("Registering singleton %s: %T", name, instance)

	if reflect.ValueOf(instance).Kind() != reflect.Ptr {
		panic(errors.New("Singleton registry has to be a pointer"))
	}
	injector.singletons[name] = instance
}

func RegisterFactory(name string, handler func() interface{}) {
	log.Printf("Registering factory %s: %s", name,
		runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name())
	injector.factorys[name] = handler
}
