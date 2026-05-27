package domain

type TaquillasRequest struct {
	IdSede    int `json:"idSede" example:"1"`
	CodModulo int `json:"codModulo" example:"4000"`
}

type TaquillaResponse struct {
	IdTaquilla     int    `json:"idTaquilla" example:"1"`
	CodTaquilla    string `json:"codTaquilla" example:"TQ01"`
	NomTaquilla    string `json:"nomTaquilla" example:"Taquilla 1"`
	EstadoTaquilla int    `json:"estadoTaquilla" example:"0"`
	NIdModulo      int    `json:"nIdModulo" example:"4000"`
}

type TaquillasResponse []TaquillaResponse

type ServiciosSigaRequest struct {
	IdSede    int `json:"idSede" example:"1"`
	CodModulo int `json:"codModulo" example:"4000"`
}

type ServicioSigaResponse struct {
	IdServicio  int    `json:"idServicio" example:"1"`
	CodServicio string `json:"codServicio" example:"SV01"`
	NomServicio string `json:"nomServicio" example:"Servicio 1"`
}

type ServiciosSigaResponse []ServicioSigaResponse

type PacienteTurno struct {
	NomTipoDocumento string `json:"nomTipoDocumento"`
	NumeroDocumento  string `json:"numeroDocumento"`
	Nombre1          string `json:"nombre1"`
	Nombre2          string `json:"nombre2"`
	Apellido1        string `json:"apellido1"`
	Apellido2        string `json:"apellido2"`
	Sexo             int    `json:"sexo"`
	FechaCumple      string `json:"fechaCumple"`
}

type TurnosDisponiblesRequest struct {
	IdTaquilla int `json:"idTaquilla"`
	IdServicio int `json:"idServicio"`
	IdSede     int `json:"idSede"`
}

type TurnoDisponibleResponse struct {
	IdSede          int           `json:"idSede"`
	NombreSede      string        `json:"nombreSede"`
	IdServicio      int           `json:"idServicio"`
	CodServicio     string        `json:"codServicio"`
	NomServicio     string        `json:"nomServicio"`
	IdTipoTurno     int           `json:"idTipoTurno"`
	CodTipoTurno    string        `json:"codTipoTurno"`
	NomTipoTurno    string        `json:"nomTipoTurno"`
	NumeroTurno     string        `json:"numeroTurno"`
	FechaHoraTurno  string        `json:"fechaHoraTurno"`
	IdTipoDocumento int           `json:"idTipoDocumento"`
	IdEmpresa       int           `json:"idEmpresa"`
	CodEmpresa      string        `json:"codEmpresa"`
	NomEmpresa      string        `json:"nomEmpresa"`
	NumeroOrden     string        `json:"numeroOrden"`
	MinutosEspera   int           `json:"minutosEspera"`
	Prioridad       int           `json:"nPrioridad"`
	Tiempo          int           `json:"tiempo"`
	Paciente        PacienteTurno `json:"paciente"`
}

type TurnosDisponiblesResponse []TurnoDisponibleResponse

type SedesUsuarioRequest struct {
	Usuario string `json:"usuario"`
}

type SedeUsuarioResponse struct {
	IdSede  int    `json:"idSede"`
	CodSede string `json:"codSede"`
	NomSede string `json:"nomSede"`
}

type SedesUsuarioResponse []SedeUsuarioResponse

type ConsumirCredencialesRequest struct {
	Usuario    string `json:"usuario"`
	Contrasena string `json:"contrasena"`
}

type ConsumirCredencialesResponse struct {
	IdUsuario int `json:"idUsuario"`
}
