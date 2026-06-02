package domain

import "time"

// SEDES
type SedeNT struct {
	ID           int       `json:"id" example:"1"`
	Code         string    `json:"code" example:"S01"`
	Name         string    `json:"name" example:"Sede 1"`
	Description  string    `json:"description" example:"Descripción de la sede 1"`
	RegisterDate time.Time `json:"registerDate" example:"1678886400"`
	State        int       `json:"state" example:"1"`
}

type SedesNTResponse struct {
	Status int      `json:"status" example:"200"`
	Data   []SedeNT `json:"data"`
}

type SiteNT struct {
	ID    int    `json:"id" example:"1"`
	Code  string `json:"code" example:"S01"`
	Name  string `json:"name" example:"Sede 1"`
	State int    `json:"state" example:"1"`
}

// SERVICIOS

type ServicioNT struct {
	ID             int      `json:"id" example:"1"`
	Code           string   `json:"code" example:"S01"`
	Name           string   `json:"name" example:"Sede 1"`
	Description    string   `json:"description" example:"Descripción de la sede 1"`
	RegisterDate   string   `json:"registerDate" example:"1678886400"`
	QualifyService bool     `json:"qualifyService" example:"true"`
	MultiCalled    bool     `json:"multiCalled" example:"true"`
	State          int      `json:"state" example:"1"`
	Site           []SiteNT `json:"site"`
}

type ServiciosNTXSedeResponse struct {
	Status int          `json:"status" example:"200"`
	Data   []ServicioNT `json:"data"`
}

// TAQUILLAS

type TaquillaNT struct {
	ID            int    `json:"id" example:"1"`
	Code          string `json:"code" example:"T01"`
	Name          string `json:"name" example:"Taquilla 1"`
	ManagePrority bool   `json:"managePriority" example:"true"`
	Service       string `json:"service" example:"Servicio 1"`
	State         int    `json:"state" example:"1"`
}

type TaquillaNTResponse struct {
	Status int          `json:"status" example:"200"`
	Data   []TaquillaNT `json:"data"`
}

// TAQUILLAS X SEDE X SERVICIO

type TaquillaxSedexServicioRequest struct {
	IdSede     int `json:"idSede" example:"1"`
	IdServicio int `json:"idServicio" example:"1"`
}

type TaquillaxSedexServicioData struct {
	Id             int    `json:"id" example:"1"`
	Name           string `json:"name" example:"Taquilla 1"`
	ManagePriority bool   `json:"managePriority" example:"true"`
}

type TaquillaxSedexServicioResponse struct {
	Status int                          `json:"status" example:"200"`
	Data   []TaquillaxSedexServicioData `json:"data"`
}

//MOTIVOS DE DESCANDO TAQUILLA

type TaquillaMotivoDescansoResponse struct {
	Status int              `json:"status" example:"1"`
	Data   []MotivoDescanso `json:"data"`
}

type MotivoDescanso struct {
	Id           int    `json:"id" example:"1"`
	Name         string `json:"name" example:"Motivo de descanso 1"`
	Description  string `json:"description" example:"Descripción del motivo de descanso 1"`
	RegisterDate string `json:"registerDate" example:"2022-01-01"`
	Type         int    `json:"type" example:"1"`
	State        int    `json:"state" example:"1"`
}

// ESTADOS DE LA TAQUILLA

type Branch struct {
	ID    int `json:"id" example:"1"`
	State int `json:"state"`
}

type Point struct {
	ID             int  `json:"id" example:"1"`
	ManagePriority bool `json:"managePriority"`
}

type User struct {
	ID       int    `json:"id" example:"123"`
	User     string `json:"user" example:"jlopez"`
	LastName string `json:"lastName" example:"Lopez"`
	Name     string `json:"name" example:"Juan"`
}

type TaquillasEstadoRequest struct {
	Action int    `json:"action" example:"1"`
	Branch Branch `json:"branch"`
	Point  Point  `json:"point"`
	User   User   `json:"user"`
}

type TaquillasEstadoResponse struct {
	Status int                    `json:"status" example:"200"`
	Data   TaquillasEstadoRequest `json:"data"`
}

type Turn struct {
	ID int `json:"id" example:"1"`
}

type Service struct {
	ID int `json:"id" example:"1"`
}

type PointOfCare struct {
	ID int `json:"id" example:"1"`
}

