// lib/pages/server_config_page.dart
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'package:my_area_flutter/widgets/auth_input_field.dart';
import 'package:my_area_flutter/widgets/auth_button.dart';
import 'package:my_area_flutter/widgets/main_app_scaffold.dart';
import 'package:my_area_flutter/core/router/route_names.dart';

class ServerConfigPage extends StatefulWidget {
  const ServerConfigPage({super.key});

  @override
  State<ServerConfigPage> createState() => _ServerConfigPageState();
}

class _ServerConfigPageState extends State<ServerConfigPage> {
  final _formKey = GlobalKey<FormState>();
  final _serverUrlController = TextEditingController();
  static const String _serverUrlKey = 'server_url';

  @override
  void initState() {
    super.initState();
    _loadSavedUrl();
  }

  void _loadSavedUrl() async {
    final prefs = await SharedPreferences.getInstance();
    final savedUrl = prefs.getString(_serverUrlKey);
    if (savedUrl != null) {
      _serverUrlController.text = savedUrl;
    }
  }

  Future<void> _saveServerUrl(String url) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString(_serverUrlKey, url);
  }

  bool _isValidUrl(String url) {
    try {
      final uri = Uri.parse(url);
      return uri.isScheme('http') || uri.isScheme('https');
    } catch (e) {
      return false;
    }
  }

  void _showSuccessStatus(bool success) {
    if (!mounted) return;

    ScaffoldMessenger.of(context).showSnackBar(SnackBar(
      content: Text(
        success ? 'Server configuration saved successfully.' : 'Invalid server URL format.',
        style: const TextStyle(fontWeight: FontWeight.w800),
      ),
      backgroundColor: success ? Colors.green : Colors.red,
    ));

    if (success) {
      context.go(RouteNames.login);
    }
  }

  void _saveConfig() async {
    if (_formKey.currentState!.validate()) {
      String url = _serverUrlController.text.trim();

      if (!url.startsWith('http://') && !url.startsWith('https://')) {
        url = 'http://$url';
      }

      if (_isValidUrl(url)) {
        await _saveServerUrl(url);
        _showSuccessStatus(true);
      } else {
        _showSuccessStatus(false);
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return MainAppScaffold(
      child: Form(
        key: _formKey,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            const Center(
              child: Text(
                'Network Config',
                style: TextStyle(
                  fontSize: 32,
                  fontWeight: FontWeight.w800,
                ),
              ),
            ),
            const SizedBox(height: 60),
            const Center(
              child: Text(
                'Please enter the server URL',
                style: TextStyle(
                  fontSize: 16,
                  color: Colors.grey,
                ),
              ),
            ),
            const SizedBox(height: 30),
            _buildServerUrlField(),
            const SizedBox(height: 25),
            _buildSaveButton(),
          ],
        ),
      ),
    );
  }

  Widget _buildServerUrlField() {
    return AuthInputField(
      controller: _serverUrlController,
      hintText: '(e.g., 10.0.2.2:8080)',
      obscureText: false,
      validator: (value) {
        if (value == null || value.isEmpty) {
          return 'Please enter a server URL.';
        }
        return null;
      },
    );
  }

  Widget _buildSaveButton() {
    return AuthButton(
      text: 'Save Configuration',
      onPressed: _saveConfig,
    );
  }

  @override
  void dispose() {
    _serverUrlController.dispose();
    super.dispose();
  }
}
