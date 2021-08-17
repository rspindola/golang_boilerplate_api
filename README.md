<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/rspindola/golang_api_rest">
    <img src="readme.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Golang Boilerplate Api</h3>

  <p align="center">
    Api rest em Go com autenticação JWT.
    <br />
    <br />
    <a href="https://github.com/rspindola/golang_api_rest">View Demo</a>
    ·
    <a href="https://github.com/rspindola/golang_api_rest/issues">Report Bug</a>
    ·
    <a href="https://github.com/rspindola/golang_api_rest/issues">Request Feature</a>
  </p>
</p>

- [💻Sobre o projeto](#sobre-o-projeto)
- [✨Tecnologias utilizadas](#tecnologias-utilizadas)
- [💻Demonstração da aplicação](#demonstração-da-aplicação)
- [🚀Como executar o projeto](#como-executar-o-projeto)
  - [Pré-requisitos](#pré-requisitos)
  - [Clonando o projeto para sua máquina](#clonando-o-projeto-para-sua-máquina)
  - [Setup](#setup)
    - [Realizando testes](#realizando-testes)
  - [Docker Setup](#docker-setup)
    - [Possiveis problemas](#possiveis-problemas)
    - [Successful Build](#successful-build)
    - [Usando Banco de Dados Usando as ferramentas GUI](#usando-banco-de-dados-usando-as-ferramentas-gui)
    - [Encerrando o aplicativo](#encerrando-o-aplicativo)
    - [Construindo uma imagem Docker de teste para o aplicativo.](#construindo-uma-imagem-docker-de-teste-para-o-aplicativo)
    - [Usando o docker-compose para construir os serviços de aplicativo de teste.](#usando-o-docker-compose-para-construir-os-serviços-de-aplicativo-de-teste)
    - [Realizando os testes](#realizando-os-testes)
- [🔨 Endpoints](#-endpoints)
  - [Login (Login) [POST]](#login-login-post)
  - [ObterUSuários (GetUsers) [GET]](#obterusuários-getusers-get)
  - [CriarUsuários (CreateUser) [POST]](#criarusuários-createuser-post)
  - [AtualizarUsuário (UpdateUser) [PUT]](#atualizarusuário-updateuser-put)
  - [ObterUSuário (GetUser) [GET]](#obterusuário-getuser-get)
  - [RemoverUSuário (DeleteUser) [DELETE]](#removerusuário-deleteuser-delete)
  - [Respostas](#respostas)
- [🌳 Estrutura do projeto](#-estrutura-do-projeto)
- [📝 Licença](#-licença)

---

## 💻Sobre o projeto

Esse repósitorio tem um boilerplate de api usando golang para acelerar futuros projetos.

---

## ✨Tecnologias utilizadas

As seguintes ferramentas foram usadas na construção do projeto:

-   [Go](https://golang.org)
-   [GORM (A Golang ORM)](https://gorm.io/index.html)
-   [JWT](https://github.com/dgrijalva/jwt-go)
-   [Gorilla Mux](https://github.com/gorilla/mux)

---

<a name="demo"></a>

## 💻Demonstração da aplicação

🚧 Em construção... 🚧

---

## 🚀Como executar o projeto

### Pré-requisitos

Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:
[Git](https://git-scm.com), [Golang](https://golang.org).
Além disto é bom ter um editor para trabalhar com o código como [VSCode](https://code.visualstudio.com/).
Também será util ter um [Docker](https://www.docker.com), caso queira trabalhar com containers, caso contrario precisa também de um banco de dados instalado ([MySql](https://www.mysql.com) ou [PostgreSQL](https://www.postgresql.org)).

### Clonando o projeto para sua máquina

```bash
# Clone este repositório
$ git clone https://github.com/rspindola/golang_api_rest.git

# Acesse a pasta do projeto no terminal/cmd
$ cd golang_api_rest
```

---

<a name="setup"></a>

### Setup

Como visto, PostgreSQL é o banco de dados padrão. Se você estiver usando MySQL, simplesmente comente o código do postgres. Observe, não execute o arquivo de composição com o postgres e o mysql sem comentários para evitar quaisquer problemas que possam roubar uma ou duas horas do seu precioso tempo. Portanto, use apenas aquele baseado em seu banco de dados.

Copie o arquivo `.env.example` ou renomei para `.env`

```sh
cp .env.example .env
# or
mv .env.example .env
```

Edite e configure o arquivo `.env`:

```sh
nano .env
```

> PostgreSQL é usado como padrão. Para usar o mysql, comente os detalhes do postgres e descomente o mysql. Ou, para dizer, remova aqueles que não atraem você.

> Observe o `DB_HOST`. Isso é obtido a partir do nome do serviço no arquivo docker-compose.yml. Se você deseja executar este aplicativo em sua máquina local, comente o ativo e descomente o host com o valor: `127.0.0.1`.

Agora estamos prontos para rodar o projeto.

Execute o comando

```go
go run main.go
```

Isso iniciará o aplicativo e executará as `migrations`. Quando completar esse processo e tudo ocorrer bem, deve exibir a mensagem `Listening to port 8000` no terminal e nossa api está pronta para uso.

#### Realizando testes

Na raiz do projeto, rode o comando

```go
go test -v ./...
```

---

<a name="docker-setup"></a>

### Docker Setup

Faça a configração do arquivo `.env` como foi comentado na sessão [setup](#setup).

Agora estamos prontos para criar e iniciar todos os contêineres/serviços listados no docker-compose.

```
docker-compose up
```

Para interromper a execução do aplicativo usando:

```sh
docker-compose down
```

Para executar em segundo plano utilize:

```sh
docker-compose up -d
```

> É aconselhável executar primeiro docker-compose up. Portanto, se houver um problema, você pode vê-los facilmente nos registros e corrigi-los.

#### Possiveis problemas

I) Problema com váriaveis de ambiente: talvez você tenha usado detalhes inválidos no arquivo .env, isso pode interromper sua compilação. Leia atentamente os registros ao executar docker-compose up. Se você identificar algum e corrigir, encerre o processo usando Ctrl + C.

Pare o processo de execução usando:

```sh
docker-compose down --remove-orphans --volumes
```

> Motivo para remover o volume: A primeira compilação já possui volumes criados para você. O volume salvou as credenciais anteriores, portanto, você deve remover esse volume anterior para permitir que os detalhes atualizados sejam refletidos. Após, rode novamente

```sh
docker-compose up --build
```

II) Qualquer outro problema: se você tiver qualquer outro problema, tente ler os registros e corrigi-los. em seguida, siga exatamente o mesmo processo para construir novamente.

#### Successful Build

Quando a construção for bem-sucedida, acesse seu navegador ou carteiro e visite:

```sh
http://localhost:8080
```

#### Usando Banco de Dados Usando as ferramentas GUI

O projeto possui ferramentas como pgadmin (para banco de dados postgres) e phpmyadmin (para banco de dados mysql).
Qualquer que seja o banco de dados usado, permite visualizar o banco de dados a partir dessas ferramentas.

-   pgadmin

    Isso é usado para banco de dados postgres. Se você iniciou seu aplicativo com o banco de dados postgres, expusemos a porta 5050 para o arquivo pgadmin (conforme visto no arquivo docker-compose.yml). Vamos usar essa porta para abrir o pgadmin.

    Vá até o navegador e acesse:

    `http://localhost:5050`

    ![pgadmin](/screenshots/pgadmin01.png)

    Clique com o botão direito nos servidores para criar um novo servidor. Escolha Criar e Servidor:

    ![pgadmin](/screenshots/pgadmin02.png)

    ![pgadmin](/screenshots/pgadmin03.png)

    Preencha o nome que desejar.

    Clique na guia de conexão:

    ![pgadmin](/screenshots/pgadmin04.png)

    Vamos obter o nome do host do banco de dados. No terminal, liste a lista de contêineres em execução:

    ![pgadmin](/screenshots/pgadmin05.png)

    Copie o ID do contêiner full_db_postgres da captura de tela acima, é: 8da1c33e418f

    Use aqui:

    `docker inspect <container_id> | grep IPAddress`

    Obtenha esse endereço IP, que é o nome do host

    ![pgadmin](/screenshots/pgadmin06.png)

    O nome de usuário e a senha são os do seu arquivo .env.

    Com o login bem-sucedido, você verá o banco de dados especificado no seu arquivo .env

-   phpmyadmin

    Se você estiver usando o banco de dados mysql, isso é para você.

    Do arquivo docker-compose.yml, expusemos a porta 9090

    Portanto, para uma construção bem-sucedida, abra em seu navegador:

    `http://localhost:9090`

    ![pgadmin](/screenshots/phpmyadmin01.png)

#### Encerrando o aplicativo

Após o teste bem-sucedido do aplicativo, você pode desligá-lo usando:

```sh
docker-compose down
```

Ou também para remover os volumes:

```sh
docker-compose down --remove-orphans --volumes
```

Se você deseja remover todas as imagens pendentes, execute:

```sh
docker system prune
```

Para remover todas as imagens não utilizadas, não apenas as pendentes, execute:

```sh
docker system prune -a
```

Para remover todas as imagens não utilizadas, não apenas as pendentes e os volumes, execute:

```sh
docker system prune -a --volumes
```

#### Construindo uma imagem Docker de teste para o aplicativo.

Criaremos um arquivo docker personalizado para não utilizar o mesmo Dockefile para teste.

Na raiz do documento rode o comando para criar o arquivo `Dockerfile.test`

```sh
touch Dockerfile.test
```

Após, cole esse conteudo dentro do arquivo.

```dockerfile
FROM golang:1.12-alpine

# Install git

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Run tests
CMD CGO_ENABLED=0 go test -v  ./...
```

#### Usando o docker-compose para construir os serviços de aplicativo de teste.

Lembre-se de que o postgres é usado como padrão. Se você estiver usando o mysql, por favor, comente os detalhes do postgres e descomente os detalhes do mysql.
Crie o arquivo

```sh
touch docker-compose.test.yml
```

```docker
version: '3'

services:
  app_test:
    container_name: full_app_test
    build:
      context: .
      dockerfile: ./Dockerfile.test
    volumes:
      - api_test:/app/src/app/
    depends_on:
      - postgres_test
      # - mysql_test
    networks:
      - fullstack_test

  postgres_test:
    image: postgres:latest
    container_name: full_db_test_postgress
    environment:
      - POSTGRES_USER=${TEST_DB_USERNAME}
      - POSTGRES_PASSWORD=${TEST_DB_PASSWORD}
      - POSTGRES_DB=${TEST_DB_DATABASE}
      - DATABASE_HOST=${TEST_DB_HOST}
    ports:
      - '5555:5432'
    volumes:
      - database_postgres_test:/var/lib/postgresql/data
    networks:
      - fullstack_test

  # mysql_test:
  #   image: mysql:5.7
  #   container_name: full_db_test_mysql
  #   ports:
  #     - 3333:3306
  #   environment:
  #     - MYSQL_DATABASE=${TEST_DB_DATABASE}
  #     - MYSQL_USER=${TEST_DB_USERNAME}
  #     - MYSQL_PASSWORD=${TEST_DB_PASSWORD}
  #     - MYSQL_ROOT_PASSWORD=${TEST_DB_PASSWORD}
  #     - DATABASE_HOST=${TEST_DB_HOST}
  #   volumes:
  #     - database_mysql_test:/var/lib/mysql
  #   networks:
  #     - fullstack_test

volumes:
  api_test:
  database_postgres_test:
  # database_mysql_test:

networks:
  fullstack_test:
    driver: bridge
```

#### Realizando os testes

Certifique-se de que sua configuração de dados de teste postgres/mysql esteja em seu arquivo .env não estejam comentadas. Após isso, rode o comando

```bash
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
```

---

<a name="endpoints"></a>

## 🔨 Endpoints

| Método | Url    | Descrição               |
| ------ | ------ | ----------------------- |
| `POST` | /login | Faz o login no sistema. |

### Login (Login) [POST]

-   Attributes (object)

    -   email: email do usuário (string, required) - limite de 60 caracteres
    -   password: senha do usuário (string, required) - limite de 16 caracteres

-   Request (application/json)

    -   Body

        ```json
        {
            "email": "joane@gmail.com",
            "password": "password"
        }
        ```

-   Response 200 (application/json)

    -   Body

        ```json
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjkxMjIyNjgsInVzZXJfaWQiOjF9.nDpILmMnhO1WUCbJ4QGgdV94lQovYrQ5IdtfmcVNRjg"
        ```

| Método   | Url         | Descrição                                |
| -------- | ----------- | ---------------------------------------- |
| `GET`    | /users      | Retorna a listagem de todos os usuários. |
| `POST`   | /users      | Cadasta um novo usuário.                 |
| `GET`    | /users/{id} | Retorna o usuário.                       |
| `PUT`    | /users/{id} | Atualiza dados de um usuário.            |
| `DELETE` | /users/{id} | Remove um usuário do sistema.            |

### ObterUSuários (GetUsers) [GET]

-   Response 200 (application/json)

    ```json
    [
        {
            "id": 1,
            "nickname": "Steven victor",
            "email": "steven@gmail.com",
            "password": "$2a$10$a0fbw8WHF2q4wGVgOsCwo.xB4EiiHMKNy/29Swf5gH608HKXsSO36",
            "created_at": "2021-08-16T09:43:35-03:00",
            "updated_at": "2021-08-16T09:43:35-03:00"
        },
        {
            "id": 2,
            "nickname": "Martin Luther",
            "email": "luther@gmail.com",
            "password": "$2a$10$fA3oxxsooeCv7Kn/3d/TDeTtigcSgYCuMge7C2L9V5z1OliuMszA2",
            "created_at": "2021-08-16T09:43:35-03:00",
            "updated_at": "2021-08-16T09:43:35-03:00"
        }
    ]
    ```

### CriarUsuários (CreateUser) [POST]

-   Attributes (object)

    -   nickname: nome do usuário (string, required) - limite 60 caracteres
    -   email: email do usuário (string, required) - limite de 60 caracteres
    -   password: senha do usuário (string, required) - limite de 16 caracteres

-   Request (application/json)

    -   Body

        ```json
        {
            "nickname": "joane",
            "email": "joane@gmail.com",
            "password": "password"
        }
        ```

-   Response 201 (application/json)

    -   Body

        ```json
        {
            "id": 3,
            "nickname": "joane",
            "email": "joane@gmail.com",
            "password": "$2a$10$4Y9iRpyNa7AlXVGeVGB2GOyjZQHOhlBDozYB/2PU0/ol6KNJ41S8m",
            "created_at": "2021-08-16T10:14:06.1137698-03:00",
            "updated_at": "2021-08-16T10:14:06.1137698-03:00"
        }
        ```

### AtualizarUsuário (UpdateUser) [PUT]

-   Attributes (object)

    -   nickname: nome do usuário (string, required) - limite 60 caracteres
    -   email: email do usuário (string, required) - limite de 60 caracteres
    -   password: senha do usuário (string, required) - limite de 16 caracteres

-   Request (application/json)

    -   Headers

              Authorization: Bearer [access_token]

    -   Body

        ```json
        {
            "nickname": "Steven victor Updated",
            "email": "steven@gmail.com",
            "password": "password"
        }
        ```

-   Response 200 (application/json)

    -   Body

        ```json
        {
            "id": 1,
            "nickname": "Steven victor Updated",
            "email": "steven@gmail.com",
            "password": "$2a$10$oDgh8KTD1ghTVdyaPMpdHe/oZBLusBNxZNIIWjD1.h6j4eA/addUW",
            "created_at": "2021-08-16T10:50:25-03:00",
            "updated_at": "2021-08-16T11:02:41-03:00"
        }
        ```

### ObterUSuário (GetUser) [GET]

-   Parameters

    -   id (required, number, `1`) ... id do usuário

-   Response 200 (application/json)

    -   Body

        ```json
        [
            {
                "id": 1,
                "nickname": "Steven victor",
                "email": "steven@gmail.com",
                "password": "$2a$10$a0fbw8WHF2q4wGVgOsCwo.xB4EiiHMKNy/29Swf5gH608HKXsSO36",
                "created_at": "2021-08-16T09:43:35-03:00",
                "updated_at": "2021-08-16T09:43:35-03:00"
            }
        ]
        ```

### RemoverUSuário (DeleteUser) [DELETE]

-   Parameters

    -   id (required, number, `1`) ... id do usuário... id do usuário

    -   Headers

              Authorization: Bearer [access_token]

-   Response 204 (application/json)

### Respostas

| Código | Descrição                                                                            |
| ------ | ------------------------------------------------------------------------------------ |
| `200`  | Requisição executada com sucesso (success).                                          |
| `204`  | Requisição de sucesso informando que não contém dados no retorno.                    |
| `400`  | Erros de validação ou os campos informados não existem no sistema.                   |
| `401`  | Dados de acesso inválidos.                                                           |
| `404`  | Registro pesquisado não encontrado (Not found).                                      |
| `405`  | Método não implementado.                                                             |
| `410`  | Registro pesquisado foi apagado do sistema e não esta mais disponível.               |
| `422`  | Dados informados estão fora do escopo definido para o campo.                         |
| `429`  | Número máximo de requisições atingido. (_aguarde alguns segundos e tente novamente_) |

---

<a name="filesystem"></a>

## 🌳 Estrutura do projeto

```
golang_api_rest
 ┣ api
 ┃ ┣ auth
 ┃ ┃ ┗ token.go
 ┃ ┣ controllers
 ┃ ┃ ┣ base.go
 ┃ ┃ ┣ home_controller.go
 ┃ ┃ ┣ login_controller.go
 ┃ ┃ ┣ routes.go
 ┃ ┃ ┗ users_controller.go
 ┃ ┣ middlewares
 ┃ ┃ ┗ middlewares.go
 ┃ ┣ models
 ┃ ┃ ┗ User.go
 ┃ ┣ responses
 ┃ ┃ ┗ json.go
 ┃ ┣ seed
 ┃ ┃ ┗ seeder.go
 ┃ ┣ utils
 ┃ ┃ ┗ formaterror
 ┃ ┃ ┃ ┗ formaterror.go
 ┃ ┗ server.go
 ┣ tests
 ┃ ┣ controllertests
 ┃ ┃ ┣ controller_test.go
 ┃ ┃ ┣ login_controller_test.go
 ┃ ┃ ┗ user_controller_test.go
 ┃ ┗ modeltests
 ┃ ┃ ┣ model_test.go
 ┃ ┃ ┗ user_model_test.go
 ┣ .env.example
 ┣ .gitignore
 ┣ Dockerfile
 ┣ README.md
 ┣ docker-compose.yml
 ┣ go.mod
 ┣ go.sum
 ┗ main.go
```

## 📝 Licença

[MIT](https://choosealicense.com/licenses/mit/)
