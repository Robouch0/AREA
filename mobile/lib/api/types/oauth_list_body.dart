// lib/api/types/oauth_list_body.dart

class OAuthListBody {
  final List<dynamic> services;

  OAuthListBody({
    required this.services
  });

  factory OAuthListBody.fromJson(Map<String, dynamic> json) {
    return OAuthListBody(services: json['services']);
  }
}
