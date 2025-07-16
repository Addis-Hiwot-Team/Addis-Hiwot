import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:mobile/core/constants/api_urls.dart';
import 'package:mobile/core/errors/exceptions.dart';
import 'package:mobile/features/auth/data/model/user_model.dart';

abstract class AuthRemoteDatasource {
  Future<(String, UserModel)> signup(
    String name,
    String username,
    String email,
    String password,
  );
  Future<(String, UserModel)> login(String identifier, String password);
  Future<void> logout();
  Future<UserModel> getme(int id, String token);
}

class AuthRemoteDatasourceImpl extends AuthRemoteDatasource {
  final http.Client _client;

  AuthRemoteDatasourceImpl({required http.Client client}) : _client = client;

  @override
  Future<(String, UserModel)> signup(
      String name, String username, String email, String password) async {
    try {
      final requestBody = {
        'Name': name,
        'Username': username,
        'Email': email,
        'password': password,
      };
      
      print('Signup URL: ${AuthApiUrls.signup()}');
      print('Signup Request Body: ${jsonEncode(requestBody)}');
      
      final response = await _client.post(
        Uri.parse(AuthApiUrls.signup()),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode(requestBody),
      );

      print('Signup Response Status: ${response.statusCode}');
      print('Signup Response Body: ${response.body}');

      final decoded = jsonDecode(response.body);

      if (response.statusCode == 201) {
        // Check if the expected structure exists
        if (decoded['data'] == null) {
          print('Error: No "data" field in response');
          throw ParsingException();
        }
        
        if (decoded['data']['token'] == null) {
          print('Error: No "token" field in data');
          throw ParsingException();
        }
        
        if (decoded['data']['user'] == null) {
          print('Error: No "user" field in data');
          throw ParsingException();
        }

        final String token = decoded['data']['token'];
        final user = UserModel.fromJson(decoded['data']['user']);
        return (token, user);
      } else if (response.statusCode == 400) {
        // Handle validation errors
        if (decoded['message'] != null) {
          throw ServerException();
        }
        throw UnexpectedException();
      } else if (response.statusCode == 401) {
        throw UnauthorizedException();
      } else if (response.statusCode == 500) {
        // Handle business logic errors that return 500
        if (decoded['error'] != null) {
          final errorMessage = decoded['error'];
          if (errorMessage == 'user already exists') {
            throw BusinessLogicException('User already exists. Please try a different username or email.');
          }
        }
        throw ServerException();
      } else if (response.statusCode >= 500) {
        throw ServerException();
      } else {
        throw UnexpectedException();
      }
    } on http.ClientException {
      throw SocketException();
    } on FormatException catch (e) {
      print('FormatException during parsing: $e');
      throw ParsingException();
    } catch (e) {
      print('Unexpected error during signup: $e');
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
        // Handle backend response with only access_token
        final String token = decoded['access_token'];
        // Create a dummy user for now
        final user = UserModel.empty(); // You may need to implement UserModel.empty()
        return (token, user);
      } else if (response.statusCode == 401) {
        throw UnauthorizedException();
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
  Future<UserModel> getme(int id, String token) async {
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
        throw UnauthorizedException();
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
