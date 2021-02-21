package model


// Atenção!!!!! use a primeira letra maiuscula no struts e nos atributos
// Pois do contrario eles não pdoeram ser acessados pos outros files
// ficaram em um estado de private 
type Usuario struct {
	Uuid string // Primeira letra maiuscula apra ficar publico e acessivel a todos
	Nome string // Primeira letra maiuscula apra ficar publico e acessivel a todos
	Email string // Primeira letra maiuscula apra ficar publico e acessivel a todos
};