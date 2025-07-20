import 'dart:convert';

import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:mobile/core/errors/exceptions.dart';
import 'package:mobile/features/auth/data/model/user_model.dart';
import 'package:shared_preferences/shared_preferences.dart';

abstract class AuthLocalDatasource {
  Future<void> storeAccessToken(String token);
  Future<String?> getAccessToken();
  Future<void> deleteAccessToken();

  Future<void> saveUser(UserModel user);
  Future<UserModel?> getUser();
  Future<void> deleteUser();
}

class AuthLocalDataSourceImpl extends AuthLocalDatasource {
  final FlutterSecureStorage _storage;
  static const String _accesstokenkey = 'ACCESS_TOKEN';
  static const String _userKey = 'USER_MODEL';

  AuthLocalDataSourceImpl({FlutterSecureStorage? storage})
      : _storage = storage ?? const FlutterSecureStorage();

  @override
  Future<void> storeAccessToken(String token) async {
    try {
      await _storage.write(key: _accesstokenkey, value: token);
    } catch (_) {
      throw CacheException();
    }
  }

  @override
  Future<String?> getAccessToken() async {
    try {
      return await _storage.read(key: _accesstokenkey);
    } catch (_) {
      throw CacheException();
    }
  }

  @override
  Future<void> deleteAccessToken() async {
    try {
      await _storage.delete(key: _accesstokenkey);
    } catch (_) {
      throw CacheException();
    }
  }

  @override
  Future<void> saveUser(UserModel user) async {
    try {
      final prefs = await SharedPreferences.getInstance();
      final userJson = jsonEncode(user.toJson());
      await prefs.setString(_userKey, userJson);
    } catch (_) {
      throw CacheException();
    }
  }

  @override
  Future<UserModel?> getUser() async {
    try {
      final prefs = await SharedPreferences.getInstance();
      final userJson = prefs.getString(_userKey);
      if (userJson == null) return null;

      final decoded = jsonDecode(userJson) as Map<String, dynamic>;
      return UserModel.fromJson(decoded);
    } catch (_) {
      throw CacheException();
    }
  }

  @override
  Future<void> deleteUser() async {
    try {
      final prefs = await SharedPreferences.getInstance();
      await prefs.remove(_userKey);
    } catch (_) {
      throw CacheException();
    }
  }
}
