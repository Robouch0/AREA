// lib/pages/login_page.dart
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import 'package:my_area_flutter/services/api/auth_service.dart';
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

  void _performLogin(String email, String pass) async {
    final success = await _authService.login(email, pass);

    if (!mounted) {
      return;
    }
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

  void _login() {
    if (_formKey.currentState!.validate()) {
      _performLogin(_emailController.text, _passwordController.text);
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
            const SizedBox(height: 5),
            _buildForgotPasswordLink(),
            const SizedBox(height: 25),
            _buildLoginButton(),
            const SizedBox(height: 25),
            _buildTextDivider('or'),
            const SizedBox(height: 15),
            _buildSignUpHereLink()
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

  Widget _buildForgotPasswordLink() {
    return TextButton(
      onPressed: () {},
      child: const Text(
        'Forgot your password?',
        style: TextStyle(
          fontSize: 16,
          fontWeight: FontWeight.w700,
          color: Colors.black,
          decoration: TextDecoration.underline,
        ),
      ),
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
