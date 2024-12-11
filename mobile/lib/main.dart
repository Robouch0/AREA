// lib/main.dart
import 'package:flutter/material.dart';
import 'package:my_area_flutter/core/router/app_router.dart';
import 'package:my_area_flutter/services/auth_service.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await AuthService.instance.initializeAuth();

  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: 'AREA',
      theme: ThemeData(
        primarySwatch: Colors.grey,
        fontFamily: 'AvenirNextCyr',
      ),
      routerConfig: AppRouter.router,
    );
  }
}