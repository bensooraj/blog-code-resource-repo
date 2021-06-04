package main

import (
	"borneo/types"
	"flag"
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "server", "supported modes: server, client")
	flag.Parse()

	switch mode {
	case "server":
		blog := types.NewBlog()

		rpc.Register(blog)
		rpc.HandleHTTP()

		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatalln("Error starting the RPC server", err)
		}
	case "client":
		client, err := rpc.DialHTTP("tcp", ":3000")
		if err != nil {
			log.Fatalln("Error creating the RPC client", err)
		}

		var post types.Post
		// Create posts
		err = client.Call("Blog.AddPost", &types.Post{Title: "post 1", Body: "Hello, world!"}, &post)
		if err != nil {
			log.Fatalln("Error creating post", err)
		}
		log.Printf("[AddPost] ID: %d | Title: %s | Body: %s\n", post.ID, post.Title, post.Body)

		client.Call("Blog.AddPost", &types.Post{Title: "post 2", Body: "Hey, there!"}, &post)
		if err != nil {
			log.Fatalln("Error creating post", err)
		}
		log.Printf("[AddPost] ID: %d | Title: %s | Body: %s\n", post.ID, post.Title, post.Body)

		client.Call("Blog.AddPost", &types.Post{Title: "post 3", Body: "Nope, bye!"}, &post)
		if err != nil {
			log.Fatalln("Error creating post", err)
		}
		log.Printf("[AddPost] ID: %d | Title: %s | Body: %s\n", post.ID, post.Title, post.Body)

		// Get post by ID
		err = client.Call("Blog.GetPostByID", 3, &post)
		if err != nil {
			log.Fatalln("Error creating post", err)
		}
		log.Printf("[GetPostByID] ID: %d | Title: %s | Body: %s\n", post.ID, post.Title, post.Body)

	default:
		log.Fatalln("mode must be one of `server` or `client`")
	}
}
