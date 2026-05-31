# QuickFeed

QuickFeed é uma plataforma de coleta de feedbacks anônimos onde usuários podem criar empresas, compartilhar páginas públicas e receber opiniões de clientes ou usuários de forma simples.

## Funcionalidades

* Cadastro de usuários
* Login com autenticação JWT
* Criação de empresas
* Geração de páginas públicas através de slugs
* Envio de feedbacks anônimos
* Listagem de feedbacks por empresa
* API REST desenvolvida em Go
* Banco de dados PostgreSQL
* Containerização com Docker

## Tecnologias Utilizadas

### Backend

* Go
* Chi Router
* PostgreSQL
* JWT

### Frontend

* HTML
* CSS
* JavaScript

### Infraestrutura

* Docker
* Docker Compose
* Nginx
* AWS EC2

## Arquitetura


Frontend (HTML/CSS/JS)
         ↓
       Nginx
         ↓
     API em Go
         ↓
    PostgreSQL


## Estrutura do Projeto


backend/
frontend/
nginx/
docker-compose.yml


## Executando o Projeto

### Clonar o repositório


git clone [URL_DO_REPOSITORIO]


### Entrar na pasta do projeto


cd quickfeed


### Subir os containers


docker compose up --build


### Acessar a aplicação


http://[IP da EC2]


## Deploy

O projeto foi publicado em uma instância AWS EC2 utilizando Docker Compose, Nginx e PostgreSQL em containers Docker.

## Variáveis de Ambiente

As variáveis presentes no `docker-compose.yml` são apenas exemplos utilizados para desenvolvimento e demonstração do projeto.

Para ambientes de produção, recomenda-se substituir os valores por credenciais próprias e armazená-los em arquivos `.env` ou serviços de gerenciamento de segredos.

## Rotas da API

### Autenticação


POST /register
POST /login
GET  /profile


### Empresas


POST /companies
GET  /companies


### Feedbacks


POST /feedback
GET  /feedbacks


## Conceitos Estudados

* APIs REST
* Autenticação JWT
* Middleware
* Context em Go
* PostgreSQL
* Relacionamentos em banco de dados
* Docker
* Docker Compose
* Reverse Proxy com Nginx
* Deploy em AWS EC2

## Screenshots

### Login

<img width="1866" height="917" alt="image" src="https://github.com/user-attachments/assets/0fa7cc9c-affa-4edc-a25b-1cde904ea628" />

### Dashboard

<img width="1863" height="916" alt="image" src="https://github.com/user-attachments/assets/fc295b25-4924-4244-a159-05d5ea2fabb8" />


### Página Pública de Feedback

<img width="1867" height="918" alt="image" src="https://github.com/user-attachments/assets/9cb2e502-7213-46ed-bdba-dd06d3241634" />


## Autor

Vinicius Malta
