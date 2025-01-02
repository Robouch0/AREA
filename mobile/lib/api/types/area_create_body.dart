// lib/api/types/area_create_body.dart

class Service {
  final String service;
  final String microservice;
  final Map<String, dynamic> ingredients;

  Service({
    required this.service,
    required this.microservice,
    required this.ingredients,
  });

  Map<String, dynamic> toJson() {
    return {
      'service': service,
      'microservice': microservice,
      'ingredients': ingredients,
    };
  }
}

class AreaCreateBody {
  final int userId;
  final Service action;
  final Service reaction;

  AreaCreateBody({
    required this.userId,
    required this.action,
    required this.reaction,
  });

  Map<String, dynamic> toJson() {
    return {
      'user_id': userId,
      'action': action.toJson(),
      'reaction': reaction.toJson(),
    };
  }
}
