import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:my_area_flutter/core/router/route_names.dart';

class BottomNavbar extends StatefulWidget {
  final Widget child;

  const BottomNavbar({
    super.key,
    required this.child,
  });

  @override
  State<BottomNavbar> createState() => _BottomNavbarState();
}

class _BottomNavbarState extends State<BottomNavbar> {
  NavigationDestinationLabelBehavior labelBehavior = NavigationDestinationLabelBehavior.onlyShowSelected;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: widget.child,
      bottomNavigationBar: NavigationBar(
        labelBehavior: labelBehavior,
        selectedIndex: _calculateSelectedIndex(GoRouterState.of(context)),
        onDestinationSelected: (index) {
          context.go(RouteNames.navigationPaths[index]);
        },
        destinations: const [
          NavigationDestination(
            icon: Icon(Icons.home_outlined),
            selectedIcon: Icon(Icons.home),
            label: 'Home',
          ),
          NavigationDestination(
            icon: Icon(Icons.add_circle_outline),
            selectedIcon: Icon(Icons.add_circle),
            label: 'Create',
          ),
          NavigationDestination(
            icon: Icon(Icons.person_outline),
            selectedIcon: Icon(Icons.person),
            label: 'Profile',
          ),
        ],
      ),
    );
  }

  int _calculateSelectedIndex(GoRouterState state) {
    final location = state.matchedLocation;

    return RouteNames.navigationPaths.indexOf(location);
  }
}
