// lib/api/types/area_body.dart
enum IngredientType { string, int, bool, time, float }

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
  final Map<String, IngredientType> ingredients;

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
        (key, value) => MapEntry(key, _stringToIngredientType(value)),
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
