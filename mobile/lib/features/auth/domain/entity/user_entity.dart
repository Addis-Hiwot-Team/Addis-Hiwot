class UserEntity {
  final int userId;
  final String name;
  final String username;
  final String email;
  final String profileImage;
  final bool isActive;
  final String role;

  UserEntity({
    required this.userId,
    required this.name,
    required this.username,
    required this.email,
    required this.isActive,
    required this.role,
    required this.profileImage,
  });
}
