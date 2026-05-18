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

// @Description Estructura que representa un tipo de documento
type TipoDocumentoResponse struct {
	IdTipoDoc  int    `json:"idTipoDoc" example:"1"`
	CodTipoDoc string `json:"codTipoDoc" example:"CC"`
	NomTipoDoc string `json:"nomTipoDoc" example:"Cédula de Ciudadanía"`
}

// @Description Listado de tipos de documento
type TipoDocumentosesResponse []TipoDocumentoResponse

// @Description Estructura que representa un request para obtener módulos
type ModuloRequest struct {
	IdSede int `json:"idSede" example:"1"`
}

// @Description Estructura que representa un módulo
type ModuloResponse struct {
	IdSede       int    `json:"idSede" example:"1"`
	CodigoModulo int    `json:"codigoModulo" example:"4000"`
	NombreModulo string `json:"nombreModulo" example:"Modulo de Prueba"`
}

// @Description Listado de módulos
type ModulosResponse []ModuloResponse

// @Description Estructura que representa la respuesta de la base de datos para módulos
type ModuloDB struct {
	Id          int    `json:"id"`
	IdSede      int    `json:"idSede"`
	JsonModulos string `json:"jsonModulos"`
}

// @Description Estructura que representa un request para obtener compañías
type CompaniaRequest struct {
	IdSede int `json:"idSede" example:"-1"`
}

// @Description Estructura que representa una compañía
type CompaniaResponse struct {
	IdCompania     int    `json:"idCompania" example:"1"`
	CodigoCompania string `json:"codigoCompania" example:"C01"`
	NombreCompania string `json:"nombreCompania" example:"Compañía de Prueba"`
}

// @Description Listado de compañías
type CompaniasResponse []CompaniaResponse

// @Description Estructura que representa la configuración de compañías
type ConfigCompaniasResponse struct {
	ManejaEmpresas int `json:"manejaEmpresas" example:"1"`
}

// @Description Estructura que representa un request para obtener tipos de servicio
type TipoServicioRequest struct {
	CodigoModulo int `json:"codigoModulo" example:"4000"`
	IdSede       int `json:"idSede" example:"1"`
	IdCompania   int `json:"idCompania" example:"-1"`
}

// @Description Estructura que representa un tipo de servicio
type TipoServicioResponse struct {
	IdServicio     int    `json:"idServicio" example:"1"`
	CodigoServicio string `json:"codigoServicio" example:"S01"`
	NombreServicio string `json:"nombreServicio" example:"Servicio de Prueba"`
}

// @Description Listado de tipos de servicio
type TipoServiciosResponse []TipoServicioResponse

// @Description Estructura que representa una sede
type SedeResponse struct {
	IdSede  int    `json:"idSede" example:"1"`
	CodSede string `json:"codSede" example:"S01"`
	NomSede string `json:"nomSede" example:"Sede Central"`
}

// @Description Listado de sedes
type SedesResponse []SedeResponse

// @Description Estructura que representa un request para obtener tipos de turno
type TipoTurnoRequest struct {
	IdServicio int `json:"idServicio" example:"1"`
	IdSede     int `json:"idSede" example:"1"`
}

// @Description Estructura que representa un tipo de turno
type TipoTurnoResponse struct {
	IdTipoTurno  int    `json:"idTipoTurno" example:"1"`
	CodTipoTurno string `json:"codTipoTurno" example:"T01"`
	NomTipoTurno string `json:"nomTipoTurno" example:"Turno Normal"`
}

// @Description Listado de tipos de turno
type TipoTurnosResponse []TipoTurnoResponse
