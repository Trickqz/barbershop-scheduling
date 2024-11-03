# Sistema de Agendamento de Barbearia

Este é um sistema de agendamento de atendimentos para uma barbearia, desenvolvido em Go utilizando o framework Gin e o ORM GORM. O sistema permite que barbeiros gerenciem seus horários e produtos, além de permitir que clientes agendem atendimentos.

## Funcionalidades

- **Registro de Usuários:** Barbeiros podem se registrar no sistema.
- **Login de Usuários:** Barbeiros podem fazer login e receber um token JWT para autenticação.
- **Agendamento de Atendimentos:** Clientes podem agendar atendimentos, e o sistema verifica a disponibilidade de horários.
- **Gerenciamento de Horários e Produtos:** Barbeiros podem criar e gerenciar seus horários e produtos disponíveis.
- **Envio de E-mails de Confirmação:** O sistema envia e-mails de confirmação para os clientes após o agendamento.

## Tecnologias Utilizadas

- **Linguagem:** Go
- **Framework:** Gin
- **ORM:** GORM
- **Banco de Dados:** PostgreSQL
- **Autenticação:** JWT (JSON Web Tokens)
- **Envio de E-mails:** Gomail

## Pré-requisitos

Antes de executar o projeto, você precisa ter o seguinte instalado:

- Go (versão 1.16 ou superior)
- PostgreSQL
- Git

## Instalação

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/seu_usuario/barbearia-agendamento.git
   cd barbearia-agendamento
   ```

2. **Instale as dependências:**

   ```bash
   go mod tidy
   ```

3. **Configure o banco de dados:**

   Crie um arquivo `.env` na raiz do projeto com a seguinte variável:

   ```plaintext
   DATABASE_URL="postgresql://usuario:senha@localhost:5432/nome_do_banco"
   ```

   Substitua `usuario`, `senha` e `nome_do_banco` pelos valores correspondentes ao seu banco de dados PostgreSQL.

4. **Execute as migrações:**

   O sistema irá automaticamente criar as tabelas necessárias no banco de dados ao iniciar.

## Execução

Para executar o projeto, use o seguinte comando:

```bash
go run main.go
```

O servidor será iniciado na porta `8080`.

## Endpoints

### Registro de Usuário

- **Método:** POST
- **URL:** `/register`
- **Corpo (JSON):**
  ```json
  {
      "nome": "Nome do Barbeiro",
      "email": "email@exemplo.com",
      "senha": "senha123",
      "tipo": "barbeiro"
  }
  ```

### Login de Usuário

- **Método:** POST
- **URL:** `/login`
- **Corpo (JSON):**
  ```json
  {
      "email": "email@exemplo.com",
      "senha": "senha123"
  }
  ```

### Agendar Atendimento

- **Método:** POST
- **URL:** `/agendar`
- **Corpo (JSON):**
  ```json
  {
      "nome": "Nome do Cliente",
      "email": "email@cliente.com",
      "horario": "2023-10-10T11:00:00Z",
      "status": "confirmado"
  }
  ```

### Listar Horários Disponíveis

- **Método:** GET
- **URL:** `/horarios-disponiveis`

### Criar Horário (Protegido)

- **Método:** POST
- **URL:** `/admin/criar-horario`
- **Cabeçalho:** `Authorization: Bearer <seu_token_jwt>`
- **Corpo (JSON):**
  ```json
  {
      "descricao": "Corte de Cabelo",
      "disponivel": true
  }
  ```

### Criar Produto (Protegido)

- **Método:** POST
- **URL:** `/admin/criar-produto`
- **Cabeçalho:** `Authorization: Bearer <seu_token_jwt>`
- **Corpo (JSON):**
  ```json
  {
      "nome": "Cortar Cabelo",
      "preco": 25.50
  }
  ```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a MIT License. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.