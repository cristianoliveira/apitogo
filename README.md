# apitogo
<img src="https://cdn.rawgit.com/cristianoliveira/apitogo/9112716a/apitogo.svg?q=1" width="200" align="center"/>

'An api to go, please." Make a fake api without a line of code for development purposes

**Working in progress**

Take a look on the desired features below. Help/suggestion are appreciated. Feel free to open an issue with any suggestion.

# Motivation

There are two motivations that made me create this project:

  - Each time I need to create some SPA for an API that os designed but still
not implemented I need to create my on server to provide the data that I need to test.

  - Each time that I have an project that uses Third Part apis I have problems to run it
locally.

That why I am working on this project so next time that I need some simple API
I going to use this tool to launch one instance to go.

# Usage

Create a collection file inside the folder, for instance, `posts.json` with the follow format:
```json
{
  "data": [
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

And that it! Here are your api to go, sir.

It runs by default at port `8080` and for each json file inside the folder
it is going to create a endpoint like that:

   - `/posts` returns all data from inside the file
   - `/posts/:id` returns an object from inside the file

Try out:

```bash
curl -XGET http://localhost:8080/posts
```
Or
```bash
curl -XGET http://localhost:8080/posts/1
```

# Future implementations

 - API
    - [x] Json endpoint from files
    - [x] Filter by id
    - [ ] Query by parameters
    - [ ] Restfull Api (GET, POST, PUT, DELETE)
    - [ ] Sort
    - [ ] Limit
    - [ ] Other Formats? (xml)
    - [ ] Follow json:api standards

 - Authentication
    - [ ] Basic (user/password)
    - [ ] oauth2 (token)
    - [ ] Custom Clients ID/Keys
    - [ ] Login Page

 - Customization
    - [ ] Routes
    - [ ] Schema

 - Distribution
    - [ ] Installation pack for brew
    - [ ] Installation Script for linux

 - Others
    - [ ] Travis

## Contributing

Any suggestion is appreciated! Feel free to open a issue.

If you want to implement a feature, follow this steps:

 - Fork it!
 - Create your feature branch: `git checkout -b my-new-feature`
 - Commit your changes: `git commit -am 'Add some feature'`
 - Push to the branch: `git push origin my-new-feature`
 - Submit a pull request

Pull Requests are really welcome! Others support also.
**Pull Request should have unit tests**

# License

This project was made under MIT License.
