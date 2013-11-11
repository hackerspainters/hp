package auth

import (
	"net/http"
	"html/template"
	"path"

	"hp/conf"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var login = template.Must(template.ParseFiles(
		path.Join(conf.Config.ProjectRoot, "templates/_base.html"),
		path.Join(conf.Config.ProjectRoot, "templates/login.html"),
	))

	type templateData struct {
		Context *conf.Context
	}

	data := templateData{conf.DefaultContext(conf.Config)}

	login.Execute(w, data)
}
