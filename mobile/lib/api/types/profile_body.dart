// lib/api/types/profile_body.dart

class UserEditBody {
  final String firstName;
  final String lastName;

  UserEditBody({
    required this.firstName,
    required this.lastName,
  });

  Map<String, dynamic> toJson() {
    return {
      'first_name': firstName,
      'last_name': lastName,
    };
  }
}

class UserInfoBody {
  final int userId;
  final String firstName;
  final String lastName;
  final String email;

  UserInfoBody({
    required this.userId,
    required this.firstName,
    required this.lastName,
    required this.email,
  });

  factory UserInfoBody.fromJson(Map<String, dynamic> json) {
    return UserInfoBody(
      userId: json['id'],
      firstName: json['first_name'],
      lastName: json['last_name'],
      email: json['email'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'user_id': userId,
      'first_name': firstName,
      'last_name': lastName,
      'email': email,
    };
  }
}
