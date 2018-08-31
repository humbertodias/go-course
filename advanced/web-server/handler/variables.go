package handler

import "html/template"

// Models armazena os modelos de pagina que serao executados pelos manipuladores
var Models = template.Must(template.ParseFiles("html/hello.html"))