type AtencionEstadoRequest struct {
	ID          int         `json:"id" example:"0"`
	State       int         `json:"state" example:"3"`
	User        User        `json:"user"`
	Turn        Turn        `json:"turn"`
	Service     Service     `json:"service"`
	PointOfCare PointOfCare `json:"pointOfCare"`
}

type AtencionEstadoResponse struct {
	Status int `json:"status" example:"200"`
	Data   int `json:"data" example:"1"`
}

type LlamadoTurnoResponse struct {
	Status int  `json:"status"`
	Data   bool `json:"data"`
}

type Response struct {
	Status int      `json:"status"`
	Data   []Reason `json:"data"`
}

type Reason struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Type  int    `json:"type"`
	State int    `json:"state"`
}

// SERVICIOS DISPONIBLES PARA TRANSFERENCIA
type ResponseTransfer struct {
	Status int              `json:"status"`
	Data   []TransferReason `json:"data"`
}

type TransferReason struct {
	ID           int         `json:"id"`
	Branch       IDReference `json:"branch"`
	ServiceOri   IDReference `json:"serviceOri"`
	ServiceDes   ServiceDes  `json:"serviceDes"`
	WaitTime     int         `json:"waitTime"`
	PendingTurns int         `json:"pendingTurns"`
	Enabled      bool        `json:"enabled"`
}

type IDReference struct {
	ID int `json:"id"`
}

type ServiceDes struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

//LLAMDO DE TURNO GET

// LLAMADO DE TURNO POST
type LlamadoTurnoPostRequest struct {
	IdTurn int `json:"idTurn"`
	Point  struct {
		Branch  Branch  `json:"branch"`
		Service Service `json:"service"`
		Point   Point   `json:"point"`
		User    User    `json:"user"`
	} `json:"point"`
	User User `json:"user"`
}

type TurnType struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Patient struct {
	ID        int    `json:"id"`
	PatientID string `json:"patientId"`
}

type ServiceInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LlamadoTurnoPostData struct {
	ID           int         `json:"id"`
	Number       string      `json:"number"`
	Priority     int         `json:"priority"`
	TurnType     TurnType    `json:"turnType"`
	Patient      Patient     `json:"patient"`
	Service      ServiceInfo `json:"service"`
	State        int         `json:"state"`
	Attended     bool        `json:"attended"`
	Transferible bool        `json:"transferible"`
	Finalizable  bool        `json:"finalizable"`
}

type LlamadoTurnoPostResponse struct {
	Status int                  `json:"status"`
	Data   LlamadoTurnoPostData `json:"data"`
}

// TURNOS DISPONIBLES MANUAL

type TurnoConsultaResponse struct {
	Status int                 `json:"status"`
	Data   []TurnoConsultaData `json:"data"`
}

type TurnoConsultaData struct {
	ID           int                   `json:"id"`
	Number       string                `json:"number"`
	Priority     int                   `json:"priority"`
	Date         int64                 `json:"date"`
	StandbyTime  int                   `json:"standbyTime"`
	TurnType     TurnoConsultaTipo     `json:"turnType"`
	Patient      TurnoConsultaPaciente `json:"patient"`
	Service      TurnoConsultaServicio `json:"service"`
	Branch       TurnoConsultaSede     `json:"branch"`
	State        int                   `json:"state"`
	Attended     bool                  `json:"attended"`
	Transferible bool                  `json:"transferible"`
	Finalizable  bool                  `json:"finalizable"`
}

type TurnoConsultaTipo struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type TurnoConsultaPaciente struct {
	ID        int    `json:"id"`
	PatientID string `json:"patientId"`
	LastName  string `json:"lastName"`
	Name      string `json:"name"`
}

type TurnoConsultaServicio struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	QualifyService bool   `json:"qualifyService"`
}

type TurnoConsultaSede struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TRANSFER
type TransferRequest struct {
	IdTurn        int            `json:"idTurn"`
	TurnsMovement []TurnMovement `json:"turnsMovement"`
	User          User           `json:"user"`
}

type TurnMovement struct {
	Turn        Turn         `json:"turn"`
	User        *User        `json:"user"`
	Service     Service      `json:"service"`
	PointOfCare *PointOfCare `json:"pointOfCare"`
	Branch      Branch       `json:"branch"`
	State       int          `json:"state"`
	Active      int          `json:"active"`
	Transfer    bool         `json:"transfer"`
}

type TransferResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
