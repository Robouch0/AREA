import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:go_router/go_router.dart';
import 'package:my_area_flutter/api/types/oauth_list_body.dart';
import 'package:my_area_flutter/api/types/user_provider_list_body.dart';
import 'package:my_area_flutter/core/router/route_names.dart';
import 'package:my_area_flutter/services/api/profile_service.dart';
import 'package:my_area_flutter/widgets/auth_button.dart';
import 'package:my_area_flutter/widgets/main_app_scaffold.dart';
import 'package:my_area_flutter/widgets/oauth_connection_button.dart';
import 'package:my_area_flutter/api/types/profile_body.dart';
import 'package:my_area_flutter/services/api/auth_service.dart';

class ProfilePage extends StatefulWidget {
  final Future<UserInfoBody> userInfo;
  final Future<OAuthListBody> oauthList;
  final Future<UserProviderListBody> userProviders;

  const ProfilePage({
    super.key,
    required this.userInfo,
    required this.oauthList,
    required this.userProviders,
  });

  @override
  State<ProfilePage> createState() => _ProfilePageState();
}

class _ProfilePageState extends State<ProfilePage> {
  bool userInfosLoaded = false;

  final _formKey = GlobalKey<FormState>();
  final _firstNameController = TextEditingController();
  final _lastNameController = TextEditingController();
  bool isEditing = false;

  late UserInfoBody userInfo;
  late OAuthListBody oauthList;
  late UserProviderListBody userProviders;

  @override
  void initState() {
    super.initState();
    _loadUserInfos();
    _loadOauthList();
  }

  Future<void> _loadUserInfos() async {
    try {
      final loadedUserInfos = await widget.userInfo;
      final loadedUserProviders = await widget.userProviders;
      setState(() {
        userInfosLoaded = true;
        userInfo = loadedUserInfos;
        userProviders = loadedUserProviders;
      });
      _firstNameController.text = userInfo.firstName;
      _lastNameController.text = userInfo.lastName;
    } catch (e) {
      debugPrint('Error loading user infos: $e');
    }
  }

  Future<void> _loadOauthList() async {
    final loadedOauthList = await widget.oauthList;
    setState(() {
      oauthList = loadedOauthList;
    });
  }

  void _toggleEdit() {
    setState(() {
      if (isEditing) {
        _firstNameController.text = userInfo.firstName;
        _lastNameController.text = userInfo.lastName;
      }
      isEditing = !isEditing;
    });
  }

