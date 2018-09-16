package handler

import "html/template"

//ModeloLocal armazena o modelos de pagina Local
var templateFullPath = "advanced/database/html/local.html"
var ModeloLocal = template.Must(template.ParseFiles(templateFullPath))
