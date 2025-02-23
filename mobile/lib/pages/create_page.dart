// lib/pages/create_page.dart
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:my_area_flutter/core/router/route_names.dart';
import 'package:my_area_flutter/services/api/area_service.dart';
import 'package:my_area_flutter/api/types/profile_body.dart';
import 'package:my_area_flutter/widgets/oauth_connection_button.dart';
import 'package:my_area_flutter/widgets/main_app_scaffold.dart';
import 'package:my_area_flutter/api/types/area_body.dart';
import 'package:my_area_flutter/api/types/area_create_body.dart';
import 'package:my_area_flutter/api/types/user_provider_list_body.dart';
import 'package:my_area_flutter/widgets/ingredient_input.dart';

class CreateAreaPage extends StatefulWidget {
  final Future<List<AreaServiceData>> services;
  final Future<UserInfoBody> userInfo;
  final Future<UserProviderListBody> userProviders;

  const CreateAreaPage(
      {super.key,
      required this.services,
      required this.userInfo,
      required this.userProviders});

  @override
  State<CreateAreaPage> createState() => _CreateAreaPageState();
}

class ActionData {
  String serviceName;
  String microServiceName;
  Map<String, dynamic> ingredients;
  Map<String, TextEditingController> controllers;

  ActionData({
    this.serviceName = '',
    this.microServiceName = '',
    Map<String, dynamic>? ingredients,
    Map<String, TextEditingController>? controllers,
  })  : ingredients = ingredients ?? {},
        controllers = controllers ?? {};

  void dispose() {
    controllers.forEach((_, controller) => controller.dispose());
  }
}

class _CreateAreaPageState extends State<CreateAreaPage> {
  List<AreaServiceData>? services;
  UserInfoBody? userInfo;
  UserProviderListBody? userProviders;
  late List<AreaServiceData> actions;
  late List<AreaServiceData> reactions;

  ActionData action = ActionData();
  List<ActionData> reactionsList = [ActionData()];

  @override
  void initState() {
    super.initState();
    _loadServices();
  }

  Future<void> _loadServices() async {
    try {
      final loadedServices = await widget.services;
      final loadedUserInfo = await widget.userInfo;
      final loadedUserProviders = await widget.userProviders;
      setState(() {
        services = loadedServices;
        userInfo = loadedUserInfo;
        userProviders = loadedUserProviders;
        userProviders?.providers.add('crypto');
        userProviders?.providers.add('dt');
        userProviders?.providers.add('weather');
        actions = _filterAreaByType(loadedServices, 'action');
        reactions = _filterAreaByType(loadedServices, 'reaction');
      });
    } catch (e) {
      debugPrint('Error loading services: $e');
    }
  }

