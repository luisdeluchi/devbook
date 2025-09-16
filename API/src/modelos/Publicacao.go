package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicação feitor por um usuario
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorID,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar a publicacao recebida
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}
	publicacao.formatar()
	return nil
}

// Validar valida se os campos estão preenchidos
func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o titulo é obrigatório e não pode ficar em branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("o conteudo é obrigatório e não pode ficar em branco")
	}

	return nil
}

// formatar faz a formatação dos campos
func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
