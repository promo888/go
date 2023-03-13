// package main

// import (
// 	"fmt"

// 	//"db"

// 	"xcoin.com/db"
// 	"xcoin.com/nodes"
// 	//"xcoin.com/node"
// 	//"xcoin.com/node"
// )

// func main() {
// 	fmt.Println("i'm main")
// 	db.Echo()
// 	nodes.Echo()
// 	n := &nodes.Node{Name: "Name"}
// 	fmt.Println("Node name: ", n.GetName())
// 	fmt.Println(n.Name)
// 	//&node.Node{Name: "test "}

// }

///////////////////////////////////////////////////

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// //If the struct variable names does not match with json attributes
// //then you can define the json attributes actual name after json:attname as shown below.
// type User struct {
// 	Name string `json:"name"`
// 	Job  string `json:"job"`
// }

// func main() {

// 	//Create user struct which need to post.
// 	user := User{
// 		Name: "Test User",
// 		Job:  "Go lang Developer",
// 	}

// 	//Convert User to byte using Json.Marshal
// 	//Ignoring error.
// 	body, _ := json.Marshal(user)

// 	//Pass new buffer for request with URL to post.
// 	//This will make a post request and will share the JSON data
// 	resp, err := http.Post("https://reqres.in/api/users", "application/json", bytes.NewBuffer(body))

// 	// An error is returned if something goes wrong
// 	if err != nil {
// 		panic(err)
// 	}
// 	//Need to close the response stream, once response is read.
// 	//Hence defer close. It will automatically take care of it.
// 	defer resp.Body.Close()

// 	//Check response code, if New user is created then read response.
// 	if resp.StatusCode == http.StatusCreated {
// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			//Failed to read response.
// 			panic(err)
// 		}

// 		//Convert bytes to String and print
// 		jsonStr := string(body)
// 		fmt.Println("Response: ", jsonStr)

// 	} else {
// 		//The status is not Created. print the error.
// 		fmt.Println("Get failed with error: ", resp.Status)
// 	}
// }

///////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	//"go/types"
	//"log"

	n "xcoin.com/nodes"
)

func main2() {
	fmt.Println("+++")
	//fmt.Println(types.MSG.MsgType)
	////types.Init()
	///fmt.Println(t.MSG.MsgType)
	fmt.Println("+++")

	//newMsgType := types.NewMsgType() //new(types.Msg)
	// fmt.Println(newMsgType)
	//fmt.Println(*newMsgType)
	// fmt.Println(&newMsgType)

	// newMsg := types.NewMsg(types.MSG_TYPE.Msg) //new(types.Msg)
	// newMsg.Sender = "testSender"
	// fmt.Println(&newMsg)
	// fmt.Println(*newMsg)
	// fmt.Println(newMsg.MsgType)
	// fmt.Println(&newMsg.MsgType)
	// fmt.Println("+++")

	var msgBytes []byte
	//// newMsg2 := types.NewMsg(types.MSG_TYPE.Transaction, "testSender1", "testSig1", &msgBytes)
	// fmt.Println(&newMsg2)
	// fmt.Println(*newMsg2)
	// fmt.Println(newMsg2.MsgType)
	//fmt.Println(&newMsg2.MsgType)
	fmt.Println("+++")
	//fmt.Println(types.NewMsg(0)) //todo to strict only enum values

	//fmt.Println(*types.Msg2JsonS(newMsg2))

	// app := fiber.New()
	// app.Server().MaxConnsPerIP = 1 //1 for public, 10 for nodes

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Welcome to Go Fiber API")
	// })
	//app.Listen(":500")
	//log.Fatal(app.Listen(":5000"))
	v1.echo()
	os.Exit(0)
	n.Start("55555") //todo init|update NetworkSettings
}

func version_test(t Test)
