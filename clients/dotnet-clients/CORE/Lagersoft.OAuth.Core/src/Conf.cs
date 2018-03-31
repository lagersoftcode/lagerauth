using Microsoft.Extensions.Configuration;

namespace Lagersoft.OAuth
{
	public static class OAuthConf
	{
		public static void ConfigureOAuthGlobals(IConfiguration configuration)
		{
			OAuthUrl = configuration["OAuthUrl"];
			ClientId = configuration["ClientId"];
			SecretKey = configuration["SecretKey"];
			CookieName = configuration["CookieName"];
		}

		public static string OAuthUrl { get; private set; }
		public static string ClientId { get; private set; }
		public static string SecretKey { get; private set; }
		public static string CookieName { get; private set; }
	}
}
