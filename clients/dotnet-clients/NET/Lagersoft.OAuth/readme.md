# Lagersoft OAuth.


After Install:
Please check your web.config and change the required appSettings keys.


Quick Documentation:

* You can display authorization links with:
  `Lagersoft.Utils.OAuthorizedActionLink(HttpContextBase, Html, ....)`
  First Parameter: HttpContextBase (use as Context from a view, HttpContext from a controller)
  Second Parameter: Html helper (use as Html from a view)
  Rest of the parameters are the same as Html.ActionLink

  Returns: Generated link or null.

  Example:
    @Lagersoft.OAuth.Utils.OAuthorizedActionLink(Context, Html, "Home", "Index")


* To check if a user is logged in:
  `Lagersoft.OAuth.Utils.IsAuthenticated(HttpContextBase)`
  
  Parameter: HttpContextBase (use Context from a view, HttpContext from controller)
  Returns: Boolean

  This could be usefull to render diferent views for example in a layout, 
  where you want to render the "login" link if the user is not logged in
  and the "logout" link when the user is logged in.

* To grab the current user object:
  `Lagersoft.OAuth.Utils.GetUser(HttpContextBase)`
   Parameter: HttpContextBase (use Context from view, HttpContext from controller)
   Returns: 
  
  This will return a User object, composed from email, and its token.

Example:
```csharp
  @if (Lagersoft.OAuth.Utils.IsAuthenticated(Context))
  {
    using (Html.BeginForm("LogOff", "OAuth", FormMethod.Post, new { id = "logoutForm", @class = "navbar-right" }))
    {
    @Html.AntiForgeryToken()

    <ul class="nav navbar-nav navbar-right">
			<li class="dropdown">
				<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
					@Lagersoft.OAuth.Utils.GetUser(Context).Email <span class="caret"></span>
				</a>
				<ul class="dropdown-menu" role="menu">
					<li class="dropdown">
					</li>
					<li class="divider"></li>
					<li><a href="javascript:document.getElementById('logoutForm').submit()">Log off</a></li>

				</ul>
			</li>

    </ul>
    }
  }
  else
  {
    <ul class="nav navbar-nav navbar-right">
        <li>@Html.ActionLink("Log in", "Login", "OAuth", routeValues: null, htmlAttributes: new { id = "loginLink" })</li>
    </ul>
  }
```

*** The rest of the package strives to be as non-intrusive as posible. ***
