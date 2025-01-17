// lib/pages/Register_page.dart
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import 'package:my_area_flutter/services/api/auth_service.dart';
import 'package:my_area_flutter/widgets/auth_input_field.dart';
import 'package:my_area_flutter/widgets/auth_button.dart';
import 'package:my_area_flutter/widgets/main_app_scaffold.dart';
import 'package:my_area_flutter/core/router/route_names.dart';

class RegisterPage extends StatefulWidget {
  const RegisterPage({super.key});

  @override
  State<RegisterPage> createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {
  final _authService = AuthService.instance;
  final _formKey = GlobalKey<FormState>();
  final _firstNameController = TextEditingController();
  final _lastNameController = TextEditingController();
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();

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
        content: Text('Register failed.',
            style: TextStyle(fontWeight: FontWeight.w800)),
        backgroundColor: Colors.red,
      ));
    }
  }

  void _register() async {
    if (_formKey.currentState!.validate()) {
      final success = await _authService.createAccount(
          _emailController.text,
          _passwordController.text,
          _firstNameController.text,
          _lastNameController.text);
      _showSuccessStatus(success);
    }
  }

  void _registerOAuth(String service) async {
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
                'Sign up',
                style: TextStyle(
                  fontSize: 32,
                  fontWeight: FontWeight.w800,
                ),
              ),
            ),
            const SizedBox(height: 30),
            _buildFirstNameField(),
            const SizedBox(height: 17),
            _buildLastNameField(),
            const SizedBox(height: 17),
            _buildEmailField(),
            const SizedBox(height: 17),
            _buildPasswordField(),
            const SizedBox(height: 25),
            _buildRegisterButton(),
            const SizedBox(height: 25),
            _buildTextDivider('or'),
            const SizedBox(height: 15),
            _buildOAuthButton('Github'),
            const SizedBox(height: 15),
            _buildOAuthButton('Discord'),
            _buildSignUpHereLink()
          ],
        ),
      ),
    );
  }

  Widget _buildFirstNameField() {
    return AuthInputField(
      controller: _firstNameController,
      hintText: 'First name',
      obscureText: false,
      validator: (value) {
        if (value == null || value.isEmpty) {
          return 'First name is required.';
        }
        return null;
      },
    );
  }

  Widget _buildLastNameField() {
    return AuthInputField(
      controller: _lastNameController,
      hintText: 'Last name',
      obscureText: false,
      validator: (value) {
        if (value == null || value.isEmpty) {
          return 'Last name is required.';
        }
        return null;
      },
    );
  }

  Widget _buildEmailField() {
    return AuthInputField(
      controller: _emailController,
      hintText: 'Email',
      obscureText: false,
      keyboardType: TextInputType.emailAddress,
      validator: (value) {
        if (value == null || value.isEmpty || !value.contains('@')) {
          return 'Invalid email address.';
        }
        return null;
      },
    );
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
          return 'Password must be at least 6 characters long.';
        }
        return null;
      },
    );
  }

  Widget _buildRegisterButton() {
    return AuthButton(text: 'Get started', onPressed: _register);
  }

  Widget _buildOAuthButton(String service) {
    return AuthButton(
        text: 'Continue with $service',
        onPressed: () {
          _registerOAuth(service);
        }
    );
  }

  Widget _buildTextDivider(String text) {
    return const Row(children: [
      Expanded(
        child: Divider(color: Colors.grey),
      ),
      Padding(
        padding: EdgeInsets.symmetric(horizontal: 8),
        child: Text(
          'or',
          style: TextStyle(
              fontSize: 14, fontWeight: FontWeight.w600, color: Colors.grey),
        ),
      ),
      Expanded(
        child: Divider(color: Colors.grey),
      ),
    ]);
  }

  Widget _buildSignUpHereLink() {
    return Row(mainAxisAlignment: MainAxisAlignment.center, children: [
      const Text(
        'Already on AREA?',
        style: TextStyle(
          fontSize: 16,
          fontWeight: FontWeight.w700,
          color: Colors.black,
        ),
      ),
      TextButton(
        onPressed: () {
          if (mounted) {
            context.push(RouteNames.login);
          }
        },
        child: const Text(
          'Log in here.',
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
