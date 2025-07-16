class ServerException implements Exception{}

class SocketException implements Exception{}

class CacheException implements Exception{}

class UnauthorizedException implements Exception {}

class TimeoutException implements Exception {}

class ParsingException implements Exception {}

class UnexpectedException implements Exception {}

class BusinessLogicException implements Exception {
  final String message;
  BusinessLogicException(this.message);
}