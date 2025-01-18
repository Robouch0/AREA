// lib/services/api/area_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'dart:developer' as developer;
import 'package:my_area_flutter/api/types/area_body.dart';
import 'package:my_area_flutter/api/types/area_create_body.dart';
import 'package:my_area_flutter/api/types/area_activation_body.dart';
import 'package:my_area_flutter/services/api/server_service.dart';
import 'package:my_area_flutter/services/storage/auth_storage.dart';

class AreaService {
  static final AreaService _instance = AreaService._internal();
  AreaService._internal();
  static AreaService get instance => _instance;

  final _authStorage = AuthStorage.instance;

  Future<List<AreaServiceData>> listAreas() async {
    try {
      final token = _authStorage.getToken();
      final apiUrl = await ServerService.getApiUrl();

      if (token == null) {
        throw Exception('Token is undefined');
      }

      final response = await http.get(
        Uri.parse('$apiUrl/areas/create/list'),
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
      final apiUrl = await ServerService.getApiUrl();

      if (token == null) {
        throw Exception('Token is undefined');
      }

      final response = await http.get(
        Uri.parse('$apiUrl/areas/list'),
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
      final apiUrl = await ServerService.getApiUrl();

      if (token == null) {
        throw Exception('Token is undefined');
      }

      final response = await http.post(
        Uri.parse('$apiUrl/areas/create/${newArea.action.service}'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
        body: json.encode(newArea.toJson()),
      );

      if (response.statusCode == 200) {
        return true;
      }
      developer.log('Error when creating area: ${response.statusCode}');
      return false;
    } catch (e) {
      developer.log('Failed to create area: $e', name: 'my_network_log');
      return false;
    }
  }

  Future<void> updateAreaActivation(int areaId, bool activated) async {
    final areaActivation = AreaActivationBody(areaId: areaId, activated: activated);

    try {
      final token = _authStorage.getToken();
      final apiUrl = await ServerService.getApiUrl();

      final response = await http.put(
        Uri.parse('$apiUrl/areas/activate'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-type': 'application/json',
        },
        body: json.encode(areaActivation.toJson())
      );
    } catch (e) {
      developer.log('Failed to update area activation: $e');
      rethrow;
    }
  }
}
