using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using OAUtils = Lagersoft.OAuth.Utils;

namespace $rootnamespace$.Controllers
{
	[AllowAnonymous]
	public class OAuthController : Controller
	{
		public IActionResult Login(string returnUrl)
		{
			return Redirect(OAUtils.GetLoginUrl(Request.Scheme + System.Uri.SchemeDelimiter + Request.Host + returnUrl));
		}

		public IActionResult Logoff()
		{
			if (OAUtils.IsAuthenticated(HttpContext))
				OAUtils.DeAuthenticate(HttpContext);

			return RedirectToAction("Index", "Home");
		}

		public new IActionResult Unauthorized()
		{
			return View();
		}

		public IActionResult OAuth(string code)
		{
			var token = OAUtils.GetToken(code);
			if (token != null)
				OAUtils.Authenticate(HttpContext, token);
			return RedirectToAction("Index", "Home");
		}
	}
}