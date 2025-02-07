package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Person struct {
    gorm.Model
    Name string
    Age  int
}

func main() {
    // Iniciar o servidor Fiber
    app := fiber.New()

    // Definindo a rota GET
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    // Definindo a rota POST
    app.Post("/people", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    // Iniciar o servidor na porta 3000
    go func() {
        log.Fatal(app.Listen(":3000"))
    }()

    // Conectar ao banco de dados MySQL
    log.Println("Starting database connection...")
    dsn := "root:root@tcp(127.0.0.1:3306)/sl_dojo?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Realizar migração automática e criar a tabela
    db.AutoMigrate(&Person{})

    // Criar um novo registro de pessoa
    person := Person{
        Name: "Natali Cordeiro",
        Age:  30,
    }
    db.Create(&person)

    log.Println("Person created:", person)
}
