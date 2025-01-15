// lib/core/router/app_router.dart
import 'package:go_router/go_router.dart';
import 'package:my_area_flutter/pages/create_page.dart';
import 'package:my_area_flutter/pages/profile_page.dart';
import 'package:my_area_flutter/pages/register_page.dart';
import 'package:my_area_flutter/services/api/area_service.dart';

import 'package:my_area_flutter/services/api/auth_service.dart';
import 'package:my_area_flutter/services/api/profile_service.dart';

import 'package:my_area_flutter/pages/login_page.dart';
import 'package:my_area_flutter/pages/user_areas_page.dart';
import 'package:my_area_flutter/pages/page_not_found.dart';
import 'package:my_area_flutter/widgets/bottom_navbar.dart';

import 'route_names.dart';

class AppRouter {
  static final GoRouter router = GoRouter(
    initialLocation: RouteNames.home,
    redirect: (context, state) {
      final authService = AuthService.instance;
      final isLoggedIn = authService.isLoggedInSync;
      final isLoginPage = state.matchedLocation == RouteNames.login;
      final isSignupPage = state.matchedLocation == RouteNames.signup;
      final isAuthPage = isLoginPage || isSignupPage;

      if (!isLoggedIn && !isAuthPage) {
        return RouteNames.login;
      }
      return null;
    },
    routes: [
      GoRoute(
        path: RouteNames.login,
        builder: (context, state) => const LoginPage(),
      ),
      GoRoute(
        path: RouteNames.signup,
        builder: (context, state) => const RegisterPage(),
      ),
      ShellRoute(
        builder: (context, state, child) {
          final isLoginPage = state.matchedLocation == RouteNames.login;
          final isSignupPage = state.matchedLocation == RouteNames.signup;
          final isAuthPage = isLoginPage || isSignupPage;

          if (isAuthPage) {
            return child;
          }
          return BottomNavbar(child: child);
        },
        routes: [
          GoRoute(
            path: RouteNames.home,
            builder: (context, state) => UserAreasPage(userAreas: AreaService.instance.listUserAreas()),
          ),
          GoRoute(
            path: RouteNames.create,
            builder: (context, state) => CreateAreaPage(services: AreaService.instance.listAreas(), userInfo: ProfileService.instance.getUserInfo()),
          ),
          GoRoute(
            path: RouteNames.profile,
            builder: (context, state) => ProfilePage(
                userInfo: ProfileService.instance.getUserInfo(),
                oauthList: ProfileService.instance.getOAuthList(),
                userProviders: ProfileService.instance.getUserProviders()
            ),
          ),
        ],
      ),
    ],
    errorBuilder: (context, state) => const NotFoundPage(),
  );
}
