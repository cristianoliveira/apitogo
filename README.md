# apitogo
<img src="https://cdn.rawgit.com/cristianoliveira/apitogo/9112716a/apitogo.svg?q=1" width="200" align="center"/>

'An api to go, please." Make a fake api without a line of code for development purposes

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
   - `/posts/id` returns an object from inside the file

## Contributing

 - Fork it!
 - Create your feature branch: `git checkout -b my-new-feature`
 - Commit your changes: `git commit -am 'Add some feature'`
 - Push to the branch: `git push origin my-new-feature`
 - Submit a pull request

Pull Requests are really welcome! Others support also.

**Pull Request should have unit tests**

# License

This project was made under MIT License.
