// lib/services/api/area_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'dart:developer' as developer;
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:my_area_flutter/api/types/area_body.dart';
import 'package:my_area_flutter/services/storage/auth_storage.dart';

class AreaService {
  static final AreaService _instance = AreaService._internal();
  AreaService._internal();
  static AreaService get instance => _instance;

  final _apiUrl = dotenv.get('NEXT_PUBLIC_GATEWAY_URL');
  final _authStorage = AuthStorage.instance;

  Future<List<AreaServiceData>> listAreas() async {
    try {
      final token = await _authStorage.getToken();

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
        return jsonData.map((json) => AreaServiceData.fromJson(json)).toList();
      }

      throw Exception('Failed to load areas: ${response.statusCode}');
    } catch (e) {
      developer.log('Failed to list areas: $e', name: 'my_network_log');
      rethrow;
    }
  }
}
