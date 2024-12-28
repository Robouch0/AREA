import 'package:flutter/material.dart';
import 'package:my_area_flutter/widgets/main_app_scaffold.dart';
import 'package:my_area_flutter/api/types/area_body.dart';

class CreateAreaPage extends StatefulWidget {
  final Future<List<AreaServiceData>> services;
  final int uid;

  const CreateAreaPage({
    super.key,
    required this.services,
    required this.uid,
  });

  @override
  State<CreateAreaPage> createState() => _CreateAreaPageState();
}

class _CreateAreaPageState extends State<CreateAreaPage> {
  List<AreaServiceData>? services;
  String actionName = '';
  String microActionName = '';
  String reactionName = '';
  String microReactionName = '';
  Map<String, String> actionIngredients = {};
  Map<String, String> reactionIngredients = {};

  late List<AreaServiceData> actions;
  late List<AreaServiceData> reactions;

  @override
  void initState() {
    super.initState();
    _loadServices();
  }

  Future<void> _loadServices() async {
    try {
      final loadedServices = await widget.services;
      setState(() {
        services = loadedServices;
        actions = _filterAreaByType(loadedServices, 'action');
        reactions = _filterAreaByType(loadedServices, 'reaction');
      });
    } catch (e) {
      debugPrint('Error loading services: $e');
    }
  }

  List<AreaServiceData> _filterAreaByType(List<AreaServiceData> services, String type) {
    return services
        .where((service) => service.microservices.any((micro) => micro.type == type))
        .map((service) => AreaServiceData(
      name: service.name,
      refName: service.refName,
      microservices: service.microservices
          .where((micro) => micro.type == type)
          .toList(),
    ))
        .toList();
  }

