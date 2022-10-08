package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
	"time"
)

type movie struct {
	id    int
	title string
}

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	app.Get("/", func(c *fiber.Ctx) error {
		//return c.SendStatus(fiber.StatusOK)
		return c.SendString("Main page")
	})
	app.Get("/text", func(c *fiber.Ctx) error {
		return c.SendString("Hello there!!")
	})
	app.Get("/user-agent", func(c *fiber.Ctx) error {
		ua := c.Get("User-Agent")
		return c.SendString(ua)
	})
	app.Get("/sid", func(c *fiber.Ctx) error {
		return c.Download("./data/sid.png")
	})
	app.Get("/sid2", func(c *fiber.Ctx) error {
		return c.SendFile("./data/sid.png")
	})
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name")
		age := c.Query("age")
		msg := fmt.Sprintf("%s is %s years old", name, age)
		return c.SendString(msg)
	})
	app.Get("/say/:name/:age/", func(c *fiber.Ctx) error {
		name := c.Params("name")
		age := c.Params("age")
		msg := fmt.Sprintf("%s is %s years old", name, age)
		return c.SendString(msg)
	})
	app.Get("/movies", func(c *fiber.Ctx) error {
		movies := map[int]string{1: "Toy story", 2: "The Raid", 3: "Hero",
			4: "Ip Man", 5: "Kung Fu Panda"}
		return c.JSON(movies)
	})
	app.Static("/static/", "./public/index.html")
	app.Get("/now", func(c *fiber.Ctx) error {
		now := time.Now()
		return c.Render("show_date", fiber.Map{
			"now": now.Format("Jan 2, 2006"),
		})
	})
	// GET /api/register
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})
	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})
	// GET /dictionary.txt
	app.Get("/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})
	// GET /john/75
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})
	// GET /john
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})
	log.Fatal(app.Listen(":8080"))
}
