// lib/services/api/area_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'dart:developer' as developer;
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:my_area_flutter/api/types/area_body.dart';
import 'package:my_area_flutter/api/types/area_create_body.dart';
import 'package:my_area_flutter/services/storage/auth_storage.dart';

class AreaService {
  static final AreaService _instance = AreaService._internal();
  AreaService._internal();
  static AreaService get instance => _instance;

  final _apiUrl = dotenv.get('NEXT_PUBLIC_GATEWAY_URL');
  final _authStorage = AuthStorage.instance;

  Future<List<AreaServiceData>> listAreas() async {
    try {
      final token = _authStorage.getToken();

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

  Future<List<UserAreaData>> listUserAreas() async {
    try {
      final token = _authStorage.getToken();

      if (token == null) {
        throw Exception('Token is undefined');
      }

      final response = await http.get(
        Uri.parse('$_apiUrl/areas/list'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final dynamic jsonData = json.decode(response.body);
        if (jsonData == null) return [];
        final areas = (jsonData as List)
            .map((json) => UserAreaData.fromJson(json))
            .toList();
        return areas;
      }

      throw Exception('Failed to load user areas: ${response.statusCode}');
    } catch (e) {
      developer.log('Failed to list user areas: $e', name: 'my_network_log');
      rethrow;
    }
  }

  Future<bool> createArea(AreaCreateBody newArea) async {
    try {
      final token = _authStorage.getToken();

      if (token == null) {
        throw Exception('Token is undefined');
      }

      final response = await http.post(
        Uri.parse('$_apiUrl/create/${newArea.action.service}'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
        body: json.encode(newArea.toJson()),
      );

      if (response.statusCode == 200) {
        return true;
      }
      return false;
    } catch (e) {
      developer.log('Failed to create area: $e', name: 'my_network_log');
      return false;
    }
  }
}
