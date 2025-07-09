// lib/success_page.dart
import 'package:flutter/material.dart';
import 'package:mobile/core/constants/constants.dart';

class SuccessPage extends StatelessWidget {
  const SuccessPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: kBackgroundColor,
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 24.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              Container(
                padding: const EdgeInsets.all(4),
                decoration: BoxDecoration(
                  shape: BoxShape.circle,
                  color: Colors.green.withOpacity(0.1),
                  border: Border.all(color: Colors.green, width: 2),
                ),
                child: const Icon(
                  Icons.check,
                  color: Colors.green,
                  size: 60,
                ),
              ),
              const SizedBox(height: 30),
              Text(
                'Successful',
                textAlign: TextAlign.center,
                style: kTitleStyle.copyWith(fontSize: 24),
              ),
              const SizedBox(height: 16),
              Text(
                'Congratulations! Your password has been changed. Click continue to login',
                textAlign: TextAlign.center,
                style: kSubtitleStyle.copyWith(height: 1.5),
              ),
              const SizedBox(height: 40),
              ElevatedButton(
                style: ElevatedButton.styleFrom(
                  backgroundColor: kPrimaryButtonColor,
                  padding: const EdgeInsets.symmetric(vertical: 16),
                  shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
                ),
                onPressed: () {
                  Navigator.pushNamed(context, '/login');
                },
                child: Text('Continue', style: kButtonTextStyle),
              ),
            ],
          ),
        ),
      ),
    );
  }
}