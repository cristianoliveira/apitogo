# apitogo
<p align="center" >
<img src="https://cdn.rawgit.com/cristianoliveira/apitogo/9112716a/apitogo.svg?q=1" width="150" align="center" />
<strong>"An api to go, please."</strong>
</p>
<p align="center">Make a fake api without a line of code for development purposes</p>


**Work in progress**

### I wanna help! 
You can help this project by:

 - Taking a look on the desired features below and implementing it.  
 
 - Help/suggestions are appreciated. Feel free to open an issue with any suggestion.

 - Using and giving feedback
 
# Motivation

There are two motivations that made me create this project:

  - Prototyping SPAs.
  
    Each time that I need to create a SPA for an API that was designed but still not implemented I need to create my on server to provide the data needed to test.

  - Third Part APIs running locally
  
    Each time that I have a project that relies on third part apis I have problems to run it locally.

That's why I am working on this app so next time that I need some simple API. I gonna just get an api to go.

# Usage

It implements a bunch of endpoits by default for you to use.

## Json Api
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

That's it! Here is your api to go, sir. Try out:

```bash
curl -XGET http://localhost:8080/posts
```

Or

```bash
curl -XGET http://localhost:8080/posts/1
```


It runs by default at port `8080` and for each json file inside the folder
it is going to create a endpoint like this:

   - `/posts` returns all data from inside the file
   - `/posts/:id` returns an object from inside the file


## Authentication API

It implements a basic oauth2 server also you can use for development

   - `/authorize` for webclients authorizations
   - `/token` for token requests

Default client_id: `1234` and client_secret: `apitogo1234`

##

# Future implementations

 - Json API
    - [x] Json endpoint from files
    - [x] Filter by id
    - [ ] Query by parameters
    - [ ] Restfull Api (GET, POST, PUT, DELETE)
    - [ ] Sort
    - [ ] Limit
    - [ ] Follow json:api standards
    - [ ] Graphql support

 - Authentication API
    - [x] Oauth2 (token)
    - [ ] Basic (user/password)
    - [ ] Custom Clients ID/Keys
    - [ ] Login Page
    
 - Websocket API
    - [ ] Create/connect channels

 - Customization
    - [ ] Routes
    - [ ] Schema

 - Distribution
    - [ ] Installation pack for brew
    - [ ] Installation Script for linux

 - Others
    - [ ] Travis

## Contributing

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
