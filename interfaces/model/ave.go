package model

// Galinha representa uma ave do tipo galinha
type Galinha interface {
	Cacareja() string
}

// Pata representa uma ave do tipo Pato
type Pata interface {
	Quack() string
}

// Ave representa um
type Ave struct {
	Nome string
}

// Cacareja representa o som feito pro uma galinha
func (a Ave) Cacareja() string {
	return "Cócóricó"
}

// Quack representa o som feito pro um pato
func (a Ave) Quack() string {
	return "Quack, quack..."
}
