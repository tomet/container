local module = "github.com/tomet/container"

local packages = {
	module.."/slices",
	module.."/list",
}

packages = table.concat(packages, " ")


return {
	name = "container",
	source_files = {"/**/*.go"},
	targets = {
		test = "go test "..packages,
		compile = "go build "..packages,
	}
}