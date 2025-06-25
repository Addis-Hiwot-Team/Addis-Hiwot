import 'dart:convert';
import 'dart:ffi';
import 'package:http/http.dart' as http;
import 'package:mobile/core/constants/api_urls.dart';
import 'package:mobile/core/errors/exceptions.dart';
import 'package:mobile/features/auth/data/model/user_model.dart';

abstract class AuthRemoteDatasource {
  Future<(String, UserModel)> signup(
    String username,
    String email,
    String password,
  );
  Future<(String, UserModel)> login(String identifier, String password);
  Future<void> logout();
  Future<UserModel> getme(Uint32 id, String token);
}

class AuthRemoteDatasourceImpl extends AuthRemoteDatasource {
  final http.Client _client;

  AuthRemoteDatasourceImpl({required http.Client client}) : _client = client;

  @override
  Future<(String, UserModel)> signup(
      String username, String email, String password) async {
    try {
      final response = await _client.post(
        Uri.parse(AuthApiUrls.signup()),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'username': username,
          'email': email,
          'password': password,
        }),
      );

      final decoded = jsonDecode(response.body);

      if (response.statusCode == 201) {
        final String token = decoded['data']['token'];
        final user = UserModel.fromJson(decoded['data']['user']);
        return (token, user);
      } else if (response.statusCode == 401) {
        throw UnauthroizedException();
      } else if (response.statusCode >= 500) {
        throw ServerException();
      } else {
        throw UnexpectedException();
      }
    } on http.ClientException {
      throw SocketException();
    } on FormatException {
      throw ParsingException();
    } catch (_) {
      throw UnexpectedException();
    }
  }

  @override
  Future<(String, UserModel)> login(String identifier, String password) async {
    try {
      final response = await _client.post(
        Uri.parse(AuthApiUrls.login()),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({'identifier': identifier, 'password': password}),
      );

      final decoded = jsonDecode(response.body);

      if (response.statusCode == 200) {
        final String token = decoded['data']['token'];
        final user = UserModel.fromJson(decoded['data']['user']);
        return (token, user);
      } else if (response.statusCode == 401) {
        throw UnauthroizedException();
      } else if (response.statusCode >= 500) {
        throw ServerException();
      } else {
        throw UnexpectedException();
      }
    } on http.ClientException {
      throw SocketException();
    } on FormatException {
      throw ParsingException();
    } catch (_) {
      throw UnexpectedException();
    }
  }

  @override
  Future<void> logout() async {
    try {
      final response = await _client.post(
        Uri.parse(AuthApiUrls.logout()),
        headers: {'Content-Type': 'application/json'},
      );

      if (response.statusCode != 200) {
        throw ServerException();
      }
    } on http.ClientException {
      throw SocketException();
    } catch (_) {
      throw UnexpectedException();
    }
  }

  @override
  Future<UserModel> getme(Uint32 id, String token) async {
    try {
      final response = await _client.get(
        Uri.parse(AuthApiUrls.getme(id)),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      final decoded = jsonDecode(response.body);
      if (response.statusCode == 200) {
        return UserModel.fromJson(decoded['data']);
      } else if (response.statusCode == 401) {
        throw UnauthroizedException();
      } else if (response.statusCode >= 500) {
        throw ServerException();
      } else {
        throw UnexpectedException();
      }
    } on http.ClientException {
      throw SocketException();
    } on FormatException {
      throw ParsingException();
    } catch (_) {
      throw UnexpectedException();
    }
  }
}
