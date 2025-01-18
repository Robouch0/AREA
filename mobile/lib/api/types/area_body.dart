// lib/api/types/area_body.dart
enum IngredientType { string, int, bool, time, float, date }

class Ingredient {
  final IngredientType type;
  final dynamic value;
  final String description;
  final bool required;

  Ingredient({
    required this.type,
    required this.value,
    required this.description,
    required this.required,
  });

  factory Ingredient.fromJson(Map<String, dynamic> json) {
    return Ingredient(
      type: _stringToIngredientType(json['type']),
      value: json['value'],
      description: json['description'],
      required: json['required'],
    );
  }
}

class AreaServiceData {
  final String name;
  final String refName;
  final List<MicroServiceBody> microservices;

  AreaServiceData({
    required this.name,
    required this.refName,
    required this.microservices,
  });

  factory AreaServiceData.fromJson(Map<String, dynamic> json) {
    var microservicesList = json['microservices'] as List;
    return AreaServiceData(
      name: json['name'] as String,
      refName: json['ref_name'] as String,
      microservices: microservicesList
          .map((m) => MicroServiceBody.fromJson(m))
          .toList(),
    );
  }
}

class MicroServiceBody {
  final String name;
  final String refName;
  final String type;
  final Map<String, Ingredient> ingredients;
  final List<dynamic> pipelines;

  MicroServiceBody({
    required this.name,
    required this.refName,
    required this.type,
    required this.ingredients,
    required this.pipelines,
  });

  factory MicroServiceBody.fromJson(Map<String, dynamic> json) {
    Map<String, Ingredient> ingredientsMap = {};
    if (json['ingredients'] != null) {
      final ingredients = json['ingredients'] as Map<String, dynamic>;
      ingredientsMap = ingredients.map(
            (key, value) => MapEntry(key, Ingredient.fromJson(value as Map<String, dynamic>)),
      );
    }

    return MicroServiceBody(
      name: json['name'] as String,
      refName: json['ref_name'] as String,
      type: json['type'] as String,
      ingredients: ingredientsMap,
      pipelines: json['pipeline_available'] != null
          ? json['pipeline_available'] as List<dynamic>
          : [],
    );
  }
}

IngredientType _stringToIngredientType(String value) {
  switch (value.toLowerCase()) {
    case 'string':
      return IngredientType.string;
    case 'int':
      return IngredientType.int;
    case 'bool':
      return IngredientType.bool;
    case 'time':
      return IngredientType.time;
    case 'float':
      return IngredientType.float;
    case 'date':
      return IngredientType.date;
    default:
      throw Exception('Unknown ingredient type: $value');
  }
}

class UserAreaData {
  final int id;
  final AreaServiceData action;
  final List<AreaServiceData> reactions;
  final bool activated;

  UserAreaData({
    required this.id,
    required this.action,
    required this.reactions,
    required this.activated,
  });

  factory UserAreaData.fromJson(Map<String, dynamic> json) {
    return UserAreaData(
      id: json['id'] as int,
      action: AreaServiceData.fromJson(json['action'] as Map<String, dynamic>),
      reactions: (json['reactions'] as List)
          .map((r) => AreaServiceData.fromJson(r as Map<String, dynamic>))
          .toList(),
      activated: json['activated'] as bool,
    );
  }
}
