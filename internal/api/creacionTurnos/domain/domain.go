package domain

// @Description Estructura que representa un request de un endpoint de la API
type CreacionTurnosRequest struct {
	// Incluir descripción del campo
	CreacionTurnos string `json:"creacionTurnos" example:"'Valor_de_ejemplo'"`
}

// @Description Estructura que representa un response de un endpoint de la API
type CreacionTurnosResponse struct {
	// Incluir descripción del campo
	CreacionTurnos string `json:"creacionTurnos" example:"'Valor_de_ejemplo'"`
}

// @Description Descripción de un conjunto de elementos CreacionTurnosResponse
type CreacionTurnosesResponse []CreacionTurnosResponse
