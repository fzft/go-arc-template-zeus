# Golang Web Template (Zeus)

This is a template for creating web applications in Go. It includes the Gin Web Framework and various database technologies for data storage and migration.

## Technologies Used

- **Gin**: A high-performance web framework for Go. It makes it simple to build robust APIs and web servers.
- **golang-migrate**: This tool is used to handle database schema migration. It can work with all the databases mentioned above.

## Getting Started

To get started with this template, you will need to have Go installed on your machine. You should also have the necessary databases installed and configured.

Follow these steps to get up and running:

1. **Clone the repository**
    ```bash
    git clone https://github.com/username/golang-web-template.git
    cd golang-web-template
    ```

2. **Install dependencies**
    ```bash
    go mod tidy
    ```

3. **Set up your database**

   You will need to create a database and modify the connection string in the configuration file.

4. **Run migrations**

   Use golang-migrate to run migrations. For example, if you're using MySQL, you can migrate up using this command:

    ```bash
   make migrate-up
    ```

5. **Run the application**
    ```bash
    go run main.go
    ```

Your application should now be running!

## Naming Conventions

This template uses the following naming conventions:

- **ENV**: Environment variables should be in all caps and use underscores to separate words. For example, `DB_MYSQL_XXX`.
- **CONFIG**: Configuration variables should be in all caps and use underscores to separate words. For example, `db.mysql.xxx`.
  - **DB** 
    1. Database configuration variables should be prefixed with `db`.
    2. Database configuration variables should be in lowercase and use underscores to separate words. For example, `db.mysql.xxx`.
    3. Database configuration variables should be fixed type name. For example, `db.mysql.host` and `db.postgres.host`.
  
- **DATABASE**: Database tables should be in all lowercase and use underscores to separate words. For example, `users`.


## Contributing

Contributions are welcome. Please submit a pull request or create an issue if you have something you would like to add or change.