// lib/auth_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'dart:developer' as developer;

class AuthService {
  final _storage = const FlutterSecureStorage();
  static const _tokenKey = 'auth_token';
  static const _apiUrl = 'http://10.0.2.2:3000';

  Future<bool> createAccount(String email, String password) async {
    try {
      final response = await http.post(
        Uri.parse('$_apiUrl/sign-up/'),
        headers: {
          'Content-type': 'application/json',
        },
        body: json.encode({
          'email': email,
          'password': password
        })
      );

      if (response.statusCode == 200) {
        developer.log('Account created successfully!');
        return true;
      }
      developer.log('Error while creating an account : ${response.statusCode}');
      return false;
    } catch (e) {
      developer.log('Failed to login: $e', name: 'my_network_log');
      return false;
    }
  }

  Future<bool> login(String email, String password) async {
    try {
      final response = await http.post(
          Uri.parse('$_apiUrl/login/'),
          headers: {
            'Content-Type': 'application/json',
          },
          body: json.encode({
            'email': email,
            'password': password
          })
      );

      if (response.statusCode == 200) {
        developer.log('Success connection!');
        final token = response.body;

        await _storage.write(key: _tokenKey, value: token);
        developer.log('Connected with token: $token', name: 'my_network_log');
        return true;
      }
      developer.log('Got invalid response statusCode: ${response.statusCode}');
      return false;
    } catch (e) {
      developer.log('Failed to login: $e', name: 'my_network_log');
      return false;
    }
  }

  Future<String?> getToken() async {
    return await _storage.read(key: _tokenKey);
  }

  Future<bool> isLoggedIn() async {
    final token = await getToken();

    return token != null;
  }

  Future<void> logout() async {
    await _storage.delete(key: _tokenKey);
  }
}