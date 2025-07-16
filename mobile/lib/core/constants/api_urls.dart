
class AuthApiUrls {
  static const String baseurl = 'https://addis-hiwot.onrender.com';

  static String login() => '$baseurl/api/v1/auth/login';
  static String signup() => '$baseurl/api/v1/auth/register';
  static String logout() => '$baseurl/api/v1/auth/logout';
  static String getme(int id) => '$baseurl/api/v1/users/$id';
}
