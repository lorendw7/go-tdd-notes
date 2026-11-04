package hello

// Greet returns a greeting for name, defaulting to the world.
func Greet(name string) string {
	if name == "" {
		name = "world"
	}
	return "Hello, " + name + "!"
}
