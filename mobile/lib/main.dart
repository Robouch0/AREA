// lib/main.dart
import 'package:flutter/material.dart';
import 'package:my_area_flutter/core/router/app_router.dart';
import 'package:my_area_flutter/services/api/auth_service.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await initializeAll();
  runApp(const MyApp());
}

Future<void> initializeAll() async {
  await AuthService.instance.initializeAuth();
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
        colorScheme: ColorScheme.fromSeed(
          seedColor: Colors.white,
          brightness: Brightness.light,
        ),
      ),
      routerConfig: AppRouter.router,
    );
  }
}
