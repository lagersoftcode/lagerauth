using JWT;
using JWT.Serializers;
using Microsoft.AspNetCore.Html;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.AspNetCore.Mvc.ViewFeatures;
using Newtonsoft.Json;
using System;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Text;
using System.Web;

namespace Lagersoft.OAuth
{
	public class Utils
	{
		private static string ClientId => OAuthConf.ClientId;
		private static string SecretKey => OAuthConf.SecretKey;
		private static string OAuthUrl => OAuthConf.OAuthUrl;
		private static string CookieName => OAuthConf.CookieName;

		#region helpers
		public static string GetToken(string code)
		{
			var tokenUrl = $"token?client_id={ClientId}&client_secret={SecretKey}&code={code}";

			ServicePointManager.ServerCertificateValidationCallback = (sender, cert, chain, sslPolicyErrors) => true;
			var httpClient = new HttpClient()
			{
				BaseAddress = new Uri(OAuthUrl),
				Timeout = new TimeSpan(0, 0, 30)
			};

			var content = new StringContent(string.Empty);
			var result = httpClient.PostAsync(tokenUrl, content).Result;
			var json = result.Content.ReadAsStringAsync().Result;
			dynamic jsonObj = JsonConvert.DeserializeObject(json);

			return jsonObj.access_token ?? null;
		}

		public static OAuthModel GetUser(HttpContext context)
		{
			if (!IsAuthenticated(context))
				return null;

			var token = context.Request.Cookies[CookieName];
			if (string.IsNullOrEmpty(token))
				return null;

			var key = Encoding.ASCII.GetBytes(SecretKey);

			IJsonSerializer serializer = new JsonNetSerializer();
			IDateTimeProvider provider = new UtcDateTimeProvider();
			IJwtValidator validator = new JwtValidator(serializer, provider);
			IJwtDecoder decoder = new JwtDecoder(serializer, validator, new JwtBase64UrlEncoder());

			var user = decoder.DecodeToObject<OAuthModel>(token, key, false);
			return user;
		}

		public static string GetJWT(HttpContext context)
		{
			// Get bearer token first hotfix for api usage:
			var authHeader = context.Request.Headers["Authorization"].ToString();
			if (!string.IsNullOrEmpty(authHeader))
			{
				return authHeader.Split(" ").Last();
			}

			if (!IsAuthenticated(context))
				return null;

			var token = context.Request.Cookies[CookieName];
			if (string.IsNullOrEmpty(token))
				return null;

			return token;
		}

		public static string GetLoginUrl(string host)
		{
			var returnUrl = $"{host}/OAuth/OAuth";
			return $"{OAuthUrl}auth?client_id={ClientId}&redirect_uri={HttpUtility.UrlEncode(returnUrl)}";
		}
		#endregion

		#region authentication
		public static bool IsAuthenticated(HttpContext context)
		{
			var authorizationHeader = context.Request.Headers["Authorization"].ToString();
			return authorizationHeader.StartsWith("Bearer") || !string.IsNullOrEmpty(context.Request.Cookies[CookieName]);
		}

		public static void Authenticate(HttpContext context, string token)
		{
			context.Response.Cookies.Append(CookieName, token);
		}

		public static void DeAuthenticate(HttpContext context)
		{
			if (IsAuthenticated(context))
			{
				Logoff(GetJWT(context));
				context.Response.Cookies.Append(CookieName, string.Empty, new CookieOptions { Expires = DateTime.Now.AddDays(-1) });
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

		public static void Logoff(string token)
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

		public static IHtmlContent OAuthorizedActionLink(HttpContext httpContext, HtmlHelper html, string linkText, string actionName, string controllerName, Object routeValues = null, Object htmlAttrubutes = null)
		{
			var token = GetJWT(httpContext);
			return Can(token, controllerName, actionName) ? html.ActionLink(linkText, actionName, controllerName, routeValues, htmlAttrubutes) : null;
		}

		#endregion

	}
}
