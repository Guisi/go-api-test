package inject

import (
	"testing"
)

func TestSingletonInjection(t *testing.T) {
	type MyType struct{}
	var myInstance = &MyType{}
	RegisterSingleton("instance", myInstance)
	var myRecoveredInstance = In("instance")
	if myInstance != myRecoveredInstance {
		t.Errorf("Failed to recover singleton: %s and %s should be equal",
			myInstance, myRecoveredInstance)
	}
}

func TestNonSingletonInjection(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Should have panicked when trying to register a non Pointer")
		}
	}()
	type MyType struct{}
	var myInstance = MyType{}
	RegisterSingleton("instance", myInstance)
}

func TestFactory(t *testing.T) {
	expectedContent := "1234"
	RegisterFactory("factory", func() interface{} {
		return expectedContent
	})
	content := In("factory").(string)
	if content != "1234" {
		t.Errorf("Failed to get from factory. Expected %s, got %s",
			expectedContent, content)
	}
}

func TestNonRegisteredInjection(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Should have panicked when trying to resolve a non Register component")
		}
	}()
	In("not-found")
}
