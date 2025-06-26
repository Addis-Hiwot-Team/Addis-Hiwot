// lib/forgot_password_page.dart
import 'package:flutter/material.dart';
import 'package:mobile/core/constants/constants.dart';

class ForgotPasswordPage extends StatelessWidget {
  const ForgotPasswordPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: kBackgroundColor,
      appBar: AppBar(
        backgroundColor: kBackgroundColor,
        elevation: 0,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back, color: kTextColor),
          onPressed: () => Navigator.of(context).pop(),
        )
      ),
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 24.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              const SizedBox(height: 20),
              Text('Forget Password', style: kTitleStyle),
              const SizedBox(height: 16),
              Text(
                'Please enter your email to reset the password',
                style: kSubtitleStyle,
              ),
              const SizedBox(height: 40),
              const TextField(
                decoration: InputDecoration(labelText: 'Email'),
              ),
              const SizedBox(height: 40),
              ElevatedButton(
                style: ElevatedButton.styleFrom(
                  backgroundColor: kPrimaryButtonColor,
                  padding: const EdgeInsets.symmetric(vertical: 16),
                  shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
                ),
                onPressed: () {
                  Navigator.pushNamed(context, '/verify_email');
                },
                child: Text('Reset Password', style: kButtonTextStyle),
              ),
            ],
          ),
        ),
      ),
    );
  }
}