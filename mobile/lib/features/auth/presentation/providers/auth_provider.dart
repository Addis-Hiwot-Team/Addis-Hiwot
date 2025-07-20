import 'package:flutter/material.dart';
import 'package:mobile/features/auth/domain/entity/user_entity.dart';
import 'package:mobile/features/auth/domain/usecase/login_usecase.dart';
import 'package:mobile/features/auth/domain/usecase/signup_usecase.dart';
import 'package:mobile/features/auth/domain/usecase/logout_usecase.dart';
import 'package:mobile/features/auth/domain/usecase/getme_usecase.dart';


class AuthProvider extends ChangeNotifier {
  final LoginUseCase loginUseCase;
  final SignupUseCase signupUseCase;
  final LogoutUseCase logoutUseCase;
  final GetMeUseCase getMeUseCase;

  AuthProvider({
    required this.loginUseCase,
    required this.signupUseCase,
    required this.logoutUseCase,
    required this.getMeUseCase,
  });

  UserEntity? _user;
  bool _isLoading = false;
  String? _error;

  UserEntity? get user => _user;
  bool get isLoading => _isLoading;
  String? get error => _error;
  bool get isLoggedIn => _user != null;

  Future<void> login(String email, String password) async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    final result = await loginUseCase(
      identifier: email,
      password: password,
    );

    result.fold(
      (failure) => _error = failure.message,
      (user) => _user = user,
    );

    _isLoading = false;
    notifyListeners();
  }

  Future<void> signup(String username, String email, String password) async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    final result = await signupUseCase(
      username: username,
      email: email,
      password: password,
    );

    result.fold(
      (failure) => _error = failure.message,
      (user) => _user = user,
    );

    _isLoading = false;
    notifyListeners();
  }

  Future<void> getMe() async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    final result = await getMeUseCase();

    result.fold(
      (failure) => _error = failure.message,
      (user) => _user = user,
    );

    _isLoading = false;
    notifyListeners();
  }

  Future<void> logout() async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    final result = await logoutUseCase();

    result.fold(
      (failure) => _error = failure.message,
      (_) => _user = null,
    );

    _isLoading = false;
    notifyListeners();
  }
}
