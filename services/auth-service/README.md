# 🔐 Auth Service

Este serviço é responsável por autenticação e autorização de usuários via JWT.

## 📋 Funcionalidades

- Registro e login de usuários
- Geração e validação de tokens JWT
- Middleware para verificação de autorização

## 📈 Diagrama de Arquitetura

```mermaid
---
config:
  flowchart:
    curve: linear
---

flowchart TD

    subgraph "🔐 auth-service"

        register[📝 /register]:::route
        

        login[🔐 /login]:::route
        
        

        profile[👤 /profile]:::route
       

        edit["✏️ /edit/{id}"]:::route
        
        

        delete["🗑️ /delete/{id}"]:::route
        
        logout[🚪 /logout]:::route
        

        register ==> |✅| db
        profile ==> |✅| db
        edit ==> |⛔| db
        delete ==> |⛔| db
        login ==> |✅| db
        logout ==> |⛔| jwt
        login ==> |✅| jwt
        db[(🗃️ Users DB)]:::db
        jwt[(🔑 JWT Service)]:::external

    end

    classDef route fill:#1f77b4,stroke:#ffffff,stroke-width:2px,color:#fff;
    classDef db fill:#343a40,stroke:#ffffff,stroke-width:2px,color:#fff;
    classDef external fill:#ffc107,stroke:#ffffff,stroke-width:2px,color:#000;

```

## Endpoints

| ID | Função       | Método  | Endpoint         | Descrição                                                                 | Requisito       |
|----|--------------|---------|------------------|---------------------------------------------------------------------------|-----------------|
| 1  | 📝 Cadastro  | POST    | `/register`      | Registra um novo usuário no sistema.                                      | ❌ Nenhum       |
| 2  | 🔐 Login     | POST    | `/login`         | Realiza o login do usuário e retorna um **token JWT** para autenticação.  | ❌ Nenhum       |
| 3  | 👤 Perfil    | GET     | `/profile`       | Retorna as informações do perfil do usuário.                              | ✅ **JWT Token** |
| 4  | ✏️ Editar    | POST    | `/edit/{id}`     | Atualiza os dados de um usuário com base no id fornecido.                 | ✅ **JWT Token** |
| 5  | 🗑️ Deletar   | DELETE  | `/delete/{id}`   | Remove o usuário identificado pelo id do sistema.                         | ✅ **JWT Token** |
| 6  | 🚪 Deslogar  | POST    | `/logout`        | Invalida o token JWT atual, efetivando o logout do usuário.               | ✅ **JWT Token** |


## 🗃️ Tabelas do Banco de Dados

### Users

| Campo       | Tipo   | Tags GORM                          | Descrição                     |
|-------------|--------|------------------------------------|-------------------------------|
| 🆔 ID       | uint   | `gorm:"primaryKey"`               | Identificador único           |
| 🕒 CreatedAt| Time   |                                    | Data de criação               |
| 🕒 UpdatedAt| Time   |                                    | Data de atualização           |
| 🗑️ DeletedAt| Time   | `gorm:"index"`                    | Data de exclusão lógica       |
| 📝 Name     | string | `gorm:"not null"`                 | Nome do usuário               |
| 📧 Email    | string | `gorm:"not null;unique"`          | Email do usuário              |
| 🔒 Password | string | `gorm:"not null;size:255"`        | Senha do usuário              |