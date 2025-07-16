import 'package:dartz/dartz.dart';
import 'package:mobile/core/errors/exceptions.dart';
import 'package:mobile/core/errors/failure.dart';
import 'package:mobile/core/utils/network_info.dart';
import 'package:mobile/features/auth/data/data_sources/local_datasource/auth_local_datasource.dart';
import 'package:mobile/features/auth/data/data_sources/remote_datasource/auth_remote_datasource.dart';
import 'package:mobile/features/auth/domain/entity/user_entity.dart';
import 'package:mobile/features/auth/domain/repository/user_repository.dart';

class UserRepositoryImpl extends UserRepository {
  final AuthLocalDataSourceImpl local;
  final AuthRemoteDatasource remote;
  final NetworkInfo networkInfo;

  UserRepositoryImpl({
    required this.local,
    required this.remote,
    required this.networkInfo,
  });

  @override
  Future<Either<Failure, UserEntity>> signup(
      String Name, String username, String email, String password) async {
    if (await networkInfo.isConnected) {
      try {
        final (token, user) = await remote.signup(Name, username, email, password);
        await local.storeAccessToken(token);
        await local.saveUser(user);
        return Right(user.toEntity());
      } on BusinessLogicException catch (e) {
        return Left(BusinessLogicFailure(message: e.message));
      } on ServerException {
        return const Left(ServerFailure(message: 'Signup failed on server.'));
      } on SocketException {
        return const Left(NetworkFailure(message: 'No internet connection.'));
      } on ParsingException {
        return const Left(ServerFailure(message: 'Error parsing response. Please check the API response format.'));
      } catch (e) {
        print('Repository signup error: $e');
        return const Left(ServerFailure(message: 'Unexpected signup error. Please try again.'));
      }
    } else {
      return const Left(NetworkFailure(message: 'No internet connection.'));
    }
  }

  @override
  Future<Either<Failure, UserEntity>> login(
      String identifier, String password) async {
    if (await networkInfo.isConnected) {
      try {
        final (token, user) = await remote.login(identifier, password);
        await local.storeAccessToken(token);
        // Skip saving the dummy user for now
        // await local.saveUser(user);
        return Right(user.toEntity());
      } on UnauthorizedException {
        return const Left(ServerFailure(message: 'Invalid credentials.'));
      } on ServerException {
        return const Left(ServerFailure(message: 'Login failed on server.'));
      } on SocketException {
        return const Left(NetworkFailure(message: 'No internet connection.'));
      } catch (_) {
        return const Left(ServerFailure(message: 'Unexpected login error.'));
      }
    } else {
      return const Left(NetworkFailure(message: 'No internet connection.'));
    }
  }

  @override
  Future<Either<Failure, void>> logout() async {
    try {
      await local.deleteAccessToken();
      await local.deleteUser();

      if (await networkInfo.isConnected) {
        await remote.logout();
      }
      return const Right(null);
    } catch (_) {
      return const Left(DatabaseFailure(message: 'Logout failed.'));
    }
  }

  @override
  Future<Either<Failure, UserEntity>> getme() async {
    try {
      final token = await local.getAccessToken();
      if (token == null) {
        return const Left(ServerFailure(message: 'User not logged in.'));
      }

      final cachedUser = await local.getUser();
      if (cachedUser == null) {
        return const Left(DatabaseFailure(message: 'No cached user found.'));
      }

      if (await networkInfo.isConnected) {
  try {
    final updatedUser = await remote.getme(cachedUser.userId, token);
    await local.saveUser(updatedUser);
    return Right(updatedUser.toEntity());
  } on ServerException {
    return Right(cachedUser.toEntity()); // fallback
  }
}

      return Right(cachedUser.toEntity());
    } catch (_) {
      return const Left(ServerFailure(message: 'Failed to get user info.'));
    }
  }
}
