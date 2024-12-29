// lib/api/types/auth_body.dart

class UserLogInfosBody {
  final String token;
  final int userId;

  UserLogInfosBody({
    required this.token,
    required this.userId,
  });

  factory UserLogInfosBody.fromJson(Map<String, dynamic> json) {
    return UserLogInfosBody(
      token: json['token'],
      userId: json['user_id'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'token': token,
      'user_id': userId,
    };
  }
}

class UserCredentials {
  final String email;
  final String password;

  UserCredentials({
    required this.email,
    required this.password,
  });

  factory UserCredentials.fromJson(Map<String, dynamic> json) {
    return UserCredentials(
      email: json['email'],
      password: json['password'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'email': email,
      'password': password,
    };
  }
}

class UserSignUpBody {
  final String email;
  final String password;
  final String firstName;
  final String lastName;

  UserSignUpBody({
    required this.email,
    required this.password,
    this.firstName = '',
    this.lastName = '',
  });

  factory UserSignUpBody.fromJson(Map<String, dynamic> json) {
    return UserSignUpBody(
      email: json['email'],
      password: json['password'],
      firstName: json['first_name'],
      lastName: json['last_name'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'email': email,
      'password': password,
      'first_name': firstName,
      'last_name': lastName,
    };
  }
}

class OAuthLoginBody {
  final String service;
  final String? code;
  final String redirectUri;

  OAuthLoginBody({
    required this.service,
    this.code,
    required this.redirectUri,
  });

  factory OAuthLoginBody.fromJson(Map<String, dynamic> json) {
    return OAuthLoginBody(
      service: json['service'],
      code: json['code'],
      redirectUri: json['redirect_uri'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'service': service,
      'code': code,
      'redirect_uri': redirectUri,
    };
  }
}
