package domain

type TaquillasRequest struct {
	IdSede    int `json:"idSede" example:"1"`
	CodModulo int `json:"codModulo" example:"4000"`
}

type TaquillaResponse struct {
	IdTaquilla    int    `json:"idTaquilla" example:"1"`
	CodTaquilla   string `json:"codTaquilla" example:"TQ01"`
	NomTaquilla   string `json:"nomTaquilla" example:"Taquilla 1"`
	EstadoTaquilla int   `json:"estadoTaquilla" example:"0"`
	NIdModulo     int    `json:"nIdModulo" example:"4000"`
}

type TaquillasResponse []TaquillaResponse
