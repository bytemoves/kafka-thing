package main

import {
	fiber
}


type Comment struct{
	Text string `form:"text" json:"text"`

} 

func main () {
	app := fiber.New()

	api := app.Group("api/v1")
	api.Post("/comments",createComment)
	app.Listen(":3000")

}


func createComment(c *fiber.Ctx) error{
	cmt := new(Comment)
	if err := c.BodyParse(cmt); err != nil{
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success":false,
			"message":err,
		})
		return err
	}

	cmtInBytes,err := json.Marshal(cmt)
	PushCommentToQueue("comments",cmtInBytes)

	err = c.JSON(&fiber.Map{
		"success":true,
		"message":comment pushed succesfully,
		"comment":cmt,

	})

	if err != nil{
		c.Status(500).JSON(&fiber.Map{
			"success":false,
			"message":"Error creating prod",
		})
		return err
	}

	return err
}