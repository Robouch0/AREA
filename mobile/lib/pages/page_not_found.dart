// lib/pages/not_found_page.dart
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../core/router/route_names.dart';
import '../widgets/main_app_scaffold.dart';

class NotFoundPage extends StatelessWidget {
  const NotFoundPage({super.key});

  @override
  Widget build(BuildContext context) {
    return MainAppScaffold(
      title: 'Page Not Found',
      child: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text(
              '404',
              style: TextStyle(fontSize: 64, fontWeight: FontWeight.bold),
            ),
            const Text(
              'Page Not Found',
              style: TextStyle(fontSize: 24),
            ),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: () => context.go(RouteNames.home),
              child: const Text('Go to Home'),
            )
          ],
        ),
      ),
    );
  }
}
