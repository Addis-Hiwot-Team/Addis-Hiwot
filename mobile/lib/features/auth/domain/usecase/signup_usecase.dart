import 'package:dartz/dartz.dart';
import 'package:mobile/core/errors/failure.dart';
import 'package:mobile/features/auth/domain/entity/user_entity.dart';
import 'package:mobile/features/auth/domain/repository/user_repository.dart';

class SignupUseCase {
  final UserRepository repository;

  SignupUseCase(this.repository);

  Future<Either<Failure, UserEntity>> call({
    required String name,
    required String username,
    required String email,
    required String password,
  }) async {
    return await repository.signup(name, username, email, password);
  }
}
