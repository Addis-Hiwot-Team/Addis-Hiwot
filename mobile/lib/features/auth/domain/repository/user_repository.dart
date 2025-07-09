import 'package:mobile/core/errors/failure.dart';
import 'package:mobile/features/auth/domain/entity/user_entity.dart';
import 'package:dartz/dartz.dart';

abstract class UserRepository {
  Future<Either<Failure,UserEntity>> signup(String username, String email, String password);
  Future<Either<Failure,UserEntity>> login(String identifier, String password);
  Future<Either<Failure,void>> logout();
  Future<Either<Failure,UserEntity>> getme();
}