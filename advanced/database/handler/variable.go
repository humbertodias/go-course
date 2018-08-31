package handler

import "html/template"

//ModeloLocal armazena o modelos de pagina Local
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
