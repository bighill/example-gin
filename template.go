package main

import "github.com/gin-contrib/multitemplate"

func initRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles(
		"default",
		"templates/default.html",
		"templates/header.html",
		"templates/footer.html",
	)
	r.AddFromFiles(
		"secondary",
		"templates/secondary.html",
		"templates/header.html",
		"templates/footer.html",
	)
	return r
}
