package domain

import "time"

// @Description Estructura que representa un request de un endpoint de la API
type SedeNT struct {
	ID           int       `json:"id" example:"1"`
	Code         string    `json:"code" example:"S01"`
	Name         string    `json:"name" example:"Sede 1"`
	Description  string    `json:"description" example:"Descripción de la sede 1"`
	RegisterDate time.Time `json:"registerDate" example:"1678886400"`
	State        int       `json:"state" example:"1"`
}

// @Description Estructura que representa un response de un endpoint de la API
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
