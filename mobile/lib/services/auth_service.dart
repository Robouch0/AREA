// lib/auth_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'dart:developer' as developer;
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:my_area_flutter/api/types/area_services.dart';

class AuthService {
  static final AuthService _instance = AuthService._internal();
  AuthService._internal();
  static AuthService get instance => _instance;

  final _storage = const FlutterSecureStorage();
  static const _tokenKey = 'auth_token';
  final _apiUrl = dotenv.get('API_URL');

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
        _isLoggedIn = true;
        developer.log('Connected with token: $token', name: 'my_network_log');
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
    _isLoggedIn = false;
  }

  Future<List<AreaService>> listAreas() async {
    try {
      final token = await _storage.read(key: _tokenKey);

      if (token == null) {
        throw Exception('Token is undefined');
      }

      final response = await http.get(
        Uri.parse('$_apiUrl/create/list'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final List<dynamic> jsonData = json.decode(response.body);
        return jsonData.map((json) => AreaService.fromJson(json)).toList();
      }

      throw Exception('Failed to load areas: ${response.statusCode}');
    } catch (e) {
      developer.log('Failed to list areas: $e', name: 'my_network_log');
      rethrow;
    }
  }
}