import 'dart:ffi';

class UserEntity {
  final Uint32 userId;
  final String username;
  final String email;
  final String profileImage;
  final bool isActive;
  final String role;

  UserEntity({
    required this.userId,
    required this.username,
    required this.email,
    required this.isActive,
    required this.role,
    required this.profileImage,
  });
}
