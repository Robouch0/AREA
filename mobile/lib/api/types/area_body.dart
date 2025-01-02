// lib/api/types/area_body.dart
enum IngredientType { string, int, bool, time, float }

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
    return AreaServiceData(
      name: json['name'],
      refName: json['ref_name'],
      microservices: (json['microservices'] as List)
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

  MicroServiceBody({
    required this.name,
    required this.refName,
    required this.type,
    required this.ingredients,
  });

  factory MicroServiceBody.fromJson(Map<String, dynamic> json) {
    return MicroServiceBody(
      name: json['name'],
      refName: json['ref_name'],
      type: json['type'],
      ingredients: (json['ingredients'] as Map<String, dynamic>).map(
          (key, value) => MapEntry(key, Ingredient.fromJson(value as Map<String, dynamic>)
        ),
      ),
    );
  }
}

IngredientType _stringToIngredientType(String value) {
  switch (value) {
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
    default:
      throw Exception('Unknown ingredient type: $value');
  }
}

class UserAreaData {
  final int id;
  final AreaServiceData action;
  final List<AreaServiceData> reactions;

  UserAreaData({
    required this.id,
    required this.action,
    required this.reactions,
  });

  factory UserAreaData.fromJson(Map<String, dynamic> json) {
    return UserAreaData(
      id: json['ID'],
      action: AreaServiceData.fromJson(json['Action']),
      reactions: (json['Reactions'] as List)
          .map((r) => AreaServiceData.fromJson(r))
          .toList(),
    );
  }
}
