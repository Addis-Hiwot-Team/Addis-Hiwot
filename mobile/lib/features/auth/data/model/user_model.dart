
import 'package:mobile/features/auth/domain/entity/user_entity.dart';

class UserModel {
  final int userId;
  final String username;
  final String email;
  final String profileImage;
  final bool isActive;
  final String role;

  UserModel({
    required this.userId,
    required this.username,
    required this.email,
    required this.profileImage,
    required this.isActive,
    required this.role,
  });

  factory UserModel.fromJson(Map<String, dynamic> json) => UserModel(
    userId: json['id'],
    username: json['username'],
    email: json['email'],
    profileImage: json['profile_image'],
    isActive: json['is_active'],
    role: json['role'],
  );

  Map<String, dynamic> toJson() => {
    'id': userId,
    'username': username,
    'email': email,
    'profile_image': profileImage,
    'is_active': isActive,
    'role': role,
  };

  UserEntity toEntity() => UserEntity(
    userId: userId,
    username: username,
    email: email,
    profileImage: profileImage,
    isActive: isActive,
    role: role,
  );
}
