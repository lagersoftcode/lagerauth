# lagerauthmw

middleware for handling lagerauth trough go

## Usage

### Configuration:
first configure the config for example:

```go
lmw := lagerauthmw.New(ClientID, SecretKey)
```

this initializes default parameters with:

OAuthURL => "https://oauth.lagersoft.com"  
CookieName => "lmwauth"  
MountPoint => "/oauth"  

you can also override defaults with functional options:
```go
lmw := lagerauthmw.New(ClientID, SecretKey, lagerauthmw.CookieNameOpt("cookiecat"), lagerauthmw.OAuthURLOpt("http://mydomain.com"), lagerauthmw.MountPointOpt("/lagerauth"))
```

you should probably read from a file, env or something instead of hardcoding it.

### Routes Registration:
then register the LagerAuthHandler like this:

```go
mux := http.NewServeMux()
// code and routes here

lmw.RegisterLagerauthHandler(mux)

http.ListenAndServe(":80", mux)
```

Login link is `{config.MountPoint}/login` for this example it should be `/oauth/login`

theres a helper `lmw.LoginLink()` so you dont have to worry about it.

### Middleware Usage:

After seting up the lmw struct you can use the `Authorize` middleware to check if user has access.

```go
mux.Handle("/", lmw.Authorize(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "HELLO FRIEND")
}))
```