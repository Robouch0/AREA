// lib/widgets/auth_input_field.dart
import 'package:flutter/material.dart';

class AuthInputField extends StatefulWidget {
  final TextEditingController controller;
  final String hintText;
  final bool obscureText;
  final TextInputType keyboardType;
  final String? Function(String?)? validator;
  final Widget? suffixIcon;

  const AuthInputField({
    super.key,
    required this.controller,
    required this.hintText,
    this.obscureText = false,
    this.keyboardType = TextInputType.text,
    this.validator,
    this.suffixIcon,
  });

  @override
  State<AuthInputField> createState() => _AuthInputFieldState();
}

class _AuthInputFieldState extends State<AuthInputField> {
  bool _isInputVisible = false;

  @override
  void initState() {
    super.initState();
    _isInputVisible = !widget.obscureText;
  }

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      style: const TextStyle(
        fontSize: 20,
        fontWeight: FontWeight.w700,
      ),
      controller: widget.controller,
      onTapOutside: (event) => FocusScope.of(context).unfocus(),
      decoration: widget.obscureText ?
        _hiddenInputDecoration(widget.hintText) :
        _inputDecoration(widget.hintText),
      obscureText: !_isInputVisible,
      keyboardType: widget.keyboardType,
      validator: widget.validator,
    );
  }

  InputDecoration _hiddenInputDecoration(String hintText) {
    return _inputDecoration(
      widget.hintText,
      suffixIcon: IconButton(
        icon: Icon(
          _isInputVisible
              ? Icons.visibility_outlined
              : Icons.visibility_off_outlined,
          color: Colors.black,
        ),
        onPressed: () {
          setState(() {
            _isInputVisible = !_isInputVisible;
          });
        },
      ),
    );
  }

  InputDecoration _inputDecoration(String hintText, {Widget? suffixIcon}) {
    return InputDecoration(
      hintText: hintText,
      suffixIcon: suffixIcon,
      contentPadding: const EdgeInsets.symmetric(vertical: 15, horizontal: 20),
      hintStyle: TextStyle(
        color: Colors.black.withOpacity(0.15),
        fontWeight: FontWeight.w800,
        fontSize: 20,
      ),
      border: OutlineInputBorder(
        borderRadius: BorderRadius.circular(8),
        borderSide: BorderSide(color: Colors.black.withOpacity(0.07), width: 4),
      ),
      enabledBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(8),
        borderSide: BorderSide(color: Colors.black.withOpacity(0.07), width: 4),
      ),
      focusedBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(8),
        borderSide: const BorderSide(color: Colors.black, width: 4),
      ),
      errorBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(8),
        borderSide: const BorderSide(color: Colors.red, width: 4),
      ),
      errorStyle: const TextStyle(color: Colors.red),
    );
  }
}