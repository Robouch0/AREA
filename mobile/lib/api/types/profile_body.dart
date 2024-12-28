// lib/api/types/profile_body.dart

class UserInfoData {
  final int userId;
  final String firstName;
  final String lastName;
  final String email;
  final String password;

  UserInfoData({
    required this.userId,
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.password
  });

  factory UserInfoData.fromJson(Map<String, dynamic> json) {
    return UserInfoData(
      userId: json['id'],
      firstName: json['first_name'],
      lastName: json['last_name'],
      email: json['email'],
      password: json['password']
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'user_id': userId,
      'first_name': firstName,
      'last_name': lastName,
      'email': email,
      'password': password
    };
  }
}