  @override
  Widget build(BuildContext context) {
    if (services == null) {
      return const MainAppScaffold(
        child: Center(child: CircularProgressIndicator()),
      );
    }

    return MainAppScaffold(
      child: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              _buildHeader(),
              const SizedBox(height: 30),
              _buildActionSection(),
              const SizedBox(height: 25),
              _buildConnector(),
              const SizedBox(height: 25),
              _buildReactionSection(),
              const SizedBox(height: 30),
              _buildCreateButton(),
              const SizedBox(height: 20),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildHeader() {
    return const Center(
      child: Text(
        'Create your AREA',
        style: TextStyle(
          fontSize: 32,
          fontWeight: FontWeight.w800,
        ),
      ),
    );
  }

  Widget _buildMicroserviceSection(bool isAction) {
    final selectedService = isAction ? actionName : reactionName;
    final selectedMicro = isAction ? microActionName : microReactionName;
    final services = isAction ? actions : reactions;

    if (selectedService.isEmpty) return const SizedBox.shrink();

    final service = services.firstWhere((s) => s.refName == selectedService);
    final selectedMicroService = service.microservices
        .firstWhere((m) => m.refName == selectedMicro, orElse: () => MicroServiceBody(
      name: '',
      refName: '',
      type: '',
      ingredients: {},
    ));

    return Column(
      crossAxisAlignment: CrossAxisAlignment.stretch,
      children: [
        _buildMicroserviceGrid(
          service: service,
          selectedMicro: selectedMicro,
          onMicroSelected: (value) {
            setState(() {
              if (isAction) {
                microActionName = value;
                actionIngredients.clear();
              } else {
                microReactionName = value;
                reactionIngredients.clear();
              }
            });
          },
        ),
        if (selectedMicro.isNotEmpty) ...[
          const SizedBox(height: 20),
          _buildIngredientsForm(
            ingredients: selectedMicroService.ingredients,
            values: isAction ? actionIngredients : reactionIngredients,
            onIngredientChanged: (key, value) {
              setState(() {
                if (isAction) {
                  actionIngredients[key] = value;
                } else {
                  reactionIngredients[key] = value;
                }
              });
            },
          ),
        ],
      ],
    );
  }

  Widget _buildActionSection() {
    return Container(
      decoration: BoxDecoration(
        color: Colors.black87,
        borderRadius: BorderRadius.circular(30),
      ),
      padding: const EdgeInsets.all(20),
      child: Column(
        children: [
          const Text(
            'Action',
            style: TextStyle(
              color: Colors.blue,
              fontSize: 32,
              fontWeight: FontWeight.bold,
            ),
          ),
          const SizedBox(height: 20),
          _buildServiceSelection(
            services: actions,
            selectedService: actionName,
            onServiceChanged: (value) {
              setState(() {
                actionName = value;
                microActionName = '';
                actionIngredients.clear();
              });
            },
          ),
          if (actionName.isNotEmpty)
            _buildMicroserviceSection(true),
        ],
      ),
    );
  }

  Widget _buildConnector() {
    return Container(
      height: 4,
      width: 100,
      margin: const EdgeInsets.symmetric(horizontal: 10),
      decoration: BoxDecoration(
        color: Colors.grey,
        borderRadius: BorderRadius.circular(2),
      ),
    );
  }

  Widget _buildReactionSection() {
    return Container(
      decoration: BoxDecoration(
        color: Colors.black87,
        borderRadius: BorderRadius.circular(30),
      ),
      padding: const EdgeInsets.all(20),
      child: Column(
        children: [
          const Text(
            'Reaction',
            style: TextStyle(
              color: Colors.red,
              fontSize: 32,
              fontWeight: FontWeight.bold,
            ),
          ),
          const SizedBox(height: 20),
          _buildServiceSelection(
            services: reactions,
            selectedService: reactionName,
            onServiceChanged: (value) {
              setState(() {
                reactionName = value;
                microReactionName = '';
                reactionIngredients.clear();
              });
            },
          ),
          if (reactionName.isNotEmpty)
            _buildMicroserviceSection(false),
        ],
      ),
    );
  }

  Widget _buildServiceSelection({
    required List<AreaServiceData> services,
    required String selectedService,
    required Function(String) onServiceChanged,
  }) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 20),
      decoration: BoxDecoration(
        color: Colors.white10,
        borderRadius: BorderRadius.circular(15),
      ),
      child: DropdownButton<String>(
        value: selectedService.isEmpty ? null : selectedService,
        hint: const Text(
          'Select a service',
          style: TextStyle(color: Colors.white70),
        ),
        isExpanded: true,
        dropdownColor: Colors.grey[800],
        underline: const SizedBox(),
        items: services.map((service) {
          return DropdownMenuItem(
            value: service.refName,
            child: Text(
              service.name,
              style: const TextStyle(color: Colors.white),
            ),
          );
        }).toList(),
        onChanged: (value) {
          if (value != null) {
            onServiceChanged(value);
          }
        },
      ),
    );
  }

  Widget _buildMicroserviceGrid({
    required AreaServiceData service,
    required String selectedMicro,
    required Function(String) onMicroSelected,
  }) {
    return Padding(
      padding: const EdgeInsets.only(top: 20),
      child: Wrap(
        spacing: 10,
        runSpacing: 10,
        children: service.microservices.map((micro) {
          final isSelected = micro.refName == selectedMicro;
          return InkWell(
            onTap: () => onMicroSelected(micro.refName),
            child: Container(
              padding: const EdgeInsets.all(12),
              decoration: BoxDecoration(
                color: isSelected ? Colors.blue[700] : Colors.blue[900],
                borderRadius: BorderRadius.circular(15),
              ),
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  Text(
                    micro.name,
                    style: const TextStyle(
                      color: Colors.white,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                  Text(
                    'Service ${micro.refName}',
                    style: const TextStyle(
                      color: Colors.white70,
                      fontSize: 12,
                    ),
                  ),
                ],
              ),
            ),
          );
        }).toList(),
      ),
    );
  }

  Widget _buildIngredientsForm({
    required Map<String, IngredientType> ingredients,
    required Map<String, String> values,
    required Function(String, String) onIngredientChanged,
  }) {
    if (ingredients.isEmpty) return const SizedBox.shrink();
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const Padding(
          padding: EdgeInsets.only(left: 4, bottom: 8),
          child: Text(
            'Parameters',
            style: TextStyle(
              color: Colors.white70,
              fontSize: 16,
              fontWeight: FontWeight.bold,
            ),
          ),
        ),
        Card(
          color: Colors.white10,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(12),
          ),
          child: Padding(
            padding: const EdgeInsets.all(16),
            child: Column(
              children: ingredients.entries.map((entry) {
                return Padding(
                  padding: const EdgeInsets.only(bottom: 12),
                  child: TextField(
                    decoration: InputDecoration(
                      labelText: entry.key,
                      hintText: 'Enter ${entry.key}',
                      labelStyle: const TextStyle(color: Colors.white70),
                      hintStyle: const TextStyle(color: Colors.white30),
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(8),
                      ),
                      filled: true,
                      fillColor: Colors.white.withOpacity(0.1),
                    ),
                    style: const TextStyle(color: Colors.white),
                    onChanged: (value) => onIngredientChanged(entry.key, value),
                    key: ValueKey(entry.key),
                    controller: null,
                    onTapOutside: (event) => FocusScope.of(context).unfocus(),
                  ),
                );
              }).toList(),
            ),
          ),
        ),
      ],
    );
  }

  Widget _buildCreateButton() {
    final bool isValid = microActionName.isNotEmpty &&
        microReactionName.isNotEmpty;

    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20),
      child: ElevatedButton(
        onPressed: isValid ? _handleSubmit : null,
        style: ElevatedButton.styleFrom(
          backgroundColor: Colors.green[500],
          padding: const EdgeInsets.symmetric(vertical: 15),
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(15),
          ),
        ),
        child: const Text(
          'Create AREA',
          style: TextStyle(
            fontSize: 20,
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
    );
  }

  void _handleSubmit() {
    final payload = {
      'user_id': widget.uid,
      'action': {
        'service': actionName,
        'microservice': microActionName,
        'ingredients': actionIngredients,
      },
      'reaction': {
        'service': reactionName,
        'microservice': microReactionName,
        'ingredients': reactionIngredients,
      },
    };

    print(payload);
  }
}
