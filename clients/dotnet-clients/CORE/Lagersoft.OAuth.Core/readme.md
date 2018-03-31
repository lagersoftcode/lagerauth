# Lagersoft OAuth.


## After Install:
Please check your app settings and add the OAuth configuration section example, fill clientId and SecretKey with the values configured on the server:

```json
 "OAuth": {
    "ClientId": "",
    "OAuthUrl": "https://oauth.lagersoft.com/",
    "SecretKey": "",
    "CookieName": "oauth"
  }
```

you also need to put this line in `Startup.cs` in the startup method:

`OAuthConf.ConfigureOAuthGlobals(Configuration.GetSection("OAuth"));`

And copy the content folder files where appropiate, while making necessary modifications.

## Quick Documentation:

* You can display authorization links with:
  `Lagersoft.OAuth.Utils.OAuthorizedActionLink`
  First Parameter: HttpContext (use as Context from a view, HttpContext from a controller)
  Second Parameter: Html helper (use as Html from a view)
  Rest of the parameters are the same as Html.ActionLink

  Returns: Generated link or null.

  Example:
    @Lagersoft.OAuth.Utils.OAuthorizedActionLink(Context, Html, "Home", "Index")


* To check if a user is logged in:
  `Lagersoft.OAuth.Utils.IsAuthenticated(HttpContext)`
  
  Parameter: HttpContext (use Context from a view, HttpContext from controller)
  Returns: Boolean

  This could be usefull to render diferent views for example in a layout, 
  where you want to render the "login" link if the user is not logged in
  and the "logout" link when the user is logged in.

* To grab the current user object:
  `Lagersoft.OAuth.Utils.GetUser(HttpContext)`
   Parameter: HttpContext (use Context from view, HttpContext from controller)
   Returns: 
  
  This will return a User object, composed from email, and its token.

Example:
```csharp
@if (Lagersoft.OAuth.Utils.IsAuthenticated(Context))
{
    <ul class="nav navbar-nav navbar-right">
        <li class="dropdown">
            <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
                @Lagersoft.OAuth.Utils.GetUser(Context).email <span class="caret"></span>
            </a>
            <ul class="dropdown-menu" role="menu">
                <li class="dropdown">
                </li>
                <li class="divider"></li>
                <li>@Html.ActionLink("Log Off", "Logoff", "OAuth")</li>

            </ul>
        </li>

    </ul>
}
else
{
    <ul class="nav navbar-nav navbar-right">
        <li>@Html.ActionLink("Log in", "Login", "OAuth", routeValues: null, htmlAttributes: new { id = "loginLink" })</li>
    </ul>
}
```

*** The rest of the package strives to be as non-intrusive as posible. ***
