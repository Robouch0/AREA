// lib/widgets/main_app_scaffold.dart
import 'package:flutter/material.dart';

class MainAppScaffold extends StatelessWidget {
  final Widget child;

  const MainAppScaffold({
    super.key,
    required this.child,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: SafeArea(
        child: SingleChildScrollView(
          padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 40),
          child: child
        ),
      ),
    );
  }
}
