package godi

import (
	"fmt"
	"reflect"
)

type IGoDI interface {
	Register(itemType interface{}, factory any, parameters ...interface{}) error
	Unregister(itemType interface{})
	Resolve(itemType interface{}) (interface{}, error)
}

type Container struct {
	factories  map[reflect.Type]interface{}
	parameters map[reflect.Type][]interface{}
}

func New() IGoDI {
	return &Container{
		factories:  make(map[reflect.Type]interface{}),
		parameters: make(map[reflect.Type][]interface{}),
	}
}

func (c *Container) Register(itemType interface{}, factory any, parameters ...interface{}) error {
	structType := reflect.TypeOf(itemType)
	factoryType := reflect.TypeOf(factory)

	if factoryType.NumIn() != len(parameters) {
		return fmt.Errorf("invalid number of parameters, %s expected %d, got %d", factoryType, factoryType.NumIn(), len(parameters))
	}
	for i := 0; i < factoryType.NumIn(); i++ {
		if factoryType.In(i) != reflect.TypeOf(parameters[i]) {
			factoryParamTypes := make([]string, factoryType.NumIn())
			parametersTypes := make([]string, len(parameters))

			for i := 0; i < factoryType.NumIn(); i++ {
				factoryParamTypes[i] = factoryType.In(i).String()
			}
			for i := 0; i < len(parameters); i++ {
				parametersTypes[i] = reflect.TypeOf(parameters[i]).String()
			}

			return fmt.Errorf("invalid parameter type, %s expected %s, got %s", factoryType, factoryParamTypes, parametersTypes)
		}
	}

	if factoryType.NumOut() != 2 {
		return fmt.Errorf("invalid number of return values, %s expected 2, got %d", factoryType, factoryType.NumOut())
	}
	factoryReturnTypes := make([]string, factoryType.NumOut())
	for i := 0; i < factoryType.NumOut(); i++ {
		factoryReturnTypes[i] = factoryType.Out(i).String()
	}

	if factoryType.Out(0).Kind() != reflect.Interface {
		return fmt.Errorf("invalid return type, %s expected [interface{}, error], got %s", factoryType, factoryReturnTypes)
	}
	if !structType.Implements(factoryType.Out(0)) {
		return fmt.Errorf("invalid return type: %s, expected: [interface{}, error], got %s", factoryType.Out(0), factoryReturnTypes)
	}
	if factoryType.Out(1) != reflect.TypeOf((*error)(nil)).Elem() {
		return fmt.Errorf("invalid return type, %s expected [interface{}, error], got %s", factoryType, factoryReturnTypes)
	}

	c.factories[structType] = factory
	c.parameters[structType] = parameters
	return nil
}

func (c *Container) Unregister(itemType interface{}) {
	structType := reflect.TypeOf(itemType)
	delete(c.factories, structType)
	delete(c.parameters, structType)
}

func (c *Container) Resolve(itemType interface{}) (interface{}, error) {
	structType := reflect.TypeOf(itemType)

	factory := c.factories[structType]
	if factory == nil {
		return nil, fmt.Errorf("item not found")
	}

	var params []reflect.Value
	for _, p := range c.parameters[structType] {
		params = append(params, reflect.ValueOf(p))
	}

	return reflect.ValueOf(factory).Call(params)[0].Interface(), nil
}
