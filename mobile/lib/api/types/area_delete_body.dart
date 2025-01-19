// lib/api/types/area_delete_body.dart

class AreaDeleteBody {
  final int areaId;

  AreaDeleteBody({
    required this.areaId,
  });

  Map<String, dynamic> toJson() {
    return {
      'area_id': areaId,
    };
  }
}
