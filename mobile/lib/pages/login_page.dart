// lib/pages/login_page.dart
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import 'package:my_area_flutter/services/api/auth_service.dart';
import 'package:my_area_flutter/services/api/server_service.dart';
import 'package:my_area_flutter/widgets/auth_input_field.dart';
import 'package:my_area_flutter/widgets/auth_button.dart';
import 'package:my_area_flutter/widgets/main_app_scaffold.dart';
import 'package:my_area_flutter/core/router/route_names.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({super.key});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final _authService = AuthService.instance;
  final _formKey = GlobalKey<FormState>();
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();

  @override
  void initState() {
    super.initState();
    _checkNetworkAvailable();
  }

  Future<void> _checkNetworkAvailable() async {
    if (await ServerService.isApiUrlDefined() == false && mounted) {
      context.push(RouteNames.serverConfig);
    }
  }

  void _showSuccessStatus(bool success) async {
    if (!mounted) return;
    if (success) {
      ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
        content:
        Text('Success.', style: TextStyle(fontWeight: FontWeight.w800)),
        backgroundColor: Colors.green,
      ));
      context.go(RouteNames.home);
    } else {
      ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
        content: Text('Login failed. Please double-check your password.',
            style: TextStyle(fontWeight: FontWeight.w800)),
        backgroundColor: Colors.red,
      ));
    }
  }

  void _login() async {
    if (_formKey.currentState!.validate()) {
      final success = await _authService.login(_emailController.text, _passwordController.text);
      _showSuccessStatus(success);
    }
  }

  void _loginOAuth(String service) async {
    final success = await _authService.loginWithOAuth(context, service);
    _showSuccessStatus(success);
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
                'Log in',
                style: TextStyle(
                  fontSize: 32,
                  fontWeight: FontWeight.w800,
                ),
              ),
            ),
            const SizedBox(height: 30),
            _buildEmailField(),
            const SizedBox(height: 17),
            _buildPasswordField(),
            const SizedBox(height: 25),
            _buildLoginButton(),
            const SizedBox(height: 10),
            _buildTextDivider('or'),
            const SizedBox(height: 15),
            _buildOAuthButton('Github'),
            const SizedBox(height: 15),
            _buildOAuthButton('Discord'),
            const SizedBox(height: 10),
            _buildSignUpHereLink(),
            const SizedBox(height: 10),
            _buildEditNetworkServer(),
          ],
        ),
      ),
    );
  }

  Widget _buildEmailField() {
    return AuthInputField(
        controller: _emailController,
        hintText: 'Email',
        obscureText: false,
        validator: (value) {
          if (value == null || value.isEmpty) {
            return 'You need to enter a valid email.';
          }
          return null;
        });
  }

  Widget _buildPasswordField() {
    return AuthInputField(
      controller: _passwordController,
      hintText: 'Password',
      obscureText: true,
      validator: (value) {
        if (value == null || value.isEmpty) {
          return 'Please enter a password.';
        }
        if (value.length < 6) {
          return 'Password must be at least 6 characters.';
        }
        return null;
      },
    );
  }

  Widget _buildLoginButton() {
    return AuthButton(text: 'Log in', onPressed: _login);
  }

  Color _getColorFromService(String service) {
    switch (service) {
      case 'Github': return const Color.fromRGBO(36, 41, 46, 1);
      case 'Discord': return const Color.fromRGBO(114, 137, 218, 1);
    }
    return Colors.black;
  }

  Widget _buildOAuthButton(String service) {
    return AuthButton(
      text: 'Continue with $service',
      onPressed: () {
        _loginOAuth(service);
      },
      backgroundColor: _getColorFromService(service),
    );
  }

  Widget _buildEditNetworkServer() {
    return AuthButton(
      text: 'Edit network server',
      onPressed: () { context.push(RouteNames.serverConfig); },
      borderRadius: 10,
    );
  }

  Widget _buildTextDivider(String text) {
    return Row(children: [
      const Expanded(
        child: Divider(color: Colors.grey),
      ),
      if (text.isNotEmpty)
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 8),
          child: Text(
            text,
            style: const TextStyle(
                fontSize: 14, fontWeight: FontWeight.w600, color: Colors.grey),
          ),
        ),
      const Expanded(
        child: Divider(color: Colors.grey),
      ),
    ]);
  }

  Widget _buildSignUpHereLink() {
    return Row(mainAxisAlignment: MainAxisAlignment.center, children: [
      const Text(
        'New to AREA?',
        style: TextStyle(
          fontSize: 16,
          fontWeight: FontWeight.w700,
          color: Colors.black,
        ),
      ),
      TextButton(
        onPressed: () {
          if (mounted) {
            context.push(RouteNames.signup);
          }
        },
        child: const Text(
          'Sign up here.',
          style: TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.w700,
            color: Colors.black,
            decoration: TextDecoration.underline,
          ),
        ),
      )
    ]);
  }
}
