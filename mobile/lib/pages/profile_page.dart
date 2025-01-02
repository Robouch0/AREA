import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:my_area_flutter/widgets/main_app_scaffold.dart';

class ProfilePage extends StatefulWidget {
  final String email;
  final String firstName;
  final String lastName;
  final String password;

  const ProfilePage({
    super.key,
    required this.email,
    required this.firstName,
    required this.lastName,
    required this.password,
  });

  @override
  State<ProfilePage> createState() => _ProfilePageState();
}

class _ProfilePageState extends State<ProfilePage> {
  bool showPassword = false;
  final List<Map<String, dynamic>> services = [
    {"name": "Github", "icon": FontAwesomeIcons.github},
    {"name": "Google", "icon": FontAwesomeIcons.google},
    {"name": "Twitter", "icon": FontAwesomeIcons.twitter},
    {"name": "Discord", "icon": FontAwesomeIcons.discord},
  ];

  @override
  Widget build(BuildContext context) {
    return MainAppScaffold(
      child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              _buildHeader(),
              _buildDivider(),
              _buildMainContent(),
              const SizedBox(height: 40),
            ],
          )
      ),
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
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.grey[700],
        borderRadius: BorderRadius.circular(30),
      ),
      child: Column(
        children: [
          _buildProfilePicture(),
          _buildInfoField('Email', widget.email),
          _buildInfoField('First name', widget.firstName),
          _buildInfoField('Last name', widget.lastName),
          _buildPasswordField(),
        ],
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
      ),
    );
  }

  Widget _buildInfoField(String label, String value) {
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
            child: Text(
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

  Widget _buildPasswordField() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const Text(
          'Password',
          style: TextStyle(
            color: Colors.white,
            fontSize: 20,
            fontWeight: FontWeight.bold,
          ),
        ),
        const SizedBox(height: 8),
        Stack(
          children: [
            Container(
              width: double.infinity,
              padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
              decoration: BoxDecoration(
                color: Colors.white.withAlpha(204),
                borderRadius: BorderRadius.circular(20),
                border: Border.all(color: Colors.black, width: 4),
              ),
              child: Text(
                showPassword ? widget.password : '•' * widget.password.length,
                style: const TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.w800,
                ),
              ),
            ),
            Positioned(
              right: 16,
              top: 12,
              child: GestureDetector(
                onTap: () => setState(() => showPassword = !showPassword),
                child: Icon(
                  showPassword ? Icons.visibility_off : Icons.visibility,
                  color: Colors.grey[600],
                  size: 24,
                ),
              ),
            ),
          ],
        ),
      ],
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
            child: ElevatedButton(
              onPressed: () {
                print('Service clicked: ${service['name']}');
              },
              child: const Text('Connect'),
            ),
          ),
        ],
      ),
    );
  }
}
