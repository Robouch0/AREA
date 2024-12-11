// lib/auth_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:flutter_web_auth/flutter_web_auth.dart';

import 'dart:developer' as developer;

class AuthService {
  static final AuthService _instance = AuthService._internal();
  AuthService._internal();
  static AuthService get instance => _instance;

  final _storage = const FlutterSecureStorage();

  static const _tokenKey = 'auth_token';
  static const _userIDKey = 'uid';
  static const _apiUrl = 'http://10.0.2.2:3000';

  static const _redirectUrl = 'com.area.epitech';

  bool _isLoggedIn = false;

  Future<void> initializeAuth() async {
    final token = await _storage.read(key: 'auth_token');
    _isLoggedIn = token != null;
  }

  bool get isLoggedInSync => _isLoggedIn;

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

  Future<bool> loginWithOAuth(String service) async {
    try {
      final response = await http.get(Uri.parse('$_apiUrl/oauth/$service'));
      final authorizationUrl = response.body;

      developer.log(authorizationUrl);

      if (response.statusCode != 200) {
        developer.log('Cannot get the authorization URL.');
        return false;
      }

      final result = await FlutterWebAuth.authenticate(
          url: authorizationUrl,
          callbackUrlScheme: _redirectUrl
      );

      developer.log(result);

      final Uri uri = Uri.parse(result);
      final code = uri.queryParameters['code'];

      if (code == null) {
        developer.log('Cannot find the code to send back.');
        return false;
      }

      final tokenResponse = await http.post(
          Uri.parse('$_apiUrl/oauth/'),
          headers: {
            'Content-Type': 'application/json',
          },
          body: json.encode({
            'service': 'github',
            'code': code
          })
      );

      developer.log('Connected with $service successfully: ${tokenResponse.body}');
      await _storage.write(key: _tokenKey, value: tokenResponse.body);
      _isLoggedIn = true;
      return true;
    } catch (e) {
      developer.log('Failed to login with $service: $e');
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
        final body = response.body.split(',');
        final token = body.first;
        final uid = body.last;

        await _storage.write(key: _tokenKey, value: token);
        await _storage.write(key: _userIDKey, value: uid);
        _isLoggedIn = true;
        developer.log('Connected with token: $token and uid: $uid', name: 'my_network_log');
        return _isLoggedIn;
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

    _isLoggedIn = token != null;
    return _isLoggedIn;
  }

  Future<void> logout() async {
    await _storage.delete(key: _tokenKey);
    await _storage.delete(key: _userIDKey);
    _isLoggedIn = false;
  }
}