// lib/services/api/area_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'dart:developer' as developer;
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:my_area_flutter/api/types/profile_body.dart';
import 'package:my_area_flutter/services/storage/auth_storage.dart';

class ProfileService {
  static final ProfileService _instance = ProfileService._internal();
  ProfileService._internal();
  static ProfileService get instance => _instance;

  final _apiUrl = dotenv.get('NEXT_PUBLIC_GATEWAY_URL');
  final _authStorage = AuthStorage.instance;

  Future<UserInfoData> getUserInfo() async {
    try {
      final token = await _authStorage.getToken();
      final userId = await _authStorage.getUserId();

      if (token == null) {
        throw Exception('Token is undefined');
      }

      final response = await http.get(
        Uri.parse('$_apiUrl/users/$userId'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      developer.log(response.body);

      developer.log('token: $token');

      if (response.statusCode == 200) {
        final Map<String, dynamic> jsonData = json.decode(response.body);
        return UserInfoData.fromJson(jsonData);
      }
      throw Exception('Failed to load user infos: ${response.statusCode}');
    } catch (e) {
      developer.log('Failed to get user infos: $e', name: 'my_network_log');
      rethrow;
    }
  }
}
