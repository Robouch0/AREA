// lib/api/types/area_activation_body.dart

class AreaActivationBody {
  final int areaId;
  final bool activated;

  AreaActivationBody({
    required this.areaId,
    required this.activated
  });

  Map<String, dynamic> toJson() {
    return {
      'area_id': areaId,
      'activated': activated,
    };
  }
}
