// lib/api/types/user_provider_list_body.dart

class UserProviderListBody {
  final List<dynamic> providers;

  UserProviderListBody({
    required this.providers
  });

  factory UserProviderListBody.fromJson(Map<String, dynamic> json) {
    return UserProviderListBody(providers: json['providers'] ?? []);
  }
}
