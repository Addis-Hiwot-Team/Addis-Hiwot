import 'package:flutter/material.dart';
import 'package:mobile/core/constants/constants.dart';

class LandingPage extends StatelessWidget {
  const LandingPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          // Background Image
          Positioned.fill(
            child: Image.asset(
              'assets/images/bg_img.png', // your exported Figma background
              fit: BoxFit.cover,
            ),
          ),

          // Foreground Content
          Center(
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 32.0),
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                   Text(
                    'Addis Hiwot',
                    style: TextStyle(
                      fontSize: 40,
                      fontFamily: 'Hurricane',         
                      color: Color(0xFF000000),
                    ),
                  ),

                  const SizedBox(height: 24),

                  Image.asset(
                    'assets/images/logo.png', // Your logo
                    width: 400,
                    height: 200,
                  ),

                  const SizedBox(height: 24),

                  const Text(
                    '"Your Recovery Ally"',
                    style: TextStyle(
                      fontSize: 16,
                      fontStyle: FontStyle.italic,
                      color: Colors.black87,
                    ),
                  ),

                  const SizedBox(height: 40),

                  ElevatedButton(
                    onPressed: () {
                      Navigator.pushNamed(context, '/login');
                    },
                    style: ElevatedButton.styleFrom(
                      backgroundColor: kPrimaryButtonColor, // Custom brown
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12),
                      ),
                      padding: const EdgeInsets.symmetric(horizontal: 40, vertical: 12),
                    ),
                    child: const Text(
                      'Get Started',
                      style: TextStyle(color: Colors.white, fontSize: 16),
                    ),
                  )
                ],
              ),
            ),
          )
        ],
      ),
    );
  }
}
