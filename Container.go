package go_dj

type Container struct {
	items map[string]Item
	cached map[string]Any
}

type Any interface {}
type ProviderFunc func(args ... Any) Any

func (c *Container) Provide(name string) (object Any, err error) {
	item, exist := c.items[name]
	if !exist {
		return nil, newError("No such item in Container: " + name)
	}
	// If Item is already cached, then return it instead of searching it
	cached, exist := c.cached[name]
	if exist {
		return cached, nil
	}
	// If item has dependencies, then aggregate them.
	if len(item.Dependencies) > 0 {
		var dependencies []Any
		// Get each dependency name and call Provide(depName) for each.
		// If some error, then return nil, newError(...)
		for i := 0; i < len(item.Dependencies); i++ {
			dependencyName := item.Dependencies[i]
			dependency, err := c.Provide(dependencyName)
			if err != nil {
				return nil, newError("Can't get dependency for " + name + " <" + dependencyName + ">\nCause: "  + err.Error())
			}
			c.cached[dependencyName] = dependency
			dependencies = append(dependencies, dependency)
		}
		// Now lets call the provider and give all the dependencies it needed
		provided := item.Provider(dependencies...)
		c.cached[name] = provided
		return provided, nil
	} else {
		// If item has no dependencies
		provided := item.Provider()
		c.cached[name] = provided
		return provided, nil
	}
}

func (c *Container) GetListOfItems() (list []string) {
	for k := range c.items {
		list = append(list, k)
	}
	return
}

func (c *Container) Register(name string, provider ProviderFunc, dependencies ... string) (err error) {
	_, exist := c.items[name]
	if exist {
		return newError("This item is already exists in Container: " + name)
	}
	var item = Item{Provider: provider, Dependencies: dependencies}
	c.items[name] = item
	return nil
}


var globalContainer *Container = nil

func NewContainer() *Container {
	return &Container{
		items: make(map[string]Item),
		cached: make(map[string]Any),
	}
}

func GlobalContainer() *Container {
	if globalContainer == nil {
		globalContainer = NewContainer()
	}
	return globalContainer
}