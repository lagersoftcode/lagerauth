namespace Lagersoft.OAuth
{
	public class OAuthModel
	{
		public string email { get; set; }
		public string token { get; set; }
	}

	public class CanModel
	{
		public string Method { get; set; }
		public string Controller { get; set; }
		public string Action { get; set; }
		public string Extra { get; set; }
	}
}
