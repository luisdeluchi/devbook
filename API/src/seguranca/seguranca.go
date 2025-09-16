package seguranca

import "golang.org/x/crypto/bcrypt"

// hash é uma função que recebe uma senha em texto simples e retorna um hash seguro da senha.
func Hash(senha string) ([]byte, error) {
	// Implementar a lógica de hash da senha
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

}

// VerificarSenha é uma função que compara uma senha em texto simples com um hash de senha armazenado.
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
