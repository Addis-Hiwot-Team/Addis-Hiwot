import 'dart:ffi';

class AuthApiUrls {
  static const String baseurl = 'h';

  static String login() => '$baseurl/login';
  static String signup() => '$baseurl/signup';
  static String logout() => '$baseurl/logout';
  static String getme(Uint32 id) => '$baseurl/id';
}
