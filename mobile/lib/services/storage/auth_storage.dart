// lib/services/storage/auth_storage.dart
import 'package:shared_preferences/shared_preferences.dart';

class AuthStorage {
  static final AuthStorage _instance = AuthStorage._internal();
  AuthStorage._internal();
  static AuthStorage get instance => _instance;

  static const _tokenKey = 'auth_token';
  static const _userIdKey = 'user_id';

  late SharedPreferences _prefs;

  Future<void> initialize() async {
    _prefs = await SharedPreferences.getInstance();
  }

  void saveToken(String token) {
    _prefs.setString(_tokenKey, token);
  }

  void saveUserId(String userId) {
    _prefs.setString(_userIdKey, userId);
  }

  String? getToken() {
    return _prefs.getString(_tokenKey);
  }

  String? getUserId() {
    return _prefs.getString(_userIdKey);
  }

  void clearAuth() {
    _prefs.remove(_tokenKey);
    _prefs.remove(_userIdKey);
  }
}
