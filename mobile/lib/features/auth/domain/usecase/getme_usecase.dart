import 'package:dartz/dartz.dart';
import 'package:mobile/core/errors/failure.dart';
import 'package:mobile/features/auth/domain/entity/user_entity.dart';
import 'package:mobile/features/auth/domain/repository/user_repository.dart';

class GetMeUseCase {
  final UserRepository repository;

  GetMeUseCase(this.repository);

  Future<Either<Failure, UserEntity>> call() async {
    return await repository.getme();
  }
}