  void _performUpdate(String firstName, String lastName) async {
    bool success = await ProfileService.instance.updateUserInfo(firstName, lastName);

    if (!mounted) return;
    if (success) {
      ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
        content: Text('Profile updated successfully'),
        backgroundColor: Colors.green,
      ));
      context.push(RouteNames.profile);
    } else {
      ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
        content: Text('Update failed'),
        backgroundColor: Colors.red,
      ));
    }
  }

  void _handleUpdate() {
    if (_formKey.currentState!.validate()) {
      _performUpdate(_firstNameController.text, _lastNameController.text);
    }
  }

  final Map<String, IconData> serviceIcons = {
    'github': FontAwesomeIcons.github,
    'google': FontAwesomeIcons.google,
    'spotify': FontAwesomeIcons.spotify,
    'discord': FontAwesomeIcons.discord,
    'gitlab': FontAwesomeIcons.gitlab
  };

  List<Map<String, dynamic>> getServicesList() {
    if (!userInfosLoaded || oauthList.services.isEmpty) {
      return [];
    }

    return oauthList.services.map((service) {
      final serviceLower = service.toLowerCase();
      return {
        "name": service,
        "icon": serviceIcons[serviceLower] ?? FontAwesomeIcons.link
      };
    }).toList();
  }

  @override
  Widget build(BuildContext context) {
    if (userInfosLoaded == false) {
      return const MainAppScaffold(
        child: Center(child: CircularProgressIndicator()),
      );
    }
    return MainAppScaffold(
      child: SingleChildScrollView(
          child: Column(
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          _buildHeader(),
          _buildDivider(),
          _buildMainContent(),
          _buildLogoutButton()
        ],
      )),
    );
  }

  Widget _buildHeader() {
    return const Center(
      child: FittedBox(
        fit: BoxFit.scaleDown,
        child: Text(
          'Account Settings',
          style: TextStyle(
            fontSize: 38,
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
    );
  }

  Widget _buildDivider() {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 24),
      child: Center(
        child: Container(
          width: MediaQuery.of(context).size.width / 3,
          height: 1,
          color: Colors.grey.withAlpha(51),
        ),
      ),
    );
  }

  Widget _buildMainContent() {
    return Container(
      width: MediaQuery.of(context).size.width,
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.grey[900],
        borderRadius: BorderRadius.circular(30),
      ),
      child: LayoutBuilder(
        builder: (context, constraints) {
          if (constraints.maxWidth > 1000) {
            return Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Expanded(child: _buildUserInfoSection()),
                const SizedBox(width: 24),
                Expanded(child: _buildLinkedAccountsSection()),
              ],
            );
          } else {
            return Column(
              children: [
                _buildUserInfoSection(),
                const SizedBox(height: 24),
                _buildLinkedAccountsSection(),
              ],
            );
          }
        },
      ),
    );
  }

  Widget _buildUserInfoSection() {
    return Form(
      key: _formKey,
      child: Container(
        padding: const EdgeInsets.all(16),
        decoration: BoxDecoration(
          color: Colors.grey[700],
          borderRadius: BorderRadius.circular(30),
        ),
        child: Column(
          children: [
            _buildProfilePicture(),
            _buildInfoField('Email', userInfo.email),
            _buildInfoField('First name', userInfo.firstName,
                controller: _firstNameController),
            _buildInfoField('Last name', userInfo.lastName,
                controller: _lastNameController),
            const SizedBox(height: 16),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Expanded(
                  child: ElevatedButton(
                    onPressed: _toggleEdit,
                    style: ElevatedButton.styleFrom(
                      backgroundColor: Colors.grey[900],
                      padding: const EdgeInsets.symmetric(vertical: 17),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(15),
                      ),
                    ),
                    child: Text(
                      isEditing ? 'Cancel' : 'Edit Profile',
                      style: TextStyle(
                          fontSize: 20,
                          fontWeight: FontWeight.w700,
                          color: isEditing ? Colors.red : Colors.white),
                    ),
                  ),
                ),
                if (isEditing) ...[
                  const SizedBox(width: 16),
                  Expanded(
                    child: ElevatedButton(
                      onPressed: _handleUpdate,
                      style: ElevatedButton.styleFrom(
                        backgroundColor: Colors.green,
                        padding: const EdgeInsets.symmetric(vertical: 17),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(15),
                        ),
                      ),
                      child: const Text(
                        'Save',
                        style: TextStyle(
                            fontSize: 20,
                            fontWeight: FontWeight.w700,
                            color: Colors.white),
                      ),
                    ),
                  ),
                ],
              ],
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildProfilePicture() {
    return Container(
      margin: const EdgeInsets.symmetric(vertical: 16),
      decoration: BoxDecoration(
        shape: BoxShape.circle,
        border: Border.all(color: Colors.black, width: 4),
      ),
      child: const CircleAvatar(
        radius: 40,
        backgroundColor: Color(0xffE6E6E6),
        child: Icon(
          Icons.person,
          color: Color(0xffCCCCCC),
          size: 80,
        ),
      ),
    );
  }

  Widget _buildInfoField(String label, String value,
      {TextEditingController? controller}) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 16),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            label,
            style: const TextStyle(
              color: Colors.white,
              fontSize: 20,
              fontWeight: FontWeight.bold,
            ),
          ),
          const SizedBox(height: 8),
          Container(
            width: double.infinity,
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
            decoration: BoxDecoration(
              color: Colors.white.withAlpha(204),
              borderRadius: BorderRadius.circular(20),
              border: Border.all(color: Colors.black, width: 4),
            ),
            child: isEditing && controller != null
                ? TextFormField(
                    controller: controller,
                    decoration: const InputDecoration(
                      border: InputBorder.none,
                      isDense: true,
                      contentPadding: EdgeInsets.zero,
                    ),
                    style: const TextStyle(
                      fontSize: 16,
                      fontWeight: FontWeight.w800,
                    ),
                    validator: (value) {
                      if (value == null || value.isEmpty) {
                        return 'This field cannot be empty';
                      }
                      return null;
                    },
                  )
                : Text(
                    value,
                    style: const TextStyle(
                      fontSize: 16,
                      fontWeight: FontWeight.w800,
                    ),
                  ),
          ),
        ],
      ),
    );
  }

  Widget _buildLinkedAccountsSection() {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.grey[700],
        borderRadius: BorderRadius.circular(30),
      ),
      child: Column(
        children: [
          const Text(
            'Linked Accounts',
            style: TextStyle(
              color: Colors.white,
              fontSize: 28,
              fontWeight: FontWeight.bold,
            ),
          ),
          const SizedBox(height: 24),
          const Text(
            'You can manage here all your external accounts linked to AREA',
            textAlign: TextAlign.center,
            style: TextStyle(
              color: Colors.white,
              fontSize: 16,
              fontWeight: FontWeight.bold,
            ),
          ),
          const SizedBox(height: 12),
          const Text(
            'Scroll through the supported services',
            style: TextStyle(
              color: Colors.white,
              fontSize: 14,
            ),
          ),
          const SizedBox(height: 16),
          _buildServicesScroll(),
        ],
      ),
    );
  }

  Widget _buildServicesScroll() {
    final services = getServicesList();

    return Container(
      height: 280,
      width: double.infinity,
      decoration: BoxDecoration(
        color: Colors.white.withAlpha(230),
        borderRadius: BorderRadius.circular(8),
        border: Border.all(color: Colors.grey[300]!),
      ),
      child: Column(
        children: [
          const Padding(
            padding: EdgeInsets.all(12),
            child: Text(
              'Services',
              style: TextStyle(
                fontSize: 20,
                fontWeight: FontWeight.bold,
              ),
            ),
          ),
          Expanded(
            child: ListView.separated(
              padding: const EdgeInsets.symmetric(horizontal: 8),
              itemCount: services.length,
              separatorBuilder: (context, index) => Divider(
                color: Colors.grey[700],
                height: 1,
              ),
              itemBuilder: (context, index) {
                final service = services[index];
                return _buildServiceItem(service);
              },
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildServiceItem(Map<String, dynamic> service) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 8, horizontal: 12),
      child: Row(
        children: [
          FaIcon(service['icon'], size: 20),
          const SizedBox(width: 12),
          Text(
            service['name'],
            style: const TextStyle(
              fontSize: 16,
              fontWeight: FontWeight.w600,
            ),
          ),
          const Spacer(),
          SizedBox(
            height: 32,
            child: OAuthConnectionButton(
              serviceName: service['name'].toString().toLowerCase(),
              initialProviders: userProviders.providers,
              onSuccess: (){},
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildLogoutButton() {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 12),
      child: AuthButton(
          text: 'Logout',
          onPressed: () async {
            final authService = AuthService.instance;
            await authService.logout();

            if (context.mounted && mounted) {
              context.go(RouteNames.login);
            }
          }
      )
    );
  }
}
