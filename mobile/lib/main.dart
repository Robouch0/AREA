// lib/main.dart
import 'package:flutter/material.dart';
import 'services/auth_service.dart';
import 'widgets/auth_input_field.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'AREA',
      theme: ThemeData(
        primarySwatch: Colors.grey,
        fontFamily: 'AvenirNextCyr',
      ),
      home: const LoginPage(),
    );
  }
}

class LoginPage extends StatefulWidget {
  const LoginPage({super.key});

  @override
  _LoginPageState createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final _authService = AuthService();
  final _formKey = GlobalKey<FormState>();
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();

  void _performLogin(String email, String pass) async {
    final success = await _authService.login(email, pass);

    if (success) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(
              content: Text(
                  'Success.',
                  style: TextStyle(fontWeight: FontWeight.w800)
              ),
              backgroundColor: Colors.green,
            )
        );
      }
    } else {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text(
                'Login failed. Please double-check your password.',
              style: TextStyle(fontWeight: FontWeight.w800)
            ),
            backgroundColor: Colors.red,
          )
        );
      }
    }
  }

  void _login() {
    if (_formKey.currentState!.validate()) {
      _performLogin(_emailController.text, _passwordController.text);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        title: const Text(
          'AREA',
          style: TextStyle(
            fontSize: 28,
            fontWeight: FontWeight.w700,
          ),
        ),
        backgroundColor: Colors.transparent,
        centerTitle: true,
      ),
      body: SafeArea(
        child: SingleChildScrollView(
          padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 70),
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
      }
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
          return 'Le mot de passe doit faire au moins 6 caractères.';
        }
        return null;
      },
    );
  }

  Widget _buildLoginButton() {
    return ElevatedButton(
      onPressed: _login,
      style: ElevatedButton.styleFrom(
        backgroundColor: Colors.black87,
        padding: const EdgeInsets.symmetric(vertical: 15),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(30),
        ),
      ),
      child: const Text(
        'Log in',
        style: TextStyle(
          fontSize: 20,
          fontWeight: FontWeight.w800,
          color: Colors.white,
        ),
      ),
    );
  }

  Widget _buildForgotPasswordLink() {
    return TextButton(
      onPressed: () {
      },
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
              fontSize: 14,
              fontWeight: FontWeight.w600,
              color: Colors.grey),
        ),
      ),
      Expanded(
        child: Divider(color: Colors.grey),
      ),
    ]);
  }

  Widget _buildSignUpHereLink() {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        const Text(
          'New to AREA?',
          style: TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.w700,
            color: Colors.black,
          ),
        ),
        TextButton(
          onPressed: () {},
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
