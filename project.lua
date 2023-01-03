local module = "github.com/tomet/container"

local packages = {
	module.."/list",
	module.."/slices",
}

packages = table.concat(packages, " ")


return {
	name = "container",
	source_files = {"/**/*.go"},
	targets = {
		test = "go test -v "..packages,
		compile = "go build "..packages,
	}
}