package handler

import (
	"html/template"
)

// Models armazena os modelos de pagina que serao executados pelos manipuladores
var templateFullPath = "advanced/web-server/html/hello.html"
var Models = template.Must(template.ParseFiles(templateFullPath))
