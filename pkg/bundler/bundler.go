package bundler

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/paradigm-network/paradigm-fn2/common"
	"github.com/paradigm-network/paradigm-fn2/pkg/utils"
)

var funcNames = map[string]string{
	"go":     "/fn2.go",
	"node":   "/fn2.js",
	"ruby":   "/fn2.rb",
	"python": "/fn2.py",
	"php":    "/fn2.php",
	"julia":  "/fn2.jl",
	"java":   "/src/main/java/fn2/Fn2.java",
	"d":      "/fn2.d",
}

var assetsMap = map[string][]string{
	"go": {
		"assets/dockerfiles/fn2/go/Dockerfile",
		"assets/dockerfiles/fn2/go/app.go",
		"assets/dockerfiles/fn2/go/fn2.go",
	},
	"java": {
		"assets/dockerfiles/fn2/java/Dockerfile",
		"assets/dockerfiles/fn2/java/pom.xml",
		"assets/dockerfiles/fn2/java/src/main/java/fn2/Fn2.java",
		"assets/dockerfiles/fn2/java/src/main/java/fn2/app.java",
	},
	"julia": {
		"assets/dockerfiles/fn2/julia/Dockerfile",
		"assets/dockerfiles/fn2/julia/REQUIRE",
		"assets/dockerfiles/fn2/julia/app.jl",
		"assets/dockerfiles/fn2/julia/deps.jl",
		"assets/dockerfiles/fn2/julia/fn2.jl",
	},
	"node": {
		"assets/dockerfiles/fn2/node/Dockerfile",
		"assets/dockerfiles/fn2/node/app.js",
		"assets/dockerfiles/fn2/node/fn2.js",
	},
	"php": {
		"assets/dockerfiles/fn2/php/Dockerfile",
		"assets/dockerfiles/fn2/php/fn2.php",
		"assets/dockerfiles/fn2/php/index.php",
	},
	"python": {
		"assets/dockerfiles/fn2/python/Dockerfile",
		"assets/dockerfiles/fn2/python/app.py",
		"assets/dockerfiles/fn2/python/fn2.py",
	},
	"ruby": {
		"assets/dockerfiles/fn2/ruby/Dockerfile",
		"assets/dockerfiles/fn2/ruby/app.rb",
		"assets/dockerfiles/fn2/ruby/fn2.rb",
	},
	"d": {
		"assets/dockerfiles/fn2/d/Dockerfile",
		"assets/dockerfiles/fn2/d/app.d",
		"assets/dockerfiles/fn2/d/fn2.d",
		"assets/dockerfiles/fn2/d/arsd/cgi.d",
	},
}

func removePrefix(lang string, filename string) (name string) {
	prefix := "assets/dockerfiles/fn2" + "/" + lang + "/"
	return strings.Split(filename, prefix)[1]
}

func isFn2FuncSource(lang string, name string) (ret bool) {
	basename := filepath.Base(name)
	nameWithoutExt := strings.TrimSuffix(basename, filepath.Ext(basename))
	return nameWithoutExt == "fn2" || nameWithoutExt == "Fn2" // Fn2 is for Java
}

//Bundle Prepare a container base image and insert the function body
func Bundle(dir string, lang string, body []byte) (err error) {
	names := assetsMap[lang]
	err = nil
	for _, name := range names {
		data, assetErr := common.Asset(name)
		if assetErr != nil {
			err = assetErr
		}

		name = removePrefix(lang, name)
		targetPath := path.Join(dir, name)

		dir := filepath.Dir(targetPath)
		utils.EnsurerDir(dir)

		if isFn2FuncSource(lang, targetPath) {
			data = body
		}

		writeErr := ioutil.WriteFile(targetPath, data, 0644)
		if writeErr != nil {
			err = writeErr
		}
	}
	return err
}
