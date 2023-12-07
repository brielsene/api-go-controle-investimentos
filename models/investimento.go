package models

type Investimento struct {
	Banco     string  `json:"banco"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
	UsuarioId uint    `json:"usuarioId"`
	Usuario   Usuario `json:"usuario"`
}
