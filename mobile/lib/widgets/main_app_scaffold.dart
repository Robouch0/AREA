// lib/widgets/main_app_scaffold.dart
import 'package:flutter/material.dart';

class MainAppScaffold extends StatelessWidget {
  final String title;
  final Widget child;

  const MainAppScaffold({
    super.key,
    this.title = 'AREA',
    required this.child,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: Text(
          title,
          style: const TextStyle(
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
            child: child),
      ),
    );
  }
}
