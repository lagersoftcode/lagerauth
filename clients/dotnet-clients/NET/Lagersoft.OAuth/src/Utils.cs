using System;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Web;
using System.Web.Configuration;
using System.Web.Mvc;
using Newtonsoft.Json;

namespace Lagersoft.OAuth
{
	public class Utils
	{
		private static string ApplicationId => WebConfigurationManager.AppSettings["OAApplicationId"];
		private static string ApplicationSecret => WebConfigurationManager.AppSettings["OAApplicationSecret"];
		private static string OAuthUrl => WebConfigurationManager.AppSettings["OAuthUrl"];
		private static string CookieName => WebConfigurationManager.AppSettings["OACookieName"];

		#region helpers
		public static string GetToken(string code)
		{
			var tokenUrl = $"token?client_id={ApplicationId}&client_secret={ApplicationSecret}&code={code}";

			ServicePointManager.ServerCertificateValidationCallback = (sender, cert, chain, sslPolicyErrors) => true;
			var httpClient = new HttpClient()
			{
				BaseAddress = new Uri(OAuthUrl),
				Timeout = new TimeSpan(0, 0, 30)
			};

			var content = new StringContent(string.Empty);
			var result = httpClient.PostAsync(tokenUrl, content).Result;
			var json = result.Content.ReadAsStringAsync().Result;
			dynamic jsonObj = System.Web.Helpers.Json.Decode(json);

			return jsonObj.access_token ?? null;
		}

		public static OAuthModel GetUser(HttpContextBase context)
		{
			if (!IsAuthenticated(context))
				return null;

			var cookie = context.Request.Cookies.Get(CookieName);
			if (cookie == null)
				return null;

			var user = JWT.JsonWebToken.DecodeToObject<OAuthModel>(cookie.Value, ApplicationSecret);
			return user;
		}

		public static string GetJWT(HttpContextBase context)
		{
			if (!IsAuthenticated(context))
				return null;

			var cookie = context.Request.Cookies.Get(CookieName);
			if (cookie == null)
				return null;

			return cookie.Value;
		}

		public static string GetLoginUrl(string host)
		{
			var returnUrl = $"{host}/OAuth/OAuth";
			return $"{OAuthUrl}auth?client_id={ApplicationId}&redirect_uri={HttpUtility.UrlEncode(returnUrl)}";
		}
		#endregion

		#region authentication
		public static bool IsAuthenticated(HttpContextBase context)
		{
			return context.Request.Cookies.AllKeys.Any(x => x == CookieName);
		}

		public static void Authenticate(HttpContextBase context, string token)
		{
			var c = new HttpCookie(CookieName, token) { Expires = DateTime.MaxValue };
			context.Response.Cookies.Add(c);
		}

		public static void DeAuthenticate(HttpContextBase context)
		{
			if (IsAuthenticated(context))
			{
				LogOff(GetJWT(context));

				var c = context.Request.Cookies.Get(CookieName);
				c.Expires = DateTime.Now.AddDays(-1);
				context.Response.Cookies.Add(c);
			}
		}
		#endregion

		#region authorization
		public static bool Can(string token, string controller = "home", string action = "index", string method = "get")
		{
			ServicePointManager.ServerCertificateValidationCallback = (sender, cert, chain, sslPolicyErrors) => true;
			var url = $"{OAuthUrl}can";
			var httpWebRequest = (HttpWebRequest)WebRequest.Create(url);
			httpWebRequest.Headers.Add("Authorization", $"Bearer {token}");
			httpWebRequest.ContentType = "application/json";
			httpWebRequest.Method = "POST";

			using (var streamWriter = new StreamWriter(httpWebRequest.GetRequestStream()))
			{
				var json = JsonConvert.SerializeObject(new CanModel { Method = method, Controller = controller, Action = action });

				streamWriter.Write(json);
				streamWriter.Flush();
				streamWriter.Close();
			}

			try
			{
				var httpResponse = (HttpWebResponse)httpWebRequest.GetResponse();
				return (httpResponse.StatusCode == HttpStatusCode.OK);
			}
			catch
			{
				return false;
			}
		}

		public static void LogOff(string token)
		{
			ServicePointManager.ServerCertificateValidationCallback = (sender, cert, chain, sslPolicyErrors) => true;
			var url = $"{OAuthUrl}logoff";
			var httpWebRequest = (HttpWebRequest)WebRequest.Create(url);
			httpWebRequest.Headers.Add("Authorization", $"Bearer {token}");
			httpWebRequest.Method = "POST";

			try
			{
				var httpResponse = (HttpWebResponse)httpWebRequest.GetResponse();
			}
			catch { }
		}

		public static MvcHtmlString OAuthorizedActionLink(HttpContextBase httpContext, HtmlHelper html, string linkText, string actionName, string controllerName, Object routeValues = null, Object htmlAttrubutes = null)
		{
			var token = GetJWT(httpContext);
			return Can(token, controllerName, actionName) ? System.Web.Mvc.Html.LinkExtensions.ActionLink(html, linkText, actionName, controllerName, routeValues, htmlAttrubutes) : null;
		}
		#endregion

	}
}

