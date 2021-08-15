package main

import (
	"log"
	"os"

	blogposts "github.com/qinchenfeng/HelloLearnGoWithTest/1_Go_fundamentals/18_Reading_files"
)

func main() {
	posts, err := blogposts.NewPostsFromFs(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
