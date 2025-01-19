// lib/services/api/area_service.dart
import 'package:shared_preferences/shared_preferences.dart';

class ServerService {
  static const String _serverUrl = 'server_url';

  static Future<bool> isApiUrlDefined() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.containsKey(_serverUrl);
  }

  static Future<String?> getApiUrl() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getString(_serverUrl);
  }

  static Future<void> setApiUrl(String newUrl) async {
    final prefs = await SharedPreferences.getInstance();
    prefs.setString(_serverUrl, newUrl);
  }
}
