# apitogo
An api to go, please. Generate a fake api without a line of code.


# Usage

Create a collection file inside the folder, for instance, `posts.json`
```json
{
  "posts": [
    { "id": 1, "title": "Some post" },
    { "id": 2, "title": "Awesome post" },
    { "id": 3, "title": "Another post" }
  ]
}
```

Run the apitogo in the same folder:
```bash
apitogo run
```

It runs by default at port 8080 and for each json file inside the folder
it is going to create a endpoint like that:

   - `/posts` returns all data from inside the file
   - `/posts/id` returns a object from inside the file

# Licence

  MIT


