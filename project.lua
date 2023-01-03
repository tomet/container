return {
	name = "container",
	source_files = {"/**/*.go"},
	targets = {
		test = "go test .",
		compile = "go build .",
	}
}