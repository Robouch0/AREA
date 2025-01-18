// lib/widgets/ingredient_input.dart
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'package:my_area_flutter/api/types/area_body.dart';

class IngredientInput extends StatelessWidget {
  final String label;
  final String? hint;
  final bool required;
  final IngredientType type;
  final TextEditingController controller;
  final Function(String) onChanged;
  final dynamic value;

  static const Color _greyInputFill = Color(0xFF303030);

  const IngredientInput({
    super.key,
    required this.label,
    this.hint,
    required this.required,
    required this.type,
    required this.controller,
    required this.onChanged,
    this.value,
  });

  @override
  Widget build(BuildContext context) {
    switch (type) {
      case IngredientType.int:
        return _buildNumberInput(context, [FilteringTextInputFormatter.digitsOnly]);
      case IngredientType.float:
        return _buildNumberInput(context, [FilteringTextInputFormatter.allow(RegExp(r'^\d+(\.\d*)?'))]);
      case IngredientType.bool:
        return _buildBoolInput();
      case IngredientType.date:
        return _buildDateInput(context);
      case IngredientType.time:
        return _buildTimeInput(context);
      case IngredientType.string:
        return _buildTextInput(context);
    }
  }

  Widget _buildNumberInput(BuildContext context, List<TextInputFormatter> textInputFormatters) {
    return TextField(
      controller: controller,
      decoration: _getInputDecoration(),
      keyboardType: TextInputType.number,
      inputFormatters: textInputFormatters,
      style: const TextStyle(color: Colors.white),
      onChanged: onChanged,
      onTapOutside: (event) => FocusScope.of(context).unfocus(),
    );
  }

  Widget _buildBoolInput() {
    bool isChecked = controller.text.toLowerCase() == 'true';

    return InkWell(
      onTap: () {
        final newValue = (!isChecked).toString();
        controller.text = newValue;
        onChanged(newValue);
      },
      child: InputDecorator(
        decoration: _getInputDecoration(),
        child: Row(
          children: [
            Checkbox(
              value: isChecked,
              onChanged: (bool? value) {
                controller.text = value.toString();
                onChanged(value.toString());
              },
            ),
            Text(
              isChecked ? 'True' : 'False',
              style: const TextStyle(color: Colors.white),
            ),
          ],
        ),
      ),
    );
  }

  ThemeData _getDarkTheme(BuildContext context) {
    return Theme.of(context).copyWith(
      colorScheme: const ColorScheme.dark(
        primary: Colors.blue,
        surface: _greyInputFill,
        onPrimary: Colors.white,
        onSurface: Colors.white,
      ),
      dialogBackgroundColor: _greyInputFill,
      timePickerTheme: const TimePickerThemeData(
        helpTextStyle: TextStyle(color: Colors.white),
        dayPeriodTextStyle: TextStyle(color: Colors.white),
        hourMinuteTextStyle: TextStyle(color: Colors.white),
      ),
    );
  }

  Widget _buildDateInput(BuildContext context) {
    return InkWell(
      onTap: () => _showDateTimePicker(context),
      child: InputDecorator(
        decoration: _getInputDecoration(),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text(
              controller.text.isEmpty ? 'Select date and time' : controller.text,
              style: TextStyle(color: controller.text.isEmpty ? Colors.white38 : Colors.white),
            ),
            const Icon(Icons.event_note, color: Colors.white70),
          ],
        ),
      ),
    );
  }

  Future<void> _showDateTimePicker(BuildContext context) async {
    final DateTime? pickedDate = await showDatePicker(
      context: context,
      initialDate: DateTime.now(),
      firstDate: DateTime(2000),
      lastDate: DateTime(2100),
      builder: (context, child) => Theme(data: _getDarkTheme(context), child: child!),
    );

    if (pickedDate != null) {
      if (!context.mounted) return;
      final TimeOfDay? pickedTime = await _showTimePicker(context);

      if (pickedTime != null) {
        final DateTime combinedDateTime = DateTime(
          pickedDate.year,
          pickedDate.month,
          pickedDate.day,
          pickedTime.hour,
          pickedTime.minute,
        );

        final String formattedDateTime = combinedDateTime.toUtc().toIso8601String();
        controller.text = formattedDateTime;
        onChanged(formattedDateTime);
      }
    }
  }

  Widget _buildTimeInput(BuildContext context) {
    return InkWell(
      onTap: () async {
        final time = await _showTimePicker(context);
        if (time != null) {
          controller.text = time.toString();
          onChanged(time.toString());
        }
      },
      child: InputDecorator(
        decoration: _getInputDecoration(),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text(
              controller.text.isEmpty ? 'Select time' : controller.text,
              style: TextStyle(color: controller.text.isEmpty ? Colors.white38 : Colors.white),
            ),
            const Icon(Icons.access_time, color: Colors.white70),
          ],
        ),
      ),
    );
  }

  Future<TimeOfDay?> _showTimePicker(BuildContext context) async {
    return showTimePicker(
      context: context,
      initialTime: TimeOfDay.now(),
      builder: (context, child) => Theme(data: _getDarkTheme(context), child: child!),
    );
  }

  Widget _buildTextInput(BuildContext context) {
    return TextField(
      controller: controller,
      decoration: _getInputDecoration(),
      style: const TextStyle(color: Colors.white),
      onChanged: onChanged,
      onTapOutside: (event) => FocusScope.of(context).unfocus(),
    );
  }

  InputDecoration _getInputDecoration() {
    return InputDecoration(
      labelText: '$label${required ? ' *' : ''}',
      hintText: hint,
      labelStyle: const TextStyle(color: Colors.white70),
      hintStyle: const TextStyle(color: Colors.white30),
      border: OutlineInputBorder(
        borderRadius: BorderRadius.circular(8),
      ),
      filled: true,
      fillColor: Colors.white.withAlpha(25),
    );
  }
}
