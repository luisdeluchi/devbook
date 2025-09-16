# Devbook

Devbook é uma aplicação web composta por duas partes principais: uma API backend em Go e um frontend WebApp também em Go, que juntos simulam uma rede social simplificada para fins de estudo e demonstração.

## Estrutura do Projeto

- **API/**: Backend REST em Go, responsável por autenticação, gerenciamento de usuários, publicações e interações.
  - `main.go`: Ponto de entrada da API.
  - `src/`: Código-fonte organizado em módulos:
    - `autenticacao/`: Lógica de autenticação e geração de tokens.
    - `banco/`: Conexão e manipulação do banco de dados.
    - `controllers/`: Handlers das rotas (login, usuários, publicações).
    - `middlewares/`: Middlewares de autenticação e autorização.
    - `modelos/`: Modelos de dados (usuário, publicação, etc).
    - `repositorios/`: Acesso e manipulação de dados no banco.
    - `respostas/`: Padronização de respostas HTTP.
    - `router/`: Definição de rotas e inicialização do servidor.
    - `seguranca/`: Funções de segurança (hash de senhas, etc).

- **WebApp/**: Frontend em Go que serve páginas HTML, CSS e JS, consumindo a API.
  - `main.go`: Inicialização do servidor web.
  - `src/`: Organização similar à API, com controllers, middlewares, modelos, etc.
  - `public/`: Arquivos estáticos (HTML, CSS, JS, imagens).
  - `views/`: Templates HTML para as páginas da aplicação.

## Como rodar o projeto

1. **Pré-requisitos:**
   - Go instalado (versão 1.18+ recomendada)
   - Banco de dados configurado (ex: MySQL)

2. **API:**
   - Acesse a pasta `API/`.
   - Configure as variáveis de ambiente ou arquivos de configuração em `src/config/`.
   - Execute:
     ```sh
     go run main.go
     ```

3. **WebApp:**
   - Acesse a pasta `WebApp/`.
   - Configure as variáveis de ambiente ou arquivos de configuração em `src/config/`.
   - Execute:
     ```sh
     go run main.go
     ```

4. **Acesse o sistema:**
   - WebApp: http://localhost:porta_configurada
   - API: http://localhost:porta_configurada

## Padrões e Convenções

- Organização modular por domínio (usuários, publicações, autenticação).
- Uso de middlewares para autenticação e autorização.
- Separação clara entre lógica de negócio (controllers), acesso a dados (repositorios) e modelos.
- Templates HTML reutilizáveis em `views/templates/`.
- Comunicação entre WebApp e API via HTTP (fetch/AJAX).

## Contribuição

Pull requests são bem-vindos! Siga o padrão de organização do projeto e adicione testes sempre que possível.

## Licença

Este projeto é open-source para fins educacionais e foi desenvolvido atravéz do curso https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa.
