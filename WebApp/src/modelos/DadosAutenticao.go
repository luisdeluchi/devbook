package modelos

// DadosAutenticacao representa os dados de autenticação do usuário
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"Token"`
}
