// lib/main.dart
import 'package:flutter/material.dart';
import 'package:my_area_flutter/core/router/app_router.dart';
import 'package:my_area_flutter/services/storage/secure_storage_service.dart';
import 'package:my_area_flutter/services/api/auth_service.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await dotenv.load(fileName: ".env");
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
