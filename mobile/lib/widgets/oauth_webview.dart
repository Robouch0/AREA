// lib/widgets/oauth_webview.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'dart:developer' as developer;
import 'package:webview_flutter/webview_flutter.dart';
import 'package:flutter/material.dart';

class OAuthWebViewPage extends StatefulWidget {
  final String service;
  final String apiUrl;
  final String redirectUrl;
  final bool isLogin;
  final String? token;

  const OAuthWebViewPage({
    super.key,
    required this.service,
    required this.apiUrl,
    required this.redirectUrl,
    required this.isLogin,
    this.token,
  });

  @override
  State<OAuthWebViewPage> createState() => _OAuthWebViewPageState();
}

class _OAuthWebViewPageState extends State<OAuthWebViewPage> {
  WebViewController? _controller;

  @override
  void initState() {
    super.initState();
    _setupWebView();
  }

  Future<void> _setupWebView() async {
    try {
      final response = await http.get(
          Uri.parse('${widget.apiUrl}/oauth/${widget.service}').replace(
              queryParameters: {'redirect_uri': widget.redirectUrl}
          )
      );

      if (response.statusCode != 200) {
        developer.log('Cannot get the authorization URL.');
        if (mounted) Navigator.of(context).pop(null);
        return;
      }

      if (!mounted) return;

      final controller = WebViewController()
        ..setUserAgent('Mozilla/5.0 (Linux; Android 10; Redmi Note 8 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.88 Mobile Safari/537.36')
        ..setJavaScriptMode(JavaScriptMode.unrestricted)
        ..setNavigationDelegate(
          NavigationDelegate(
            onNavigationRequest: (NavigationRequest request) {
              if (request.url.startsWith(widget.redirectUrl)) {
                _handleRedirect(request.url);
                return NavigationDecision.prevent;
              }
              return NavigationDecision.navigate;
            },
          ),
        );

      await controller.loadRequest(Uri.parse(response.body));

      if (!mounted) return;
      setState(() {
        _controller = controller;
      });
    } catch (e) {
      developer.log('Failed to initialize WebView: $e');
      if (mounted) Navigator.of(context).pop(null);
    }
  }

  Future<void> _handleRedirect(String url) async {
    try {
      final Uri uri = Uri.parse(url);
      final code = uri.queryParameters['code'];

      if (code == null) {
        developer.log('Cannot find the code to send back.');
        if (mounted) Navigator.of(context).pop(null);
        return;
      }

      final endpoint = widget.isLogin ? 'oauth/' : 'oauth/connect/';
      final headers = {
        'Content-Type': 'application/json',
        if (!widget.isLogin) 'Authorization': 'Bearer ${widget.token}',
      };

      final tokenResponse = await http.post(
          Uri.parse('${widget.apiUrl}/$endpoint'),
          headers: headers,
          body: json.encode({
            'service': widget.service,
            'code': code,
            'redirect_uri': widget.redirectUrl,
          })
      );

      if (tokenResponse.statusCode == 200) {
        developer.log('response: ${tokenResponse.body}');
        if (mounted) Navigator.of(context).pop(tokenResponse.body);
      } else {
        if (mounted) Navigator.of(context).pop(null);
      }
    } catch (e) {
      developer.log('Error handling redirect: $e');
      if (mounted) Navigator.of(context).pop(null);
    }
  }

  @override
  void dispose() {
    if (_controller != null) {
      _controller?.clearCache();
    }
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.isLogin
            ? 'Connect with ${widget.service}'
            : 'Connect ${widget.service}'),
        leading: IconButton(
          icon: const Icon(Icons.close),
          onPressed: () => Navigator.of(context).pop(null),
        ),
      ),
      body: _controller == null
          ? const SizedBox.shrink()
          : WebViewWidget(controller: _controller!),
    );
  }
}
