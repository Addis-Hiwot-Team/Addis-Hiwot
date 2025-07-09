import 'package:dartz/dartz.dart';
import 'package:mobile/core/errors/failure.dart';
import 'package:mobile/features/auth/domain/entity/user_entity.dart';
import 'package:mobile/features/auth/domain/repository/user_repository.dart';

class LoginUseCase {
  final UserRepository repository;

  LoginUseCase(this.repository);

  Future<Either<Failure, UserEntity>> call({
    required String identifier,
    required String password,
  }) async {
    return await repository.login(identifier, password);
  }
}
