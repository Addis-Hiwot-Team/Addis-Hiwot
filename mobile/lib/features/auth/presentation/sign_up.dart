// lib/signup_page.dart
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:provider/provider.dart';
import 'package:mobile/features/auth/presentation/providers/auth_provider.dart';
import 'package:mobile/core/constants/constants.dart';

class SignUpPage extends StatefulWidget {
  const SignUpPage({super.key});

  @override
  State<SignUpPage> createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {
  final TextEditingController nameController = TextEditingController();
  final TextEditingController usernameController = TextEditingController();
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();

  @override
  void dispose() {
    nameController.dispose();
    usernameController.dispose();
    emailController.dispose();
    passwordController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final authProvider = context.watch<AuthProvider>();

    return Scaffold(
      backgroundColor: kBackgroundColor,
      appBar: AppBar(
        backgroundColor: kBackgroundColor,
        elevation: 0,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back, color: kTextColor),
          onPressed: () => Navigator.of(context).pop(),
        ),
      ),
      body: SafeArea(
        child: SingleChildScrollView(
          padding: const EdgeInsets.symmetric(horizontal: 48.0, vertical: 24),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text('Sign Up', style: kTitleStyle),
                ],
              ),
              const SizedBox(height: 40),

              // Name TextField
              TextField(
                controller: nameController,
                decoration: const InputDecoration(
                  labelText: 'Full Name',
                  border: OutlineInputBorder(),
                ),
                onChanged: (value) {
                  // Trigger rebuild to update UI
                  setState(() {});
                },
              ),
              const SizedBox(height: 20),

              // Username TextField
              TextField(
                controller: usernameController,
                decoration: const InputDecoration(
                  labelText: 'Username',
                  border: OutlineInputBorder(),
                ),
                onChanged: (value) {
                  // Trigger rebuild to update UI
                  setState(() {});
                },
              ),
              const SizedBox(height: 20),

              // Email TextField
              TextField(
                controller: emailController,
                keyboardType: TextInputType.emailAddress,
                decoration: const InputDecoration(
                  labelText: 'Email',
                  border: OutlineInputBorder(),
                ),
                onChanged: (value) {
                  // Trigger rebuild to update UI
                  setState(() {});
                },
              ),
              const SizedBox(height: 20),

              // Password TextField
              TextField(
                controller: passwordController,
                obscureText: true,
                enableSuggestions: false,
                autocorrect: false,
                decoration: const InputDecoration(
                  labelText: 'Password',
                  border: OutlineInputBorder(),
                  suffixIcon: Icon(Icons.visibility_off_outlined),
                ),
                onChanged: (value) {
                  // Trigger rebuild to update UI
                  setState(() {});
                },
              ),
              const SizedBox(height: 30),

              // Sign Up Button with loading state
              ElevatedButton(
                onPressed: authProvider.isLoading
                    ? null
                    : () async {
                        final success = await authProvider.signup(
                          nameController.text.trim(),
                          usernameController.text.trim(),
                          emailController.text.trim(),
                          passwordController.text,
                        );
                        
                        if (success && mounted) {
                          // Navigate to dashboard or show success message
                          Navigator.of(context).pushReplacementNamed('/dashboard');
                        } else if (mounted && authProvider.error != null) {
                          // Show error message
                          ScaffoldMessenger.of(context).showSnackBar(
                            SnackBar(
                              content: Text(authProvider.error!),
                              backgroundColor: Colors.red,
                            ),
                          );
                        }
                      },
                style: ElevatedButton.styleFrom(
                  backgroundColor: kPrimaryButtonColor,
                  padding: const EdgeInsets.symmetric(vertical: 16),
                  shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(12)),
                ),
                
                child: authProvider.isLoading
                    ? const SizedBox(
                        height: 20,
                        width: 20,
                        child: CircularProgressIndicator(
                          color: Colors.white,
                          strokeWidth: 2,
                        ),
                      )
                    : Text('Sign Up', style: kButtonTextStyle),
              ),

              const SizedBox(height: 30),
              const Row(
                children: [
                  Expanded(child: Divider()),
                  Padding(
                    padding: EdgeInsets.symmetric(horizontal: 8.0),
                    child: Text('Or login with'),
                  ),
                  Expanded(child: Divider()),
                ],
              ),
              const SizedBox(height: 30),

              _buildSocialLoginButton(
                icon: FontAwesomeIcons.facebook,
                text: 'Login with Facebook',
                color: Colors.blue.shade800,
                onPressed: () {},
              ),
              const SizedBox(height: 16),
              _buildSocialLoginButton(
                icon: FontAwesomeIcons.google,
                text: 'Login with Google',
                color: Colors.red.shade700,
                onPressed: () {},
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildSocialLoginButton({
    required IconData icon,
    required String text,
    required Color color,
    required VoidCallback onPressed,
  }) {
    return OutlinedButton.icon(
      icon: FaIcon(icon, color: color, size: 20),
      label: Text(text, style: const TextStyle(color: kTextColor)),
      onPressed: onPressed,
      style: OutlinedButton.styleFrom(
        backgroundColor: kSecondaryButtonColor,
        padding: const EdgeInsets.symmetric(vertical: 14),
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
        side: BorderSide(color: Colors.grey.shade300),
      ),
    );
  }
}
