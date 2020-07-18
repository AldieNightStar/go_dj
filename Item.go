package go_dj

type Item struct {
	Provider ProviderFunc
	Dependencies []string
}