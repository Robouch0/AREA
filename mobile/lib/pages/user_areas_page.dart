// lib/pages/user_areas_page.dart
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:my_area_flutter/widgets/main_app_scaffold.dart';
import 'package:my_area_flutter/api/types/area_body.dart';

class UserAreasPage extends StatefulWidget {
  final Future<List<UserAreaData>> userAreas;

  const UserAreasPage({super.key, required this.userAreas});

  @override
  State<UserAreasPage> createState() => _UserAreasPageState();
}

class _UserAreasPageState extends State<UserAreasPage> {
  List<UserAreaData> areas = [];
  bool isLoading = true;
  String? errorMessage;

  @override
  void initState() {
    super.initState();
    _loadAreas();
  }

  Future<void> _loadAreas() async {
    try {
      final loadedAreas = await widget.userAreas;
      setState(() {
        areas = loadedAreas;
        isLoading = false;
      });
    } catch (e) {
      setState(() {
        errorMessage = 'Failed to load areas: ${e.toString()}';
        isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    if (isLoading) {
      return const MainAppScaffold(
        child: Center(child: CircularProgressIndicator()),
      );
    }

    if (errorMessage != null) return MainAppScaffold(child: _buildError());
    return MainAppScaffold(child: _buildContent());
  }

  Widget _buildError() {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(
            errorMessage!,
            textAlign: TextAlign.center,
            style: const TextStyle(color: Colors.red),
          ),
          const SizedBox(height: 16),
          ElevatedButton(
            onPressed: () {
              context.push(GoRouterState.of(context).matchedLocation);
            },
            child: const Text('Retry'),
          ),
        ],
      ),
    );
  }

  Widget _buildContent() {
    return SingleChildScrollView(
      padding: const EdgeInsets.all(16),
      child: Column(
        children: [
          const Text(
            'Your AREAs',
            style: TextStyle(fontSize: 38, fontWeight: FontWeight.bold),
            textAlign: TextAlign.center,
          ),
          const SizedBox(height: 16),
          if (areas.isEmpty)
            const EmptyState()
          else
            ...areas.map(
              (area) => Padding(
                padding: const EdgeInsets.only(bottom: 8),
                child: AreaTile(area: area),
              )
            ),
        ],
      ),
    );
  }
}

class EmptyState extends StatelessWidget {
  const EmptyState({super.key});

  @override
  Widget build(BuildContext context) {
    return const Center(
      child: Padding(
        padding: EdgeInsets.all(32),
        child: Text(
          'No AREAs yet - Create your first automation!',
          textAlign: TextAlign.center,
        ),
      ),
    );
  }
}

class AreaTile extends StatelessWidget {
  final UserAreaData area;

  const AreaTile({super.key, required this.area});

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                Expanded(
                  child: Text(
                    '${area.action.name} → ${area.reactions.length > 1
                      ? '${area.reactions.length} reactions'
                      : area.reactions[0].name}',
                    style: const TextStyle(
                        fontSize: 16, fontWeight: FontWeight.bold
                    ),
                  ),
                ),
                IconButton(
                  icon: const Icon(Icons.delete_outline),
                  onPressed: () => _showDeleteDialog(context),
                ),
              ],
            ),
            ServiceInfo(service: area.action, isAction: true),
            ...area.reactions.map(
                (reaction) => ServiceInfo(service: reaction, isAction: false)),
          ],
        ),
      ),
    );
  }

  void _showDeleteDialog(BuildContext context) {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('Delete AREA'),
        content: const Text('Are you sure you want to delete this AREA?'),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('Cancel'),
          ),
          TextButton(
            onPressed: () {
              // TODO: Implement delete API call
              Navigator.pop(context);
            },
            style: TextButton.styleFrom(foregroundColor: Colors.red),
            child: const Text('Delete'),
          ),
        ],
      ),
    );
  }
}

class ServiceInfo extends StatelessWidget {
  final AreaServiceData service;
  final bool isAction;

  const ServiceInfo({
    super.key,
    required this.service,
    required this.isAction,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 4),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            '${isAction ? "If" : "Then"}: ${service.microservices.map((m) => m.name).join(", ")}',
            style: TextStyle(
              color: isAction ? Colors.blue : Colors.red,
              fontWeight: FontWeight.w500,
            ),
          ),
        ],
      ),
    );
  }
}
