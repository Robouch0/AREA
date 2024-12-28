// lib/services/storage/auth_storage.dart
import 'package:my_area_flutter/services/storage/secure_storage_service.dart';

class AuthStorage {
  static final AuthStorage _instance = AuthStorage._internal();
  AuthStorage._internal();
  static AuthStorage get instance => _instance;

  final _storage = SecureStorageService.instance;
  static const _tokenKey = 'auth_token';
  static const _userIdKey = 'user_id';

  Future<void> saveToken(String token) async {
    await _storage.write(_tokenKey, token);
  }

  Future<void> saveUserId(String userId) async {
    await _storage.write(_userIdKey, userId);
  }

  Future<String?> getToken() async {
    return await _storage.read(_tokenKey);
  }

  Future<String?> getUserId() async {
    return await _storage.read(_userIdKey);
  }

  Future<void> clearAuth() async {
    await _storage.delete(_tokenKey);
    await _storage.delete(_userIdKey);
  }
}
