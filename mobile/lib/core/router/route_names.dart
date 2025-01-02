// lib/core/router/route_names.dart

class RouteNames {
  static const login = '/login';
  static const signup = '/signup';

  static const List<String> navigationPaths = [
    '/home',
    '/create',
    '/profile'
  ];

  static String get home => navigationPaths[0];
  static String get create => navigationPaths[1];
  static String get profile => navigationPaths[2];
}
