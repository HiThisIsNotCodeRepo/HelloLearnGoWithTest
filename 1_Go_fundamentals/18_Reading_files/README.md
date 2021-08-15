# Reading files

## Test dependency injection

```
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFs(fs)
```

## Production dependency injection

```
	posts, err := blogposts.NewPostsFromFs(os.DirFS("posts"))
```