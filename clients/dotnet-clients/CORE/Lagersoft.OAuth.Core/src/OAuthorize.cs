using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using System;

namespace Lagersoft.OAuth
{
	public class OAuthorize : Attribute, IResourceFilter, IOrderedFilter
	{
		public int Order => 1000;

		public void OnResourceExecuting(ResourceExecutingContext context)
		{
			var httpContext = context.HttpContext;
			if (!Utils.IsAuthenticated(httpContext))
			{
				context.Result = new RedirectResult("/OAuth/Unauthorized");
				return;
			}

			var token = Utils.GetJWT(httpContext);

			var routeData = context.RouteData;
			var currentController = routeData.Values["controller"].ToString().ToLower();
			var currentAction = routeData.Values["action"].ToString().ToLower();
			var currentMethod = httpContext.Request.Method.ToLower();

			var can = Utils.Can(token, currentController, currentAction, currentMethod);
			if (!can)
				context.Result = new RedirectResult("/OAuth/Unauthorized");
		}

		public void OnResourceExecuted(ResourceExecutedContext context)
		{
			// Do nothing special after request.
		}
	}
}
