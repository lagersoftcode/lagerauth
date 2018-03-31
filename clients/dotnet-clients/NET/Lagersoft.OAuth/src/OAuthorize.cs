using System.Web;
using System.Web.Mvc;

namespace Lagersoft.OAuth
{
	public class OAuthorize : AuthorizeAttribute
	{
		protected override bool AuthorizeCore(HttpContextBase httpContext)
		{

			if (!Utils.IsAuthenticated(httpContext)) return false;

			var token = Utils.GetJWT(httpContext);

			var routeData = httpContext.Request.RequestContext.RouteData;
			var currentController = routeData.GetRequiredString("controller").ToLower();
			var currentAction = routeData.GetRequiredString("action").ToLower();
			var currentMethod = httpContext.Request.HttpMethod.ToLower();

			return Utils.Can(token, currentController, currentAction, currentMethod);
		}

		protected override void HandleUnauthorizedRequest(AuthorizationContext filterContext)
		{
			filterContext.Controller.TempData["error"] = "You are not authorized, contact your administrator or ";
			filterContext.Controller.TempData["uri"] = filterContext.HttpContext.Request.Url.PathAndQuery;

			filterContext.Result = new RedirectResult("/OAuth");
		}
	}
}
