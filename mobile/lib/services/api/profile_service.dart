// lib/services/api/area_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'dart:developer' as developer;
import 'package:my_area_flutter/api/types/profile_body.dart';
import 'package:my_area_flutter/api/types/oauth_list_body.dart';
import 'package:my_area_flutter/api/types/user_provider_list_body.dart';
import 'package:my_area_flutter/services/api/server_service.dart';
import 'package:my_area_flutter/services/storage/auth_storage.dart';

class ProfileService {
  static final ProfileService _instance = ProfileService._internal();
  ProfileService._internal();
  static ProfileService get instance => _instance;

  final _authStorage = AuthStorage.instance;

  Future<bool> updateUserInfo(String firstName, String lastName) async {
    final UserEditBody userEditBody = UserEditBody(
      firstName: firstName,
      lastName: lastName,
    );

    try {
      final apiUrl = await ServerService.getApiUrl();
      final token = _authStorage.getToken();

      if (token == null) {
        return false;
      }

      final response = await http.put(
        Uri.parse('$apiUrl/user/me'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-type': 'application/json',
        },
        body: json.encode(userEditBody.toJson())
      );

      if (response.statusCode == 200) {
        developer.log('Account edited successfully!');
        return true;
      }
      return false;
    } catch (e) {
      developer.log('Failed to edit user infos: $e', name: 'my_network_log');
      return false;
    }
  }

  Future<UserInfoBody> getUserInfo() async {
    try {
      final token = _authStorage.getToken();
      final apiUrl = await ServerService.getApiUrl();

      if (token == null) {
        throw Exception('Token is undefined');
      }

      final response = await http.get(
        Uri.parse('$apiUrl/user/me'),
        headers: {
          'Authorization': 'Bearer $token',
        },
      );

      if (response.statusCode == 200) {
        final Map<String, dynamic> jsonData = json.decode(response.body);
        return UserInfoBody.fromJson(jsonData);
      }
      throw Exception('Failed to load user infos: ${response.statusCode}');
    } catch (e) {
      developer.log('Failed to get user infos: $e', name: 'my_network_log');
      rethrow;
    }
  }

  Future<UserProviderListBody> getUserProviders() async {
    try {
      final token = _authStorage.getToken();
      final apiUrl = await ServerService.getApiUrl();

      final response = await http.get(
        Uri.parse('$apiUrl/token/'),
        headers: {
          'Authorization': 'Bearer $token',
        }
      );

      if (response.statusCode == 200) {
        final Map<String, dynamic> jsonData = json.decode(response.body);
        return UserProviderListBody.fromJson(jsonData);
      }
      throw Exception('Failed to load users providers list: ${response.statusCode}');
    } catch (e) {
      rethrow;
    }
  }

  Future<OAuthListBody> getOAuthList() async {
    try {
      final apiUrl = await ServerService.getApiUrl();

      final response = await http.get(
        Uri.parse('$apiUrl/oauth/list'),
        headers: {
          'Content-type': 'application/json',
        }
      );

      if (response.statusCode == 200) {
        final Map<String, dynamic> jsonData = json.decode(response.body);
        return OAuthListBody.fromJson(jsonData);
      }
      throw Exception('Failed to load oauth list: ${response.statusCode}');
    } catch (e) {
      rethrow;
    }
  }
}