  List<AreaServiceData> _filterAreaByType(
      List<AreaServiceData> services, String type) {
    return services
        .where((service) =>
            service.microservices.any((micro) => micro.type == type))
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
  void dispose() {
    action.dispose();
    for (var reaction in reactionsList) {
      reaction.dispose();
    }
    super.dispose();
  }

  void _initializeControllers(
      ActionData actionData, Map<String, Ingredient> ingredients) {
    actionData.dispose();
    actionData.controllers.clear();
    actionData.ingredients.clear();

    actionData.ingredients.forEach((key, ingredient) {
      final controller =
          TextEditingController(text: ingredient.value?.toString() ?? '');
      actionData.controllers[key] = controller;
      actionData.ingredients[key] = ingredient.value;
    });
  }

  bool _validateIngredients(
      Map<String, Ingredient> ingredients, Map<String, dynamic> values) {
    return ingredients.entries.every((entry) {
      final ingredient = entry.value;
      if (ingredient.required) {
        final value = values[entry.key];
        return value != null && value.toString().isNotEmpty;
      }
      return true;
    });
  }

  dynamic convertIngredientValue(String value, Ingredient ingredient) {
    if (value.isEmpty) return null;

    switch (ingredient.type) {
      case IngredientType.int:
        return int.tryParse(value) ?? ingredient.value;
      case IngredientType.float:
        return double.tryParse(value) ?? ingredient.value;
      case IngredientType.bool:
        return value.toLowerCase() == 'true';
      case IngredientType.time:
        return value;
      case IngredientType.string:
        return value;
      case IngredientType.date:
        return value;
    }
  }

  void _handleIngredientChange(
      ActionData actionData, String key, String value, Ingredient ingredient) {
    final convertedValue = convertIngredientValue(value, ingredient);
    setState(() {
      actionData.ingredients[key] = convertedValue;
    });
  }

  void _handleSubmit() async {
    final reactions = reactionsList
        .map((reaction) => Service(
              service: reaction.serviceName,
              microservice: reaction.microServiceName,
              ingredients: reaction.ingredients,
            ))
        .toList();

    final newArea = AreaCreateBody(
      userId: userInfo!.userId,
      action: Service(
        service: action.serviceName,
        microservice: action.microServiceName,
        ingredients: action.ingredients,
      ),
      reactions: reactions,
    );

    final success = await AreaService.instance.createArea(newArea);
    _displayScaffoldStatus(success);
  }

  void _displayScaffoldStatus(bool success) {
    if (success) {
      ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
        content: Text('Area created successfully.',
            style: TextStyle(fontWeight: FontWeight.w800)),
        backgroundColor: Colors.green,
      ));
      context.go(RouteNames.home);
    } else {
      ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
        content: Text('Area creation failed.',
            style: TextStyle(fontWeight: FontWeight.w800)),
        backgroundColor: Colors.red,
      ));
    }
  }

  void _addReaction() {
    setState(() {
      reactionsList.add(ActionData());
    });
  }

  void _removeReaction(int index) {
    setState(() {
      reactionsList[index].dispose();
      reactionsList.removeAt(index);
    });
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
              _buildReactionsSection(),
              const SizedBox(height: 10),
              _buildAddReactionButton(),
              const SizedBox(height: 20),
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
            selectedService: action.serviceName,
            onServiceChanged: (value) {
              setState(() {
                action.serviceName = value;
                action.microServiceName = '';
                action.ingredients.clear();
                action.dispose();
                action.controllers.clear();
              });
            },
          ),
          if (action.serviceName.isNotEmpty)
            _buildMicroserviceSection(action, actions),
        ],
      ),
    );
  }

  Widget _buildReactionsSection() {
    return Column(
      children: List.generate(reactionsList.length, (index) {
        return Column(
          children: [
            if (index > 0) const SizedBox(height: 20),
            _buildReactionCard(index),
          ],
        );
      }),
    );
  }

  List<Widget> _buildMultipleReactions(int index) {
    const TextStyle textStyle = TextStyle(
      color: Colors.red,
      fontSize: 32,
      fontWeight: FontWeight.bold,
    );

    if (reactionsList.length == 1) {
      return [const Text('Reaction', style: textStyle)];
    }
    return [
      Text('Reaction #${index + 1}', style: textStyle),
      IconButton(
        icon: const Icon(Icons.delete, color: Colors.red),
        onPressed: () => _removeReaction(index),
      ),
    ];
  }

  Widget _buildReactionCard(int index) {
    final reaction = reactionsList[index];
    return Stack(
      children: [
        Container(
          decoration: BoxDecoration(
            color: Colors.black87,
            borderRadius: BorderRadius.circular(30),
          ),
          padding: const EdgeInsets.all(20),
          child: Column(
            children: [
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: _buildMultipleReactions(index),
              ),
              const SizedBox(height: 20),
              _buildServiceSelection(
                services: reactions,
                selectedService: reaction.serviceName,
                onServiceChanged: (value) {
                  setState(() {
                    reaction.serviceName = value;
                    reaction.microServiceName = '';
                    reaction.ingredients.clear();
                    reaction.controllers
                        .forEach((_, controller) => controller.dispose());
                    reaction.controllers.clear();
                  });
                },
              ),
              if (reaction.serviceName.isNotEmpty)
                _buildMicroserviceSection(reactionsList[index], reactions),
            ],
          ),
        ),
      ],
    );
  }

  Widget _buildConnector() {
    return Center(
      child: Container(
        height: 4,
        width: 100,
        margin: const EdgeInsets.symmetric(horizontal: 10),
        decoration: BoxDecoration(
          color: Colors.grey,
          borderRadius: BorderRadius.circular(2),
        ),
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

  Widget _buildMicroserviceSection(
      ActionData actionData, List<AreaServiceData> actionsData) {
    final selectedService = actionData.serviceName;
    final selectedMicro = actionData.microServiceName;
    final services = actionsData;

    if (selectedService.isEmpty) return const SizedBox.shrink();

    if (userProviders!.providers.contains(selectedService) == false) {
      return Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const SizedBox(height: 10),
            Text(
              'Link your $selectedService account',
              style: const TextStyle(
                color: Colors.red,
                fontSize: 24,
                fontWeight: FontWeight.bold,
              ),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 10),
            OAuthConnectionButton(
              serviceName: selectedService,
              initialProviders: userProviders!.providers,
              onSuccess: () {
                setState(() {
                  userProviders!.providers.add(selectedService);
                });
              },
            ),
          ],
        ),
      );
    }

    final service = services.firstWhere((s) => s.refName == selectedService);
    final selectedMicroService =
    service.microservices.firstWhere((m) => m.refName == selectedMicro,
        orElse: () => MicroServiceBody(
          name: '',
          refName: '',
          type: '',
          ingredients: {},
          pipelines: [],
        ));

    return Column(
      crossAxisAlignment: CrossAxisAlignment.stretch,
      children: [
        _buildMicroserviceGrid(
          service: service,
          selectedMicro: selectedMicro,
          onMicroSelected: (value) {
            setState(() {
              actionData.microServiceName = value;
              actionData.ingredients.clear();
            });
            final selectedMicroService =
            service.microservices.firstWhere((m) => m.refName == value);
            _initializeControllers(
                actionData, selectedMicroService.ingredients);
          },
        ),
        if (selectedMicro.isNotEmpty) ...[
          const SizedBox(height: 15),
          _buildIngredientsForm(
              ingredients: selectedMicroService.ingredients,
              values: actionData.ingredients,
              actionData: actionData),
        ],
      ],
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
        children: service.microservices
            .where((micro) => selectedMicro.isEmpty || micro.refName == selectedMicro)
            .map((micro) => InkWell(
          onTap: () => onMicroSelected(micro.refName == selectedMicro ? '' : micro.refName),
          child: Container(
            width: MediaQuery.of(context).size.width,
            padding: const EdgeInsets.all(12),
            decoration: BoxDecoration(
              color: micro.refName == selectedMicro ? Colors.blue[700] : Colors.blue[900],
              borderRadius: BorderRadius.circular(15),
            ),
            child: Column(
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
        ))
            .toList(),
      ),
    );
  }

  Widget _buildIngredientsFormTitle(String title, double fontSize) {
    return Text(
      title,
      style: TextStyle(
        color: Colors.white70,
        fontSize: fontSize,
        fontWeight: FontWeight.bold,
      ),
    );
  }

  Widget _buildIngredientsForm({
    required Map<String, Ingredient> ingredients,
    required Map<String, dynamic> values,
    required ActionData actionData,
  }) {
    List<dynamic> availablePipelines = [];
    if (actionData.serviceName.isNotEmpty && actionData.microServiceName.isNotEmpty) {
      final service = services!.firstWhere((s) => s.refName == actionData.serviceName);
      final microservice = service.microservices
          .firstWhere((m) => m.refName == actionData.microServiceName);
      availablePipelines = microservice.pipelines;
    }
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        if (ingredients.isNotEmpty) ...[
        Padding(
          padding: const EdgeInsets.only(left: 4, bottom: 8),
          child: _buildIngredientsFormTitle('Parameters', 16),
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
                final ingredient = entry.value;
                final controllers = actionData.controllers;
                final controller = controllers[entry.key] ??
                    TextEditingController(
                        text: ingredient.value?.toString() ?? '');

                if (!controllers.containsKey(entry.key)) {
                  controllers[entry.key] = controller;
                }

                return Padding(
                  padding: const EdgeInsets.only(bottom: 12),
                  child: IngredientInput(
                    label: entry.key,
                    hint: ingredient.description,
                    required: ingredient.required,
                    type: ingredient.type,
                    controller: controller,
                    value: values[entry.key],
                    onChanged: (value) => _handleIngredientChange(
                      actionData,
                      entry.key,
                      value,
                      ingredient,
                    ),
                  ),
                );
              }).toList(),
            ),
          ),
        ),
        ],
        if (availablePipelines.isNotEmpty) ...[
          Padding(
            padding: const EdgeInsets.only(top: 10),
            child: _buildIngredientsFormTitle('Variables available in next reactions', 15),
          ),
          Padding(
            padding: const EdgeInsets.only(top: 10),
            child: Container(
              padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 8),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(8),
                color: Colors.white.withAlpha(25),
              ),
              child: Text(
                availablePipelines.map((p) => '{{.$p}}').join(', '),
                style: const TextStyle(
                  color: Colors.blue,
                  fontSize: 12,
                  fontWeight: FontWeight.bold
                ),
              ),
            ),
          )
        ],
      ],
    );
  }

  Widget _buildAddReactionButton() {
    return Center(
      child: TextButton.icon(
        onPressed: _addReaction,
        icon: const Icon(Icons.add_circle, color: Colors.blue),
        label: const Text(
          'Add Reaction',
          style: TextStyle(color: Colors.blue, fontSize: 16),
        ),
      ),
    );
  }

  bool get isValid {
    if (action.microServiceName.isEmpty) return false;
    final actionService =
        services!.firstWhere((s) => s.refName == action.serviceName);
    final actionMicroservice = actionService.microservices
        .firstWhere((m) => m.refName == action.microServiceName);
    if (!_validateIngredients(
        actionMicroservice.ingredients, action.ingredients)) {
      return false;
    }

    for (var reaction in reactionsList) {
      if (reaction.microServiceName.isEmpty) return false;
      final reactionService =
          services!.firstWhere((s) => s.refName == reaction.serviceName);
      final reactionMicroservice = reactionService.microservices
          .firstWhere((m) => m.refName == reaction.microServiceName);
      if (!_validateIngredients(
          reactionMicroservice.ingredients, reaction.ingredients)) {
        return false;
      }
    }
    return true;
  }

  Widget _buildCreateButton() {
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
            fontWeight: FontWeight.w800,
            color: Colors.white
          ),
        ),
      ),
    );
  }
}
