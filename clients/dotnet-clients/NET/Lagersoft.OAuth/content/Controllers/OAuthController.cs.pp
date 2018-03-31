using System.Web.Mvc;

using OAUtils = Lagersoft.OAuth.Utils;

namespace $rootnamespace$.Controllers
{
	[AllowAnonymous]
	public class OAuthController : Controller
	{
		public ActionResult Index()
		{
			return View();
		}
	
		public ActionResult Login(string returnUrl)
		{
			var uri = HttpContext.Request.Url;
			var url = uri.Scheme + System.Uri.SchemeDelimiter + uri.Authority;
			return Redirect(OAUtils.GetLoginUrl(url + returnUrl));
		}

		public ActionResult OAuth(string code)
		{
			var token = OAUtils.GetToken(code);
			if (token != null)
				OAUtils.Authenticate(HttpContext, token);
			return RedirectToAction("Index", "Home");
		}

		[ValidateAntiForgeryToken]
		[HttpPost]
		public ActionResult Logoff()
		{
			if (OAUtils.IsAuthenticated(HttpContext))
				OAUtils.DeAuthenticate(HttpContext);
			return RedirectToAction("Index", "Home");
		}
	}
}