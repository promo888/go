package nodes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	t "xcoin.com/v1/types"
)

type Node struct {
	Name string
}

func (node *Node) GetName() string {
	return node.Name
}

func Echo() {
	fmt.Println("i'm static echo from node")
}

func handlePostRequest(ctx *fiber.Ctx) error {
	if string(ctx.Method()) != "POST" {
		return ctx.Status(fiber.StatusMethodNotAllowed).SendString("METHOD NOT ALLOWED")
	}

	msg, err := t.UnmarshalMsgPack(ctx.Body())
	if err != nil {
		ctx.Context().SetStatusCode(fiber.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// fmt.Printf("Received msg: %+v\n", msg)

	// Use the BodyParser to parse the incoming JSON data
	// msg := new(t.Msg)
	// if err := ctx.BodyParser(msg); err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())

	// }
	// // Return a response with the parsed data
	// //return ctx.JSON(msg)
	// //fmt.Printf("Received msg: %+v\n", msg)
	// //return ctx.JSON(msg)
	ctx.JSON(msg)
	//fmt.Printf("Received msg: %+v\n", msg)

	return ctx.SendStatus(fiber.StatusOK) //.SendString("OK")
	//ctx.Status(fiber.StatusOK).SendString("OK") //fiber.NewError(200, "OK") //ctx.SendStatus(200) //ctx.SendString("OK") // ctx.JSON(msg) //nil //ctx.Status(200).SendString("OK")

}

func Start(port string) {
	port = ":" + port
	//log.Default("Starting node... host:port  endpointPortsList[kv]")
	app := fiber.New()
	//app.Server().MaxConnsPerIP = 1 //1 for public, 10 for nodes ?

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go Fiber API")
	})

	app.Post("/", handlePostRequest)

	app.Listen(port) //todo from config
	//log.Fatal(app.Listen(":5000"))
}
