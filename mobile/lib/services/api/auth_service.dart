// lib/services/api/auth_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'dart:developer' as developer;
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:my_area_flutter/api/types/auth_body.dart';
import 'package:my_area_flutter/services/storage/auth_storage.dart';

class AuthService {
  static final AuthService _instance = AuthService._internal();
  AuthService._internal();
  static AuthService get instance => _instance;

  final _authStorage = AuthStorage.instance;
  final _apiUrl = dotenv.get('NEXT_PUBLIC_GATEWAY_URL');
  bool _isLoggedIn = false;
  bool get isLoggedInSync => _isLoggedIn;

  Future<void> initializeAuth() async {
    await AuthStorage.instance.initialize();

    if (await checkAuthentification()) {
      _isLoggedIn = true;
    }
  }

  Future<bool> createAccount(String email, String password, String firstName, String lastName) async {
    final UserSignUpBody signUpBody = UserSignUpBody(
      email: email,
      password: password,
      firstName: firstName,
      lastName: lastName,
    );

    try {
      final response = await http.post(
          Uri.parse('$_apiUrl/sign-up/'),
          headers: {
            'Content-type': 'application/json',
          },
          body: json.encode(signUpBody.toJson())
      );

      if (response.statusCode == 200) {
        developer.log('Account created successfully!');
        return login(email, password);
      }
      developer.log('Error while creating an account : ${response.statusCode}');
      return false;
    } catch (e) {
      developer.log('Failed to login: $e', name: 'my_network_log');
      return false;
    }
  }

  Future<bool> login(String email, String password) async {
    final UserCredentials loginBody = UserCredentials(
        email: email,
        password: password
    );

    try {
      final response = await http.post(
          Uri.parse('$_apiUrl/login/'),
          headers: {
            'Content-Type': 'application/json',
          },
          body: json.encode(loginBody.toJson())
      );

      if (response.statusCode == 200) {
        developer.log('Success connection!');

        final Map<String, dynamic> jsonData = json.decode(response.body);
        final UserLogInfosBody responseData = UserLogInfosBody.fromJson(jsonData);

        _authStorage.saveToken(responseData.token);
        _authStorage.saveUserId(responseData.userId.toString());

        _isLoggedIn = true;
        developer.log('Connected with token: ${responseData.token}, userId: ${responseData.userId}', name: 'my_network_log');
        return _isLoggedIn;
      }
      developer.log('Got invalid response statusCode: ${response.statusCode}');
      return false;
    } catch (e) {
      developer.log('Failed to login: $e', name: 'my_network_log');
      return false;
    }
  }

  Future<bool> checkAuthentification() async {
    try {
      final token = _authStorage.getToken();
      final response = await http.get(
        Uri.parse('$_apiUrl/ping'),
        headers: {
          'Authorization': 'Bearer $token'
        },
      );

      if (response.statusCode == 200) {
        return true;
      }
      return false;
    } catch (e) {
      developer.log('Failed to check authentification: $e');
      return false;
    }
  }

  Future<void> logout() async {
    _authStorage.clearAuth();
    _isLoggedIn = false;
  }
}
