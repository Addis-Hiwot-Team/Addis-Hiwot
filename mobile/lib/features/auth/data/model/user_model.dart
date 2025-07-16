
import 'package:mobile/features/auth/domain/entity/user_entity.dart';

class UserModel {
  final int userId;
  final String name;
  final String username;
  final String email;
  final String profileImage;
  final bool isActive;
  final String role;

  UserModel({
    required this.userId,
    required this.name,
    required this.username,
    required this.email,
    required this.profileImage,
    required this.isActive,
    required this.role,
  });

  factory UserModel.fromJson(Map<String, dynamic> json) {
    try {
      return UserModel(
        userId: json['id'] ?? 0,
        name: json['name'] ?? '',
        username: json['username'] ?? '',
        email: json['email'] ?? '',
        profileImage: json['profile_image'] ?? '',
        isActive: json['is_active'] ?? false,
        role: json['role'] ?? 'user',
      );
    } catch (e) {
      print('Error parsing UserModel from JSON: $e');
      print('JSON data: $json');
      rethrow;
    }
  }

  Map<String, dynamic> toJson() => {
    'id': userId,
    'name': name,
    'username': username,
    'email': email,
    'profile_image': profileImage,
    'is_active': isActive,
    'role': role,
  };

  UserEntity toEntity() => UserEntity(
    userId: userId,
    name: name,
    username: username,
    email: email,
    profileImage: profileImage,
    isActive: isActive,
    role: role,
  );

  static UserModel empty() {
    return UserModel(
      userId: 0,
      name: '',
      username: '',
      email: '',
      profileImage: '',
      isActive: false,
      role: 'user',
    );
  }
}
