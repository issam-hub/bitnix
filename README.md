# bitnix

## Getting Started
1- Clone this repo:
```
git clone https://github.com/issam-hub/bitnix
cd backend
go mod download
```
2- Setup some required tools: 
- **air**: for hot-reloading, see this [link](https://github.com/air-verse/air) 
-  **go-migrate**: for database migrations, see this [link](https://github.com/golang-migrate/migrate) 
-  **echo-swagger**: for swagger docs generation, see this [link](https://github.com/swaggo/echo-swagger)

3- Create `.env` file, check `.env.example` to get an idea
4- add the environment variable `BITNIX_DSN` that contains the DSN (Data Source Name) for the database
```bash
# .zshrc / .bashrc / .zsh_profile / .bash_profile
BITNIX_DSN="postgres://[db_username]:[db_password]@localhost/[db_name]"
```

## Notice
- To visit swagger docs for the REST API, visit `http://localhost:[the used port]/swagger/index.html`
- Make sure to check the makefile to utilize the automated commands for migration and test
