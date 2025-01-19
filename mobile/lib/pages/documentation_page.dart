// lib/pages/documentation_page.dart
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_markdown/flutter_markdown.dart';
import 'package:go_router/go_router.dart';

class DocumentationPage extends StatelessWidget {
  const DocumentationPage({super.key});

  Future<String> _loadMarkdownContent() async {
    return await rootBundle.loadString('assets/userDocumentation.md');
  }

  Widget _buildImage(Uri uri, String? title, String? alt) {
    final name = uri.path;
    final imagePath = 'assets/$name';

    return Image.asset(
      imagePath,
      errorBuilder: (context, error, stackTrace) {
        return Container(
          padding: const EdgeInsets.all(8.0),
          decoration: BoxDecoration(
            color: Theme.of(context).colorScheme.errorContainer,
            borderRadius: BorderRadius.circular(4),
          ),
          child: Text(
            'Image not found: $name',
            style: TextStyle(
              color: Theme.of(context).colorScheme.onErrorContainer,
            ),
          ),
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Documentation'),
      ),
      body: FutureBuilder<String>(
        future: _loadMarkdownContent(),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          }
          if (snapshot.hasError) {
            context.pop();
          }
          return Markdown(
            data: snapshot.data!,
            imageBuilder: _buildImage,
          );
        },
      ),
    );
  }
}
