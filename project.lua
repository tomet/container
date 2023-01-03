return {
	name = "container",
	source_files = {"/**/*.go"},
	targets = {
		run = "go test .",
		compile = "go build .",
	}
}