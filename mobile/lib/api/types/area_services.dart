enum IngredientType { string, int, bool, time }

class AreaService {
  final String name;
  final String refName;
  final List<MicroService> microservices;

  AreaService({
    required this.name,
    required this.refName,
    required this.microservices,
  });

  factory AreaService.fromJson(Map<String, dynamic> json) {
    return AreaService(
      name: json['name'],
      refName: json['ref_name'],
      microservices: (json['microservices'] as List)
          .map((m) => MicroService.fromJson(m))
          .toList(),
    );
  }
}

class MicroService {
  final String name;
  final String refName;
  final String type;
  final Map<String, IngredientType> ingredients;

  MicroService({
    required this.name,
    required this.refName,
    required this.type,
    required this.ingredients,
  });

  factory MicroService.fromJson(Map<String, dynamic> json) {
    return MicroService(
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
    case 'string': return IngredientType.string;
    case 'int': return IngredientType.int;
    case 'bool': return IngredientType.bool;
    case 'time': return IngredientType.time;
    default: throw Exception('Unknown ingredient type: $value');
  }
}