// lib/core/router/app_router.dart
import 'package:go_router/go_router.dart';
import 'package:my_area_flutter/pages/register_page.dart';

import 'package:my_area_flutter/services/auth_service.dart';
import 'package:my_area_flutter/pages/login_page.dart';
import 'package:my_area_flutter/pages/home_page.dart';
import 'package:my_area_flutter/pages/page_not_found.dart';

import 'route_names.dart';

class AppRouter {
  static final GoRouter router = GoRouter(
    initialLocation: '/home',
    redirect: (context, state) {
      final authService = AuthService.instance;
      final isLoggedIn = authService.isLoggedInSync;
      final isLoginPage = state.matchedLocation == RouteNames.login;
      final isSignupPage = state.matchedLocation == RouteNames.signup;

      if (!isLoggedIn && !isLoginPage && !isSignupPage) {
        return RouteNames.login;
      }
      if (isLoggedIn && (isLoginPage || isSignupPage)) {
        return RouteNames.home;
      }
      return null;
    },
    routes: [
      GoRoute(
        path: RouteNames.login,
        builder: (context, state) => const LoginPage(),
      ),
      GoRoute(
        path: RouteNames.home,
        builder: (context, state) => const HomePage(),
      ),
      GoRoute(
        path: RouteNames.signup,
        builder: (context, state) => const RegisterPage(),
      )
    ],
    errorBuilder: (context, state) => const NotFoundPage(),
  );
}