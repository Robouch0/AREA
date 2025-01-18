import 'package:flutter/material.dart';
import 'package:my_area_flutter/services/api/auth_service.dart';

class OAuthConnectionButton extends StatefulWidget {
  final String serviceName;
  final List<dynamic> initialProviders;
  final VoidCallback onSuccess;

  const OAuthConnectionButton({
    super.key,
    required this.serviceName,
    required this.initialProviders,
    required this.onSuccess,
  });

  @override
  State<OAuthConnectionButton> createState() => _OAuthConnectionButtonState();
}

class _OAuthConnectionButtonState extends State<OAuthConnectionButton> {
  late bool isLinked;

  @override
  void initState() {
    super.initState();
    isLinked = widget.initialProviders.contains(widget.serviceName);
  }

  Future<void> _performOAuth() async {
    final success = await AuthService.instance.connectOAuthService(context, widget.serviceName);

    if (!mounted) return;
    if (success) {
      setState(() => isLinked = true);
      ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
        content: Text('Service linked successfully.', style: TextStyle(fontWeight: FontWeight.w800)),
        backgroundColor: Colors.green,
      ));
      widget.onSuccess();
    } else {
      ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
        content: Text('Service linking failed.', style: TextStyle(fontWeight: FontWeight.w800)),
        backgroundColor: Colors.red,
      ));
    }
  }

  Future<void> _unlinkOAuth() async {
    await AuthService.instance.unlinkOAuthService(widget.serviceName);

    if (!mounted) return;
    setState(() => isLinked = false);
    ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
      content: Text('Service unlinked successfully.', style: TextStyle(fontWeight: FontWeight.w800)),
      backgroundColor: Colors.green,
    ));
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 32,
      child: ElevatedButton(
        onPressed: isLinked ? _unlinkOAuth : _performOAuth,
        style: ElevatedButton.styleFrom(
          backgroundColor: isLinked ? Colors.red : Colors.green,
          minimumSize: const Size(100, 0),
          padding: const EdgeInsets.symmetric(),
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
          alignment: Alignment.center,
        ),
        child: Text(
          isLinked ? 'Unlink' : 'Link',
          style: const TextStyle(
            color: Colors.black,
            fontSize: 20,
            fontWeight: FontWeight.w800,
          ),
        ),
      ),
    );
  }
}
